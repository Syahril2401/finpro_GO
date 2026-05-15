<template>
  <DashboardLayout :showBuddy="false">
    <div class="flex h-[calc(100vh-6rem)] -m-10 overflow-hidden bg-white rounded-tl-3xl shadow-inner border-t border-l border-slate-100">
      
      <!-- Left Mini Sidebar -->
      <NotesSidebar 
        :notes="notes"
        :selectedNoteId="selectedNote?.id"
        v-model:searchQuery="filters.search"
        v-model:dimensionFilter="filters.focusDimension"
        :isLoading="isLoadingList"
        @create="openTemplatePicker"
        @select="selectNote"
        @toggle-pin="togglePin"
      />

      <!-- Main Editor -->
      <div v-if="selectedNote" class="flex-1 flex flex-col relative overflow-hidden">
        <NoteEditor 
          v-model:title="selectedNote.title"
          v-model:blocks="editorBlocks"
          :saveStatus="saveStatus"
          @save="saveNote"
        />
      </div>

      <!-- Empty Editor State -->
      <div v-else class="flex-1 flex flex-col items-center justify-center bg-[#FDFDFF]">
        <div class="w-20 h-20 bg-indigo-50 text-[#3D3ACE] rounded-full flex items-center justify-center text-3xl mb-6 shadow-inner border border-indigo-100">📝</div>
        <h3 class="text-xl font-black text-[#1E1B4B] mb-2">No note selected</h3>
        <p class="text-slate-500 font-medium max-w-sm mx-auto mb-8 text-center">Select a note from the sidebar or create a new reflection to start writing.</p>
        <button @click="openTemplatePicker" class="bg-[#3D3ACE] hover:bg-[#322fb0] text-white px-6 py-3 rounded-2xl font-bold shadow-xl shadow-indigo-200 transition-all active:scale-95 flex items-center gap-2">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M12 4v16m8-8H4"/></svg>
          Create New Note
        </button>
      </div>

      <!-- Right Properties Panel -->
      <NotePropertiesPanel 
        v-if="selectedNote"
        :note="selectedNote"
        @update="updateProperty"
        @delete="deleteNote"
        @toggle-archive="toggleArchive"
      />

    </div>

    <!-- Modals -->
    <NoteTemplatePicker 
      :isOpen="isTemplatePickerOpen"
      :isLoading="isLoadingTemplates"
      :templates="templates"
      @close="isTemplatePickerOpen = false"
      @select="createNoteFromTemplate"
    />

  </DashboardLayout>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import { usePage } from '@inertiajs/vue3';
import DashboardLayout from '@/Layouts/DashboardLayout.vue';
import NotesSidebar from '@/Components/Notes/NotesSidebar.vue';
import NoteEditor from '@/Components/Notes/NoteEditor.vue';
import NotePropertiesPanel from '@/Components/Notes/NotePropertiesPanel.vue';
import NoteTemplatePicker from '@/Components/Notes/NoteTemplatePicker.vue';
import { notesApi, setAuthToken } from '@/services/notesApi';

// State
const notes = ref([]);
const selectedNote = ref(null);
const editorBlocks = ref([]);
const templates = ref([]);

const filters = ref({
  search: '',
  focusDimension: 'all'
});

const isLoadingList = ref(false);
const isLoadingTemplates = ref(false);
const isTemplatePickerOpen = ref(false);
const saveStatus = ref('saved'); // 'saved', 'saving', 'unsaved', 'error'
let autosaveTimeout = null;

// Query Params
const page = usePage();
const targetId = new URLSearchParams(window.location.search).get('target_id');
const sessionId = new URLSearchParams(window.location.search).get('session_id');

onMounted(async () => {
  setAuthToken(page.props.go_token);
  await fetchNotes();
  
  if (targetId || sessionId) {
    // Auto open new note modal if arriving from another page
    openTemplatePicker();
  }
});

async function fetchNotes() {
  isLoadingList.value = true;
  try {
    notes.value = await notesApi.getNotes();
  } catch (err) {
    console.error('Failed to fetch notes', err);
  } finally {
    isLoadingList.value = false;
  }
}

async function fetchTemplates() {
  if (templates.value.length > 0) return;
  isLoadingTemplates.value = true;
  try {
    templates.value = await notesApi.getTemplates();
  } catch (err) {
    console.error('Failed to fetch templates', err);
    // Fallback
    templates.value = [{ id: 'blank', title: 'Blank Note', content_json: '{"blocks":[]}' }];
  } finally {
    isLoadingTemplates.value = false;
  }
}

async function openTemplatePicker() {
  isTemplatePickerOpen.value = true;
  await fetchTemplates();
}

async function createNoteFromTemplate(tpl) {
  isTemplatePickerOpen.value = false;
  
  const blocks = tpl.content_json ? JSON.parse(tpl.content_json).blocks : [];
  
  const payload = {
    title: tpl.title === 'Blank Note' ? 'Untitled' : tpl.title,
    content_json: JSON.stringify({ blocks }),
    content_text: tpl.content_text || '',
    focus_dimension: 'General',
    target_id: targetId || null,
    planner_session_id: sessionId || null
  };
  
  try {
    const newNote = await notesApi.createNote(payload);
    notes.value.unshift(newNote);
    selectNote(newNote.id);
  } catch (err) {
    console.error('Create failed', err);
    alert('Failed to create note.');
  }
}

async function selectNote(id) {
  const note = notes.value.find(n => n.id === id);
  if (!note) return;
  
  // Try to fetch full details if needed, but list might have it all.
  try {
    const fullNote = await notesApi.getNote(id);
    selectedNote.value = fullNote;
    editorBlocks.value = fullNote.content_json ? JSON.parse(fullNote.content_json).blocks || [] : [];
    saveStatus.value = 'saved';
  } catch (err) {
    console.error('Failed to load note details', err);
  }
}

function updateProperty(field, value) {
  if (!selectedNote.value) return;
  selectedNote.value[field] = value;
  triggerAutosave();
}

async function togglePin(id) {
  try {
    const updated = await notesApi.togglePin(id);
    const index = notes.value.findIndex(n => n.id === id);
    if (index !== -1) notes.value[index] = updated;
    if (selectedNote.value?.id === id) selectedNote.value.is_pinned = updated.is_pinned;
  } catch (err) {
    console.error('Pin toggle failed', err);
  }
}

async function toggleArchive() {
  if (!selectedNote.value) return;
  try {
    const updated = await notesApi.toggleArchive(selectedNote.value.id);
    const index = notes.value.findIndex(n => n.id === updated.id);
    if (index !== -1) notes.value[index] = updated;
    selectedNote.value = null; // deselect after archive
  } catch (err) {
    console.error('Archive toggle failed', err);
  }
}

async function deleteNote() {
  if (!selectedNote.value) return;
  if (!confirm('Are you sure you want to delete this note?')) return;
  
  try {
    await notesApi.deleteNote(selectedNote.value.id);
    notes.value = notes.value.filter(n => n.id !== selectedNote.value.id);
    selectedNote.value = null;
  } catch (err) {
    console.error('Delete failed', err);
    alert('Failed to delete note.');
  }
}

function extractText(blocks) {
  return blocks.map(b => b.content).join('\n');
}

function triggerAutosave() {
  if (saveStatus.value !== 'saving') {
    saveStatus.value = 'unsaved';
  }
  
  if (autosaveTimeout) clearTimeout(autosaveTimeout);
  
  autosaveTimeout = setTimeout(() => {
    saveNote();
  }, 1000);
}

async function saveNote() {
  if (!selectedNote.value) return;
  
  saveStatus.value = 'saving';
  
  const payload = {
    title: selectedNote.value.title,
    content_json: JSON.stringify({ blocks: editorBlocks.value }),
    content_text: extractText(editorBlocks.value),
    focus_dimension: selectedNote.value.focus_dimension,
    tags: selectedNote.value.tags,
    mood: selectedNote.value.mood,
    confidence_level: selectedNote.value.confidence_level ? parseInt(selectedNote.value.confidence_level) : null,
    target_id: selectedNote.value.target_id,
    planner_session_id: selectedNote.value.planner_session_id
  };

  try {
    const updated = await notesApi.updateNote(selectedNote.value.id, payload);
    const index = notes.value.findIndex(n => n.id === updated.id);
    if (index !== -1) notes.value[index] = updated;
    selectedNote.value = updated;
    saveStatus.value = 'saved';
  } catch (err) {
    console.error('Save failed', err);
    saveStatus.value = 'error';
  }
}

watch(() => selectedNote.value?.title, () => {
  if (selectedNote.value) triggerAutosave();
});

watch(editorBlocks, () => {
  if (selectedNote.value) triggerAutosave();
}, { deep: true });

</script>
