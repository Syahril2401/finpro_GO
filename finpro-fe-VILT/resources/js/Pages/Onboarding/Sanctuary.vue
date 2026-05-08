<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-50 to-indigo-50 flex items-center justify-center px-4 py-16">
    <!-- Background decoration -->
    <div class="absolute top-0 right-0 w-72 h-72 bg-indigo-100 rounded-full blur-3xl opacity-60 pointer-events-none"></div>
    <div class="absolute bottom-0 left-0 w-64 h-64 bg-violet-100 rounded-full blur-3xl opacity-50 pointer-events-none"></div>

    <div class="relative z-10 w-full max-w-lg">
      <!-- Logo -->
      <div class="text-center mb-10">
        <Link :href="route('landing')">
          <span class="text-2xl font-extrabold text-indigo-600 tracking-tight">Lumora</span>
        </Link>
      </div>

      <!-- Card -->
      <div class="bg-white rounded-3xl shadow-2xl shadow-indigo-100 border border-slate-100 overflow-hidden">
        <!-- Header gradient -->
        <div class="bg-gradient-to-r from-indigo-600 to-violet-600 px-8 py-10 text-center">
          <div class="w-16 h-16 bg-white/20 backdrop-blur rounded-2xl flex items-center justify-center mx-auto mb-5 text-4xl shadow-lg">
            🌿
          </div>
          <h1 class="text-2xl font-extrabold text-white leading-snug">
            Welcome to your<br>
            <span class="text-yellow-300">Sanctuary of Focus.</span>
          </h1>
          <p class="text-indigo-200 mt-3 text-sm leading-relaxed max-w-sm mx-auto">
            Sebelum membangun rencana belajarmu, kami perlu mengenal cara belajarmu.
            Ikuti assessment singkat 3-5 menit untuk menemukan profil belajarmu.
          </p>
        </div>

        <!-- Steps -->
        <div class="px-8 py-7">
          <p class="text-xs font-semibold text-slate-400 uppercase tracking-widest mb-5 text-center">Yang akan kamu lakukan</p>
          <div class="space-y-4">
            <div v-for="(step, i) in steps" :key="i"
              class="flex items-start gap-4">
              <div class="w-9 h-9 rounded-xl flex items-center justify-center flex-shrink-0 text-lg"
                :class="step.bg">
                {{ step.icon }}
              </div>
              <div>
                <p class="font-semibold text-slate-800 text-sm">{{ step.title }}</p>
                <p class="text-slate-500 text-xs mt-0.5">{{ step.desc }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- CTA -->
        <div class="px-8 pb-8 flex flex-col gap-3">
          <Link :href="route('onboarding.questionnaire')"
            id="sanctuary-start-btn"
            class="block text-center bg-indigo-600 text-white font-bold py-3.5 rounded-2xl hover:bg-indigo-700 active:scale-[0.98] transition-all shadow-lg shadow-indigo-200">
            Mulai Assessment →
          </Link>
        </div>
      </div>

      <!-- Trust note -->
      <p class="text-center text-xs text-slate-400 mt-5">
        🔒 Jawabanmu hanya digunakan untuk personalisasi, tidak disebarkan.
      </p>
    </div>
  </div>
</template>

<script setup>
import { Link } from '@inertiajs/vue3'

const steps = [
  {
    icon: '📋',
    bg: 'bg-indigo-50',
    title: 'Isi 20 pertanyaan singkat',
    desc: 'Dibagi dalam 4 segmen: Planning, Time Management, Cognitive, & Reflection.',
  },
  {
    icon: '🧠',
    bg: 'bg-violet-50',
    title: 'Dapatkan profil belajarmu',
    desc: 'AI kami menganalisis jawabanmu untuk membuat profil yang unik untukmu.',
  },
  {
    icon: '🚀',
    bg: 'bg-emerald-50',
    title: 'Mulai belajar lebih efektif',
    desc: 'Gunakan rekomendasi strategi yang sudah dipersonalisasi.',
  },
]
</script>
