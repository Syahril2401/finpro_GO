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
        <Link v-for="item in navItems" :key="item.label"
          :href="item.href"
          class="w-full flex items-center gap-4 px-5 py-4 rounded-2xl text-[14px] font-bold transition-all group"
          :class="$page.url === item.path || $page.url.startsWith(item.path + '/') && item.path !== '/dashboard'
            ? 'bg-[#3D3ACE] text-white shadow-xl shadow-indigo-100'
            : 'text-slate-400 hover:bg-slate-50 hover:text-slate-600'">
          <svg class="w-5 h-5 transition-transform group-hover:scale-110" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" :d="item.iconPath" />
          </svg>
          {{ item.label }}
          <div v-if="item.badge" class="ml-auto bg-red-500 text-white text-[10px] px-2 py-0.5 rounded-full font-black">
              {{ item.badge }}
          </div>
        </Link>
      </nav>

      <!-- Bottom User Section -->
      <div class="p-6 border-t border-slate-50">
        <div class="bg-slate-50 p-4 rounded-3xl flex items-center gap-4 group cursor-pointer hover:bg-slate-100 transition-all">
            <div class="w-12 h-12 rounded-2xl overflow-hidden shadow-md">
                <img src="https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?auto=format&fit=crop&q=80&w=200&h=200" alt="Profile" class="w-full h-full object-cover">
            </div>
            <div class="flex-1 min-w-0">
                <p class="text-sm font-black text-[#1E1B4B] truncate">{{ user.name }}</p>
                <p class="text-[11px] font-bold text-slate-400 uppercase tracking-wider truncate">User</p>
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
                <div class="group relative flex items-center cursor-pointer">
                    <svg class="absolute left-5 w-5 h-5 text-slate-400 group-hover:text-[#3D3ACE] transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/></svg>
                    <div class="w-full bg-slate-50 border border-slate-100 rounded-2xl py-3.5 pl-14 pr-5 text-sm font-bold text-slate-400 group-hover:bg-white group-hover:border-[#3D3ACE]/20 transition-all flex items-center">
                        Search anything...
                        <span class="ml-auto text-[10px] bg-white px-2 py-1 rounded-lg border border-slate-100 shadow-sm font-black text-slate-400 tracking-tighter">⌘ K</span>
                    </div>
                </div>
            </div>

            <div class="flex items-center gap-6 ml-10">
                <button class="relative w-12 h-12 bg-white border border-slate-100 rounded-2xl flex items-center justify-center text-slate-400 hover:text-[#3D3ACE] hover:border-indigo-100 hover:shadow-lg hover:shadow-indigo-50 transition-all">
                    <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"/></svg>
                    <div class="absolute top-3 right-3 w-2.5 h-2.5 bg-red-500 border-2 border-white rounded-full"></div>
                </button>
                <div class="h-8 w-[1px] bg-slate-100"></div>
                <div class="flex items-center gap-4">
                    <div class="text-right max-md:hidden">
                        <p class="text-sm font-black text-[#1E1B4B]">{{ user.name }}</p>
                        <p class="text-[10px] font-bold text-[#059669] uppercase tracking-widest flex items-center justify-end gap-1.5">
                            <span class="w-1.5 h-1.5 bg-emerald-500 rounded-full animate-pulse"></span>
                            Online
                        </p>
                    </div>
                </div>
            </div>
        </header>

        <!-- Page Content -->
        <div class="flex-1 overflow-y-auto bg-[#FDFDFF] p-10 scrollbar-hide">
            <slot />
        </div>
    </main>

    <!-- Right Sidebar: AI Study Buddy -->
    <aside v-if="showBuddy" class="w-[400px] bg-white border-l border-slate-100 flex flex-col z-40 max-2xl:hidden">
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
                    <p class="text-[10px] font-black text-[#3D3ACE] uppercase tracking-widest mb-3">{{ buddyContextTitle }}</p>
                    <p class="text-sm font-bold text-[#1E1B4B] leading-relaxed">
                        {{ buddyContextMessage }}
                    </p>
                </div>
                <div class="absolute top-0 right-0 p-2 opacity-5">
                    <svg class="w-20 h-20" fill="currentColor" viewBox="0 0 24 24"><path d="M13 10V3L4 14h7v7l9-11h-7z"/></svg>
                </div>
            </div>
        </div>

        <!-- Chat History -->
        <div class="flex-1 overflow-y-auto p-8 space-y-8 scrollbar-hide">
            <div class="text-center py-10">
                <p class="text-slate-400 text-xs font-bold">No recent messages. Start a conversation!</p>
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
  </div>
</template>

<script setup>
import { Link, usePage } from '@inertiajs/vue3'
import { computed } from 'vue'

const props = defineProps({
    showBuddy: { type: Boolean, default: true },
    buddyContextTitle: { type: String, default: 'Daily Insight' },
    buddyContextMessage: { type: String, default: 'Focus on completing your most important tasks today to stay ahead.' },
})

const page = usePage()
const user = computed(() => page.props.auth?.user || { name: 'User' })

const navItems = [
    { label: 'Dashboard', path: '/dashboard', href: route('dashboard'), iconPath: 'M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z' },
    { label: 'Planner', path: '/dashboard/planner', href: route('planner'), iconPath: 'M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z' },
    { label: 'Weekly Targets', path: '/dashboard/weekly-targets', href: route('targets'), iconPath: 'M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z' },
    { label: 'Notes', path: '/dashboard/notes', href: route('notes'), iconPath: 'M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z' },
    { label: 'Progress', path: '/dashboard/progress', href: route('progress'), iconPath: 'M13 7h8m0 0v8m0-8l-8 8-4-4-6 6' },
    { label: 'Settings', path: '/dashboard/settings', href: route('settings'), iconPath: 'M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z' },
]
</script>
