<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-50 to-indigo-50 flex items-center justify-center px-4 py-14">
    <div class="fixed top-0 right-0 w-96 h-96 bg-indigo-100 rounded-full blur-3xl opacity-40 pointer-events-none"></div>
    <div class="fixed bottom-0 left-0 w-80 h-80 bg-violet-100 rounded-full blur-3xl opacity-40 pointer-events-none"></div>

    <div class="relative z-10 w-full max-w-3xl">
      <!-- Logo -->
      <div class="text-center mb-6">
        <span class="text-xl font-extrabold text-indigo-600">Lumora</span>
      </div>

      <div class="bg-white rounded-3xl shadow-2xl shadow-indigo-100 border border-slate-100 overflow-hidden">

        <!-- Header -->
        <div class="bg-gradient-to-r from-indigo-600 via-violet-600 to-purple-600 px-8 py-10 text-center relative overflow-hidden">
          <!-- Decorative circles -->
          <div class="absolute -top-8 -right-8 w-32 h-32 bg-white/10 rounded-full"></div>
          <div class="absolute -bottom-6 -left-6 w-24 h-24 bg-white/10 rounded-full"></div>

          <div class="relative">
            <p class="text-indigo-200 text-xs font-bold uppercase tracking-widest mb-2">✨ Profil Belajarmu</p>
            <h1 class="text-3xl font-extrabold text-white leading-tight">
              The <span class="text-yellow-300">{{ profileTitle }}</span>
            </h1>
            <p class="text-indigo-200 mt-3 text-sm max-w-lg mx-auto leading-relaxed">
              {{ profileDesc }}
            </p>

            <!-- Total score badge -->
            <div class="inline-flex items-center gap-2 bg-white/20 backdrop-blur border border-white/30 text-white rounded-2xl px-5 py-2.5 mt-5 text-sm font-semibold">
              <span class="text-yellow-300 text-lg">⭐</span>
              Total Skor: <span class="font-black text-white text-base">{{ totalScore }} / 100</span>
            </div>
          </div>
        </div>

        <div class="p-8 grid md:grid-cols-2 gap-8">

          <!-- Dimension Scores -->
          <div>
            <h2 class="font-extrabold text-slate-900 mb-5 flex items-center gap-2">
              <span class="text-lg">📊</span> Skor per Dimensi
            </h2>
            <div class="space-y-5">
              <div v-for="dim in dimensions" :key="dim.label">
                <div class="flex justify-between items-center mb-1.5">
                  <div class="flex items-center gap-2">
                    <span>{{ dim.icon }}</span>
                    <span class="text-sm font-semibold text-slate-700">{{ dim.label }}</span>
                  </div>
                  <div class="flex items-center gap-2">
                    <span class="text-sm font-black text-slate-800">{{ dim.score }}/25</span>
                    <span class="text-xs font-bold px-2 py-0.5 rounded-full"
                      :class="getCategoryClass(dim.category)">
                      {{ dim.category }}
                    </span>
                  </div>
                </div>
                <div class="h-3 bg-slate-100 rounded-full overflow-hidden">
                  <div class="h-full rounded-full transition-all duration-1000 ease-out"
                    :class="dim.barColor"
                    :style="{ width: (dim.score / 25 * 100) + '%' }">
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Core Strengths -->
          <div>
            <h2 class="font-extrabold text-slate-900 mb-5 flex items-center gap-2">
              <span class="text-lg">💪</span> Kekuatanmu
            </h2>
            <ul class="space-y-3">
              <li v-for="s in strengths" :key="s"
                class="flex items-start gap-3 text-sm text-slate-700">
                <span class="w-5 h-5 rounded-full bg-emerald-100 text-emerald-600 flex items-center justify-center text-xs flex-shrink-0 mt-0.5 font-bold">✓</span>
                {{ s }}
              </li>
            </ul>

            <h2 class="font-extrabold text-slate-900 mt-7 mb-4 flex items-center gap-2">
              <span class="text-lg">🎯</span> Area Pengembangan
            </h2>
            <ul class="space-y-3">
              <li v-for="w in weaknesses" :key="w"
                class="flex items-start gap-3 text-sm text-slate-700">
                <span class="w-5 h-5 rounded-full bg-amber-100 text-amber-600 flex items-center justify-center text-xs flex-shrink-0 mt-0.5 font-bold">!</span>
                {{ w }}
              </li>
            </ul>
          </div>

          <!-- AI Recommendation -->
          <div class="md:col-span-2 bg-gradient-to-br from-indigo-50 to-violet-50 rounded-2xl p-6 border border-indigo-100">
            <div class="flex items-center gap-3 mb-4">
              <div class="w-10 h-10 bg-indigo-600 rounded-xl flex items-center justify-center text-xl">🧠</div>
              <div>
                <p class="text-xs font-bold text-indigo-500 uppercase tracking-wider">AI-Powered</p>
                <h2 class="font-extrabold text-slate-900">Rekomendasi Strategi Belajar</h2>
              </div>
            </div>

            <!-- Loading state -->
            <div v-if="loadingAI" class="flex items-center gap-3 py-4">
              <svg class="animate-spin h-5 w-5 text-indigo-600" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"/>
              </svg>
              <p class="text-sm text-slate-500 italic">Gemini AI sedang menganalisis profilmu...</p>
            </div>

            <!-- AI content -->
            <div v-else class="space-y-4">
              <div v-for="(rec, i) in aiRecommendations" :key="i"
                class="flex items-start gap-3">
                <span class="w-6 h-6 bg-indigo-600 text-white text-xs font-black rounded-lg flex items-center justify-center flex-shrink-0">
                  {{ i + 1 }}
                </span>
                <p class="text-sm text-slate-700 leading-relaxed">{{ rec }}</p>
              </div>
            </div>

            <!-- Tags -->
            <div class="flex flex-wrap gap-2 mt-5">
              <span v-for="tag in strategyTags" :key="tag"
                class="bg-indigo-100 text-indigo-700 text-xs font-semibold px-3 py-1.5 rounded-full">
                {{ tag }}
              </span>
            </div>
          </div>
        </div>

        <!-- Footer actions -->
        <div class="px-8 pb-8 flex flex-col sm:flex-row gap-3 justify-center">
          <Link :href="route('dashboard')"
            id="result-dashboard-btn"
            class="flex-1 sm:flex-none text-center bg-indigo-600 text-white font-bold px-10 py-3.5 rounded-2xl hover:bg-indigo-700 active:scale-95 transition-all shadow-lg shadow-indigo-200">
            Buka Dashboard →
          </Link>
          <button
            @click="printResult"
            class="flex-1 sm:flex-none text-center border-2 border-slate-200 text-slate-700 font-semibold px-8 py-3.5 rounded-2xl hover:border-indigo-300 hover:text-indigo-600 transition-all">
            📄 Unduh Laporan
          </button>
        </div>

        <!-- Pagination dots -->
        <div class="pb-6 flex justify-center gap-2">
          <div class="w-3 h-1.5 bg-emerald-400 rounded-full"></div>
          <div class="w-3 h-1.5 bg-emerald-400 rounded-full"></div>
          <div class="w-3 h-1.5 bg-emerald-400 rounded-full"></div>
          <div class="w-8 h-1.5 bg-indigo-600 rounded-full"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { Link } from '@inertiajs/vue3'

// ─── Props from controller ─────────────────────────────────────────────────
const props = defineProps({
  result: { type: Object, default: () => ({}) },
})

// ─── State ────────────────────────────────────────────────────────────────
const loadingAI          = ref(true)
const aiRecommendations  = ref([])
const strategyTags       = ref([])

// ─── Scores: use props from server or fall back to sessionStorage ──────────
const resultData = computed(() => {
  if (props.result && props.result.planning_score !== undefined) return props.result
  try {
    const stored = sessionStorage.getItem('lumora_result')
    return stored ? JSON.parse(stored) : {}
  } catch { return {} }
})

const planningScore       = computed(() => resultData.value.planning_score       ?? 0)
const timeManagementScore = computed(() => resultData.value.time_management_score ?? 0)
const cognitiveScore      = computed(() => resultData.value.cognitive_score       ?? 0)
const reflectionScore     = computed(() => resultData.value.reflection_score      ?? 0)
const totalScore          = computed(() => resultData.value.total_score           ?? (planningScore.value + timeManagementScore.value + cognitiveScore.value + reflectionScore.value))

// ─── Dimensions ───────────────────────────────────────────────────────────
const dimensions = computed(() => [
  {
    icon: '🟡', label: 'Planning',
    score: planningScore.value,
    category: getCategory(planningScore.value),
    barColor: 'bg-yellow-400',
  },
  {
    icon: '🟢', label: 'Time Management',
    score: timeManagementScore.value,
    category: getCategory(timeManagementScore.value),
    barColor: 'bg-emerald-500',
  },
  {
    icon: '🔵', label: 'Cognitive',
    score: cognitiveScore.value,
    category: getCategory(cognitiveScore.value),
    barColor: 'bg-blue-500',
  },
  {
    icon: '🔴', label: 'Reflection',
    score: reflectionScore.value,
    category: getCategory(reflectionScore.value),
    barColor: 'bg-red-400',
  },
])

// ─── Category logic ───────────────────────────────────────────────────────
function getCategory(score) {
  if (score <= 12) return 'Low'
  if (score <= 18) return 'Medium'
  return 'High'
}

function getCategoryClass(cat) {
  return cat === 'High'   ? 'bg-emerald-100 text-emerald-700' :
         cat === 'Medium' ? 'bg-amber-100 text-amber-700' :
                            'bg-red-100 text-red-600'
}

// ─── Profile naming ───────────────────────────────────────────────────────
const profileTitle = computed(() => {
  const avg = totalScore.value / 4
  if (avg >= 19) return 'Focused Achiever'
  if (avg >= 13) return 'Growing Learner'
  return 'Rising Explorer'
})

const profileDesc = computed(() => {
  const avg = totalScore.value / 4
  if (avg >= 19) return 'Kamu memiliki kemampuan belajar yang sangat terstruktur dan termotivasi tinggi. Kamu unggul dalam merencanakan, mengelola waktu, dan merefleksi diri.'
  if (avg >= 13) return 'Kamu sudah memiliki fondasi yang baik dan terus berkembang. Dengan sedikit penyesuaian pada beberapa area, kamu bisa mencapai hasil yang luar biasa.'
  return 'Kamu baru memulai perjalanan belajar yang lebih terstruktur. Lumora akan membantumu membangun kebiasaan yang kuat langkah demi langkah.'
})

// ─── Dynamic strengths/weaknesses based on scores ─────────────────────────
const strengths = computed(() => {
  const s = []
  if (planningScore.value >= 19)       s.push('Perencanaan belajar yang sangat terstruktur')
  if (timeManagementScore.value >= 19) s.push('Manajemen waktu yang efisien dan konsisten')
  if (cognitiveScore.value >= 19)      s.push('Strategi kognitif yang kuat dan beragam')
  if (reflectionScore.value >= 19)     s.push('Kemampuan refleksi dan evaluasi diri yang tinggi')
  if (s.length === 0) s.push('Semangat untuk terus belajar dan berkembang', 'Terbuka terhadap feedback dan perbaikan diri')
  return s
})

const weaknesses = computed(() => {
  const w = []
  if (planningScore.value <= 12)       w.push('Perlu lebih konsisten dalam merencanakan sesi belajar')
  if (timeManagementScore.value <= 12) w.push('Manajemen waktu masih perlu ditingkatkan')
  if (cognitiveScore.value <= 12)      w.push('Eksplorasi berbagai metode belajar yang berbeda')
  if (reflectionScore.value <= 12)     w.push('Jadikan refleksi harian sebagai kebiasaan rutin')
  if (w.length === 0) w.push('Pertahankan konsistensi yang sudah kamu bangun')
  return w
})

// ─── Parse AI recommendation from server ──────────────────────────────────
onMounted(() => {
  const raw = resultData.value.category_result
  if (raw) {
    try {
      // Parse JSON from Go AI response
      let parsed = typeof raw === 'string' ? JSON.parse(raw) : raw

      // 1. Recommendations (from Gemini JSON structure)
      const recs = []
      if (Array.isArray(parsed.recommendations)) {
        recs.push(...parsed.recommendations)
      } else if (typeof parsed.recommendations === 'string') {
        recs.push(...parsed.recommendations.split('\n').filter(Boolean))
      }

      aiRecommendations.value = recs.length > 0 ? recs.slice(0, 5) : getDefaultRecommendations()

      // 2. Strategy Tags
      if (Array.isArray(parsed.strategy_tags)) {
        strategyTags.value = parsed.strategy_tags
      } else if (parsed.style) {
        strategyTags.value.push(parsed.style)
      }

      // 3. Optional: Override profile title if AI provides a better one
      if (parsed.profile_title) {
        // You can use this to make the UI even more dynamic
      }

    } catch (e) {
      console.warn("Failed to parse AI JSON:", e)
      aiRecommendations.value = getDefaultRecommendations()
    }
  } else {
    aiRecommendations.value = getDefaultRecommendations()
  }

  // Add strategy tags based on scores
  if (strategyTags.value.length === 0) {
    if (planningScore.value <= 14)       strategyTags.value.push('Time Blocking')
    if (cognitiveScore.value >= 17)      strategyTags.value.push('Active Recall')
    if (reflectionScore.value >= 17)     strategyTags.value.push('Reflective Journaling')
    if (timeManagementScore.value >= 17) strategyTags.value.push('Pomodoro Technique')
    strategyTags.value.push('Spaced Repetition')
  }

  loadingAI.value = false
})

function getDefaultRecommendations() {
  return [
    'Buat jadwal belajar mingguan yang realistis dan patuhi dengan konsisten.',
    'Gunakan teknik Pomodoro: belajar 25 menit, istirahat 5 menit.',
    'Tulis catatan refleksi singkat setiap selesai belajar — apa yang dipahami dan apa yang belum.',
    'Pecah materi besar menjadi topik-topik kecil yang lebih mudah dicerna.',
    'Review ulang materi lama secara berkala menggunakan prinsip spaced repetition.',
  ]
}

function printResult() {
  window.print()
}
</script>
