<template>
  <div class="min-h-screen bg-white font-sans">
    <!-- ─── Navbar ─────────────────────────────────────────────────────── -->
    <nav class="fixed top-0 inset-x-0 z-50 bg-white/80 backdrop-blur-md border-b border-slate-100">
      <div class="max-w-7xl mx-auto px-6 py-4 flex items-center justify-between">
        <span class="text-xl font-bold text-slate-900 tracking-tight">
          Lumora
        </span>
        <ul class="hidden md:flex gap-8 text-sm font-medium text-slate-600">
          <li><a href="#features" class="hover:text-indigo-600 transition-colors">Features</a></li>
          <li><a href="#how" class="hover:text-indigo-600 transition-colors">How it works</a></li>
          <li><a href="#testimonials" class="hover:text-indigo-600 transition-colors">Testimonials</a></li>
        </ul>
        <div class="flex items-center gap-3">
          <Link :href="route('login')" class="text-sm font-medium text-slate-700 hover:text-indigo-600 transition-colors">
            Log in
          </Link>
          <Link :href="route('register')"
            class="bg-indigo-600 text-white text-sm font-semibold px-5 py-2.5 rounded-full hover:bg-indigo-700 transition-all shadow-md shadow-indigo-200">
            Get started
          </Link>
        </div>
      </div>
    </nav>

    <!-- ─── Hero ──────────────────────────────────────────────────────── -->
    <section class="relative pt-32 pb-24 overflow-hidden">
      <!-- Gradient blobs -->
      <div class="absolute -top-32 -left-32 w-[600px] h-[600px] bg-indigo-100 rounded-full blur-3xl opacity-50"></div>
      <div class="absolute -bottom-20 right-0 w-[400px] h-[400px] bg-violet-100 rounded-full blur-3xl opacity-40"></div>

      <div class="relative max-w-7xl mx-auto px-6 grid md:grid-cols-2 gap-12 items-center">
        <!-- Left copy -->
        <div>
          <div class="inline-flex items-center gap-2 bg-indigo-50 text-indigo-700 text-xs font-semibold px-3 py-1.5 rounded-full mb-6">
            <span class="w-1.5 h-1.5 bg-indigo-500 rounded-full animate-pulse"></span>
            Self-Regulated Learning Platform
          </div>
          <h1 class="text-5xl md:text-6xl font-extrabold text-slate-900 leading-tight mb-4">
            Build better<br>study habits<br>with
            <span class="text-indigo-600"> clarity.</span>
          </h1>
          <p class="text-slate-500 text-lg leading-relaxed max-w-md mb-8">
            Lumora helps you understand how you learn, then gives you AI-powered tools to study smarter every day.
          </p>
          <div class="flex flex-wrap gap-4">
            <Link :href="route('register')"
              class="bg-indigo-600 text-white font-semibold px-8 py-3 rounded-full hover:bg-indigo-700 transition-all shadow-lg shadow-indigo-200">
              Get started free
            </Link>
            <a href="#how"
              class="flex items-center gap-2 text-slate-700 font-medium px-6 py-3 rounded-full border border-slate-200 hover:border-indigo-300 hover:text-indigo-600 transition-all">
              See how it works
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
              </svg>
            </a>
          </div>
        </div>

        <!-- Right visual -->
        <div class="relative">
          <div class="bg-gradient-to-br from-indigo-600 to-violet-600 rounded-3xl p-1 shadow-2xl shadow-indigo-200">
            <div class="bg-slate-900 rounded-[22px] p-6">
              <!-- Mini dashboard preview -->
              <div class="flex items-center gap-2 mb-5">
                <div class="w-2.5 h-2.5 bg-red-400 rounded-full"></div>
                <div class="w-2.5 h-2.5 bg-yellow-400 rounded-full"></div>
                <div class="w-2.5 h-2.5 bg-green-400 rounded-full"></div>
                <div class="ml-auto text-xs text-slate-500 font-mono">lumora.app</div>
              </div>
              <div class="space-y-3">
                <div class="bg-slate-800 rounded-xl p-4">
                  <p class="text-slate-400 text-xs mb-1">Study streak</p>
                  <p class="text-white font-bold text-2xl">14 days 🔥</p>
                </div>
                <div class="grid grid-cols-2 gap-3">
                  <div class="bg-indigo-600/20 rounded-xl p-3 border border-indigo-600/30">
                    <p class="text-indigo-300 text-xs mb-1">Sessions today</p>
                    <p class="text-white font-bold text-lg">3 / 4</p>
                  </div>
                  <div class="bg-violet-600/20 rounded-xl p-3 border border-violet-600/30">
                    <p class="text-violet-300 text-xs mb-1">Focus score</p>
                    <p class="text-white font-bold text-lg">87%</p>
                  </div>
                </div>
                <div class="bg-slate-800 rounded-xl p-4">
                  <p class="text-slate-400 text-xs mb-2">AI Strategy</p>
                  <p class="text-white text-sm font-medium">🧠 Pomodoro + Spaced Repetition</p>
                  <p class="text-slate-500 text-xs mt-1">Based on your Focused Achiever profile</p>
                </div>
              </div>
            </div>
          </div>
          <!-- Badge floating -->
          <div class="absolute -bottom-4 -left-4 bg-white rounded-2xl shadow-xl px-4 py-3 border border-slate-100">
            <p class="text-xs text-slate-500">Used by students at</p>
            <p class="text-sm font-bold text-slate-800">500+ universities</p>
          </div>
        </div>
      </div>
    </section>

    <!-- ─── The Old Way vs New Way ─────────────────────────────────────── -->
    <section class="py-20 bg-slate-50">
      <div class="max-w-7xl mx-auto px-6">
        <p class="text-center text-slate-400 text-sm font-medium mb-2">The end of academic chaos.</p>
        <h2 class="text-center text-3xl font-extrabold text-slate-900 mb-12">
          Study smarter, not harder.
        </h2>
        <div class="grid md:grid-cols-2 gap-6 max-w-4xl mx-auto">
          <div class="bg-white border border-red-100 rounded-2xl p-6">
            <p class="text-red-500 font-bold text-sm mb-4">❌ The Old Way</p>
            <ul class="space-y-3 text-slate-600 text-sm">
              <li class="flex gap-2 items-start"><span class="text-red-300 mt-0.5">•</span> Study randomly without a plan</li>
              <li class="flex gap-2 items-start"><span class="text-red-300 mt-0.5">•</span> Waste hours on ineffective methods</li>
              <li class="flex gap-2 items-start"><span class="text-red-300 mt-0.5">•</span> No insight into your learning style</li>
              <li class="flex gap-2 items-start"><span class="text-red-300 mt-0.5">•</span> Burnout before exams</li>
            </ul>
          </div>
          <div class="bg-indigo-600 rounded-2xl p-6">
            <p class="text-indigo-200 font-bold text-sm mb-4">✅ The Lumora Way</p>
            <ul class="space-y-3 text-white text-sm">
              <li class="flex gap-2 items-start"><span class="text-indigo-300 mt-0.5">•</span> Personalized AI study strategy</li>
              <li class="flex gap-2 items-start"><span class="text-indigo-300 mt-0.5">•</span> Intelligent planner & habit tracker</li>
              <li class="flex gap-2 items-start"><span class="text-indigo-300 mt-0.5">•</span> Know your learning profile</li>
              <li class="flex gap-2 items-start"><span class="text-indigo-300 mt-0.5">•</span> Stay focused with deep work sessions</li>
            </ul>
          </div>
        </div>
      </div>
    </section>

    <!-- ─── Features ──────────────────────────────────────────────────── -->
    <section id="features" class="py-24">
      <div class="max-w-7xl mx-auto px-6">
        <p class="text-center text-indigo-600 text-sm font-semibold mb-2">AI-POWERED FEATURES</p>
        <h2 class="text-center text-3xl font-extrabold text-slate-900 mb-16">Everything you need to excel.</h2>
        <div class="grid md:grid-cols-3 gap-8">
          <div v-for="feature in features" :key="feature.title"
            class="group bg-white border border-slate-100 rounded-2xl p-6 hover:border-indigo-200 hover:shadow-lg hover:shadow-indigo-50 transition-all duration-300">
            <div class="w-12 h-12 rounded-xl flex items-center justify-center text-2xl mb-4"
              :class="feature.bg">
              {{ feature.icon }}
            </div>
            <h3 class="font-bold text-slate-900 mb-2">{{ feature.title }}</h3>
            <p class="text-slate-500 text-sm leading-relaxed">{{ feature.desc }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- ─── How It Works ──────────────────────────────────────────────── -->
    <section id="how" class="py-24 bg-slate-50">
      <div class="max-w-7xl mx-auto px-6">
        <p class="text-center text-indigo-600 text-sm font-semibold mb-2">MASTERY IN THREE STEPS</p>
        <h2 class="text-center text-3xl font-extrabold text-slate-900 mb-4">Be the master of your learning.</h2>
        <p class="text-center text-slate-500 mb-16 max-w-xl mx-auto">
          Lumora's process is simple but powerful — rooted in cognitive science and personalized to you.
        </p>
        <div class="grid md:grid-cols-3 gap-10">
          <div v-for="(step, idx) in steps" :key="step.title" class="text-center">
            <div class="w-12 h-12 bg-indigo-600 text-white font-black rounded-full flex items-center justify-center text-lg mx-auto mb-4">
              {{ idx + 1 }}
            </div>
            <h3 class="font-bold text-slate-900 text-lg mb-2">{{ step.title }}</h3>
            <p class="text-slate-500 text-sm leading-relaxed">{{ step.desc }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- ─── CTA ───────────────────────────────────────────────────────── -->
    <section class="py-24">
      <div class="max-w-7xl mx-auto px-6">
        <div class="bg-gradient-to-br from-indigo-600 to-violet-600 rounded-3xl px-8 py-16 text-center relative overflow-hidden">
          <div class="absolute inset-0 bg-[url('data:image/svg+xml,...')] opacity-5"></div>
          <h2 class="text-3xl md:text-4xl font-extrabold text-white mb-4">
            Ready to transform your study habits?
          </h2>
          <p class="text-indigo-200 mb-8 max-w-xl mx-auto">
            AI-powered · Personalized · Built for students who aim higher.
          </p>
          <Link :href="route('register')"
            class="inline-block bg-white text-indigo-700 font-bold px-10 py-4 rounded-full hover:bg-indigo-50 transition-all shadow-lg">
            Get started — it's free
          </Link>
          <p class="text-indigo-300 text-xs mt-4">No credit card required</p>
        </div>
      </div>
    </section>

    <!-- ─── Footer ────────────────────────────────────────────────────── -->
    <footer class="border-t border-slate-100 py-10">
      <div class="max-w-7xl mx-auto px-6 flex flex-col md:flex-row items-center justify-between gap-4">
        <span class="text-slate-900 font-bold">Lumora</span>
        <div class="flex gap-6 text-sm text-slate-500">
          <a href="#" class="hover:text-indigo-600 transition-colors">Privacy Policy</a>
          <a href="#" class="hover:text-indigo-600 transition-colors">Terms of Service</a>
          <a href="#" class="hover:text-indigo-600 transition-colors">Contact Us</a>
        </div>
        <p class="text-xs text-slate-400">© 2024 Lumora. All rights reserved.</p>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { Link } from '@inertiajs/vue3'

const features = [
  {
    icon: '📅',
    bg: 'bg-indigo-50',
    title: 'Intelligent Planner',
    desc: 'Smart scheduling that adapts to your learning rhythm and energy levels throughout the day.',
  },
  {
    icon: '🎯',
    bg: 'bg-violet-50',
    title: 'Deep Focus Mode',
    desc: 'Distraction-free study sessions with built-in Pomodoro timer and focus analytics.',
  },
  {
    icon: '🧠',
    bg: 'bg-pink-50',
    title: 'AI Strategies',
    desc: 'Personalized study strategies generated from your learning profile and past performance.',
  },
  {
    icon: '📊',
    bg: 'bg-emerald-50',
    title: 'Progress Tracker',
    desc: 'Visual insights into your study habits, streaks, and improvement over time.',
  },
  {
    icon: '📝',
    bg: 'bg-amber-50',
    title: 'Smart Notes',
    desc: 'Organize your study materials and notes in one place, linked to your planner.',
  },
  {
    icon: '🏆',
    bg: 'bg-blue-50',
    title: 'Learning Profile',
    desc: 'Discover your unique learner type through our science-backed assessment.',
  },
]

const steps = [
  {
    title: 'Assess',
    desc: 'Take our quick self-regulated learning assessment to reveal your unique study style and habits.',
  },
  {
    title: 'Plan',
    desc: 'Lumora builds a personalized study plan based on your profile, goals, and schedule.',
  },
  {
    title: 'Improve',
    desc: 'Track your progress, follow AI strategies, and see real improvement week over week.',
  },
]
</script>
