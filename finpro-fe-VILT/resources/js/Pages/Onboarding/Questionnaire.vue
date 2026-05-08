<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-50 to-indigo-50 flex flex-col items-center justify-center px-4 py-10">
    <!-- BG blobs -->
    <div class="fixed top-0 right-0 w-96 h-96 bg-indigo-100 rounded-full blur-3xl opacity-40 pointer-events-none"></div>
    <div class="fixed bottom-0 left-0 w-80 h-80 bg-violet-100 rounded-full blur-3xl opacity-40 pointer-events-none"></div>

    <div class="relative z-10 w-full max-w-2xl">
      <!-- Top bar -->
      <div class="flex items-center justify-between mb-6">
        <Link :href="route('onboarding.sanctuary')" class="flex items-center gap-1.5 text-sm text-slate-500 hover:text-indigo-600 transition-colors">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/></svg>
          Kembali
        </Link>
        <span class="text-xs font-bold text-indigo-600">{{ answeredCount }} / 20 dijawab</span>
      </div>

      <!-- Overall progress bar -->
      <div class="bg-white rounded-2xl p-4 shadow-sm border border-slate-100 mb-6">
        <div class="flex items-center justify-between mb-2">
          <p class="text-xs font-semibold text-slate-500 uppercase tracking-wider">Progress Keseluruhan</p>
          <p class="text-xs font-bold text-indigo-600">{{ Math.round((answeredCount / 20) * 100) }}%</p>
        </div>
        <div class="h-2.5 bg-slate-100 rounded-full overflow-hidden">
          <div class="h-full bg-gradient-to-r from-indigo-500 to-violet-500 rounded-full transition-all duration-500 ease-out"
            :style="{ width: (answeredCount / 20 * 100) + '%' }"></div>
        </div>
        <!-- Segment dots -->
        <div class="flex justify-between mt-3">
          <div v-for="(seg, i) in segments" :key="i"
            class="flex items-center gap-1.5 text-xs"
            :class="currentSegment === i ? 'text-indigo-700 font-bold' : isSegmentDone(i) ? 'text-emerald-600 font-semibold' : 'text-slate-400'">
            <span class="w-5 h-5 rounded-full flex items-center justify-center text-[10px]"
              :class="isSegmentDone(i) ? 'bg-emerald-100' : currentSegment === i ? 'bg-indigo-100' : 'bg-slate-100'">
              {{ isSegmentDone(i) ? '✓' : seg.icon }}
            </span>
            <span class="hidden sm:inline">{{ seg.shortTitle }}</span>
          </div>
        </div>
      </div>

      <!-- Segment card -->
      <Transition name="slide" mode="out-in">
        <div :key="currentSegment" class="bg-white rounded-3xl shadow-xl border border-slate-100 overflow-hidden">
          <!-- Segment header -->
          <div class="px-8 py-6 border-b border-slate-50"
            :style="{ background: segments[currentSegment].gradient }">
            <div class="flex items-center gap-4">
              <div class="w-12 h-12 rounded-2xl flex items-center justify-center text-2xl shadow-sm"
                :class="segments[currentSegment].iconBg">
                {{ segments[currentSegment].icon }}
              </div>
              <div>
                <p class="text-xs font-bold uppercase tracking-widest"
                  :class="segments[currentSegment].labelColor">
                  Segmen {{ currentSegment + 1 }} dari 4
                </p>
                <h2 class="text-lg font-extrabold text-slate-900">{{ segments[currentSegment].title }}</h2>
                <p class="text-xs text-slate-500 mt-0.5">{{ segments[currentSegment].subtitle }}</p>
              </div>
            </div>
          </div>

          <!-- Questions -->
          <div class="px-6 py-6 space-y-6">
            <div v-for="(q, qi) in segments[currentSegment].questions" :key="q.id"
              class="rounded-2xl border-2 p-5 transition-all duration-200"
              :class="answers[q.id] ? 'border-indigo-200 bg-indigo-50/40' : 'border-slate-100 bg-white'">
              <p class="text-sm font-semibold text-slate-800 mb-4">
                <span class="inline-block w-6 h-6 bg-indigo-100 text-indigo-700 text-xs font-bold rounded-lg text-center leading-6 mr-2">
                  {{ q.number }}
                </span>
                {{ q.text }}
              </p>

              <!-- Likert scale -->
              <div class="grid grid-cols-5 gap-2">
                <button
                  v-for="val in 5"
                  :key="val"
                  :id="`q${q.number}-v${val}`"
                  @click="setAnswer(q.id, val)"
                  class="flex flex-col items-center gap-1 py-2.5 px-1 rounded-xl border-2 transition-all duration-150 hover:scale-105 active:scale-95"
                  :class="answers[q.id] === val
                    ? `border-${segments[currentSegment].activeColor} bg-${segments[currentSegment].activeBg} scale-105 shadow-md`
                    : 'border-slate-200 hover:border-slate-300 bg-white'">
                  <span class="font-black text-base"
                    :class="answers[q.id] === val ? segments[currentSegment].activeTextClass : 'text-slate-600'">
                    {{ val }}
                  </span>
                  <span class="text-[9px] text-center leading-tight"
                    :class="answers[q.id] === val ? segments[currentSegment].activeTextClass : 'text-slate-400'">
                    {{ likertLabels[val - 1] }}
                  </span>
                </button>
              </div>
            </div>
          </div>

          <!-- Navigation -->
          <div class="px-6 pb-6 flex justify-between items-center gap-3">
            <button
              v-if="currentSegment > 0"
              @click="currentSegment--"
              class="flex items-center gap-2 text-sm font-medium text-slate-500 hover:text-indigo-600 transition-colors px-4 py-2.5 rounded-xl border border-slate-200 hover:border-indigo-200">
              ← Segmen sebelumnya
            </button>
            <div v-else></div>

            <button
              @click="nextSegment"
              :disabled="!isSegmentComplete || submitting"
              id="questionnaire-next-btn"
              class="flex items-center gap-2 font-bold py-3 px-6 rounded-2xl transition-all shadow-lg"
              :class="isSegmentComplete
                ? 'bg-indigo-600 text-white hover:bg-indigo-700 shadow-indigo-200 hover:-translate-y-0.5'
                : 'bg-slate-100 text-slate-400 cursor-not-allowed shadow-none'">
              <span v-if="submitting" class="flex items-center gap-2">
                <svg class="animate-spin h-4 w-4" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"/>
                </svg>
                Menganalisis...
              </span>
              <span v-else-if="currentSegment < 3">
                Segmen Berikutnya →
              </span>
              <span v-else>
                Lihat Hasilku 🎯
              </span>
            </button>
          </div>

          <!-- Segment progress pills -->
          <div class="px-6 pb-5 flex justify-center gap-2">
            <div v-for="(seg, i) in segments" :key="i"
              class="h-1.5 rounded-full transition-all duration-300"
              :class="[
                i === currentSegment ? 'w-8 bg-indigo-500' :
                isSegmentDone(i) ? 'w-5 bg-emerald-400' :
                'w-4 bg-slate-200'
              ]"></div>
          </div>
        </div>
      </Transition>

      <!-- Reminder -->
      <p class="text-center text-xs text-slate-400 mt-5">
        Jawablah sejujurnya — tidak ada jawaban yang benar atau salah ✨
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { Link, router } from '@inertiajs/vue3'
import axios from 'axios'

// ─── Data ────────────────────────────────────────────────────────────────────

const props = defineProps({
  questions: { type: Array, default: () => [] },
})

const currentSegment = ref(0)
const answers        = ref({})
const submitting     = ref(false)

// Hardcoded 20 questions by segment (matches Go backend questions by category)
const segments = [
  {
    icon: '🟡',
    iconBg: 'bg-yellow-50',
    shortTitle: 'Planning',
    title: 'Planning Score',
    subtitle: 'Seberapa baik kamu merencanakan belajarmu?',
    gradient: 'linear-gradient(135deg, #fefce8 0%, #fef9c3 100%)',
    labelColor: 'text-yellow-700',
    activeColor: 'yellow-400',
    activeBg: 'yellow-50',
    activeTextClass: 'text-yellow-700',
    questions: [
      { id: '11111111-0001-0001-0001-000000000001', number: 1,  text: 'Saya membuat jadwal belajar sebelum mulai belajar' },
      { id: '11111111-0001-0001-0001-000000000002', number: 2,  text: 'Saya menetapkan target belajar yang jelas' },
      { id: '11111111-0001-0001-0001-000000000003', number: 3,  text: 'Saya merencanakan materi yang akan dipelajari' },
      { id: '11111111-0001-0001-0001-000000000004', number: 4,  text: 'Saya menentukan waktu belajar secara rutin' },
      { id: '11111111-0001-0001-0001-000000000005', number: 5,  text: 'Saya mempersiapkan kebutuhan belajar sebelum mulai' },
    ],
  },
  {
    icon: '🟢',
    iconBg: 'bg-emerald-50',
    shortTitle: 'Time Mgmt',
    title: 'Time Management Score',
    subtitle: 'Seberapa efektif kamu mengelola waktu belajarmu?',
    gradient: 'linear-gradient(135deg, #f0fdf4 0%, #dcfce7 100%)',
    labelColor: 'text-emerald-700',
    activeColor: 'emerald-500',
    activeBg: 'emerald-50',
    activeTextClass: 'text-emerald-700',
    questions: [
      { id: '11111111-0001-0001-0002-000000000006', number: 6,  text: 'Saya mengatur waktu belajar dengan baik' },
      { id: '11111111-0001-0001-0002-000000000007', number: 7,  text: 'Saya menyelesaikan tugas tepat waktu' },
      { id: '11111111-0001-0001-0002-000000000008', number: 8,  text: 'Saya jarang menunda pekerjaan' },
      { id: '11111111-0001-0001-0002-000000000009', number: 9,  text: 'Saya memprioritaskan tugas yang penting' },
      { id: '11111111-0001-0001-0002-000000000010', number: 10, text: 'Saya konsisten dengan jadwal belajar saya' },
    ],
  },
  {
    icon: '🔵',
    iconBg: 'bg-blue-50',
    shortTitle: 'Cognitive',
    title: 'Cognitive Score',
    subtitle: 'Seberapa baik strategi belajar (cara berpikirmu)?',
    gradient: 'linear-gradient(135deg, #eff6ff 0%, #dbeafe 100%)',
    labelColor: 'text-blue-700',
    activeColor: 'blue-500',
    activeBg: 'blue-50',
    activeTextClass: 'text-blue-700',
    questions: [
      { id: '11111111-0001-0001-0003-000000000011', number: 11, text: 'Saya menggunakan metode belajar tertentu (mencatat, merangkum, dll)' },
      { id: '11111111-0001-0001-0003-000000000012', number: 12, text: 'Saya mencoba berbagai cara belajar untuk menemukan yang efektif' },
      { id: '11111111-0001-0001-0003-000000000013', number: 13, text: 'Saya memahami materi, bukan hanya menghafal' },
      { id: '11111111-0001-0001-0003-000000000014', number: 14, text: 'Saya mengulang materi untuk memperkuat pemahaman' },
      { id: '11111111-0001-0001-0003-000000000015', number: 15, text: 'Saya menghubungkan materi dengan pengetahuan sebelumnya' },
    ],
  },
  {
    icon: '🔴',
    iconBg: 'bg-red-50',
    shortTitle: 'Reflection',
    title: 'Reflection Score',
    subtitle: 'Seberapa sering kamu mengevaluasi diri sendiri?',
    gradient: 'linear-gradient(135deg, #fff1f2 0%, #ffe4e6 100%)',
    labelColor: 'text-red-600',
    activeColor: 'red-400',
    activeBg: 'red-50',
    activeTextClass: 'text-red-600',
    questions: [
      { id: '11111111-0001-0001-0004-000000000016', number: 16, text: 'Saya mengecek apakah saya memahami materi' },
      { id: '11111111-0001-0001-0004-000000000017', number: 17, text: 'Saya menyadari ketika saya tidak memahami sesuatu' },
      { id: '11111111-0001-0001-0004-000000000018', number: 18, text: 'Saya mengevaluasi cara belajar saya' },
      { id: '11111111-0001-0001-0004-000000000019', number: 19, text: 'Saya memperbaiki strategi belajar jika kurang efektif' },
      { id: '11111111-0001-0001-0004-000000000020', number: 20, text: 'Saya belajar dari kesalahan sebelumnya' },
    ],
  },
]

const likertLabels = [
  'Sangat\nTidak Setuju',
  'Tidak\nSetuju',
  'Netral',
  'Setuju',
  'Sangat\nSetuju',
]

const STORAGE_KEY = 'lumora_survey_progress'

// ─── Computed ─────────────────────────────────────────────────────────────────

const answeredCount = computed(() =>
  segments.flatMap(s => s.questions).filter(q => answers.value[q.id]).length
)

const isSegmentComplete = computed(() => {
  return segments[currentSegment.value].questions.every(q => answers.value[q.id])
})

function isSegmentDone(i) {
  return segments[i].questions.every(q => answers.value[q.id]) && i < currentSegment.value
}

// ─── Persistence ──────────────────────────────────────────────────────────────

function saveProgress() {
  try {
    localStorage.setItem(STORAGE_KEY, JSON.stringify({
      answers:        answers.value,
      currentSegment: currentSegment.value,
    }))
  } catch { /* ignore */ }
}

function loadProgress() {
  try {
    const saved = localStorage.getItem(STORAGE_KEY)
    if (!saved) return
    const { answers: savedAnswers, currentSegment: savedSegment } = JSON.parse(saved)
    if (savedAnswers)  answers.value        = savedAnswers
    if (savedSegment !== undefined) currentSegment.value = savedSegment
  } catch { /* ignore */ }
}

function clearProgress() {
  try { localStorage.removeItem(STORAGE_KEY) } catch { /* ignore */ }
}

// Load saved progress on mount
onMounted(loadProgress)

// Watch for changes and auto-save
watch([answers, currentSegment], saveProgress, { deep: true })

// ─── Methods ──────────────────────────────────────────────────────────────────

function setAnswer(questionId, value) {
  answers.value[questionId] = value
}

async function nextSegment() {
  if (!isSegmentComplete.value || submitting.value) return

  if (currentSegment.value < 3) {
    currentSegment.value++
    window.scrollTo({ top: 0, behavior: 'smooth' })
    return
  }

  // Last segment → submit
  submitting.value = true
  try {
    // q.id is now the real UUID from the seeded questions table
    const allQuestions = segments.flatMap(s => s.questions)
    const payload = allQuestions.map(q => ({
      question_id:  q.id,
      answer_value: answers.value[q.id],
    }))

    const response = await axios.post('/onboarding/submit', {
      answers: payload,
    })

    // Store result in sessionStorage for Result page
    sessionStorage.setItem('lumora_result', JSON.stringify(response.data?.data ?? {}))
    clearProgress() // ← wipe saved progress so retaking starts fresh
    router.visit(route('onboarding.result'))
  } catch (err) {
    console.error('Submit error:', err)
    // Even on error, navigate to result (will show fallback)
    router.visit(route('onboarding.result'))
  } finally {
    submitting.value = false
  }
}
</script>

<style scoped>
.slide-enter-active,
.slide-leave-active {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}
.slide-enter-from {
  opacity: 0;
  transform: translateX(30px);
}
.slide-leave-to {
  opacity: 0;
  transform: translateX(-30px);
}
</style>
