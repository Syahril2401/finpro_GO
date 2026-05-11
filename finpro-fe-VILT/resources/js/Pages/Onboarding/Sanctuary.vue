<template>
  <div class="min-h-screen bg-[#F0F2F9] flex items-center justify-center px-4 py-16 relative overflow-hidden font-sans">
    <!-- Background decorations -->
    <div class="absolute top-[-10%] right-[-5%] w-[500px] h-[500px] bg-indigo-200/20 rounded-full blur-[120px]"></div>
    <div class="absolute bottom-[-10%] left-[-5%] w-[400px] h-[400px] bg-purple-200/20 rounded-full blur-[120px]"></div>

    <!-- Card Container -->
    <div class="relative z-10 w-full max-w-[1000px] bg-white rounded-[40px] shadow-2xl shadow-indigo-100 flex flex-col md:flex-row overflow-hidden border border-white/50">
      
      <!-- Left Section: Visuals & Quote -->
      <div class="w-full md:w-[45%] bg-[#F9FAFB] p-12 flex flex-col justify-between relative overflow-hidden">
        <div class="relative z-10">
          <div class="text-2xl font-black text-[#1E1B4B] mb-8">Lumora</div>
          <div class="w-12 h-1 bg-[#3D3ACE] rounded-full mb-2"></div>
          <div class="w-8 h-1 bg-slate-200 rounded-full"></div>
        </div>

        <div class="relative z-10 flex flex-col items-center">
            <div class="w-full max-w-[320px] aspect-square rounded-[40px] overflow-hidden shadow-2xl shadow-indigo-200 border-4 border-white mb-12 rotate-[-2deg]">
                <img src="/image/welcome image.png" alt="Welcome to Lumora" class="w-full h-full object-cover">
            </div>
            <p class="text-slate-500 font-medium italic text-center text-sm leading-relaxed max-w-[280px]">
                "The capacity to learn is a gift; the ability to learn is a skill; the willingness to learn is a choice."
            </p>
        </div>
        
        <!-- Subtle pattern overlay -->
        <div class="absolute inset-0 opacity-[0.03] pointer-events-none">
            <svg class="w-full h-full" viewBox="0 0 100 100" preserveAspectRatio="none">
                <defs><pattern id="grid" width="10" height="10" patternUnits="userSpaceOnUse"><path d="M 10 0 L 0 0 0 10" fill="none" stroke="currentColor" stroke-width="0.5"/></pattern></defs>
                <rect width="100" height="100" fill="url(#grid)" />
            </svg>
        </div>
      </div>

      <!-- Right Section: Welcome Content -->
      <div class="w-full md:w-[55%] p-12 md:p-16 flex flex-col justify-center">
        <div class="mb-10">
            <p class="text-[11px] font-black text-[#3D3ACE] uppercase tracking-[0.2em] mb-4">Step 01 / 03</p>
            <h1 class="text-[44px] font-extrabold text-[#1E1B4B] leading-[1.1] mb-6">
                Welcome to your Sanctuary of Focus.
            </h1>
            <p class="text-slate-500 text-lg leading-relaxed font-medium">
                Lumora is designed to help you transform your academic journey through AI-driven strategies and intentional habit building.
            </p>
        </div>

        <!-- Features -->
        <div class="space-y-8 mb-12">
            <div class="flex items-start gap-6">
                <div class="w-14 h-14 bg-[#D1FAE5] text-[#059669] rounded-2xl flex items-center justify-center flex-shrink-0 shadow-sm shadow-emerald-100">
                    <svg class="w-7 h-7" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0012 18.75c-1.03 0-1.9-.4-2.593-1.012l-.547-.547z" />
                    </svg>
                </div>
                <div>
                    <h3 class="text-lg font-bold text-[#1E1B4B] mb-1">Intelligent Assistance</h3>
                    <p class="text-slate-500 font-medium">Personalized AI models that adapt to your unique study patterns.</p>
                </div>
            </div>

            <div class="flex items-start gap-6">
                <div class="w-14 h-14 bg-[#EEF2FF] text-[#3D3ACE] rounded-2xl flex items-center justify-center flex-shrink-0 shadow-sm shadow-indigo-100">
                    <svg class="w-7 h-7" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
                    </svg>
                </div>
                <div>
                    <h3 class="text-lg font-bold text-[#1E1B4B] mb-1">Focused Sessions</h3>
                    <p class="text-slate-500 font-medium">Minimize distractions with curated productivity environments.</p>
                </div>
            </div>
        </div>

        <div class="flex items-center justify-between">
            <Link :href="route('onboarding.questionnaire')" 
                class="bg-[#3D3ACE] hover:bg-[#322fb0] text-white font-bold px-10 py-5 rounded-[20px] flex items-center gap-3 shadow-xl shadow-indigo-100 transition-all active:scale-[0.98]">
                Let's Begin
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3" />
                </svg>
            </Link>
        </div>
      </div>
    </div>

    <!-- Simple Footer -->
    <div class="absolute bottom-10 inset-x-0 text-center">
        <p class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] mb-4">Secured by Lumora Intelligence</p>
        <div class="flex justify-center gap-8 text-[13px] font-semibold text-slate-500">
            <a href="#" class="hover:text-slate-900 transition-colors">Privacy Policy</a>
            <a href="#" class="hover:text-slate-900 transition-colors">Terms of Use</a>
            <a href="#" class="hover:text-slate-900 transition-colors">Contact Support</a>
        </div>
    </div>
  </div>
</template>

<script setup>
import { Link } from '@inertiajs/vue3'
</script>

