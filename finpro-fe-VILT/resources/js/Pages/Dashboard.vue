<template>
  <div class="min-h-screen bg-[#FDFDFF] font-sans text-slate-900 flex overflow-hidden">
    
    <!-- Left Sidebar -->
    <aside class="w-72 bg-white border-r border-slate-100 flex flex-col z-40 relative">
      <div class="p-8 pb-10 flex items-center gap-3">
        <div class="w-10 h-10 bg-[#3D3ACE] rounded-xl flex items-center justify-center text-white shadow-lg shadow-indigo-100">
            <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24"><path d="M12 2L4.5 20.29l.71.71L12 18l6.79 3 .71-.71z"/></svg>
        </div>
        <div>
            <span class="text-xl font-black text-[#1E1B4B] tracking-tight block">Lumora</span>
            <span class="text-[10px] font-black text-[#3D3ACE] uppercase tracking-widest bg-indigo-50 px-2 py-0.5 rounded-md">BETA v1.0</span>
        </div>
      </div>

      <nav class="flex-1 px-4 space-y-2">
        <button v-for="item in navItems" :key="item.label"
          @click="activeNav = item.label"
          class="w-full flex items-center gap-4 px-5 py-4 rounded-2xl text-[14px] font-bold transition-all group"
          :class="activeNav === item.label
            ? 'bg-[#3D3ACE] text-white shadow-xl shadow-indigo-100'
            : 'text-slate-400 hover:bg-slate-50 hover:text-slate-600'">
          <svg class="w-5 h-5 transition-transform group-hover:scale-110" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" :d="item.iconPath" />
          </svg>
          {{ item.label }}
          <div v-if="item.badge" class="ml-auto bg-red-500 text-white text-[10px] px-2 py-0.5 rounded-full font-black">
              {{ item.badge }}
          </div>
        </button>
      </nav>

      <!-- Bottom User Section -->
      <div class="p-6 border-t border-slate-50">
        <div class="bg-slate-50 p-4 rounded-3xl flex items-center gap-4 group cursor-pointer hover:bg-slate-100 transition-all">
            <div class="w-12 h-12 rounded-2xl overflow-hidden shadow-md">
                <img src="https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?auto=format&fit=crop&q=80&w=200&h=200" alt="Profile" class="w-full h-full object-cover">
            </div>
            <div class="flex-1 min-w-0">
                <p class="text-sm font-black text-[#1E1B4B] truncate">{{ dashboardData.user.name }}</p>
                <p class="text-[11px] font-bold text-slate-400 uppercase tracking-wider truncate">{{ profileTitle }}</p>
            </div>
            <svg class="w-5 h-5 text-slate-300 group-hover:text-slate-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/></svg>
        </div>
      </div>
    </aside>

    <!-- Main Content Area -->
    <main class="flex-1 flex flex-col relative">
        
        <!-- Top Bar -->
        <header class="h-24 px-10 border-b border-slate-100 flex items-center justify-between bg-white/50 backdrop-blur-xl z-30">
            <div class="flex-1 max-w-xl">
                <div @click="searchOpen = true" class="group relative flex items-center cursor-pointer">
                    <svg class="absolute left-5 w-5 h-5 text-slate-400 group-hover:text-[#3D3ACE] transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/></svg>
                    <div class="w-full bg-slate-50 border border-slate-100 rounded-2xl py-3.5 pl-14 pr-5 text-sm font-bold text-slate-400 group-hover:bg-white group-hover:border-[#3D3ACE]/20 transition-all flex items-center">
                        Search anything...
                        <span class="ml-auto text-[10px] bg-white px-2 py-1 rounded-lg border border-slate-100 shadow-sm font-black text-slate-400 tracking-tighter">⌘ K</span>
                    </div>
                </div>
            </div>

            <div class="flex items-center gap-6 ml-10">
                <button @click="notificationsOpen = !notificationsOpen" class="relative w-12 h-12 bg-white border border-slate-100 rounded-2xl flex items-center justify-center text-slate-400 hover:text-[#3D3ACE] hover:border-indigo-100 hover:shadow-lg hover:shadow-indigo-50 transition-all">
                    <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"/></svg>
                    <div class="absolute top-3 right-3 w-2.5 h-2.5 bg-red-500 border-2 border-white rounded-full"></div>
                </button>
                <div class="h-8 w-[1px] bg-slate-100"></div>
                <div class="flex items-center gap-4">
                    <div class="text-right max-md:hidden">
                        <p class="text-sm font-black text-[#1E1B4B]">{{ dashboardData.user.name }}</p>
                        <p class="text-[10px] font-bold text-[#059669] uppercase tracking-widest flex items-center justify-end gap-1.5">
                            <span class="w-1.5 h-1.5 bg-emerald-500 rounded-full animate-pulse"></span>
                            {{ dashboardData.user.status }}
                        </p>
                    </div>
                </div>
            </div>
        </header>

        <!-- Page Content -->
        <div class="flex-1 overflow-y-auto bg-[#FDFDFF] p-10 scrollbar-hide">
            
            <!-- Greeting Hero -->
            <div class="mb-12">
                <p class="text-[11px] font-black text-[#3D3ACE] uppercase tracking-[0.3em] mb-3">Your Sanctuary Dashboard</p>
                <h1 class="text-4xl font-black text-[#1E1B4B] mb-4">Morning, {{ dashboardData.user.name }}! Ready to focus today?</h1>
                <p class="text-slate-500 font-medium max-w-2xl leading-relaxed">
                    Welcome back. Based on your profile, we've optimized your learning sanctuary. You have <span class="text-[#3D3ACE] font-black">{{ focusSessionsCount }} focus sessions</span> completed so far.
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
                            <span class="text-[10px] font-black text-emerald-500 bg-emerald-50 px-2 py-1 rounded-lg">+12%</span>
                        </div>
                        <p class="text-3xl font-black text-[#1E1B4B] mb-1">{{ stat.value }}</p>
                        <p class="text-xs font-bold text-slate-400 uppercase tracking-widest">{{ stat.label }}</p>
                    </div>
                </div>

                <!-- Main Activity Chart -->
                <div class="col-span-12 lg:col-span-8">
                    <div class="bg-white p-10 rounded-[48px] shadow-sm border border-slate-50 h-full flex flex-col relative overflow-hidden group">
                        <div class="flex items-center justify-between mb-10 relative z-10">
                            <div>
                                <h3 class="text-xl font-black text-[#1E1B4B]">Your Focus Map</h3>
                                <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mt-1">Weekly Activity Tracking</p>
                            </div>
                            <div class="flex bg-slate-50 p-1.5 rounded-2xl border border-slate-100">
                                <button class="px-4 py-2 text-[11px] font-black text-slate-400 rounded-xl hover:text-slate-600">Month</button>
                                <button class="px-4 py-2 text-[11px] font-black text-[#1E1B4B] bg-white rounded-xl shadow-sm border border-slate-100">Week</button>
                            </div>
                        </div>
                        
                        <div class="flex-1 flex items-end justify-between px-2 gap-4 relative z-10">
                            <div v-for="(stat, i) in weeklyStats" :key="i" class="flex-1 flex flex-col items-center gap-4 group/bar">
                                <div class="w-full bg-slate-50 rounded-2xl relative overflow-hidden flex items-end" style="height: 200px">
                                    <div class="w-full bg-indigo-50 group-hover/bar:bg-indigo-100 transition-all absolute bottom-0 left-0" :style="{ height: stat.value + '%' }"></div>
                                    <div class="w-full bg-[#3D3ACE] rounded-t-2xl transition-all duration-1000 relative z-10 group-hover/bar:shadow-[0_0_20px_rgba(61,58,206,0.4)]" 
                                        :style="{ height: (stat.value * 0.7) + '%' }">
                                        <div class="absolute -top-10 left-1/2 -translate-x-1/2 bg-[#1E1B4B] text-white text-[10px] px-2 py-1 rounded opacity-0 group-hover/bar:opacity-100 transition-opacity font-black">
                                            {{ stat.value }}%
                                        </div>
                                    </div>
                                </div>
                                <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">{{ stat.day }}</span>
                            </div>
                        </div>

                        <!-- BG Decoration -->
                        <div class="absolute bottom-0 right-0 opacity-[0.03] rotate-12 group-hover:scale-110 group-hover:rotate-6 transition-all duration-1000">
                            <svg class="w-64 h-64" fill="currentColor" viewBox="0 0 24 24"><path d="M12 2L4.5 20.29l.71.71L12 18l6.79 3 .71-.71z"/></svg>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Bottom Sections -->
            <div class="grid grid-cols-12 gap-8">
                <!-- Focus Sessions List -->
                <div class="col-span-12 lg:col-span-7">
                    <div class="bg-white p-10 rounded-[48px] shadow-sm border border-slate-50 h-full">
                        <div class="flex items-center justify-between mb-10">
                            <h3 class="text-xl font-black text-[#1E1B4B]">Recent Focus Sessions</h3>
                            <button class="text-[11px] font-black text-[#3D3ACE] uppercase tracking-widest border-b-2 border-[#3D3ACE]">View Archive</button>
                        </div>
                        <div class="space-y-4">
                            <div v-if="dashboardData.focus_sessions.length === 0" class="text-center py-10">
                                <p class="text-slate-400 font-bold">No sessions today. Start your first deep work cycle!</p>
                            </div>
                            <div v-for="item in dashboardData.focus_sessions" :key="item.title"
                                class="flex items-center gap-6 p-6 rounded-[32px] border border-slate-50 hover:border-[#3D3ACE]/20 hover:shadow-xl hover:shadow-indigo-50 transition-all cursor-pointer group">
                                <div class="w-14 h-14 rounded-2xl flex items-center justify-center text-2xl transition-transform group-hover:scale-110"
                                    :class="item.status === 'done' ? 'bg-emerald-50 text-emerald-500' : 'bg-indigo-50 text-[#3D3ACE]'">
                                    {{ item.status === 'done' ? '📂' : '🧠' }}
                                </div>
                                <div class="flex-1">
                                    <h4 class="text-base font-black text-[#1E1B4B] group-hover:text-[#3D3ACE] transition-colors">{{ item.title }}</h4>
                                    <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mt-1">{{ item.time }} • {{ item.duration }}</p>
                                </div>
                                <div class="flex flex-col items-end gap-2">
                                    <span v-if="item.status === 'done'" class="text-[10px] font-black text-emerald-500 bg-emerald-50 px-3 py-1 rounded-full uppercase tracking-widest">Completed</span>
                                    <span v-else class="text-[10px] font-black text-[#3D3ACE] bg-indigo-50 px-3 py-1 rounded-full uppercase tracking-widest">{{ item.status }}</span>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                <!-- AI Insight Buddy -->
                <div class="col-span-12 lg:col-span-5 flex flex-col gap-8">
                    <div class="bg-indigo-600 p-10 rounded-[48px] text-white shadow-2xl shadow-indigo-100 relative overflow-hidden group h-full">
                        <div class="relative z-10">
                            <div class="bg-white/10 p-4 rounded-3xl w-fit mb-8 border border-white/10 backdrop-blur-xl group-hover:scale-110 transition-transform">
                                <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M13 5l7 7m0 0l-7 7m7-7H3" /></svg>
                            </div>
                            <h3 class="text-2xl font-black mb-4 leading-tight">Your AI Buddy Insight</h3>
                            <p class="text-indigo-100 font-medium mb-10 leading-relaxed text-sm">
                                {{ latestAIInsight }}
                            </p>
                            <button class="bg-white text-indigo-600 px-8 py-4 rounded-[20px] font-black text-sm shadow-xl hover:shadow-white/20 hover:-translate-y-1 transition-all active:scale-[0.98]">
                                Start Deep Work
                            </button>
                        </div>
                        <div class="absolute bottom-[-60px] right-[-60px] w-64 h-64 bg-white/10 rounded-full blur-3xl transition-all group-hover:bg-white/20"></div>
                    </div>
                </div>
            </div>
        </div>
    </main>

    <!-- Right Sidebar: AI Study Buddy -->
    <aside class="w-[400px] bg-white border-l border-slate-100 flex flex-col z-40 max-2xl:hidden">
        <div class="p-8 pb-4 border-b border-slate-50">
            <div class="flex items-center justify-between mb-8">
                <h3 class="text-xl font-black text-[#1E1B4B]">AI Study Buddy</h3>
                <div class="w-10 h-10 bg-indigo-50 text-[#3D3ACE] rounded-2xl flex items-center justify-center">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4"/></svg>
                </div>
            </div>

            <!-- AI Insight Card -->
            <div class="bg-slate-50 p-6 rounded-[32px] border border-slate-100 mb-8 relative overflow-hidden group">
                <div class="relative z-10">
                    <p class="text-[10px] font-black text-[#3D3ACE] uppercase tracking-widest mb-3">Daily Insight</p>
                    <p class="text-sm font-bold text-[#1E1B4B] leading-relaxed">
                        {{ latestAIInsight }}
                    </p>
                </div>
                <div class="absolute top-0 right-0 p-2 opacity-5">
                    <svg class="w-20 h-20" fill="currentColor" viewBox="0 0 24 24"><path d="M13 10V3L4 14h7v7l9-11h-7z"/></svg>
                </div>
            </div>
        </div>

        <!-- Chat History -->
        <div class="flex-1 overflow-y-auto p-8 space-y-8 scrollbar-hide">
            <div v-if="dashboardData.ai_messages.length === 0" class="text-center py-10">
                <p class="text-slate-400 text-xs font-bold">No recent messages. Start a conversation!</p>
            </div>
            <div v-for="(chat, i) in dashboardData.ai_messages" :key="i" class="flex flex-col gap-3">
                <div class="flex items-center gap-3">
                    <div class="w-8 h-8 rounded-xl bg-indigo-50 text-[#3D3ACE] flex items-center justify-center text-[10px] font-black">AI</div>
                    <span class="text-[10px] font-black text-slate-400 uppercase tracking-widest">Lumora • {{ chat.timestamp }}</span>
                </div>
                <div class="bg-indigo-50 p-6 rounded-[28px] rounded-tl-none border border-indigo-100/30">
                    <p class="text-xs font-bold text-slate-700 leading-relaxed">{{ chat.content }}</p>
                </div>
            </div>
        </div>

        <!-- Chat Input -->
        <div class="p-8 pt-0">
            <div class="relative group">
                <input type="text" placeholder="Ask your buddy anything..." 
                    class="w-full bg-slate-50 border border-slate-100 rounded-[28px] py-6 pl-8 pr-16 text-sm font-bold text-[#1E1B4B] outline-none focus:bg-white focus:border-[#3D3ACE]/20 transition-all shadow-sm">
                <button class="absolute right-3 top-3 w-12 h-12 bg-[#3D3ACE] text-white rounded-2xl flex items-center justify-center shadow-lg shadow-indigo-100 hover:scale-105 active:scale-95 transition-all">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M13 5l7 7m0 0l-7 7m7-7H3" /></svg>
                </button>
            </div>
        </div>
    </aside>

    <!-- Search Overlay -->
    <Transition name="fade">
        <div v-if="searchOpen" class="fixed inset-0 z-[100] flex items-start justify-center pt-32 px-6">
            <div class="absolute inset-0 bg-[#1E1B4B]/80 backdrop-blur-md" @click="searchOpen = false"></div>
            <div class="relative w-full max-w-2xl bg-white rounded-[40px] shadow-2xl border border-slate-100 overflow-hidden">
                <div class="p-8 border-b border-slate-100 relative">
                    <svg class="absolute left-10 top-1/2 -translate-y-1/2 w-6 h-6 text-[#3D3ACE]" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/></svg>
                    <input type="text" autofocus placeholder="Type to search..." 
                        class="w-full bg-transparent pl-14 pr-10 py-2 text-xl font-black text-[#1E1B4B] outline-none placeholder-slate-300">
                </div>
                <div class="p-8 space-y-8">
                    <p class="text-[11px] font-black text-slate-400 uppercase tracking-widest">Recent Searches</p>
                    <div class="space-y-3">
                        <div v-for="s in ['Deep work', 'Consistency trends']" :key="s" class="p-4 rounded-2xl hover:bg-slate-50 cursor-pointer font-bold text-sm text-slate-600">{{ s }}</div>
                    </div>
                </div>
            </div>
        </div>
    </Transition>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { router } from '@inertiajs/vue3'

const props = defineProps({
    dashboard: {
        type: Object,
        default: () => ({
            user: { name: 'User', status: 'online' },
            metrics: { focus_sessions: 0, task_efficiency: 0, deep_work_hours: '0.0', consistency: 0, retention: 0 },
            focus_sessions: [],
            ai_messages: [],
            profile_title: 'Focused Achiever'
        })
    }
})

const activeNav = ref('Dashboard')
const searchOpen = ref(false)
const notificationsOpen = ref(false)

const dashboardData = computed(() => props.dashboard)

const profileTitle = computed(() => dashboardData.value.profile_title || 'Focused Achiever')
const focusSessionsCount = computed(() => dashboardData.value.metrics.focus_sessions)
const latestAIInsight = computed(() => {
    if (dashboardData.value.ai_messages.length > 0) {
        return dashboardData.value.ai_messages[0].content
    }
    return "Complete your first deep work session to unlock personalized AI insights!"
})

const metrics = computed(() => [
    { label: 'Focus Sessions', value: dashboardData.value.metrics.focus_sessions, icon: '🎯', bg: 'bg-indigo-50 text-indigo-600' },
    { label: 'Task Efficiency', value: dashboardData.value.metrics.task_efficiency + '%', icon: '⚡', bg: 'bg-emerald-50 text-emerald-600' },
    { label: 'Deep Work Hours', value: dashboardData.value.metrics.deep_work_hours + 'h', icon: '🔋', bg: 'bg-violet-50 text-violet-600' },
    { label: 'Consistency', value: dashboardData.value.metrics.consistency + '%', icon: '📈', bg: 'bg-amber-50 text-amber-600' },
    { label: 'Retention', value: dashboardData.value.metrics.retention + '%', icon: '🧠', bg: 'bg-rose-50 text-rose-600' },
])

const weeklyStats = [
    { day: 'MON', value: 40 },
    { day: 'TUE', value: 65 },
    { day: 'WED', value: 30 },
    { day: 'THU', value: 85 },
    { day: 'FRI', value: 45 },
    { day: 'SAT', value: 90 },
    { day: 'SUN', value: 60 },
]

const navItems = [
    { label: 'Dashboard', iconPath: 'M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10' },
    { label: 'Focus Sessions', iconPath: 'M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z' },
    { label: 'Analytics', iconPath: 'M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z' },
    { label: 'Settings', iconPath: 'M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z M15 12a3 3 0 11-6 0 3 3 0 016 0z' },
]

onMounted(() => {
    // Initial fetch if needed, though props should already be there
    console.log('Dashboard mounted with data:', props.dashboard)
})
</script>

<style scoped>
.scrollbar-hide::-webkit-scrollbar { display: none; }
.scrollbar-hide { -ms-overflow-style: none; scrollbar-width: none; }
.fade-enter-active, .fade-leave-active { transition: opacity 0.3s ease; }
.fade-enter-from, .fade-leave-to { opacity: 0; }
</style>
