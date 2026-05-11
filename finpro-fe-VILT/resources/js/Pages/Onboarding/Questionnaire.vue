<template>
  <div class="min-h-screen bg-[#F8F9FE] font-sans text-slate-900 pb-20">
    <!-- Navbar -->
    <nav class="fixed top-0 inset-x-0 z-50 bg-white/70 backdrop-blur-xl border-b border-slate-100 h-16 flex items-center justify-between px-8">
      <span class="text-xl font-black text-[#3D3ACE] tracking-tight">Lumora</span>
      <button class="text-sm font-bold text-slate-500 hover:text-slate-900 transition-colors">Save & Exit</button>
    </nav>

    <div class="max-w-[1200px] mx-auto pt-32 px-6 flex flex-col lg:flex-row gap-12">
      <!-- Main Content -->
      <div class="flex-1">
        <!-- Progress Stepper -->
        <div class="mb-16">
            <p class="text-[11px] font-black text-slate-400 uppercase tracking-widest mb-2">Progress</p>
            <h2 class="text-lg font-bold text-[#3D3ACE] mb-8">Step 0{{ currentSegmentIndex + 1 }}: {{ currentSegment.title }}</h2>
            
            <div class="relative">
                <div class="absolute top-1/2 left-0 w-full h-[2px] bg-slate-200 -translate-y-1/2"></div>
                <div class="absolute top-1/2 left-0 h-[2px] bg-[#3D3ACE] -translate-y-1/2 transition-all duration-500" 
                    :style="{ width: (currentSegmentIndex / (segments.length - 1) * 100) + '%' }"></div>
                
                <div class="relative flex justify-between">
                    <div v-for="(seg, i) in segments" :key="i" class="flex flex-col items-center">
                        <div class="w-4 h-4 rounded-full border-2 transition-all duration-500 z-10"
                            :class="i <= currentSegmentIndex ? 'bg-[#3D3ACE] border-[#3D3ACE]' : 'bg-white border-slate-300'">
                        </div>
                        <span class="absolute mt-6 text-[10px] font-bold uppercase tracking-widest text-slate-400 whitespace-nowrap"
                            :class="{ 'text-[#3D3ACE]': i === currentSegmentIndex }">
                            {{ seg.shortTitle }}
                        </span>
                    </div>
                </div>
            </div>
        </div>

        <!-- Question Title -->
        <Transition name="fade" mode="out-in">
          <div :key="currentQuestion.id" class="text-center mb-12 min-h-[140px]">
              <h1 class="text-[28px] md:text-[32px] font-extrabold text-[#1E1B4B] leading-tight mb-4 px-4">
                  {{ currentQuestion.text }}
              </h1>
              <p class="text-slate-500 font-medium max-w-xl mx-auto">
                  Understanding your rhythm helps us tailor the sanctuary to your natural productivity flow.
              </p>
          </div>
        </Transition>

        <!-- Options Grid -->
        <div class="space-y-4 max-w-xl mx-auto">
            <button v-for="(opt, idx) in likertOptions" :key="idx"
                @click="setAnswer(currentQuestion.id, opt.value)"
                class="w-full flex items-center gap-6 p-5 rounded-[24px] bg-white border-2 transition-all duration-300 group"
                :class="answers[currentQuestion.id] === opt.value ? 'border-[#3D3ACE] shadow-xl shadow-indigo-50' : 'border-transparent hover:border-slate-200'">
                
                <div class="w-12 h-12 rounded-2xl flex items-center justify-center transition-all group-hover:scale-110"
                    :class="answers[currentQuestion.id] === opt.value ? 'bg-[#3D3ACE] text-white' : 'bg-slate-50 text-slate-400'">
                    <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="opt.iconPath" />
                    </svg>
                </div>

                <div class="text-left flex-1">
                    <h3 class="text-lg font-bold transition-colors" :class="answers[currentQuestion.id] === opt.value ? 'text-[#3D3ACE]' : 'text-slate-900'">
                        {{ opt.title }}
                    </h3>
                    <p class="text-sm font-medium text-slate-400">{{ opt.desc }}</p>
                </div>

                <div v-if="answers[currentQuestion.id] === opt.value" class="w-6 h-6 bg-[#3D3ACE] rounded-full flex items-center justify-center text-white">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7" />
                    </svg>
                </div>
            </button>
        </div>

        <!-- Navigation -->
        <div class="mt-20 flex items-center justify-between max-w-xl mx-auto">
            <button @click="prevQuestion" v-if="!isFirstQuestion" class="flex items-center gap-2 text-slate-400 font-bold hover:text-slate-900 transition-colors">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7 7-7"/></svg>
                Previous Question
            </button>
            <div v-else></div>

            <button @click="handleNext"
                :disabled="!answers[currentQuestion.id] || submitting"
                class="bg-[#3D3ACE] hover:bg-[#322fb0] disabled:bg-slate-200 disabled:cursor-not-allowed text-white font-bold px-12 py-5 rounded-[20px] shadow-xl shadow-indigo-100 transition-all active:scale-[0.98]">
                {{ isLastQuestionInAll ? (submitting ? 'Analyzing...' : 'View Results') : 'Continue' }}
            </button>
        </div>
      </div>

      <!-- Sidebar -->
      <div class="w-full lg:w-[320px] space-y-8">
        <div class="bg-white rounded-[32px] p-8 border border-slate-100 shadow-sm relative overflow-hidden group">
            <div class="w-12 h-12 bg-indigo-50 text-[#3D3ACE] rounded-2xl flex items-center justify-center mb-6">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0012 18.75c-1.03 0-1.9-.4-2.593-1.012l-.547-.547z" /></svg>
            </div>
            <h3 class="text-xl font-extrabold text-[#1E1B4B] mb-3">Reflective Learning</h3>
            <p class="text-slate-500 text-[14px] leading-relaxed font-medium">
                Taking a moment to reflect on your habits is the first step toward a more focused academic life. We use these insights to calibrate your Lumora study cycles.
            </p>
            <div class="absolute top-[-20px] right-[-20px] w-24 h-24 bg-indigo-50 rounded-full blur-2xl group-hover:bg-indigo-100 transition-colors"></div>
        </div>

        <div class="rounded-[32px] overflow-hidden shadow-2xl shadow-indigo-50 rotate-2">
            <img src="https://images.unsplash.com/photo-1484480974693-6ca0a78fb36b?auto=format&fit=crop&q=80&w=300&h=400" alt="Reflection" class="w-full h-[280px] object-cover">
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { router } from '@inertiajs/vue3'
import axios from 'axios'

// ─── Data & Logic ────────────────────────────────────────────────────────────

const props = defineProps({
  questions: { type: Array, default: () => [] }
})

const currentSegmentIndex = ref(0)
const currentQuestionIndex = ref(0)
const answers = ref({})
const submitting = ref(false)

// Segments matching DB categories
const segments = [
    { id: 'planning', shortTitle: 'Planning', title: 'Planning Habits' },
    { id: 'time_management', shortTitle: 'Time Mgmt', title: 'Time Management' },
    { id: 'cognitive', shortTitle: 'Cognitive', title: 'Cognitive Strategies' },
    { id: 'reflection', shortTitle: 'Reflection', title: 'Self-Reflection' },
]

// Exact 20 questions from database image
const allQuestions = [
    // Planning
    { id: '11111111-0001-0001-0001-000000000001', segId: 'planning', text: 'Saya membuat jadwal belajar sebelum mulai belajar' },
    { id: '11111111-0001-0001-0001-000000000002', segId: 'planning', text: 'Saya menetapkan target belajar yang jelas' },
    { id: '11111111-0001-0001-0001-000000000003', segId: 'planning', text: 'Saya merencanakan materi yang akan dipelajari' },
    { id: '11111111-0001-0001-0001-000000000004', segId: 'planning', text: 'Saya menentukan waktu belajar secara rutin' },
    { id: '11111111-0001-0001-0001-000000000005', segId: 'planning', text: 'Saya mempersiapkan kebutuhan belajar sebelum mulai' },
    
    // Time Management
    { id: '11111111-0001-0001-0002-000000000006', segId: 'time_management', text: 'Saya mengatur waktu belajar dengan baik' },
    { id: '11111111-0001-0001-0002-000000000007', segId: 'time_management', text: 'Saya menyelesaikan tugas tepat waktu' },
    { id: '11111111-0001-0001-0002-000000000008', segId: 'time_management', text: 'Saya jarang menunda pekerjaan' },
    { id: '11111111-0001-0001-0002-000000000009', segId: 'time_management', text: 'Saya memprioritaskan tugas yang penting' },
    { id: '11111111-0001-0001-0002-000000000010', segId: 'time_management', text: 'Saya konsisten dengan jadwal belajar saya' },
    
    // Cognitive
    { id: '11111111-0001-0001-0003-000000000011', segId: 'cognitive', text: 'Saya menggunakan metode belajar tertentu (mencatat, merangkum, dll)' },
    { id: '11111111-0001-0001-0003-000000000012', segId: 'cognitive', text: 'Saya mencoba berbagai cara belajar untuk menemukan yang efektif' },
    { id: '11111111-0001-0001-0003-000000000013', segId: 'cognitive', text: 'Saya memahami materi, bukan hanya menghafal' },
    { id: '11111111-0001-0001-0003-000000000014', segId: 'cognitive', text: 'Saya mengulang materi untuk memperkuat pemahaman' },
    { id: '11111111-0001-0001-0003-000000000015', segId: 'cognitive', text: 'Saya menghubungkan materi dengan pengetahuan sebelumnya' },
    
    // Reflection
    { id: '11111111-0001-0001-0004-000000000016', segId: 'reflection', text: 'Saya mengecek apakah saya memahami materi' },
    { id: '11111111-0001-0001-0004-000000000017', segId: 'reflection', text: 'Saya menyadari ketika saya tidak memahami sesuatu' },
    { id: '11111111-0001-0001-0004-000000000018', segId: 'reflection', text: 'Saya mengevaluasi cara belajar saya' },
    { id: '11111111-0001-0001-0004-000000000019', segId: 'reflection', text: 'Saya memperbaiki strategi belajar jika kurang efektif' },
    { id: '11111111-0001-0001-0004-000000000020', segId: 'reflection', text: 'Saya belajar dari kesalahan sebelumnya' },
]

const likertOptions = [
    { value: 1, title: 'Never', desc: 'Sangat Tidak Setuju / Tidak Pernah', iconPath: 'M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728A9 9 0 015.636 5.636m12.728 12.728L5.636 5.636' },
    { value: 2, title: 'Rarely', desc: 'Tidak Setuju / Jarang', iconPath: 'M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z' },
    { value: 3, title: 'Sometimes', desc: 'Netral / Kadang-kadang', iconPath: 'M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z' },
    { value: 4, title: 'Often', desc: 'Setuju / Sering', iconPath: 'M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4' },
    { value: 5, title: 'Always', desc: 'Sangat Setuju / Selalu', iconPath: 'M5 3v4M3 5h4M6 17v4m-2-2h4m5-16l2.286 6.857L21 12l-7.714 2.143L11 21l-2.286-6.857L1 12l7.714-2.143L11 3z' },
]

// ─── Computed ─────────────────────────────────────────────────────────────

const currentSegment = computed(() => segments[currentSegmentIndex.value] || segments[0])

const currentSegmentQuestions = computed(() => {
    return allQuestions.filter(q => q.segId === currentSegment.value.id)
})

const currentQuestion = computed(() => {
    return currentSegmentQuestions.value[currentQuestionIndex.value] || { text: 'Loading...', id: 'loading' }
})

const isFirstQuestion = computed(() => {
    return currentSegmentIndex.value === 0 && currentQuestionIndex.value === 0
})

const isLastQuestionInSegment = computed(() => {
    return currentQuestionIndex.value === currentSegmentQuestions.value.length - 1
})

const isLastQuestionInAll = computed(() => {
    return currentSegmentIndex.value === segments.length - 1 && isLastQuestionInSegment.value
})

// ─── Methods ───────────────────────────────────────────────────────────────

function setAnswer(questionId, value) {
    if (questionId === 'loading') return
    answers.value[questionId] = value
}

function handleNext() {
    if (!answers.value[currentQuestion.value.id]) return

    if (isLastQuestionInSegment.value) {
        if (currentSegmentIndex.value < segments.length - 1) {
            currentSegmentIndex.value++
            currentQuestionIndex.value = 0
            window.scrollTo({ top: 0, behavior: 'smooth' })
        } else {
            submitAll()
        }
    } else {
        currentQuestionIndex.value++
    }
}

function prevQuestion() {
    if (currentQuestionIndex.value > 0) {
        currentQuestionIndex.value--
    } else if (currentSegmentIndex.value > 0) {
        currentSegmentIndex.value--
        // Go to last question of previous segment
        const prevSegId = segments[currentSegmentIndex.value].id
        currentQuestionIndex.value = allQuestions.filter(q => q.segId === prevSegId).length - 1
    }
}

async function submitAll() {
    submitting.value = true
    try {
        const payload = Object.keys(answers.value).map(id => ({
            question_id: id,
            answer_value: answers.value[id]
        }))

        router.post(route('onboarding.submit'), {
            answers: payload
        }, {
            onStart: () => { submitting.value = true },
            onFinish: () => { submitting.value = false },
            onSuccess: (page) => {
                // Cache the result for other components like Dashboard
                // page.props.result should be updated by the redirect to result page
                if (page.props.result) {
                    sessionStorage.setItem('lumora_result', JSON.stringify(page.props.result))
                }
                localStorage.removeItem('lumora_survey_progress')
            },
            onError: (errors) => {
                console.error('Submit error:', errors)
            }
        })
    } catch (err) {
        console.error('Unexpected error:', err)
    }
}

// ─── Persistence ───────────────────────────────────────────────────────────

const STORAGE_KEY = 'lumora_survey_progress'

function saveProgress() {
    localStorage.setItem(STORAGE_KEY, JSON.stringify({
        answers: answers.value,
        segIdx: currentSegmentIndex.value,
        qIdx: currentQuestionIndex.value
    }))
}

onMounted(() => {
    const saved = localStorage.getItem(STORAGE_KEY)
    if (saved) {
        try {
            const parsed = JSON.parse(saved)
            answers.value = parsed.answers || {}
            currentSegmentIndex.value = parsed.segIdx || 0
            currentQuestionIndex.value = parsed.qIdx || 0
        } catch (e) {
            console.error('Error parsing progress:', e)
        }
    }
})

watch([answers, currentSegmentIndex, currentQuestionIndex], saveProgress, { deep: true })

</script>

<style scoped>
.fade-enter-active, .fade-leave-active { transition: opacity 0.3s ease, transform 0.3s ease; }
.fade-enter-from { opacity: 0; transform: translateY(10px); }
.fade-leave-to { opacity: 0; transform: translateY(-10px); }
</style>



