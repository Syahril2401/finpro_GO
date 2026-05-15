<template>
  <DashboardLayout 
    :showBuddy="true"
    buddyContextTitle="Planner Insight" 
    :buddyContextMessage="insightMessage"
  >
    <!-- Page Header -->
    <div class="mb-10 flex flex-col md:flex-row md:items-end justify-between gap-6">
      <div>
        <h1 class="text-3xl font-black text-[#1E1B4B] mb-2">Planner</h1>
        <p class="text-slate-500 font-medium">Design your study rhythm for this week.</p>
      </div>
      <button @click="openModal()" class="bg-[#3D3ACE] hover:bg-[#322fb0] text-white px-6 py-3 rounded-2xl font-bold shadow-xl shadow-indigo-200 transition-all active:scale-95 whitespace-nowrap flex items-center gap-2">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M12 4v16m8-8H4"/></svg>
        Add Study Session
      </button>
    </div>

    <!-- Weekly Calendar View -->
    <div class="bg-white p-6 rounded-[32px] border border-slate-100 shadow-sm mb-10 overflow-x-auto">
      <div class="flex items-center justify-between min-w-[700px]">
        <div v-for="day in weekDays" :key="day.dateStr" 
          @click="selectDay(day)"
          class="flex flex-col items-center p-4 rounded-2xl transition-colors cursor-pointer flex-1"
          :class="day.dateStr === selectedDate ? 'bg-[#3D3ACE] text-white shadow-lg shadow-indigo-100' : 'hover:bg-slate-50 text-[#1E1B4B]'">
          <span class="text-[10px] font-black uppercase tracking-widest mb-1" :class="day.dateStr === selectedDate ? 'text-indigo-200' : 'text-slate-400'">{{ day.name }}</span>
          <span class="text-2xl font-black mb-3">{{ day.date }}</span>
          <div class="w-full flex justify-center gap-1">
            <div v-for="i in Math.min(day.count, 3)" :key="i" class="w-1.5 h-1.5 rounded-full" :class="day.dateStr === selectedDate ? 'bg-white' : 'bg-emerald-400'"></div>
            <div v-if="day.count === 0" class="w-1.5 h-1.5 rounded-full bg-slate-200"></div>
          </div>
        </div>
      </div>
    </div>

    <!-- Sessions for Selected Day -->
    <div class="mb-6 flex items-center justify-between">
      <h3 class="text-xl font-black text-[#1E1B4B]">{{ selectedDateLabel }}</h3>
      <span class="text-[11px] font-black text-slate-400 uppercase tracking-widest">{{ daySessions.length }} Scheduled</span>
    </div>

    <!-- Loading -->
    <div v-if="isLoading" class="space-y-4">
      <div v-for="i in 2" :key="i" class="h-24 bg-slate-100 rounded-[28px] animate-pulse"></div>
    </div>

    <!-- Empty State -->
    <div v-else-if="daySessions.length === 0" class="bg-slate-50 border-2 border-dashed border-slate-200 p-12 rounded-[40px] text-center">
      <div class="text-3xl mb-4">📅</div>
      <p class="text-slate-500 font-bold">No sessions scheduled for this day.</p>
      <button @click="openModal()" class="mt-4 text-[#3D3ACE] font-bold hover:underline">+ Create one</button>
    </div>

    <!-- Session Cards -->
    <div v-else class="space-y-4">
      <div v-for="session in daySessions" :key="session.id" class="bg-white p-6 rounded-[28px] border border-slate-100 shadow-sm hover:border-[#3D3ACE]/30 hover:shadow-lg transition-all flex flex-col md:flex-row md:items-center gap-6 group">
        <div class="w-16 h-16 rounded-2xl bg-indigo-50 text-[#3D3ACE] flex flex-col items-center justify-center flex-shrink-0">
          <span class="text-xs font-black">{{ session.start_time }}</span>
        </div>
        <div class="flex-1">
          <div class="flex items-center gap-2 mb-1 flex-wrap">
            <span class="text-[10px] font-black uppercase tracking-widest px-2 py-1 rounded-md"
              :class="statusClass(session.status)">
              {{ session.status }}
            </span>
            <span class="text-[10px] font-black bg-indigo-50 text-[#3D3ACE] uppercase tracking-widest px-2 py-1 rounded-md">{{ session.focus_dimension }}</span>
          </div>
          <h4 class="text-lg font-black text-[#1E1B4B] group-hover:text-[#3D3ACE] transition-colors">{{ session.title }}</h4>
          <p class="text-sm font-medium text-slate-500 mt-1">{{ session.duration_minutes }} mins{{ session.description ? ' • ' + session.description : '' }}</p>
        </div>
        <div class="flex items-center gap-3">
          <button v-if="session.status === 'planned'" @click="completeSession(session.id)" title="Complete" class="w-10 h-10 rounded-xl bg-emerald-50 text-emerald-500 hover:bg-emerald-500 hover:text-white flex items-center justify-center transition-colors">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7"/></svg>
          </button>
          <button v-if="session.status === 'planned'" @click="skipSession(session.id)" title="Skip" class="w-10 h-10 rounded-xl bg-amber-50 text-amber-500 hover:bg-amber-500 hover:text-white flex items-center justify-center transition-colors">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 5l7 7-7 7M5 5l7 7-7 7"/></svg>
          </button>
          <button @click="openModal(session)" title="Edit" class="w-10 h-10 rounded-xl bg-slate-50 text-slate-400 hover:bg-indigo-50 hover:text-[#3D3ACE] flex items-center justify-center transition-colors">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/></svg>
          </button>
          <a :href="`/dashboard/notes?session_id=${session.id}`" title="Write Reflection" class="w-10 h-10 rounded-xl bg-slate-50 text-slate-400 hover:bg-purple-50 hover:text-purple-500 flex items-center justify-center transition-colors">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z"/></svg>
          </a>
          <button @click="deleteSession(session.id)" title="Delete" class="w-10 h-10 rounded-xl bg-slate-50 text-slate-400 hover:bg-rose-50 hover:text-rose-500 flex items-center justify-center transition-colors">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/></svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Add/Edit Modal -->
    <Teleport to="body">
      <div v-if="showModal" class="fixed inset-0 bg-black/40 backdrop-blur-sm z-50 flex items-center justify-center p-4" @click.self="showModal = false">
        <div class="bg-white rounded-[32px] p-8 w-full max-w-lg shadow-2xl">
          <h2 class="text-xl font-black text-[#1E1B4B] mb-6">{{ editingSession ? 'Edit Session' : 'Add Study Session' }}</h2>
          <form @submit.prevent="saveSession" class="space-y-5">
            <div>
              <label class="block text-[11px] font-black text-slate-400 uppercase tracking-widest mb-2">Title *</label>
              <input v-model="form.title" required class="w-full bg-slate-50 border border-slate-200 rounded-2xl py-3 px-5 text-sm font-bold text-[#1E1B4B] focus:border-[#3D3ACE] focus:ring-2 focus:ring-indigo-100 outline-none transition-all" placeholder="e.g. Deep Work: React">
            </div>
            <div>
              <label class="block text-[11px] font-black text-slate-400 uppercase tracking-widest mb-2">Description</label>
              <textarea v-model="form.description" rows="2" class="w-full bg-slate-50 border border-slate-200 rounded-2xl py-3 px-5 text-sm font-bold text-[#1E1B4B] focus:border-[#3D3ACE] focus:ring-2 focus:ring-indigo-100 outline-none transition-all resize-none" placeholder="Optional details..."></textarea>
            </div>
            <div class="grid grid-cols-3 gap-4">
              <div>
                <label class="block text-[11px] font-black text-slate-400 uppercase tracking-widest mb-2">Date *</label>
                <input v-model="form.date" type="date" required class="w-full bg-slate-50 border border-slate-200 rounded-2xl py-3 px-4 text-sm font-bold text-[#1E1B4B] focus:border-[#3D3ACE] outline-none transition-all">
              </div>
              <div>
                <label class="block text-[11px] font-black text-slate-400 uppercase tracking-widest mb-2">Start *</label>
                <input v-model="form.start_time" type="time" required class="w-full bg-slate-50 border border-slate-200 rounded-2xl py-3 px-4 text-sm font-bold text-[#1E1B4B] focus:border-[#3D3ACE] outline-none transition-all">
              </div>
              <div>
                <label class="block text-[11px] font-black text-slate-400 uppercase tracking-widest mb-2">End *</label>
                <input v-model="form.end_time" type="time" required class="w-full bg-slate-50 border border-slate-200 rounded-2xl py-3 px-4 text-sm font-bold text-[#1E1B4B] focus:border-[#3D3ACE] outline-none transition-all">
              </div>
            </div>
            <div>
              <label class="block text-[11px] font-black text-slate-400 uppercase tracking-widest mb-2">Focus Dimension</label>
              <select v-model="form.focus_dimension" class="w-full bg-slate-50 border border-slate-200 rounded-2xl py-3 px-5 text-sm font-bold text-[#1E1B4B] focus:border-[#3D3ACE] outline-none transition-all">
                <option v-for="d in dimensions" :key="d" :value="d">{{ d }}</option>
              </select>
            </div>
            <div class="flex gap-3 pt-4">
              <button type="button" @click="showModal = false" class="flex-1 bg-slate-100 text-slate-600 py-3 rounded-2xl font-bold hover:bg-slate-200 transition-colors">Cancel</button>
              <button type="submit" :disabled="isSaving" class="flex-1 bg-[#3D3ACE] text-white py-3 rounded-2xl font-bold hover:bg-[#322fb0] shadow-lg shadow-indigo-200 transition-all disabled:opacity-50">
                {{ isSaving ? 'Saving...' : (editingSession ? 'Update' : 'Create') }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </Teleport>

    <!-- Error Toast -->
    <Teleport to="body">
      <div v-if="toast" class="fixed bottom-6 right-6 z-50 bg-white border border-slate-200 shadow-2xl rounded-2xl p-4 flex items-center gap-3 animate-slide-up">
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
import { ref, computed, onMounted } from 'vue'
import { usePage } from '@inertiajs/vue3'
import DashboardLayout from '@/Layouts/DashboardLayout.vue'
import { plannerApi, setWorkspaceToken } from '@/services/workspaceApi'

const page = usePage()
const isLoading = ref(true)
const isSaving = ref(false)
const showModal = ref(false)
const editingSession = ref(null)
const sessions = ref([])
const selectedDate = ref('')
const toast = ref(null)

const dimensions = ['General', 'Planning', 'Time Management', 'Cognitive Strategy', 'Reflection']

const form = ref({
  title: '', description: '', date: '', start_time: '', end_time: '',
  focus_dimension: 'General', target_id: null
})

// Initialize
const today = new Date()
selectedDate.value = formatDate(today)

// Check for target_id from URL
const urlTargetId = new URLSearchParams(window.location.search).get('target_id')

// Week calendar
const weekDays = computed(() => {
  const monday = getMonday(new Date(selectedDate.value || today))
  return Array.from({ length: 7 }, (_, i) => {
    const d = new Date(monday)
    d.setDate(d.getDate() + i)
    const dateStr = formatDate(d)
    const count = sessions.value.filter(s => s.date === dateStr).length
    return {
      name: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'][i],
      date: d.getDate(),
      dateStr,
      count
    }
  })
})

const daySessions = computed(() => sessions.value.filter(s => s.date === selectedDate.value))
const selectedDateLabel = computed(() => {
  const d = new Date(selectedDate.value)
  return d.toLocaleDateString('en-US', { weekday: 'long', month: 'long', day: 'numeric' })
})

const insightMessage = computed(() => {
  const completed = sessions.value.filter(s => s.status === 'completed').length
  if (sessions.value.length === 0) return 'Create at least 2 focused study sessions this week to build consistency.'
  if (completed >= 2) return `Great progress! You've completed ${completed} sessions this week. Keep the momentum going.`
  return 'You already have study sessions planned. Complete them to improve your consistency.'
})

onMounted(async () => {
  setWorkspaceToken(page.props.go_token)
  await fetchSessions()
})

async function fetchSessions() {
  isLoading.value = true
  try {
    sessions.value = await plannerApi.getSessions(selectedDate.value) || []
  } catch (err) {
    console.error('Failed to fetch sessions', err)
    sessions.value = []
  } finally {
    isLoading.value = false
  }
}

function selectDay(day) {
  selectedDate.value = day.dateStr
}

function openModal(session = null) {
  editingSession.value = session
  if (session) {
    form.value = { ...session }
  } else {
    form.value = {
      title: '', description: '', date: selectedDate.value,
      start_time: '09:00', end_time: '10:00',
      focus_dimension: 'General', target_id: urlTargetId || null
    }
  }
  showModal.value = true
}

async function saveSession() {
  isSaving.value = true
  try {
    if (editingSession.value) {
      const updated = await plannerApi.updateSession(editingSession.value.id, form.value)
      const idx = sessions.value.findIndex(s => s.id === updated.id)
      if (idx !== -1) sessions.value[idx] = updated
      showToast('Session updated!', 'success')
    } else {
      const created = await plannerApi.createSession(form.value)
      sessions.value.push(created)
      showToast('Session created!', 'success')
    }
    showModal.value = false
  } catch (err) {
    console.error('Save failed', err)
    showToast('Failed to save session.', 'error')
  } finally {
    isSaving.value = false
  }
}

async function completeSession(id) {
  try {
    const updated = await plannerApi.completeSession(id)
    const idx = sessions.value.findIndex(s => s.id === id)
    if (idx !== -1) sessions.value[idx] = updated
    showToast('Session completed!', 'success')
  } catch (err) {
    showToast('Failed to complete session.', 'error')
  }
}

async function skipSession(id) {
  try {
    const updated = await plannerApi.skipSession(id)
    const idx = sessions.value.findIndex(s => s.id === id)
    if (idx !== -1) sessions.value[idx] = updated
    showToast('Session skipped.', 'success')
  } catch (err) {
    showToast('Failed to skip session.', 'error')
  }
}

async function deleteSession(id) {
  if (!confirm('Delete this session?')) return
  try {
    await plannerApi.deleteSession(id)
    sessions.value = sessions.value.filter(s => s.id !== id)
    showToast('Session deleted.', 'success')
  } catch (err) {
    showToast('Failed to delete session.', 'error')
  }
}

function showToast(message, type = 'success') {
  toast.value = { message, type }
  setTimeout(() => toast.value = null, 3000)
}

function statusClass(status) {
  const map = {
    completed: 'bg-emerald-50 text-emerald-500',
    planned: 'bg-slate-50 text-slate-500',
    in_progress: 'bg-blue-50 text-blue-500',
    skipped: 'bg-amber-50 text-amber-500'
  }
  return map[status] || 'bg-slate-50 text-slate-500'
}

function formatDate(d) {
  return d.toISOString().split('T')[0]
}

function getMonday(d) {
  const day = d.getDay()
  const diff = d.getDate() - day + (day === 0 ? -6 : 1)
  return new Date(d.setDate(diff))
}
</script>
