<template>
  <div v-if="isOpen" class="fixed inset-0 bg-slate-900/40 backdrop-blur-sm z-[100] flex items-center justify-center p-6">
    <div class="bg-white rounded-[32px] p-8 w-full max-w-3xl shadow-2xl flex flex-col max-h-[80vh]">
      <div class="flex items-center justify-between mb-6">
        <h2 class="text-2xl font-black text-[#1E1B4B]">Choose a Template</h2>
        <button @click="$emit('close')" class="text-slate-400 hover:text-rose-500 transition-colors">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/></svg>
        </button>
      </div>

      <div v-if="isLoading" class="flex-1 flex items-center justify-center">
        <div class="w-8 h-8 border-4 border-indigo-200 border-t-[#3D3ACE] rounded-full animate-spin"></div>
      </div>

      <div v-else class="flex-1 overflow-y-auto grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 p-2 scrollbar-hide">
        <div 
          v-for="tpl in templates" 
          :key="tpl.id"
          @click="$emit('select', tpl)"
          class="bg-slate-50 border border-slate-100 p-6 rounded-3xl cursor-pointer hover:border-[#3D3ACE] hover:shadow-xl hover:shadow-indigo-50 hover:-translate-y-1 transition-all group flex flex-col h-40"
        >
          <div class="w-10 h-10 bg-white rounded-xl shadow-sm flex items-center justify-center text-xl mb-4 group-hover:scale-110 transition-transform">
            {{ getTemplateIcon(tpl.id) }}
          </div>
          <h3 class="text-sm font-black text-[#1E1B4B] group-hover:text-[#3D3ACE] transition-colors mb-1">{{ tpl.title }}</h3>
          <p class="text-[10px] font-bold text-slate-400 line-clamp-2">{{ getTemplateDesc(tpl.id) }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  isOpen: Boolean,
  isLoading: Boolean,
  templates: Array
});

defineEmits(['close', 'select']);

function getTemplateIcon(id) {
  const icons = {
    'blank': '📄',
    'daily': '🌅',
    'session': '⏱️',
    'weekly': '📅',
    'exam': '🎓',
    'cognitive': '🧠'
  };
  return icons[id] || '📝';
}

function getTemplateDesc(id) {
  const descs = {
    'blank': 'Start from scratch with a blank page.',
    'daily': 'Quick reflection on your day\'s learning.',
    'session': 'Review what you achieved in a focus session.',
    'weekly': 'Analyze your progress over the past week.',
    'exam': 'Plan and review your exam preparation.',
    'cognitive': 'Analyze the effectiveness of a learning strategy.'
  };
  return descs[id] || 'Use this template to structure your note.';
}
</script>
