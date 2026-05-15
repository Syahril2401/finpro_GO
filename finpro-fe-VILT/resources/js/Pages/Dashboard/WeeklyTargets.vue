<template>
  <DashboardLayout 
    :showBuddy="true"
    buddyContextTitle="Target Insight" 
    :buddyContextMessage="insightMessage"
  >
    <!-- Header -->
    <div class="mb-10 flex flex-col md:flex-row md:items-end justify-between gap-6">
      <div>
        <h1 class="text-3xl font-black text-[#1E1B4B] mb-2">Weekly Targets</h1>
        <p class="text-slate-500 font-medium">Set measurable goals and track completion.</p>
      </div>
      <button @click="openTargetModal()" class="bg-[#3D3ACE] hover:bg-[#322fb0] text-white px-6 py-3 rounded-2xl font-bold shadow-xl shadow-indigo-200 transition-all active:scale-95 whitespace-nowrap flex items-center gap-2">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M12 4v16m8-8H4"/></svg>
        Add Weekly Target
      </button>
    </div>

    <!-- Summary Cards -->
    <div class="grid grid-cols-2 lg:grid-cols-4 gap-4 mb-10">
      <div class="bg-white p-6 rounded-[28px] border border-slate-100 shadow-sm">
        <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">Total Targets</p>
        <p class="text-3xl font-black text-[#1E1B4B]">{{ summary.total_targets || 0 }}</p>
      </div>
      <div class="bg-white p-6 rounded-[28px] border border-slate-100 shadow-sm">
        <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">Completed</p>
        <p class="text-3xl font-black text-emerald-500">{{ summary.completed_targets || 0 }}</p>
      </div>
      <div class="bg-white p-6 rounded-[28px] border border-slate-100 shadow-sm">
        <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">Completion Rate</p>
        <p class="text-3xl font-black text-[#3D3ACE]">{{ Math.round(summary.completion_rate || 0) }}%</p>
      </div>
      <div class="bg-white p-6 rounded-[28px] border border-slate-100 shadow-sm">
        <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">Primary Focus</p>
        <p class="text-lg font-black text-[#1E1B4B] truncate">{{ summary.primary_focus || 'No focus yet' }}</p>
      </div>
    </div>

    <!-- Loading -->
    <div v-if="isLoading" class="space-y-4">
      <div v-for="i in 3" :key="i" class="h-32 bg-slate-100 rounded-[28px] animate-pulse"></div>
    </div>

    <!-- Empty State -->
    <div v-else-if="targets.length === 0" class="bg-slate-50 border-2 border-dashed border-slate-200 p-12 rounded-[40px] text-center">
      <div class="text-3xl mb-4">🎯</div>
      <p class="text-slate-500 font-bold mb-2">No targets set yet.</p>
      <p class="text-slate-400 text-sm">Create your first weekly target to start tracking your learning progress.</p>
      <button @click="openTargetModal()" class="mt-4 text-[#3D3ACE] font-bold hover:underline">+ Create Target</button>
    </div>

    <!-- Target Cards -->
    <div v-else class="space-y-6">
      <div v-for="target in targets" :key="target.id" class="bg-white p-6 rounded-[28px] border border-slate-100 shadow-sm hover:shadow-lg transition-all">
        <!-- Target Header -->
        <div class="flex items-start justify-between mb-4">
          <div class="flex-1">
            <div class="flex items-center gap-2 mb-2 flex-wrap">
              <span class="text-[10px] font-black uppercase tracking-widest px-2 py-1 rounded-md" :class="priorityClass(target.priority)">{{ target.priority }}</span>
              <span class="text-[10px] font-black bg-indigo-50 text-[#3D3ACE] uppercase tracking-widest px-2 py-1 rounded-md">{{ target.focus_dimension }}</span>
              <span class="text-[10px] font-black uppercase tracking-widest px-2 py-1 rounded-md" :class="statusBadge(target.status)">{{ target.status?.replace('_', ' ') }}</span>
            </div>
            <h4 class="text-lg font-black text-[#1E1B4B]">{{ target.title }}</h4>
            <p v-if="target.description" class="text-sm font-medium text-slate-500 mt-1">{{ target.description }}</p>
            <p v-if="target.due_date" class="text-xs font-bold text-slate-400 mt-1">Due: {{ target.due_date }}</p>
          </div>
          <div class="flex items-center gap-2 ml-4">
            <button @click="openTargetModal(target)" class="w-9 h-9 rounded-xl bg-slate-50 text-slate-400 hover:bg-indigo-50 hover:text-[#3D3ACE] flex items-center justify-center transition-colors">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/></svg>
            </button>
            <a :href="`/dashboard/planner?target_id=${target.id}`" class="w-9 h-9 rounded-xl bg-slate-50 text-slate-400 hover:bg-blue-50 hover:text-blue-500 flex items-center justify-center transition-colors" title="Schedule Session">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"/></svg>
            </a>
            <a :href="`/dashboard/notes?target_id=${target.id}`" class="w-9 h-9 rounded-xl bg-slate-50 text-slate-400 hover:bg-purple-50 hover:text-purple-500 flex items-center justify-center transition-colors" title="Add Reflection">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z"/></svg>
            </a>
            <button @click="deleteTarget(target.id)" class="w-9 h-9 rounded-xl bg-slate-50 text-slate-400 hover:bg-rose-50 hover:text-rose-500 flex items-center justify-center transition-colors">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/></svg>
            </button>
          </div>
        </div>

        <!-- Progress Bar -->
        <div class="mb-4">
          <div class="flex items-center justify-between mb-1">
            <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Progress</span>
            <span class="text-sm font-black text-[#1E1B4B]">{{ target.progress || 0 }}%</span>
          </div>
          <div class="w-full h-2 bg-slate-100 rounded-full overflow-hidden">
            <div class="h-full rounded-full transition-all duration-500" :class="target.progress >= 100 ? 'bg-emerald-500' : 'bg-[#3D3ACE]'" :style="{ width: `${target.progress || 0}%` }"></div>
          </div>
        </div>

        <!-- Subtasks -->
        <div class="space-y-2">
          <div v-for="sub in target.subtasks" :key="sub.id" class="flex items-center gap-3 group">
            <button @click="toggleSubtask(target.id, sub.id)" class="w-5 h-5 rounded-md border-2 flex items-center justify-center flex-shrink-0 transition-all"
              :class="sub.is_completed ? 'bg-emerald-500 border-emerald-500' : 'border-slate-300 hover:border-[#3D3ACE]'">
              <svg v-if="sub.is_completed" class="w-3 h-3 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7"/></svg>
            </button>
            <span class="text-sm font-bold flex-1" :class="sub.is_completed ? 'text-slate-400 line-through' : 'text-[#1E1B4B]'">{{ sub.title }}</span>
            <button @click="deleteSubtask(target.id, sub.id)" class="opacity-0 group-hover:opacity-100 w-6 h-6 rounded text-slate-400 hover:text-rose-500 flex items-center justify-center transition-all">
              <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/></svg>
            </button>
          </div>

          <!-- Add Subtask -->
          <div class="flex items-center gap-3 pt-2">
            <input v-model="newSubtask[target.id]" @keydown.enter="addSubtask(target.id)" placeholder="Add subtask..." class="flex-1 text-sm font-bold text-[#1E1B4B] bg-transparent border-b border-dashed border-slate-200 py-1 focus:border-[#3D3ACE] outline-none transition-colors placeholder:text-slate-300">
            <button v-if="newSubtask[target.id]" @click="addSubtask(target.id)" class="text-[#3D3ACE] font-black text-xs hover:underline">Add</button>
          </div>
        </div>
      </div>
    </div>

    <!-- Target Modal -->
    <Teleport to="body">
      <div v-if="showTargetModal" class="fixed inset-0 bg-black/40 backdrop-blur-sm z-50 flex items-center justify-center p-4" @click.self="showTargetModal = false">
        <div class="bg-white rounded-[32px] p-8 w-full max-w-lg shadow-2xl max-h-[90vh] overflow-y-auto">
          <h2 class="text-xl font-black text-[#1E1B4B] mb-6">{{ editingTarget ? 'Edit Target' : 'Add Weekly Target' }}</h2>
          <form @submit.prevent="saveTarget" class="space-y-5">
            <div>
              <label class="block text-[11px] font-black text-slate-400 uppercase tracking-widest mb-2">Title *</label>
              <input v-model="targetForm.title" required class="w-full bg-slate-50 border border-slate-200 rounded-2xl py-3 px-5 text-sm font-bold text-[#1E1B4B] focus:border-[#3D3ACE] outline-none transition-all" placeholder="e.g. Complete Chapter 5">
            </div>
            <div>
              <label class="block text-[11px] font-black text-slate-400 uppercase tracking-widest mb-2">Description</label>
              <textarea v-model="targetForm.description" rows="2" class="w-full bg-slate-50 border border-slate-200 rounded-2xl py-3 px-5 text-sm font-bold text-[#1E1B4B] focus:border-[#3D3ACE] outline-none resize-none transition-all"></textarea>
            </div>
            <div class="grid grid-cols-2 gap-4">
              <div>
                <label class="block text-[11px] font-black text-slate-400 uppercase tracking-widest mb-2">Focus Dimension</label>
                <select v-model="targetForm.focus_dimension" class="w-full bg-slate-50 border border-slate-200 rounded-2xl py-3 px-5 text-sm font-bold text-[#1E1B4B] focus:border-[#3D3ACE] outline-none">
                  <option v-for="d in dimensions" :key="d" :value="d">{{ d }}</option>
                </select>
              </div>
              <div>
                <label class="block text-[11px] font-black text-slate-400 uppercase tracking-widest mb-2">Priority</label>
                <select v-model="targetForm.priority" class="w-full bg-slate-50 border border-slate-200 rounded-2xl py-3 px-5 text-sm font-bold text-[#1E1B4B] focus:border-[#3D3ACE] outline-none">
                  <option value="low">Low</option>
                  <option value="medium">Medium</option>
                  <option value="high">High</option>
                </select>
              </div>
            </div>
            <div>
              <label class="block text-[11px] font-black text-slate-400 uppercase tracking-widest mb-2">Due Date</label>
              <input v-model="targetForm.due_date" type="date" class="w-full bg-slate-50 border border-slate-200 rounded-2xl py-3 px-5 text-sm font-bold text-[#1E1B4B] focus:border-[#3D3ACE] outline-none transition-all">
            </div>
            <!-- Initial Subtasks (create mode only) -->
            <div v-if="!editingTarget">
              <label class="block text-[11px] font-black text-slate-400 uppercase tracking-widest mb-2">Initial Subtasks</label>
              <div v-for="(st, i) in targetForm.subtasks" :key="i" class="flex items-center gap-2 mb-2">
                <input v-model="st.title" class="flex-1 bg-slate-50 border border-slate-200 rounded-xl py-2 px-4 text-sm font-bold text-[#1E1B4B] focus:border-[#3D3ACE] outline-none" :placeholder="`Subtask ${i + 1}`">
                <button type="button" @click="targetForm.subtasks.splice(i, 1)" class="text-slate-400 hover:text-rose-500">&times;</button>
              </div>
              <button type="button" @click="targetForm.subtasks.push({ title: '' })" class="text-sm text-[#3D3ACE] font-bold hover:underline">+ Add subtask</button>
            </div>
            <div class="flex gap-3 pt-4">
              <button type="button" @click="showTargetModal = false" class="flex-1 bg-slate-100 text-slate-600 py-3 rounded-2xl font-bold hover:bg-slate-200">Cancel</button>
              <button type="submit" :disabled="isSaving" class="flex-1 bg-[#3D3ACE] text-white py-3 rounded-2xl font-bold hover:bg-[#322fb0] shadow-lg shadow-indigo-200 disabled:opacity-50">
                {{ isSaving ? 'Saving...' : (editingTarget ? 'Update' : 'Create') }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </Teleport>

    <!-- Toast -->
    <Teleport to="body">
      <div v-if="toast" class="fixed bottom-6 right-6 z-50 bg-white border border-slate-200 shadow-2xl rounded-2xl p-4 flex items-center gap-3">
        <div class="w-8 h-8 rounded-lg flex items-center justify-center" :class="toast.type === 'error' ? 'bg-rose-50 text-rose-500' : 'bg-emerald-50 text-emerald-500'">
          <svg v-if="toast.type === 'error'" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/></svg>
          <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7"/></svg>
        </div>
        <span class="text-sm font-bold text-[#1E1B4B]">{{ toast.message }}</span>
      </div>
    </Teleport>

  </DashboardLayout>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { usePage } from '@inertiajs/vue3'
import DashboardLayout from '@/Layouts/DashboardLayout.vue'
import { targetsApi, setWorkspaceToken } from '@/services/workspaceApi'

const page = usePage()
const isLoading = ref(true)
const isSaving = ref(false)
const showTargetModal = ref(false)
const editingTarget = ref(null)
const targets = ref([])
const summary = ref({})
const newSubtask = reactive({})
const toast = ref(null)

const dimensions = ['General', 'Planning', 'Time Management', 'Cognitive Strategy', 'Reflection']

const targetForm = ref({
  title: '', description: '', focus_dimension: 'General',
  priority: 'medium', due_date: '', subtasks: []
})

const insightMessage = computed(() => {
  if (targets.value.length === 0) return 'Create your first weekly target to make your learning goals measurable.'
  const totalSubs = targets.value.reduce((a, t) => a + (t.subtasks?.length || 0), 0)
  if (totalSubs < 3) return 'Break large goals into smaller subtasks to make progress easier to track.'
  return `You have ${targets.value.length} targets with ${totalSubs} subtasks. Keep completing them!`
})

onMounted(async () => {
  setWorkspaceToken(page.props.go_token)
  await fetchTargets()
})

async function fetchTargets() {
  isLoading.value = true
  try {
    const res = await targetsApi.getTargets()
    targets.value = res.data || []
    summary.value = res.summary || {}
  } catch (err) {
    console.error('Failed to fetch targets', err)
    targets.value = []
  } finally {
    isLoading.value = false
  }
}

function openTargetModal(target = null) {
  editingTarget.value = target
  if (target) {
    targetForm.value = { title: target.title, description: target.description, focus_dimension: target.focus_dimension, priority: target.priority, due_date: target.due_date || '', subtasks: [] }
  } else {
    targetForm.value = { title: '', description: '', focus_dimension: 'General', priority: 'medium', due_date: '', subtasks: [{ title: '' }] }
  }
  showTargetModal.value = true
}

async function saveTarget() {
  isSaving.value = true
  try {
    if (editingTarget.value) {
      await targetsApi.updateTarget(editingTarget.value.id, targetForm.value)
      showToast('Target updated!')
    } else {
      await targetsApi.createTarget(targetForm.value)
      showToast('Target created!')
    }
    showTargetModal.value = false
    await fetchTargets()
  } catch (err) {
    showToast('Failed to save target.', 'error')
  } finally {
    isSaving.value = false
  }
}

async function deleteTarget(id) {
  if (!confirm('Delete this target and all its subtasks?')) return
  try {
    await targetsApi.deleteTarget(id)
    targets.value = targets.value.filter(t => t.id !== id)
    showToast('Target deleted.')
    await fetchTargets()
  } catch (err) {
    showToast('Failed to delete.', 'error')
  }
}

async function toggleSubtask(targetId, subtaskId) {
  try {
    await targetsApi.toggleSubtask(targetId, subtaskId)
    await fetchTargets()
  } catch (err) {
    showToast('Failed to toggle subtask.', 'error')
  }
}

async function addSubtask(targetId) {
  const title = newSubtask[targetId]?.trim()
  if (!title) return
  try {
    await targetsApi.createSubtask(targetId, { title })
    newSubtask[targetId] = ''
    await fetchTargets()
    showToast('Subtask added!')
  } catch (err) {
    showToast('Failed to add subtask.', 'error')
  }
}

async function deleteSubtask(targetId, subtaskId) {
  try {
    await targetsApi.deleteSubtask(targetId, subtaskId)
    await fetchTargets()
  } catch (err) {
    showToast('Failed to delete subtask.', 'error')
  }
}

function showToast(message, type = 'success') {
  toast.value = { message, type }
  setTimeout(() => toast.value = null, 3000)
}

function priorityClass(p) {
  const map = { high: 'bg-rose-50 text-rose-500', medium: 'bg-amber-50 text-amber-500', low: 'bg-slate-50 text-slate-400' }
  return map[p] || 'bg-slate-50 text-slate-400'
}

function statusBadge(s) {
  const map = { completed: 'bg-emerald-50 text-emerald-500', in_progress: 'bg-blue-50 text-blue-500', not_started: 'bg-slate-50 text-slate-400', paused: 'bg-amber-50 text-amber-500' }
  return map[s] || 'bg-slate-50 text-slate-400'
}
</script>
