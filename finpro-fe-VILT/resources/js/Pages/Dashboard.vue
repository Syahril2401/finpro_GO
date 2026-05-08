<template>
  <div class="min-h-screen bg-slate-50 font-sans">
    <!-- Sidebar -->
    <aside class="fixed inset-y-0 left-0 w-64 bg-white border-r border-slate-100 flex flex-col z-40">
      <div class="px-6 py-5 border-b border-slate-100 flex items-center gap-3">
        <span class="text-xl font-extrabold text-slate-900">Lumora</span>
        <span class="text-xs bg-indigo-100 text-indigo-700 font-semibold px-2 py-0.5 rounded-full">Beta</span>
      </div>

      <nav class="flex-1 px-4 py-6 space-y-1">
        <a v-for="item in navItems" :key="item.label"
          :href="item.href"
          class="flex items-center gap-3 px-3 py-2.5 rounded-xl text-sm font-medium transition-all"
          :class="item.active
            ? 'bg-indigo-600 text-white shadow-md shadow-indigo-200'
            : 'text-slate-600 hover:bg-slate-50 hover:text-slate-900'">
          <span class="text-lg">{{ item.icon }}</span>
          {{ item.label }}
        </a>
      </nav>

      <!-- User -->
      <div class="px-4 py-4 border-t border-slate-100">
        <div class="flex items-center gap-3 px-3 py-2">
          <div class="w-8 h-8 rounded-full bg-gradient-to-br from-indigo-500 to-violet-500 flex items-center justify-center text-white text-xs font-bold">
            S
          </div>
          <div class="flex-1 min-w-0">
            <p class="text-sm font-semibold text-slate-800 truncate">Student</p>
            <p class="text-xs text-slate-400 truncate">{{ profileTitle }}</p>
          </div>
        </div>
        <form method="POST" :action="route('logout')" class="mt-2">
          <input type="hidden" name="_token" :value="csrfToken">
          <button type="submit" class="w-full text-left px-3 py-2 text-xs text-slate-500 hover:text-red-500 transition-colors rounded-lg hover:bg-red-50">
            Sign out
          </button>
        </form>
      </div>
    </aside>

    <!-- Main -->
    <main class="ml-64 p-8">
      <!-- Header -->
      <div class="mb-8 flex items-start justify-between">
        <div>
          <h1 class="text-2xl font-extrabold text-slate-900">Good afternoon 👋</h1>
          <p class="text-slate-500 text-sm mt-1">Here's your study overview for today.</p>
        </div>
        <!-- Chatbot toggle -->
        <button @click="chatOpen = true"
          id="open-chatbot-btn"
          class="flex items-center gap-2 bg-indigo-600 text-white text-sm font-semibold px-4 py-2.5 rounded-xl shadow-lg shadow-indigo-200 hover:bg-indigo-700 transition-all hover:-translate-y-0.5">
          🧠 Tanya Lumora AI
        </button>
      </div>

      <!-- Stats row -->
      <div class="grid grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
        <div v-for="stat in stats" :key="stat.label"
          class="bg-white rounded-2xl p-5 border border-slate-100 hover:border-indigo-200 hover:shadow-md transition-all">
          <p class="text-2xl font-black" :class="stat.color">{{ stat.value }}</p>
          <p class="text-slate-500 text-xs mt-1">{{ stat.label }}</p>
        </div>
      </div>

      <!-- SRL Score row (from result) -->
      <div v-if="hasSrlData" class="bg-white rounded-2xl border border-slate-100 p-6 mb-6">
        <div class="flex items-center justify-between mb-4">
          <h2 class="font-bold text-slate-900 flex items-center gap-2">
            <span>📊</span> Profil SRL Kamu
          </h2>
          <Link :href="route('onboarding.result')" class="text-xs text-indigo-500 hover:text-indigo-700 font-medium">
            Lihat detail →
          </Link>
        </div>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <div v-for="dim in srlDimensions" :key="dim.label">
            <div class="flex justify-between text-xs text-slate-500 mb-1">
              <span class="flex items-center gap-1">{{ dim.icon }} {{ dim.label }}</span>
              <span class="font-bold text-slate-700">{{ dim.score }}/25</span>
            </div>
            <div class="h-2 bg-slate-100 rounded-full overflow-hidden">
              <div class="h-full rounded-full transition-all duration-700" :class="dim.bar"
                :style="{ width: (dim.score / 25 * 100) + '%' }"></div>
            </div>
            <p class="text-[10px] mt-1 font-semibold" :class="getCategoryClass(dim.score)">
              {{ getCategory(dim.score) }}
            </p>
          </div>
        </div>
      </div>

      <!-- Middle row: Planner + AI Strategy -->
      <div class="grid lg:grid-cols-2 gap-6 mb-6">
        <!-- Planner -->
        <div class="bg-white rounded-2xl border border-slate-100 p-6">
          <h2 class="font-bold text-slate-900 mb-4">📅 Today's Planner</h2>
          <div class="space-y-3">
            <div v-for="item in plannerItems" :key="item.title"
              class="flex items-center gap-3 p-3 rounded-xl"
              :class="item.done ? 'bg-slate-50' : 'bg-indigo-50'">
              <div class="w-2 h-2 rounded-full flex-shrink-0"
                :class="item.done ? 'bg-slate-300' : 'bg-indigo-500'"></div>
              <div class="flex-1 min-w-0">
                <p class="text-sm font-medium text-slate-800 truncate"
                  :class="{ 'line-through text-slate-400': item.done }">{{ item.title }}</p>
                <p class="text-xs text-slate-400">{{ item.time }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- AI Strategy -->
        <div class="bg-gradient-to-br from-indigo-600 to-violet-600 rounded-2xl p-6 text-white">
          <div class="flex items-center gap-3 mb-4">
            <div class="w-10 h-10 bg-white/20 rounded-xl flex items-center justify-center">🧠</div>
            <div>
              <p class="text-indigo-200 text-xs font-semibold">AI-POWERED</p>
              <h2 class="font-bold">Today's Strategy</h2>
            </div>
          </div>
          <p class="font-bold text-lg mb-1">Pomodoro Technique</p>
          <p class="text-indigo-200 text-sm leading-relaxed">
            Work in 25-minute focused sessions. Take a 5-minute break, then repeat.
            After 4 sessions, take a longer 15-30 minute break.
          </p>
          <div class="mt-4 flex gap-2 flex-wrap">
            <span class="bg-white/20 text-white text-xs px-3 py-1 rounded-full">25 min focus</span>
            <span class="bg-white/20 text-white text-xs px-3 py-1 rounded-full">5 min break</span>
          </div>
        </div>
      </div>

      <!-- Weekly Progress -->
      <div class="bg-white rounded-2xl border border-slate-100 p-6">
        <h2 class="font-bold text-slate-900 mb-4">📈 Weekly Progress</h2>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <div v-for="p in progressItems" :key="p.label" class="text-center">
            <div class="relative w-16 h-16 mx-auto mb-2">
              <svg class="w-16 h-16 -rotate-90" viewBox="0 0 64 64">
                <circle cx="32" cy="32" r="26" fill="none" stroke="#f1f5f9" stroke-width="6"/>
                <circle cx="32" cy="32" r="26" fill="none"
                  stroke="#6366f1" stroke-width="6"
                  :stroke-dasharray="`${p.pct * 163.4 / 100} 163.4`"
                  stroke-linecap="round"/>
              </svg>
              <span class="absolute inset-0 flex items-center justify-center text-xs font-bold text-slate-800">{{ p.pct }}%</span>
            </div>
            <p class="text-xs text-slate-500 font-medium">{{ p.label }}</p>
          </div>
        </div>
      </div>
    </main>

    <!-- ─── AI Chatbot Panel ──────────────────────────────────────────────── -->
    <Transition name="chat-panel">
      <div v-if="chatOpen"
        class="fixed bottom-6 right-6 w-96 bg-white rounded-3xl shadow-2xl border border-slate-100 flex flex-col z-50"
        style="height: 520px;">
        <!-- Header -->
        <div class="flex items-center gap-3 p-4 border-b border-slate-100">
          <div class="w-9 h-9 bg-gradient-to-br from-indigo-500 to-violet-500 rounded-xl flex items-center justify-center text-lg">🧠</div>
          <div class="flex-1">
            <p class="font-bold text-slate-900 text-sm">Lumora AI</p>
            <p class="text-xs text-emerald-500 font-medium">● Online</p>
          </div>
          <button @click="chatOpen = false" class="text-slate-400 hover:text-slate-600 transition-colors">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
            </svg>
          </button>
        </div>

        <!-- Messages -->
        <div ref="chatContainer" class="flex-1 overflow-y-auto p-4 space-y-3">
          <div v-for="(msg, i) in chatMessages" :key="i"
            class="flex gap-2"
            :class="msg.role === 'user' ? 'justify-end' : 'justify-start'">
            <!-- AI avatar -->
            <div v-if="msg.role === 'ai'"
              class="w-7 h-7 bg-indigo-100 rounded-full flex items-center justify-center text-sm flex-shrink-0 mt-1">
              🧠
            </div>
            <div class="max-w-[75%] px-3.5 py-2.5 rounded-2xl text-sm leading-relaxed"
              :class="msg.role === 'user'
                ? 'bg-indigo-600 text-white rounded-br-sm'
                : 'bg-slate-100 text-slate-800 rounded-bl-sm'">
              {{ msg.content }}
            </div>
          </div>

          <!-- Typing indicator -->
          <div v-if="chatLoading" class="flex gap-2 justify-start">
            <div class="w-7 h-7 bg-indigo-100 rounded-full flex items-center justify-center text-sm flex-shrink-0">🧠</div>
            <div class="bg-slate-100 px-4 py-3 rounded-2xl rounded-bl-sm">
              <div class="flex gap-1">
                <span class="w-1.5 h-1.5 bg-slate-400 rounded-full animate-bounce" style="animation-delay:0ms"></span>
                <span class="w-1.5 h-1.5 bg-slate-400 rounded-full animate-bounce" style="animation-delay:150ms"></span>
                <span class="w-1.5 h-1.5 bg-slate-400 rounded-full animate-bounce" style="animation-delay:300ms"></span>
              </div>
            </div>
          </div>
        </div>

        <!-- Input -->
        <div class="p-3 border-t border-slate-100">
          <div class="flex gap-2 bg-slate-50 rounded-2xl px-4 py-2.5 border border-slate-200 focus-within:border-indigo-400 focus-within:ring-2 focus-within:ring-indigo-100 transition-all">
            <input
              v-model="chatInput"
              @keydown.enter.prevent="sendChat"
              id="chat-input"
              type="text"
              placeholder="Tanya sesuatu tentang belajarmu..."
              class="flex-1 bg-transparent text-sm text-slate-800 outline-none placeholder-slate-400"
              :disabled="chatLoading">
            <button @click="sendChat"
              id="chat-send-btn"
              :disabled="!chatInput.trim() || chatLoading"
              class="text-indigo-600 hover:text-indigo-800 disabled:opacity-30 transition-all">
              <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                <path d="M2.01 21L23 12 2.01 3 2 10l15 2-15 2z"/>
              </svg>
            </button>
          </div>
        </div>
      </div>
    </Transition>

    <!-- Floating chatbot button (when panel is closed) -->
    <Transition name="fade">
      <button v-if="!chatOpen"
        @click="chatOpen = true"
        id="floating-chat-btn"
        class="fixed bottom-6 right-6 w-14 h-14 bg-indigo-600 text-white rounded-2xl shadow-xl shadow-indigo-300 hover:bg-indigo-700 hover:scale-110 active:scale-95 transition-all flex items-center justify-center text-2xl z-50">
        🧠
      </button>
    </Transition>
  </div>
</template>

<script setup>
import { ref, computed, nextTick } from 'vue'
import { Link } from '@inertiajs/vue3'
import axios from 'axios'

// ─── Props ────────────────────────────────────────────────────────────────────
const props = defineProps({
  result: { type: Object, default: () => ({}) },
})

const csrfToken = document.querySelector('meta[name="csrf-token"]')?.content

// ─── SRL scores from last result ──────────────────────────────────────────────
const lastResult = computed(() => {
  if (props.result && Object.keys(props.result).length > 0) return props.result
  try { return JSON.parse(sessionStorage.getItem('lumora_result') ?? '{}') } catch { return {} }
})

const hasSrlData = computed(() =>
  (lastResult.value.planning_score ?? 0) > 0 || (lastResult.value.total_score ?? 0) > 0
)

const srlDimensions = computed(() => [
  { icon: '🟡', label: 'Planning',      score: lastResult.value.planning_score        ?? 0, bar: 'bg-yellow-400' },
  { icon: '🟢', label: 'Time Mgmt',     score: lastResult.value.time_management_score ?? 0, bar: 'bg-emerald-500' },
  { icon: '🔵', label: 'Cognitive',     score: lastResult.value.cognitive_score       ?? 0, bar: 'bg-blue-500' },
  { icon: '🔴', label: 'Reflection',    score: lastResult.value.reflection_score      ?? 0, bar: 'bg-red-400' },
])

const profileTitle = computed(() => {
  const t = lastResult.value.total_score ?? 0
  if (t >= 75) return 'Focused Achiever'
  if (t >= 50) return 'Growing Learner'
  return 'Rising Explorer'
})

function getCategory(score) {
  if (score <= 12) return 'Low'
  if (score <= 18) return 'Medium'
  return 'High'
}
function getCategoryClass(score) {
  if (score <= 12) return 'text-red-500'
  if (score <= 18) return 'text-amber-500'
  return 'text-emerald-600'
}

// ─── Nav & Static data ────────────────────────────────────────────────────────
const navItems = [
  { icon: '📊', label: 'Dashboard', href: '#', active: true },
  { icon: '📅', label: 'Planner',   href: '#', active: false },
  { icon: '📝', label: 'Notes',     href: '#', active: false },
  { icon: '🎯', label: 'Progress',  href: '#', active: false },
  { icon: '⚙️', label: 'Settings',  href: '#', active: false },
]

const stats = [
  { label: 'Study streak',   value: '14 🔥', color: 'text-orange-500' },
  { label: 'Sessions today', value: '3/4',   color: 'text-indigo-600' },
  { label: 'Focus score',    value: '87%',   color: 'text-emerald-600' },
  { label: 'Weekly target',  value: '72%',   color: 'text-violet-600' },
]

const plannerItems = [
  { title: 'Read Chapter 5 — Algorithms',    time: '08:00–09:00', done: true },
  { title: 'Practice problem set 3',         time: '10:00–11:30', done: true },
  { title: 'Study group: Data Structures',   time: '14:00–15:30', done: false },
  { title: 'Review flashcards',              time: '16:00–16:30', done: false },
]

const progressItems = [
  { label: 'Self-Monitoring', pct: 87 },
  { label: 'Goal Setting',    pct: 92 },
  { label: 'Strategy Use',    pct: 74 },
  { label: 'Self-Efficacy',   pct: 81 },
]

// ─── Chatbot ──────────────────────────────────────────────────────────────────
const chatOpen      = ref(false)
const chatInput     = ref('')
const chatLoading   = ref(false)
const chatContainer = ref(null)
const chatMessages  = ref([
  {
    role: 'ai',
    content: 'Halo! Saya Lumora AI 👋 Saya siap membantu kamu belajar lebih efektif. Ada yang bisa saya bantu?',
  },
])

async function sendChat() {
  const msg = chatInput.value.trim()
  if (!msg || chatLoading.value) return

  chatInput.value = ''
  chatMessages.value.push({ role: 'user', content: msg })
  chatLoading.value = true

  await nextTick()
  scrollChat()

  try {
    const resp = await axios.post('/chat', { message: msg })
    chatMessages.value.push({
      role: 'ai',
      content: resp.data?.reply ?? 'Maaf, saya tidak bisa menjawab saat ini.',
    })
  } catch {
    chatMessages.value.push({
      role: 'ai',
      content: 'Maaf, terjadi kendala teknis. Coba lagi sebentar ya! 🙏',
    })
  } finally {
    chatLoading.value = false
    await nextTick()
    scrollChat()
  }
}

function scrollChat() {
  if (chatContainer.value) {
    chatContainer.value.scrollTop = chatContainer.value.scrollHeight
  }
}
</script>

<style scoped>
.chat-panel-enter-active,
.chat-panel-leave-active { transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1); }
.chat-panel-enter-from,
.chat-panel-leave-to { opacity: 0; transform: translateY(16px) scale(0.97); }

.fade-enter-active,
.fade-leave-active { transition: all 0.2s ease; }
.fade-enter-from,
.fade-leave-to { opacity: 0; transform: scale(0.8); }
</style>
