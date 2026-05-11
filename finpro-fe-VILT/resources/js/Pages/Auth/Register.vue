<template>
  <div class="min-h-screen bg-[#F0F2F9] flex flex-col items-center justify-center px-4 relative overflow-hidden font-sans py-20">
    <!-- Decorative background elements -->
    <div class="absolute top-[-10%] left-[-5%] w-[500px] h-[500px] bg-indigo-200/30 rounded-full blur-[100px]"></div>
    <div class="absolute bottom-[-10%] right-[-5%] w-[400px] h-[400px] bg-purple-200/30 rounded-full blur-[100px]"></div>

    <!-- Main Content -->
    <div class="relative z-10 w-full max-w-[500px] flex flex-col items-center">
      <!-- Logo & Header -->
      <div class="text-center mb-10">
        <Link :href="route('landing')" class="flex items-center justify-center gap-3 mb-4">
          <div class="w-10 h-10 bg-[#3D3ACE] rounded-xl flex items-center justify-center shadow-lg shadow-indigo-200">
             <svg class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 24 24">
               <path d="M12 2L4.5 20.29l.71.71L12 18l6.79 3 .71-.71z"/>
             </svg>
          </div>
          <span class="text-2xl font-black text-[#1E1B4B] tracking-tight">Lumora</span>
        </Link>
        <p class="text-slate-500 font-medium">Create your account to start your journey.</p>
      </div>

      <!-- Register Card -->
      <div class="w-full bg-white rounded-[40px] p-10 md:p-12 shadow-2xl shadow-indigo-100 border border-white/50">
        <form @submit.prevent="submit" class="space-y-6">
          <!-- Full Name Field -->
          <div class="space-y-3">
            <label class="block text-[11px] font-black text-slate-400 uppercase tracking-widest ml-1">Full Name</label>
            <input 
              v-model="form.name"
              type="text" 
              placeholder="Enter your name"
              class="w-full bg-[#F3F4F6] border-transparent focus:bg-white focus:border-[#3D3ACE] focus:ring-4 focus:ring-indigo-100 rounded-2xl py-4 px-6 text-slate-700 font-medium transition-all duration-300"
              required
            >
            <p v-if="errors.name" class="text-xs text-red-500 ml-1">{{ errors.name }}</p>
          </div>

          <!-- Email Field -->
          <div class="space-y-3">
            <label class="block text-[11px] font-black text-slate-400 uppercase tracking-widest ml-1">Email Address</label>
            <input 
              v-model="form.email"
              type="email" 
              placeholder="example@lumora.edu"
              class="w-full bg-[#F3F4F6] border-transparent focus:bg-white focus:border-[#3D3ACE] focus:ring-4 focus:ring-indigo-100 rounded-2xl py-4 px-6 text-slate-700 font-medium transition-all duration-300"
              required
            >
            <p v-if="errors.email" class="text-xs text-red-500 ml-1">{{ errors.email }}</p>
          </div>

          <!-- Password Field -->
          <div class="space-y-3">
            <label class="block text-[11px] font-black text-slate-400 uppercase tracking-widest ml-1">Password</label>
            <div class="relative group">
              <span class="absolute left-5 top-1/2 -translate-y-1/2 text-slate-400">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                </svg>
              </span>
              <input 
                v-model="form.password"
                type="password" 
                placeholder="Create a password"
                class="w-full bg-[#F3F4F6] border-transparent focus:bg-white focus:border-[#3D3ACE] focus:ring-4 focus:ring-indigo-100 rounded-2xl py-4 pl-14 pr-6 text-slate-700 font-medium transition-all duration-300"
                required
              >
            </div>
            <p v-if="errors.password" class="text-xs text-red-500 ml-1">{{ errors.password }}</p>
          </div>

          <!-- Confirm Password Field -->
          <div class="space-y-3">
            <label class="block text-[11px] font-black text-slate-400 uppercase tracking-widest ml-1">Confirm Password</label>
            <div class="relative group">
              <span class="absolute left-5 top-1/2 -translate-y-1/2 text-slate-400">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                </svg>
              </span>
              <input 
                v-model="form.password_confirmation"
                type="password" 
                placeholder="Repeat your password"
                class="w-full bg-[#F3F4F6] border-transparent focus:bg-white focus:border-[#3D3ACE] focus:ring-4 focus:ring-indigo-100 rounded-2xl py-4 pl-14 pr-6 text-slate-700 font-medium transition-all duration-300"
                required
              >
            </div>
          </div>

          <!-- Terms -->
          <p class="text-center text-xs text-slate-400 leading-relaxed px-4">
            By registering, you agree to our 
            <a href="#" class="text-[#3D3ACE] font-bold hover:underline">Terms of Service</a> and 
            <a href="#" class="text-[#3D3ACE] font-bold hover:underline">Privacy Policy</a>.
          </p>

          <!-- Register Button -->
          <button 
            type="submit"
            :disabled="form.processing"
            class="w-full bg-[#3D3ACE] hover:bg-[#322fb0] text-white font-bold py-5 rounded-2xl shadow-xl shadow-indigo-100 transition-all duration-300 transform active:scale-[0.98] disabled:opacity-70"
          >
            <span v-if="form.processing">Creating account...</span>
            <span v-else>Register</span>
          </button>
        </form>

        <p class="text-center mt-10 text-[15px] font-medium text-slate-500">
          Already have an account? 
          <Link :href="route('login')" class="text-[#3D3ACE] font-bold hover:underline">Login</Link>
        </p>
      </div>

      <!-- Decorative shape -->
      <div class="mt-12 opacity-20">
          <svg class="w-16 h-16 text-[#3D3ACE]" fill="currentColor" viewBox="0 0 24 24">
             <path d="M12 2L4.5 20.29l.71.71L12 18l6.79 3 .71-.71z"/>
          </svg>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useForm, Link } from '@inertiajs/vue3'

defineProps({
  errors: { type: Object, default: () => ({}) },
})

const form = useForm({
  name: '',
  email: '',
  password: '',
  password_confirmation: '',
})

function submit() {
  form.post(route('register'), {
    onError: () => form.reset('password', 'password_confirmation'),
  })
}
</script>

