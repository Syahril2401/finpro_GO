<template>
  <div class="min-h-screen bg-[#F8F9FE] font-sans text-slate-900 pb-32 pt-16 px-6">

    <!-- Loading State -->
    <div v-if="loading" class="max-w-[960px] mx-auto py-32 flex flex-col items-center justify-center text-center">
        <div class="w-20 h-20 border-4 border-indigo-100 border-t-[#3D3ACE] rounded-full animate-spin mb-8"></div>
        <h2 class="text-2xl font-black text-[#1E1B4B] mb-4">Finalizing Your Sanctuary...</h2>
        <p class="text-slate-500 font-medium max-w-md">Our AI is meticulously analyzing your study habits to build the perfect environment for your growth.</p>
    </div>

    <!-- Main Content -->
    <div v-else class="max-w-[960px] mx-auto">

      <!-- ═══ Hero Section ═══ -->
      <div class="text-center mb-14">
        <div class="inline-flex items-center gap-2 bg-[#D1FAE5] text-[#059669] px-4 py-2 rounded-full text-[11px] font-black uppercase tracking-widest mb-8 border border-emerald-100 shadow-sm shadow-emerald-50">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
            Assessment Complete
        </div>
        <h1 class="text-[36px] md:text-[44px] font-black text-[#1E1B4B] leading-tight mb-5">
            Your Study Profile: <span class="text-[#3D3ACE]">{{ profileTitle }}</span>
        </h1>
        <p class="text-slate-500 font-medium text-base md:text-lg max-w-2xl mx-auto leading-relaxed">
            We've analyzed your responses. You possess a rare ability to enter deep work states, though your tactical planning has room for optimization.
        </p>
      </div>

      <!-- ═══ Top Row: Core Strengths + Cognitive Style ═══ -->
      <div class="grid grid-cols-1 md:grid-cols-5 gap-6 mb-6">

        <!-- Core Strengths (wider, 3/5) -->
        <div class="md:col-span-3 bg-white rounded-[32px] p-8 shadow-lg shadow-indigo-50/50 border border-slate-100/80">
            <h3 class="text-lg font-black text-[#1E1B4B] mb-6">Core Strengths</h3>

            <div class="space-y-5 mb-8">
                <div v-for="(s, idx) in strengths" :key="idx" class="flex items-start gap-4">
                    <div class="w-10 h-10 bg-indigo-50 text-[#3D3ACE] rounded-xl flex items-center justify-center flex-shrink-0 mt-0.5">
                        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="getIconPath(s.icon)" />
                        </svg>
                    </div>
                    <div class="min-w-0">
                        <h4 class="text-[15px] font-bold text-[#1E1B4B] mb-1">{{ s.title }}</h4>
                        <p class="text-slate-400 font-medium text-[13px] leading-relaxed">{{ s.desc }}</p>
                    </div>
                </div>
            </div>

            <!-- Deep Work Capacity -->
            <div class="pt-6 border-t border-slate-100">
                <div class="flex justify-between items-end mb-3">
                    <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Deep Work Capacity</span>
                    <span class="text-lg font-black text-[#3D3ACE]">{{ deepWorkCapacity }}%</span>
                </div>
                <div class="h-2 bg-slate-100 rounded-full overflow-hidden">
                    <div class="h-full bg-[#3D3ACE] rounded-full transition-all duration-1000 shadow-[0_0_10px_rgba(61,58,206,0.3)]" :style="{ width: deepWorkCapacity + '%' }"></div>
                </div>
            </div>
        </div>

        <!-- Cognitive Style (narrower, 2/5) -->
        <div class="md:col-span-2 bg-white rounded-[32px] p-8 shadow-lg shadow-indigo-50/50 border border-slate-100/80 text-center flex flex-col items-center justify-center relative overflow-hidden">
            <div class="w-20 h-20 rounded-full overflow-hidden mb-5 border-4 border-white shadow-xl rotate-[-3deg] relative z-10">
                <img src="https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?auto=format&fit=crop&q=80&w=200&h=200" alt="Profile" class="w-full h-full object-cover">
            </div>
            <h3 class="text-lg font-black text-[#1E1B4B] mb-1 relative z-10">Your Cognitive Style</h3>
            <p class="text-slate-400 font-bold text-[11px] uppercase tracking-widest mb-6 relative z-10">{{ cognitiveStyle }}</p>

            <div class="grid grid-cols-2 gap-4 w-full relative z-10">
                <div class="bg-slate-50 p-5 rounded-2xl border border-slate-100/50">
                    <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest mb-2">Consistency</p>
                    <p class="text-2xl font-black text-[#1E1B4B]">{{ consistencyScore }}%</p>
                </div>
                <div class="bg-slate-50 p-5 rounded-2xl border border-slate-100/50">
                    <p class="text-[9px] font-black text-slate-400 uppercase tracking-widest mb-2">Retention</p>
                    <p class="text-2xl font-black text-[#1E1B4B]">{{ retentionScore }}%</p>
                </div>
            </div>
        </div>
      </div>

      <!-- ═══ Bottom Row: Areas for Growth + AI Strategy ═══ -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-10">

        <!-- Areas for Growth -->
        <div class="bg-white rounded-[32px] p-8 shadow-lg shadow-indigo-50/50 border border-slate-100/80">
            <h3 class="text-lg font-black text-[#1E1B4B] mb-6 flex items-center gap-2">
                <span class="text-xl">📈</span>
                Areas for Growth
            </h3>

            <div class="space-y-4">
                <div v-for="(w, idx) in weaknesses" :key="idx" class="p-5 bg-slate-50/80 rounded-2xl border border-slate-100 group hover:border-indigo-100 transition-all hover:bg-white hover:shadow-lg hover:shadow-indigo-50">
                    <h4 class="text-[15px] font-bold text-[#1E1B4B] mb-1.5">{{ w.title }}</h4>
                    <p class="text-[13px] font-medium text-slate-400 leading-relaxed">{{ w.desc }}</p>
                </div>
            </div>
        </div>

        <!-- AI Strategy -->
        <div class="bg-[#3D3ACE] rounded-[32px] p-8 shadow-xl shadow-indigo-200 text-white relative overflow-hidden group flex flex-col justify-between">
            <div class="relative z-10">
                <div class="inline-flex items-center gap-2 bg-white/10 px-3 py-1.5 rounded-full text-[10px] font-black uppercase tracking-widest mb-6 border border-white/10">
                    AI Recommendation
                </div>
                <h3 class="text-2xl font-black mb-4 leading-tight">{{ aiStrategy.title }}</h3>
                <p class="text-indigo-100 font-medium text-sm leading-relaxed mb-8">
                    {{ aiStrategy.desc }}
                </p>
                <button class="flex items-center gap-2 font-bold text-sm group/btn hover:text-white transition-colors">
                    Explore Strategy
                    <svg class="w-4 h-4 transition-transform group-hover/btn:translate-x-1" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M14 5l7 7m0 0l-7 7m7-7H3" /></svg>
                </button>
            </div>

            <!-- Abstract shapes -->
            <div class="absolute bottom-[-40px] right-[-40px] w-48 h-48 bg-white/10 rounded-full blur-3xl group-hover:bg-white/20 transition-colors duration-700"></div>
            <div class="absolute top-8 right-8 opacity-10 group-hover:opacity-20 group-hover:scale-110 transition-all duration-700">
                <svg class="w-24 h-24" fill="currentColor" viewBox="0 0 24 24"><path d="M12 2L4.5 20.29l.71.71L12 18l6.79 3 .71-.71z"/></svg>
            </div>
        </div>
      </div>

      <!-- ═══ Action Buttons ═══ -->
      <div class="flex flex-col items-center pt-6 gap-8">
        <p class="text-slate-400 font-semibold italic text-sm">"Small habits are the architecture of great minds." — Lumora AI</p>

        <div class="flex flex-col sm:flex-row gap-4 w-full max-w-[460px]">
            <Link :href="route('dashboard')" class="flex-1 bg-[#3D3ACE] hover:bg-[#322fb0] text-white font-bold py-5 rounded-2xl text-center shadow-xl shadow-indigo-200 transition-all active:scale-[0.98] hover:shadow-indigo-300">
                Go to Dashboard
            </Link>
            <button class="flex-1 bg-slate-200 hover:bg-slate-300 text-slate-700 font-bold py-5 rounded-2xl transition-all active:scale-[0.98]">
                Download Report
            </button>
        </div>

        <!-- Footer markers -->
        <div class="flex items-center gap-6 text-slate-300/60 pb-4">
            <div class="w-1.5 h-1.5 rounded-full bg-current"></div>
            <div class="w-2 h-2 rounded-full bg-indigo-200"></div>
            <div class="w-1.5 h-1.5 rounded-full bg-current"></div>
        </div>
      </div>
    </div>

    <!-- Live indicator -->
    <div class="fixed bottom-8 right-8 flex items-center gap-3 bg-white/90 backdrop-blur-xl px-6 py-3 rounded-full shadow-[0_15px_40px_rgba(0,0,0,0.08)] border border-slate-100 z-50">
        <div class="relative w-2 h-2">
            <div class="absolute inset-0 bg-emerald-500 rounded-full animate-ping opacity-75"></div>
            <div class="relative w-2 h-2 bg-emerald-500 rounded-full"></div>
        </div>
        <span class="text-[10px] font-black text-slate-500 uppercase tracking-[0.15em]">AI Assistant is active</span>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { Link } from '@inertiajs/vue3'

const props = defineProps({
  result: { type: Object, default: () => ({}) }
})

// ─── Loading State ──────────────────────────────────────────────────────────

const loading = ref(true)

// ─── Data Parsing ───────────────────────────────────────────────────────────
// The Go backend returns ResultSummary with PascalCase keys (no json tags).
// The AI JSON is stored as a string in `CategoryResult`.
// We need to: 1) find the CategoryResult field, 2) JSON.parse it.

function tryParseAI(obj) {
    if (!obj || typeof obj !== 'object') return null

    // Go sends PascalCase (no json tags on ResultSummary)
    const raw = obj.CategoryResult || obj.category_result
    if (raw && typeof raw === 'string') {
        try {
            const parsed = JSON.parse(raw)
            if (parsed && parsed.profile_title) return parsed
        } catch (e) {
            console.warn('Failed to parse CategoryResult:', e)
        }
    }

    // Maybe it's already a flat AI object
    if (obj.profile_title) return obj

    return null
}

const aiData = computed(() => {
    // 1. Try from Inertia props (server-side session)
    const fromProps = tryParseAI(props.result)
    if (fromProps) return fromProps

    // 2. Try from sessionStorage (client-side cache from submit)
    try {
        const stored = sessionStorage.getItem('lumora_result')
        if (stored) {
            const parsed = JSON.parse(stored)
            const fromStored = tryParseAI(parsed)
            if (fromStored) return fromStored
        }
    } catch {}

    // 3. Return empty → triggers fallback values
    return {}
})

// ─── Computed Fields with Fallbacks ─────────────────────────────────────────

const profileTitle = computed(() => aiData.value.profile_title || 'The Focused Achiever')
const deepWorkCapacity = computed(() => aiData.value.deep_work_capacity || 92)
const cognitiveStyle = computed(() => aiData.value.cognitive_style || 'Architectural & Analytical')
const consistencyScore = computed(() => aiData.value.consistency_score || 64)
const retentionScore = computed(() => aiData.value.retention_score || 88)

// Default fallback icon paths
const defaultIcons = [
    'M10 21h7a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v11m10 5l4-4',
    'M13 10V3L4 14h7v7l9-11h-7z',
    'M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z'
]

function getIconPath(icon) {
    if (icon && typeof icon === 'string' && icon.startsWith('M')) return icon
    return defaultIcons[0]
}

const strengths = computed(() => {
    const raw = aiData.value.strengths
    if (Array.isArray(raw) && raw.length > 0) {
        return raw.map((s, i) => ({
            title: s.title || `Strength ${i + 1}`,
            desc: s.desc || s.description || '',
            icon: s.icon || defaultIcons[i % defaultIcons.length]
        }))
    }
    return [
        { title: 'Strong Concentration', desc: 'You naturally filter out digital noise and maintain high-quality focus for sessions exceeding 90 minutes.', icon: defaultIcons[0] },
        { title: 'Conceptual Mastery', desc: 'You excel at connecting complex ideas rather than just memorizing isolated facts.', icon: defaultIcons[1] }
    ]
})

const weaknesses = computed(() => {
    // Prioritize areas_for_growth as requested in new prompt, fallback to weaknesses
    const raw = aiData.value.areas_for_growth || aiData.value.weaknesses
    if (Array.isArray(raw) && raw.length > 0) {
        return raw.map((w, i) => ({
            title: w.title || `Area ${i + 1}`,
            desc: w.desc || w.description || ''
        }))
    }
    return [
        { title: 'Planning Consistency', desc: 'Your performance peaks in bursts. Stabilizing your daily routine will prevent late-night fatigue.' },
        { title: 'Recall Verification', desc: 'Transition from passive reading to active testing to solidify long-term memory structures.' }
    ]
})

const aiStrategy = computed(() => {
    // Prioritize ai_strategy, fallback to first recommendation
    const raw = aiData.value.ai_strategy || (aiData.value.recommendations && aiData.value.recommendations[0])
    if (raw && raw.title) return raw
    return {
        title: 'The "Interval" Strategy',
        desc: 'Based on your high concentration scores, we recommend 90-minute Deep Cycles followed by 15-minute sensory breaks. Use the first 5 minutes to map your targets.'
    }
})

// ─── Lifecycle & Reactivity ──────────────────────────────────────────────────
import { watch } from 'vue'

watch(() => props.result, (newVal) => {
    if (newVal && (newVal.profile_title || newVal.CategoryResult)) {
        loading.value = false
    }
}, { immediate: true })

onMounted(() => {
    // If we already have data from props, loading is handled by watch immediate
    // Otherwise, wait up to 5 seconds before showing fallback to give AI time
    if (!aiData.value.profile_title) {
        setTimeout(() => {
            loading.value = false
        }, 5000)
    } else {
        loading.value = false
    }
})
</script>
