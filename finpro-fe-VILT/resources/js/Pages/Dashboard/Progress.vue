<template>
  <DashboardLayout 
    :showBuddy="true"
    buddyContextTitle="Progress Insight" 
    :buddyContextMessage="insightMessage"
  >
    <!-- Header -->
    <div class="mb-10">
      <h1 class="text-3xl font-black text-[#1E1B4B] mb-2">Progress</h1>
      <p class="text-slate-500 font-medium">Track your learning development over time.</p>
    </div>

    <!-- Loading -->
    <div v-if="isLoading" class="space-y-6">
      <div class="grid grid-cols-2 lg:grid-cols-4 gap-4">
        <div v-for="i in 4" :key="i" class="h-28 bg-slate-100 rounded-[28px] animate-pulse"></div>
      </div>
      <div class="h-64 bg-slate-100 rounded-[28px] animate-pulse"></div>
    </div>

    <template v-else>
      <!-- Metric Cards -->
      <div class="grid grid-cols-2 lg:grid-cols-4 gap-4 mb-10">
        <!-- Current Profile -->
        <div class="bg-white p-6 rounded-[28px] border border-slate-100 shadow-sm">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">Current Profile</p>
          <p v-if="progress.latest_result" class="text-lg font-black text-[#1E1B4B] truncate">{{ profileTitle }}</p>
          <p v-else class="text-sm font-bold text-slate-400">No profile yet</p>
          <a v-if="!progress.latest_result" :href="route('onboarding.sanctuary')" class="text-xs text-[#3D3ACE] font-bold mt-2 inline-block hover:underline">Take Assessment →</a>
        </div>
        <!-- Consistency -->
        <div class="bg-white p-6 rounded-[28px] border border-slate-100 shadow-sm">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">Consistency</p>
          <p class="text-3xl font-black text-[#3D3ACE]">{{ Math.round(progress.consistency || 0) }}%</p>
          <p class="text-xs font-bold text-slate-400 mt-1">active days / week</p>
        </div>
        <!-- Target Completion -->
        <div class="bg-white p-6 rounded-[28px] border border-slate-100 shadow-sm">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">Target Completion</p>
          <p class="text-3xl font-black text-emerald-500">{{ Math.round(progress.target_completion || 0) }}%</p>
          <p class="text-xs font-bold text-slate-400 mt-1">subtasks done</p>
        </div>
        <!-- Reflections -->
        <div class="bg-white p-6 rounded-[28px] border border-slate-100 shadow-sm">
          <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">Reflections</p>
          <p class="text-3xl font-black text-purple-500">{{ progress.reflections_count || 0 }}</p>
          <p class="text-xs font-bold text-slate-400 mt-1">this month</p>
        </div>
      </div>

      <!-- SRL Score Trend -->
      <div class="bg-white p-8 rounded-[32px] border border-slate-100 shadow-sm mb-10">
        <h3 class="text-lg font-black text-[#1E1B4B] mb-6">SRL Score Trend</h3>
        <div v-if="trendData.length === 0" class="text-center py-10">
          <p class="text-slate-400 font-bold">No assessment data yet. Take your first assessment to see trends.</p>
        </div>
        <div v-else class="space-y-3">
          <div v-for="entry in trendData" :key="entry.ResultID" class="flex items-center gap-4">
            <span class="text-xs font-bold text-slate-400 w-24 flex-shrink-0">{{ formatDate(entry.CreatedAt) }}</span>
            <div class="flex-1 flex items-center gap-2">
              <div class="flex-1 h-3 bg-slate-50 rounded-full overflow-hidden">
                <div class="h-full bg-gradient-to-r from-[#3D3ACE] to-indigo-400 rounded-full transition-all duration-500" :style="{ width: `${avgScore(entry)}%` }"></div>
              </div>
              <span class="text-sm font-black text-[#1E1B4B] w-12 text-right">{{ avgScore(entry) }}%</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Dimension Analysis -->
      <div class="bg-white p-8 rounded-[32px] border border-slate-100 shadow-sm mb-10">
        <h3 class="text-lg font-black text-[#1E1B4B] mb-6">Dimension Analysis</h3>
        <div v-if="!progress.dimension_delta" class="text-center py-10">
          <p class="text-slate-400 font-bold">Take another assessment later to unlock comparative dimension analysis.</p>
        </div>
        <div v-else class="grid grid-cols-2 lg:grid-cols-4 gap-4">
          <div v-for="(dim, key) in dimensionLabels" :key="key" class="bg-slate-50 p-5 rounded-2xl">
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">{{ dim }}</p>
            <div class="flex items-center gap-2">
              <span class="text-2xl font-black" :class="deltaColor(progress.dimension_delta[key])">
                {{ progress.dimension_delta[key] > 0 ? '+' : '' }}{{ progress.dimension_delta[key] || 0 }}
              </span>
              <svg v-if="progress.dimension_delta[key] > 0" class="w-4 h-4 text-emerald-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 15l7-7 7 7"/></svg>
              <svg v-else-if="progress.dimension_delta[key] < 0" class="w-4 h-4 text-rose-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M19 9l-7 7-7-7"/></svg>
              <span v-else class="text-xs font-bold text-slate-400">—</span>
            </div>
          </div>
        </div>
      </div>

      <!-- CTA Buttons -->
      <div class="flex flex-wrap gap-4">
        <a :href="route('onboarding.sanctuary')" class="bg-[#3D3ACE] hover:bg-[#322fb0] text-white px-6 py-3 rounded-2xl font-bold shadow-xl shadow-indigo-200 transition-all active:scale-95 flex items-center gap-2">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/></svg>
          Retake Assessment
        </a>
      </div>
    </template>

  </DashboardLayout>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { usePage } from '@inertiajs/vue3'
import DashboardLayout from '@/Layouts/DashboardLayout.vue'
import { progressApi, setWorkspaceToken } from '@/services/workspaceApi'

const page = usePage()
const isLoading = ref(true)
const progress = ref({})

const dimensionLabels = {
  planning: 'Planning',
  time_management: 'Time Management',
  cognitive: 'Cognitive Strategy',
  reflection: 'Reflection'
}

const trendData = computed(() => progress.value.assessment_trend || [])

const profileTitle = computed(() => {
  const r = progress.value.latest_result
  if (!r) return ''
  try {
    if (r.CategoryResult) {
      const parsed = JSON.parse(r.CategoryResult)
      return parsed.profile_title || 'Learner'
    }
  } catch {}
  return 'Learner'
})

const insightMessage = computed(() => {
  if (trendData.value.length < 2) return 'Take another assessment later to unlock comparison and dimension analysis.'
  if ((progress.value.target_completion || 0) < 50) return 'Your progress becomes more accurate when you complete targets and write reflections consistently.'
  return 'Great consistency! Keep refining your learning strategies to improve your SRL scores.'
})

onMounted(async () => {
  setWorkspaceToken(page.props.go_token)
  try {
    progress.value = await progressApi.getProgress()
  } catch (err) {
    console.error('Failed to fetch progress', err)
  } finally {
    isLoading.value = false
  }
})

function avgScore(entry) {
  const scores = [entry.PlanningScore, entry.TimeManagementScore, entry.CognitiveScore, entry.ReflectionScore].filter(s => s != null)
  if (scores.length === 0) return 0
  const avg = scores.reduce((a, b) => a + b, 0) / scores.length
  return Math.round(avg / 5 * 100) // Assuming max score 5 per dimension
}

function formatDate(dateStr) {
  if (!dateStr) return ''
  return new Date(dateStr).toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })
}

function deltaColor(val) {
  if (val > 0) return 'text-emerald-500'
  if (val < 0) return 'text-rose-500'
  return 'text-slate-400'
}
</script>
