<template>
  <aside class="w-80 bg-slate-50 border-r border-slate-100 flex flex-col h-full z-30">
    <div class="p-6 pb-4 border-b border-slate-100 bg-white">
      <div class="flex items-center justify-between mb-4">
        <h2 class="text-lg font-black text-[#1E1B4B]">My Notes</h2>
        <button @click="$emit('create')" class="bg-[#3D3ACE] text-white w-8 h-8 rounded-xl flex items-center justify-center hover:bg-[#322fb0] hover:shadow-lg transition-all" title="New Note">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M12 4v16m8-8H4"/></svg>
        </button>
      </div>
      
      <!-- Search -->
      <div class="relative mb-3">
        <svg class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/></svg>
        <input 
          type="text" 
          placeholder="Search notes..." 
          :value="searchQuery"
          @input="$emit('update:searchQuery', $event.target.value)"
          class="w-full bg-slate-50 border border-slate-200 rounded-xl pl-9 pr-3 py-2 text-sm font-bold text-[#1E1B4B] focus:bg-white focus:border-[#3D3ACE] outline-none transition-colors"
        >
      </div>

      <!-- Filters -->
      <div class="flex gap-2">
        <select 
          :value="dimensionFilter"
          @change="$emit('update:dimensionFilter', $event.target.value)"
          class="flex-1 bg-slate-50 border border-slate-200 rounded-xl px-3 py-2 text-xs font-bold text-slate-600 focus:border-[#3D3ACE] outline-none cursor-pointer"
        >
          <option value="all">All Dimensions</option>
          <option value="Planning">Planning</option>
          <option value="Time Management">Time Management</option>
          <option value="Cognitive Strategy">Cognitive Strategy</option>
          <option value="Reflection">Reflection</option>
          <option value="General">General</option>
        </select>
      </div>
    </div>

    <!-- Notes List -->
    <div class="flex-1 overflow-y-auto p-4 space-y-6 scrollbar-hide">
      
      <div v-if="isLoading" class="space-y-3">
        <div v-for="i in 3" :key="i" class="h-24 bg-slate-200/50 rounded-2xl animate-pulse"></div>
      </div>

      <div v-else-if="filteredNotes.length === 0" class="text-center py-10">
        <p class="text-slate-400 text-xs font-bold mb-3">No notes found.</p>
        <button @click="$emit('create')" class="text-[#3D3ACE] text-xs font-black uppercase tracking-widest border-b border-[#3D3ACE]">Create Note</button>
      </div>

      <template v-else>
        <!-- Pinned Notes -->
        <div v-if="pinnedNotes.length > 0">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-3 pl-2">Pinned</p>
          <div class="space-y-3">
            <NoteCard 
              v-for="note in pinnedNotes" 
              :key="note.id" 
              :note="note"
              :isActive="selectedNoteId === note.id"
              @click="$emit('select', note.id)"
              @toggle-pin="$emit('toggle-pin', note.id)"
            />
          </div>
        </div>

        <!-- Recent Notes -->
        <div v-if="unpinnedNotes.length > 0">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-3 pl-2 mt-6">Recent</p>
          <div class="space-y-3">
            <NoteCard 
              v-for="note in unpinnedNotes" 
              :key="note.id" 
              :note="note"
              :isActive="selectedNoteId === note.id"
              @click="$emit('select', note.id)"
              @toggle-pin="$emit('toggle-pin', note.id)"
            />
          </div>
        </div>
      </template>

    </div>
  </aside>
</template>

<script setup>
import { computed } from 'vue';
import NoteCard from './NoteCard.vue';

const props = defineProps({
  notes: Array,
  selectedNoteId: String,
  searchQuery: String,
  dimensionFilter: String,
  isLoading: Boolean
});

defineEmits(['create', 'select', 'update:searchQuery', 'update:dimensionFilter', 'toggle-pin']);

const filteredNotes = computed(() => {
  return props.notes.filter(note => {
    // Exclude archived from main view unless we add an archive filter later
    if (note.is_archived) return false;
    
    // Dimension filter
    if (props.dimensionFilter !== 'all' && note.focus_dimension !== props.dimensionFilter) return false;
    
    // Search filter
    if (props.searchQuery) {
      const q = props.searchQuery.toLowerCase();
      const title = (note.title || '').toLowerCase();
      const content = (note.content_text || '').toLowerCase();
      const tags = (note.tags || '').toLowerCase();
      if (!title.includes(q) && !content.includes(q) && !tags.includes(q)) return false;
    }
    
    return true;
  });
});

const pinnedNotes = computed(() => filteredNotes.value.filter(n => n.is_pinned));
const unpinnedNotes = computed(() => filteredNotes.value.filter(n => !n.is_pinned));
</script>
