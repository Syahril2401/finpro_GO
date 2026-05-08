<template>
  <div class="min-h-screen bg-gradient-to-br from-indigo-950 via-slate-900 to-slate-950 flex items-center justify-center px-4">
    <!-- Background blobs -->
    <div class="absolute top-0 left-0 w-96 h-96 bg-indigo-600/20 rounded-full blur-3xl"></div>
    <div class="absolute bottom-0 right-0 w-80 h-80 bg-violet-600/20 rounded-full blur-3xl"></div>

    <div class="relative z-10 w-full max-w-md">
      <!-- Logo -->
      <div class="text-center mb-8">
        <Link :href="route('landing')" class="inline-block">
          <span class="text-3xl font-extrabold text-white tracking-tight">Lumora</span>
        </Link>
        <p class="text-slate-400 mt-2 text-sm">Welcome back. Let's keep learning.</p>
      </div>

      <!-- Card -->
      <div class="bg-white/5 backdrop-blur-xl border border-white/10 rounded-3xl p-8 shadow-2xl">
        <h1 class="text-2xl font-bold text-white mb-6">Sign in to your account</h1>

        <!-- Error -->
        <div v-if="errors.email" class="bg-red-500/10 border border-red-500/30 text-red-300 text-sm rounded-xl px-4 py-3 mb-5">
          {{ errors.email }}
        </div>

        <form @submit.prevent="submit" class="space-y-5">
          <!-- Email -->
          <div>
            <label for="email" class="block text-sm font-medium text-slate-300 mb-1.5">Email</label>
            <input
              id="email"
              v-model="form.email"
              type="email"
              autocomplete="email"
              required
              placeholder="you@example.com"
              class="w-full bg-white/5 border border-white/10 text-white placeholder-slate-500 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition-all"
            />
          </div>

          <!-- Password -->
          <div>
            <label for="password" class="block text-sm font-medium text-slate-300 mb-1.5">Password</label>
            <input
              id="password"
              v-model="form.password"
              type="password"
              autocomplete="current-password"
              required
              placeholder="Enter your password"
              class="w-full bg-white/5 border border-white/10 text-white placeholder-slate-500 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition-all"
            />
          </div>

          <!-- Submit -->
          <button
            type="submit"
            :disabled="form.processing"
            id="login-submit-btn"
            class="w-full bg-indigo-600 hover:bg-indigo-500 disabled:opacity-60 text-white font-semibold py-3 rounded-xl transition-all shadow-lg shadow-indigo-900/50"
          >
            <span v-if="form.processing">Signing in...</span>
            <span v-else>Sign in</span>
          </button>
        </form>

        <p class="text-center text-sm text-slate-400 mt-6">
          Don't have an account?
          <Link :href="route('register')" class="text-indigo-400 hover:text-indigo-300 font-medium transition-colors">
            Sign up
          </Link>
        </p>
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
  email: '',
  password: '',
})

function submit() {
  form.post(route('login'), {
    onError: () => form.reset('password'),
  })
}
</script>
