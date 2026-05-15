<template>
  <DashboardLayout 
    :showBuddy="true"
    buddyContextTitle="Study Insight" 
    :buddyContextMessage="insightMessage"
  >
    <!-- Greeting Hero -->
    <div class="mb-12">
        <p class="text-[11px] font-black text-[#3D3ACE] uppercase tracking-[0.3em] mb-3">Your Sanctuary Dashboard</p>
        <h1 class="text-4xl font-black text-[#1E1B4B] mb-4">{{ greeting }}, {{ userName }}! Ready to focus today?</h1>
        <p class="text-slate-500 font-medium max-w-2xl leading-relaxed">
            Welcome back. You have <span class="text-[#3D3ACE] font-black">{{ data.focus_sessions || 0 }} focus sessions</span> completed this week.
        </p>
    </div>

    <!-- Dashboard Grid -->
    <div class="grid grid-cols-12 gap-8 mb-12">
        <!-- Stats Cards -->
        <div class="col-span-12 lg:col-span-4 flex flex-col gap-6">
            <div v-for="stat in metrics" :key="stat.label" 
                class="bg-white p-8 rounded-[36px] shadow-sm border border-slate-50 hover:shadow-xl hover:shadow-indigo-50 transition-all group">
                <div class="flex items-center justify-between mb-4">
                    <div class="w-12 h-12 rounded-2xl flex items-center justify-center text-xl transition-transform group-hover:scale-110" :class="stat.bg">
                        {{ stat.icon }}
                    </div>
                </div>
                <p class="text-3xl font-black text-[#1E1B4B] mb-1">{{ stat.value }}</p>
                <p class="text-xs font-bold text-slate-400 uppercase tracking-widest">{{ stat.label }}</p>
            </div>
        </div>

        <!-- Focus Map -->
        <div class="col-span-12 lg:col-span-8">
            <div class="bg-white p-10 rounded-[48px] shadow-sm border border-slate-50 h-full flex flex-col relative overflow-hidden group">
                <div class="flex items-center justify-between mb-10 relative z-10">
                    <div>
                        <h3 class="text-xl font-black text-[#1E1B4B]">Your Focus Map</h3>
                        <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mt-1">Weekly Activity Tracking</p>
                    </div>
                </div>
                
                <div class="flex-1 flex items-end justify-between px-2 gap-4 relative z-10">
                    <div v-for="(stat, i) in weeklyStats" :key="i" class="flex-1 flex flex-col items-center gap-4 group/bar">
                        <div class="w-full bg-slate-50 rounded-2xl relative overflow-hidden flex items-end" style="height: 200px">
                            <div class="w-full bg-indigo-50 group-hover/bar:bg-indigo-100 transition-all absolute bottom-0 left-0" :style="{ height: stat.value + '%' }"></div>
                            <div class="w-full bg-[#3D3ACE] rounded-t-2xl transition-all duration-1000 relative z-10 group-hover/bar:shadow-[0_0_20px_rgba(61,58,206,0.4)]" 
                                :style="{ height: (stat.value * 0.7) + '%' }">
                                <div class="absolute -top-10 left-1/2 -translate-x-1/2 bg-[#1E1B4B] text-white text-[10px] px-2 py-1 rounded opacity-0 group-hover/bar:opacity-100 transition-opacity font-black whitespace-nowrap">
                                    {{ stat.count }} sessions
                                </div>
                            </div>
                        </div>
                        <span class="text-[10px] font-black uppercase tracking-widest" :class="stat.isToday ? 'text-[#3D3ACE]' : 'text-slate-400'">{{ stat.day }}</span>
                    </div>
                </div>

                <div class="absolute bottom-0 right-0 opacity-[0.03] rotate-12 group-hover:scale-110 group-hover:rotate-6 transition-all duration-1000">
                    <svg class="w-64 h-64" fill="currentColor" viewBox="0 0 24 24"><path d="M12 2L4.5 20.29l.71.71L12 18l6.79 3 .71-.71z"/></svg>
                </div>
            </div>
        </div>
    </div>

    <!-- Bottom Sections -->
    <div class="grid grid-cols-12 gap-8">
        <!-- Recent Sessions -->
        <div class="col-span-12 lg:col-span-7">
            <div class="bg-white p-10 rounded-[48px] shadow-sm border border-slate-50 h-full">
                <div class="flex items-center justify-between mb-10">
                    <h3 class="text-xl font-black text-[#1E1B4B]">Recent Focus Sessions</h3>
                    <Link :href="route('planner')" class="text-[11px] font-black text-[#3D3ACE] uppercase tracking-widest border-b-2 border-[#3D3ACE]">View All</Link>
                </div>
                <div class="space-y-4">
                    <div v-if="!recentSessions.length" class="text-center py-10">
                        <p class="text-slate-400 font-bold">No sessions yet. Start your first deep work cycle!</p>
                        <Link :href="route('planner')" class="text-[#3D3ACE] font-bold mt-2 inline-block hover:underline">Go to Planner →</Link>
                    </div>
                    <div v-for="item in recentSessions.slice(0, 5)" :key="item.id"
                        class="flex items-center gap-6 p-6 rounded-[32px] border border-slate-50 hover:border-[#3D3ACE]/20 hover:shadow-xl hover:shadow-indigo-50 transition-all cursor-pointer group">
                        <div class="w-14 h-14 rounded-2xl flex items-center justify-center text-2xl transition-transform group-hover:scale-110"
                            :class="item.status === 'completed' ? 'bg-emerald-50 text-emerald-500' : 'bg-indigo-50 text-[#3D3ACE]'">
                            {{ item.status === 'completed' ? '✅' : '🧠' }}
                        </div>
                        <div class="flex-1">
                            <h4 class="text-base font-black text-[#1E1B4B] group-hover:text-[#3D3ACE] transition-colors">{{ item.title }}</h4>
                            <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mt-1">{{ item.start_time }} • {{ item.duration_minutes }}m • {{ item.focus_dimension }}</p>
                        </div>
                        <span class="text-[10px] font-black px-3 py-1 rounded-full uppercase tracking-widest"
                            :class="item.status === 'completed' ? 'text-emerald-500 bg-emerald-50' : 'text-[#3D3ACE] bg-indigo-50'">
                            {{ item.status }}
                        </span>
                    </div>
                </div>
            </div>
        </div>

        <!-- Quick Actions -->
        <div class="col-span-12 lg:col-span-5 flex flex-col gap-8">
            <div class="bg-indigo-600 p-10 rounded-[48px] text-white shadow-2xl shadow-indigo-100 relative overflow-hidden group h-full">
                <div class="relative z-10">
                    <div class="bg-white/10 p-4 rounded-3xl w-fit mb-8 border border-white/10 backdrop-blur-xl group-hover:scale-110 transition-transform">
                        <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M13 5l7 7m0 0l-7 7m7-7H3" /></svg>
                    </div>
                    <h3 class="text-2xl font-black mb-4 leading-tight">Quick Actions</h3>
                    <div class="space-y-3 mb-8">
                      <Link :href="route('planner')" class="flex items-center gap-3 text-indigo-100 hover:text-white font-bold text-sm transition-colors">
                        <span class="w-2 h-2 bg-white/50 rounded-full"></span> Create Study Session
                      </Link>
                      <Link :href="route('targets')" class="flex items-center gap-3 text-indigo-100 hover:text-white font-bold text-sm transition-colors">
                        <span class="w-2 h-2 bg-white/50 rounded-full"></span> Set Weekly Target
                      </Link>
                      <Link :href="route('notes')" class="flex items-center gap-3 text-indigo-100 hover:text-white font-bold text-sm transition-colors">
                        <span class="w-2 h-2 bg-white/50 rounded-full"></span> Write Reflection
                      </Link>
                      <Link :href="route('progress')" class="flex items-center gap-3 text-indigo-100 hover:text-white font-bold text-sm transition-colors">
                        <span class="w-2 h-2 bg-white/50 rounded-full"></span> View Progress
                      </Link>
                    </div>
                    <Link :href="route('planner')" class="inline-block bg-white text-indigo-600 px-8 py-4 rounded-[20px] font-black text-sm shadow-xl hover:shadow-white/20 hover:-translate-y-1 transition-all active:scale-[0.98]">
                        Start Deep Work
                    </Link>
                </div>
                <div class="absolute bottom-[-60px] right-[-60px] w-64 h-64 bg-white/10 rounded-full blur-3xl transition-all group-hover:bg-white/20"></div>
            </div>
        </div>
    </div>
  </DashboardLayout>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { Link, usePage } from '@inertiajs/vue3'
import DashboardLayout from '@/Layouts/DashboardLayout.vue'
import { dashboardApi, setWorkspaceToken } from '@/services/workspaceApi'

const page = usePage()
const data = ref({})
const recentSessions = ref([])
const isLoading = ref(true)

const userName = computed(() => page.props.auth?.user?.name || 'User')

const greeting = computed(() => {
  const hour = new Date().getHours()
  if (hour < 12) return 'Good morning'
  if (hour < 18) return 'Good afternoon'
  return 'Good evening'
})

const metrics = computed(() => [
  { label: 'Focus Sessions', value: data.value.focus_sessions || 0, icon: '🎯', bg: 'bg-indigo-50 text-indigo-600' },
  { label: 'Task Efficiency', value: Math.round(data.value.task_efficiency || 0) + '%', icon: '⚡', bg: 'bg-emerald-50 text-emerald-600' },
  { label: 'Deep Work Hours', value: (data.value.deep_work_hours || 0).toFixed(1) + 'h', icon: '🔋', bg: 'bg-violet-50 text-violet-600' },
  { label: 'Consistency', value: Math.round(data.value.consistency || 0) + '%', icon: '📈', bg: 'bg-amber-50 text-amber-600' },
  { label: 'Learning Continuity', value: Math.round(data.value.learning_continuity || 0) + '%', icon: '🧠', bg: 'bg-rose-50 text-rose-600' },
])

const weeklyStats = computed(() => {
  const days = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
  const today = new Date()
  const todayDay = today.getDay()
  const monday = new Date(today)
  monday.setDate(today.getDate() - (todayDay === 0 ? 6 : todayDay - 1))

  return days.map((name, i) => {
    const d = new Date(monday)
    d.setDate(d.getDate() + i)
    const dateStr = d.toISOString().split('T')[0]
    const count = (data.value.weekly_focus_map && data.value.weekly_focus_map[dateStr]) || 0
    const maxCount = Math.max(...Object.values(data.value.weekly_focus_map || { x: 1 }), 1)
    return {
      day: name.toUpperCase().substring(0, 3),
      count,
      value: Math.round((count / maxCount) * 100),
      isToday: dateStr === today.toISOString().split('T')[0]
    }
  })
})

const insightMessage = computed(() => {
  const fs = data.value.focus_sessions || 0
  const te = data.value.task_efficiency || 0
  if (fs === 0 && te === 0) return 'Start with one study session, one weekly target, and one reflection note.'
  if (fs < 2) return 'Create at least 2 focused study sessions this week to build consistency.'
  return `Great momentum! You've completed ${fs} sessions with ${Math.round(te)}% task efficiency.`
})

onMounted(async () => {
  setWorkspaceToken(page.props.go_token)
  try {
    data.value = await dashboardApi.getMetrics()
    recentSessions.value = data.value.recent_sessions || []
  } catch (err) {
    console.error('Failed to load dashboard', err)
  } finally {
    isLoading.value = false
  }
})
</script>
