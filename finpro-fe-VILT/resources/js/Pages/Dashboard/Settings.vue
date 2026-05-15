<template>
  <DashboardLayout :showBuddy="false">
    <!-- Page Header -->
    <div class="mb-10">
      <h1 class="text-3xl font-black text-[#1E1B4B] mb-2">Settings</h1>
      <p class="text-slate-500 font-medium">Manage your account, preferences, and privacy.</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
      <!-- Settings Nav -->
      <div class="lg:col-span-3 space-y-2">
        <button v-for="tab in tabs" :key="tab.id" @click="activeTab = tab.id"
          class="w-full text-left px-5 py-4 rounded-2xl text-sm font-bold transition-all"
          :class="activeTab === tab.id ? 'bg-white text-[#3D3ACE] shadow-sm border border-slate-100' : 'text-slate-500 hover:bg-slate-50'">
          {{ tab.label }}
        </button>
      </div>

      <!-- Settings Content -->
      <div class="lg:col-span-9 bg-white p-8 md:p-12 rounded-[40px] border border-slate-100 shadow-sm min-h-[500px]">
        
        <!-- Profile Tab -->
        <div v-if="activeTab === 'profile'" class="max-w-2xl">
          <h2 class="text-2xl font-black text-[#1E1B4B] mb-8">Profile Settings</h2>
          
          <div class="flex items-center gap-6 mb-10">
            <div class="w-24 h-24 rounded-[32px] overflow-hidden shadow-md">
                <img src="https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?auto=format&fit=crop&q=80&w=200&h=200" alt="Profile" class="w-full h-full object-cover">
            </div>
            <div>
              <button class="bg-slate-50 hover:bg-slate-100 text-[#1E1B4B] px-5 py-2.5 rounded-xl text-sm font-bold border border-slate-200 transition-colors mb-2">Change Avatar</button>
              <p class="text-xs text-slate-400 font-medium">JPG, GIF or PNG. Max size of 800K</p>
            </div>
          </div>

          <form @submit.prevent="saveSettings" class="space-y-6">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-xs font-black text-slate-400 uppercase tracking-widest mb-2">Full Name</label>
                <input type="text" v-model="form.name" class="w-full bg-slate-50 border border-slate-100 rounded-xl px-4 py-3 text-sm font-bold focus:bg-white focus:border-[#3D3ACE] transition-colors outline-none">
              </div>
              <div>
                <label class="block text-xs font-black text-slate-400 uppercase tracking-widest mb-2">Email Address</label>
                <input type="email" v-model="form.email" disabled class="w-full bg-slate-100 border border-slate-100 rounded-xl px-4 py-3 text-sm font-bold text-slate-400 cursor-not-allowed outline-none">
              </div>
            </div>
            
            <div class="pt-6">
              <button type="submit" class="bg-[#3D3ACE] hover:bg-[#322fb0] text-white px-8 py-3.5 rounded-2xl font-bold shadow-xl shadow-indigo-200 transition-all active:scale-95 flex items-center gap-2">
                <svg v-if="isSaving" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path></svg>
                Save Changes
              </button>
            </div>
          </form>
        </div>

        <!-- Preferences Tab -->
        <div v-if="activeTab === 'preferences'" class="max-w-2xl">
          <h2 class="text-2xl font-black text-[#1E1B4B] mb-8">Learning Preferences</h2>
          <div class="space-y-8">
            <div>
              <label class="block text-xs font-black text-slate-400 uppercase tracking-widest mb-4">Preferred Study Duration</label>
              <div class="flex flex-wrap gap-3">
                <button v-for="d in [25, 45, 60, 90]" :key="d" 
                  @click="form.duration = d"
                  class="px-6 py-3 rounded-xl text-sm font-bold transition-all border"
                  :class="form.duration === d ? 'bg-indigo-50 border-indigo-200 text-[#3D3ACE]' : 'bg-white border-slate-200 text-slate-500 hover:border-slate-300'">
                  {{ d }} Mins
                </button>
              </div>
            </div>
            <div>
              <label class="block text-xs font-black text-slate-400 uppercase tracking-widest mb-4">Default Focus Dimension</label>
              <select v-model="form.dimension" class="w-full bg-slate-50 border border-slate-100 rounded-xl px-4 py-3 text-sm font-bold focus:bg-white focus:border-[#3D3ACE] transition-colors outline-none">
                <option value="planning">Planning & Structuring</option>
                <option value="time">Time Management</option>
                <option value="cognitive">Cognitive Strategy</option>
                <option value="reflection">Reflection & Evaluation</option>
              </select>
            </div>
          </div>
        </div>

        <!-- Notifications Tab -->
        <div v-if="activeTab === 'notifications'" class="max-w-2xl">
          <h2 class="text-2xl font-black text-[#1E1B4B] mb-8">Notification Settings</h2>
          <div class="space-y-6">
            <div v-for="(notif, key) in form.notifications" :key="key" class="flex items-center justify-between p-4 bg-slate-50 rounded-2xl border border-slate-100">
              <div>
                <p class="text-sm font-bold text-[#1E1B4B]">{{ notif.label }}</p>
                <p class="text-[11px] font-medium text-slate-500 mt-0.5">{{ notif.desc }}</p>
              </div>
              <button @click="notif.active = !notif.active" 
                class="w-12 h-6 rounded-full transition-colors relative"
                :class="notif.active ? 'bg-emerald-500' : 'bg-slate-200'">
                <div class="absolute top-1 left-1 bg-white w-4 h-4 rounded-full transition-transform shadow-sm"
                  :class="notif.active ? 'translate-x-6' : 'translate-x-0'"></div>
              </button>
            </div>
          </div>
        </div>

        <!-- Privacy Tab -->
        <div v-if="activeTab === 'privacy'" class="max-w-2xl">
          <h2 class="text-2xl font-black text-[#1E1B4B] mb-8">Privacy & Data</h2>
          <div class="bg-indigo-50 border border-indigo-100 p-6 rounded-2xl mb-8">
            <h4 class="text-sm font-bold text-indigo-900 mb-2">How Lumora uses your data</h4>
            <p class="text-xs text-indigo-700/80 leading-relaxed">
              Lumora strictly uses your learning data (assessments, plans, notes) to generate personalized study recommendations and AI insights. Your data is not shared with third-party advertisers.
            </p>
          </div>
          <div class="space-y-4 border-t border-slate-100 pt-8">
            <button class="w-full text-left px-5 py-4 rounded-2xl text-sm font-bold text-[#1E1B4B] bg-slate-50 hover:bg-slate-100 transition-colors flex justify-between items-center">
              Export my data as JSON
              <svg class="w-4 h-4 text-slate-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"/></svg>
            </button>
            <button class="w-full text-left px-5 py-4 rounded-2xl text-sm font-bold text-rose-600 bg-rose-50 hover:bg-rose-100 transition-colors flex justify-between items-center">
              Delete account and all data
            </button>
          </div>
        </div>

      </div>
    </div>
  </DashboardLayout>
</template>

<script setup>
import { ref } from 'vue'
import { usePage } from '@inertiajs/vue3'
import DashboardLayout from '@/Layouts/DashboardLayout.vue'

const page = usePage()
const user = page.props.auth?.user || { name: 'User', email: 'user@example.com' }

const activeTab = ref('profile')
const isSaving = ref(false)

const tabs = [
  { id: 'profile', label: 'Profile Settings' },
  { id: 'preferences', label: 'Learning Preferences' },
  { id: 'notifications', label: 'Notifications' },
  { id: 'privacy', label: 'Privacy & Data' }
]

const form = ref({
  name: user.name,
  email: user.email,
  duration: 45,
  dimension: 'planning',
  notifications: {
    study: { label: 'Study Reminders', desc: 'Get notified when a planned session is about to start.', active: true },
    target: { label: 'Target Alerts', desc: 'Weekly target completion reminders.', active: true },
    reflection: { label: 'Reflection Prompts', desc: 'Reminders to write notes after sessions.', active: false },
    ai: { label: 'AI Insights', desc: 'Notify when new personalized strategy is ready.', active: true }
  }
})

function saveSettings() {
  isSaving.value = true
  setTimeout(() => {
    isSaving.value = false
    // Show success toast here
  }, 1000)
}
</script>
