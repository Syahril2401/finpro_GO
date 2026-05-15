<template>
  <div class="min-h-screen bg-[#FDFDFF] font-sans text-slate-900 flex flex-col items-center justify-center p-6">
    <div class="max-w-2xl text-center">
      <div class="w-20 h-20 bg-emerald-50 text-emerald-600 rounded-full flex items-center justify-center mx-auto mb-6 text-3xl shadow-lg">
        💡
      </div>
      <h1 class="text-4xl font-black text-[#1E1B4B] mb-4">Recommendation Detail</h1>
      <p class="text-slate-500 text-lg mb-10 leading-relaxed">
        Deep dive into specific action plans tailored to your study profile. This detailed view is coming in the next sprint.
      </p>

      <div class="bg-white border border-slate-100 shadow-xl shadow-slate-100/50 p-8 rounded-[32px] mb-8">
         <h3 class="text-xl font-bold text-slate-400">Recommendation ID: {{ id }}</h3>
      </div>

      <Link :href="route('dashboard')" class="inline-flex items-center gap-2 bg-[#3D3ACE] hover:bg-[#322fb0] text-white font-bold py-4 px-8 rounded-xl shadow-lg shadow-indigo-200 transition-all active:scale-[0.98]">
        ← Back to Dashboard
      </Link>
    </div>
  </div>
</template>

<script setup>
import { Link } from '@inertiajs/vue3'

defineProps({
    id: String
})
</script>
