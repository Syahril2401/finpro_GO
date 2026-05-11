<template>
  <div class="min-h-screen bg-[#F9FAFB] font-sans text-slate-900">
    <!-- ─── Navbar ─────────────────────────────────────────────────────── -->
    <nav class="fixed top-0 inset-x-0 z-50 bg-white/70 backdrop-blur-xl border-b border-slate-100">
      <div class="max-w-7xl mx-auto px-6 h-20 flex items-center justify-between">
        <div class="flex items-center gap-12">
          <span class="text-2xl font-black text-[#3D3ACE] tracking-tight">
            Lumora
          </span>
          <ul class="hidden lg:flex gap-8 text-[15px] font-medium text-slate-500">
            <li><a href="#features" class="hover:text-slate-900 transition-colors">Features</a></li>
            <li><a href="#how" class="hover:text-slate-900 transition-colors">How it Works</a></li>
            <li><a href="#pricing" class="hover:text-slate-900 transition-colors">Pricing</a></li>
            <li><a href="#testimonials" class="hover:text-slate-900 transition-colors">Testimonials</a></li>
          </ul>
        </div>
        <div class="flex items-center gap-6">
          <Link :href="route('login')" class="text-[15px] font-semibold text-slate-600 hover:text-slate-900 transition-colors">
            Log In
          </Link>
          <Link :href="route('register')"
            class="bg-[#3D3ACE] text-white text-[15px] font-bold px-7 py-3 rounded-xl hover:bg-[#322fb0] transition-all shadow-lg shadow-indigo-100">
            Get Started
          </Link>
        </div>
      </div>
    </nav>

    <!-- ─── Hero ──────────────────────────────────────────────────────── -->
    <section class="relative pt-44 pb-32 bg-white">
      <div class="max-w-7xl mx-auto px-6 grid lg:grid-cols-2 gap-16 items-center">
        <!-- Left Side -->
        <div class="max-w-xl">
          <div class="inline-flex items-center bg-[#EEF2FF] text-[#3D3ACE] text-[11px] font-black uppercase tracking-widest px-4 py-2 rounded-lg mb-8">
            The Intelligent Sanctuary
          </div>
          <h1 class="text-[64px] font-extrabold text-slate-900 leading-[1.1] tracking-tight mb-8">
            Build better study habits with <span class="text-[#3D3ACE]">clarity.</span>
          </h1>
          <p class="text-slate-500 text-lg leading-relaxed mb-10">
            Lumora helps students plan, track, reflect, and improve their learning process in one integrated workspace. Designed for the focused mind.
          </p>
          <div class="flex flex-wrap gap-4">
            <Link :href="route('register')"
              class="bg-[#3D3ACE] text-white font-bold px-9 py-4 rounded-full hover:bg-[#322fb0] transition-all shadow-xl shadow-indigo-100">
              Get Started for Free
            </Link>
            <button class="flex items-center gap-3 bg-[#E5E7EB]/50 text-slate-700 font-bold px-7 py-4 rounded-full hover:bg-[#E5E7EB] transition-all">
              <div class="w-6 h-6 bg-white rounded-full flex items-center justify-center shadow-sm">
                <svg class="w-3 h-3 text-slate-600 fill-current ml-0.5" viewBox="0 0 24 24">
                  <path d="M8 5v14l11-7z"/>
                </svg>
              </div>
              See how it works
            </button>
          </div>
        </div>

        <!-- Right Side -->
        <div class="relative">
          <div class="relative rounded-[40px] overflow-hidden shadow-2xl shadow-indigo-100 border-8 border-white">
            <img src="/image/hero-image-square-ball.png" alt="Lumora AI Visualization" class="w-full h-auto">
          </div>
        </div>
      </div>
    </section>

    <!-- ─── Comparison Section ────────────────────────────────────────── -->
    <section class="py-32 bg-[#F9FAFB]">
      <div class="max-w-7xl mx-auto px-6">
        <div class="text-center mb-20">
          <h2 class="text-[40px] font-extrabold text-slate-900 mb-4 tracking-tight">The end of academic chaos.</h2>
          <p class="text-slate-500 text-lg max-w-2xl mx-auto">
            Traditional productivity tools are built for work, not learning. Lumora is designed for the specific cognitive demands of students.
          </p>
        </div>

        <div class="grid md:grid-cols-2 gap-8 max-w-5xl mx-auto">
          <!-- The Old Way -->
          <div class="bg-white rounded-[32px] p-10 border border-slate-100 shadow-sm">
            <div class="w-12 h-12 bg-red-50 rounded-xl flex items-center justify-center mb-8">
              <svg class="w-6 h-6 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
              </svg>
            </div>
            <h3 class="text-xl font-extrabold mb-6">The Old Way</h3>
            <ul class="space-y-4">
              <li v-for="item in oldWay" :key="item" class="flex gap-3 items-start text-slate-500 font-medium text-[15px]">
                <span class="text-red-400 mt-1">✕</span> {{ item }}
              </li>
            </ul>
          </div>

          <!-- The Lumora Way -->
          <div class="bg-white rounded-[32px] p-10 border border-slate-100 shadow-sm">
            <div class="w-12 h-12 bg-[#EEF2FF] rounded-xl flex items-center justify-center mb-8">
              <svg class="w-6 h-6 text-[#3D3ACE]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
            </div>
            <h3 class="text-xl font-extrabold mb-6">The Lumora Way</h3>
            <ul class="space-y-4">
              <li v-for="item in lumoraWay" :key="item" class="flex gap-3 items-start text-slate-500 font-medium text-[15px]">
                <span class="text-[#3D3ACE] mt-1 font-bold">✓</span> {{ item }}
              </li>
            </ul>
          </div>
        </div>
      </div>
    </section>

    <!-- ─── Bento Grid Features ────────────────────────────────────────── -->
    <section id="features" class="py-32 bg-white">
      <div class="max-w-7xl mx-auto px-6">
        <div class="grid lg:grid-cols-3 gap-8">
          <!-- Intelligent Planner -->
          <div class="lg:col-span-2 bg-[#F9FAFB] rounded-[32px] p-12 overflow-hidden flex flex-col justify-between group">
            <div class="max-w-md">
              <div class="text-[#3D3ACE] text-[11px] font-black uppercase tracking-widest mb-4">Visual Clarity</div>
              <h3 class="text-3xl font-extrabold mb-4">Intelligent Planner</h3>
              <p class="text-slate-500 font-medium mb-12">
                Sync your syllabi and let Lumora map out your semester. It adjusts automatically as priorities shift.
              </p>
            </div>
            <div class="relative mt-auto">
              <div class="bg-white rounded-t-2xl shadow-2xl p-4 transform translate-y-6 group-hover:translate-y-2 transition-transform duration-500 overflow-hidden">
                 <img src="/image/Planner Preview.png" alt="Intelligent Planner Preview" class="w-full h-auto rounded-lg">
              </div>
            </div>
          </div>

          <!-- Deep Notes -->
          <div class="bg-[#F9FAFB] rounded-[32px] p-12 flex flex-col items-center text-center group">
            <div class="w-12 h-12 bg-[#D1FAE5] text-[#065F46] rounded-xl flex items-center justify-center mb-8">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
              </svg>
            </div>
            <h3 class="text-2xl font-extrabold mb-4">Deep Notes</h3>
            <p class="text-slate-500 font-medium mb-12">
              Structured Markdown support with bidirectional linking to your study sessions.
            </p>
            <div class="w-full mt-auto">
              <div class="bg-white rounded-2xl shadow-lg p-6 flex gap-2 items-end h-32">
                <div class="flex-1 bg-[#D1FAE5] rounded-t-md h-1/2"></div>
                <div class="flex-1 bg-[#D1FAE5] rounded-t-md h-3/4"></div>
                <div class="flex-1 bg-[#65A30D]/60 rounded-t-md h-full"></div>
              </div>
            </div>
          </div>

          <!-- AI Strategies Banner -->
          <div class="lg:col-span-3 bg-[#3D3ACE] rounded-[32px] p-12 flex flex-col md:flex-row items-center justify-between text-white overflow-hidden relative">
            <div class="max-w-xl z-10">
              <div class="text-indigo-200 text-[11px] font-black uppercase tracking-widest mb-4">Future-Proof Learning</div>
              <h3 class="text-3xl font-extrabold mb-4">AI-Powered Strategies</h3>
              <p class="text-indigo-100 font-medium">
                Our AI doesn't just answer questions; it teaches you how to study. Get personalized active recall suggestions and spaced repetition schedules based on your actual performance.
              </p>
            </div>
            <div class="relative mt-8 md:mt-0 z-10">
              <div class="w-32 h-32 bg-white/10 backdrop-blur-md rounded-full flex items-center justify-center border border-white/20">
                <svg class="w-12 h-12 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0012 18.75c-1.03 0-1.9-.4-2.593-1.012l-.547-.547z" />
                </svg>
              </div>
            </div>
            <!-- Decorative circle -->
            <div class="absolute right-0 bottom-0 w-96 h-96 bg-white/5 rounded-full translate-x-1/3 translate-y-1/3"></div>
          </div>
        </div>
      </div>
    </section>

    <!-- ─── Three Steps ──────────────────────────────────────────────── -->
    <section id="how" class="py-32 bg-[#F9FAFB]">
      <div class="max-w-7xl mx-auto px-6 text-center">
        <h2 class="text-[40px] font-extrabold text-slate-900 mb-4">Mastery in three steps.</h2>
        <p class="text-slate-500 text-lg mb-20">We've codified the science of learning into a simple workflow.</p>
        
        <div class="relative max-w-6xl mx-auto">
          <!-- Connecting Line -->
          <div class="hidden md:block absolute top-12 left-[15%] right-[15%] h-0.5 bg-slate-200"></div>
          
          <div class="grid md:grid-cols-3 gap-16 relative">
            <div v-for="(step, idx) in steps" :key="step.title">
              <div class="w-24 h-24 bg-white rounded-[32px] shadow-sm flex items-center justify-center text-3xl font-black text-[#3D3ACE] mx-auto mb-10 relative z-10 border border-slate-100">
                {{ idx + 1 }}
              </div>
              <h3 class="text-2xl font-extrabold mb-4">{{ step.title }}</h3>
              <p class="text-slate-500 font-medium leading-relaxed">{{ step.desc }}</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- ─── Testimonials ──────────────────────────────────────────────── -->
    <section id="testimonials" class="py-32 bg-white">
      <div class="max-w-7xl mx-auto px-6">
        <h2 class="text-center text-[32px] font-extrabold text-slate-900 mb-20">Loved by students at 500+ universities.</h2>
        
        <div class="grid md:grid-cols-3 gap-8">
          <div v-for="t in testimonials" :key="t.name" class="bg-[#F9FAFB] rounded-[32px] p-10 border border-slate-100">
            <div class="flex text-amber-400 gap-1 mb-6 text-sm">
              <span v-for="n in 5" :key="n">★</span>
            </div>
            <p class="text-slate-600 font-medium italic mb-10 leading-relaxed">
              "{{ t.quote }}"
            </p>
            <div class="flex items-center gap-4">
              <div class="w-12 h-12 rounded-full bg-[#EEF2FF] flex items-center justify-center text-[#3D3ACE] font-bold text-sm">
                {{ t.initials }}
              </div>
              <div>
                <p class="font-bold text-slate-900">{{ t.name }}</p>
                <p class="text-xs text-slate-500 font-medium">{{ t.role }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- ─── Final CTA ─────────────────────────────────────────────────── -->
    <section class="py-32 bg-white">
      <div class="max-w-5xl mx-auto px-6">
        <div class="bg-white rounded-[48px] p-16 text-center border border-slate-100 shadow-2xl shadow-indigo-50 relative overflow-hidden">
          <h2 class="text-[40px] font-extrabold text-slate-900 mb-6 tracking-tight">Ready to transform your study habits?</h2>
          <p class="text-slate-500 text-lg mb-10 max-w-xl mx-auto">
            Join thousands of students who have found their sanctuary of productivity.
          </p>
          
          <Link :href="route('register')"
            class="inline-block bg-[#3D3ACE] text-white font-black px-12 py-5 rounded-2xl hover:bg-[#322fb0] transition-all shadow-xl shadow-indigo-100 mb-6">
            Get Started for Free
          </Link>
          
          <p class="text-slate-400 text-xs font-medium uppercase tracking-widest">No credit card required. Free tier available forever.</p>
        </div>
      </div>
    </section>

    <!-- ─── Footer ────────────────────────────────────────────────────── -->
    <footer class="bg-white py-20 border-t border-slate-50">
      <div class="max-w-7xl mx-auto px-6 text-center">
        <div class="text-xl font-black text-slate-900 mb-10">Lumora</div>
        <div class="flex flex-wrap justify-center gap-8 text-[15px] font-medium text-slate-500 mb-12">
          <a href="#" class="hover:text-slate-900">Privacy Policy</a>
          <a href="#" class="hover:text-slate-900">Terms of Service</a>
          <a href="#" class="hover:text-slate-900">Contact Support</a>
        </div>
        <p class="text-slate-400 text-sm font-medium">© 2024 Lumora AI. Designed for focused minds.</p>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { Link } from '@inertiajs/vue3'

const oldWay = [
  "Scattered notes across five different apps.",
  "Feeling overwhelmed by monolithic to-do lists.",
  "No feedback loop on study effectiveness."
]

const lumoraWay = [
  "One unified sanctuary for every lecture and task.",
  "AI-driven breaking down of complex goals.",
  "Reflective analytics that help you learn faster."
]

const steps = [
  {
    title: 'Assess',
    desc: 'Upload your syllabus and curriculum goals. Lumora identifies critical focus areas using semantic analysis.',
  },
  {
    title: 'Plan',
    desc: 'Get a customized, flexible study calendar that balances high-intensity work with essential rest periods.',
  },
  {
    title: 'Improve',
    desc: 'Weekly reflection sessions help you adjust your habits based on real performance data and focus metrics.',
  },
]

const testimonials = [
  {
    name: 'Elena J.',
    role: 'Medical Student, Johns Hopkins',
    initials: 'EJ',
    quote: 'Lumora changed how I approach my med school exams. The AI-suggested study schedules actually made me feel human again.'
  },
  {
    name: 'Marcus T.',
    role: 'Law Student, Yale',
    initials: 'MT',
    quote: 'The visual planner is a lifesaver. Being able to see my entire semester and where my focus needs to be is so empowering.'
  },
  {
    name: 'Sarah C.',
    role: 'Ph.D. Candidate, Stanford',
    initials: 'SC',
    quote: 'Finally, a tool that focuses on the learning process, not just checking boxes. It\'s my daily sanctuary.'
  }
]
</script>
