<template>
  <div class="fixed z-[100] w-64 bg-white rounded-2xl shadow-2xl border border-slate-100 overflow-hidden py-2 text-sm font-bold">
    <div class="px-3 py-2 text-[10px] uppercase tracking-widest text-slate-400 font-black">Basic Blocks</div>
    
    <button @click="select('paragraph')" class="w-full text-left px-4 py-2 hover:bg-slate-50 flex items-center gap-3 text-[#1E1B4B]">
      <span class="w-6 h-6 flex items-center justify-center bg-slate-100 rounded text-slate-500">T</span> Text
    </button>
    <button @click="select('heading1')" class="w-full text-left px-4 py-2 hover:bg-slate-50 flex items-center gap-3 text-[#1E1B4B]">
      <span class="w-6 h-6 flex items-center justify-center bg-slate-100 rounded text-slate-500">H1</span> Heading 1
    </button>
    <button @click="select('heading2')" class="w-full text-left px-4 py-2 hover:bg-slate-50 flex items-center gap-3 text-[#1E1B4B]">
      <span class="w-6 h-6 flex items-center justify-center bg-slate-100 rounded text-slate-500">H2</span> Heading 2
    </button>
    <button @click="select('bullet')" class="w-full text-left px-4 py-2 hover:bg-slate-50 flex items-center gap-3 text-[#1E1B4B]">
      <span class="w-6 h-6 flex items-center justify-center bg-slate-100 rounded text-slate-500">•</span> Bulleted List
    </button>
    <button @click="select('number')" class="w-full text-left px-4 py-2 hover:bg-slate-50 flex items-center gap-3 text-[#1E1B4B]">
      <span class="w-6 h-6 flex items-center justify-center bg-slate-100 rounded text-slate-500">1.</span> Numbered List
    </button>
    <button @click="select('todo')" class="w-full text-left px-4 py-2 hover:bg-slate-50 flex items-center gap-3 text-[#1E1B4B]">
      <span class="w-6 h-6 flex items-center justify-center bg-slate-100 rounded text-slate-500">✓</span> To-do List
    </button>
    <button @click="select('quote')" class="w-full text-left px-4 py-2 hover:bg-slate-50 flex items-center gap-3 text-[#1E1B4B]">
      <span class="w-6 h-6 flex items-center justify-center bg-slate-100 rounded text-slate-500">"</span> Quote
    </button>
    <button @click="select('divider')" class="w-full text-left px-4 py-2 hover:bg-slate-50 flex items-center gap-3 text-[#1E1B4B]">
      <span class="w-6 h-6 flex items-center justify-center bg-slate-100 rounded text-slate-500">—</span> Divider
    </button>
    <button @click="select('callout')" class="w-full text-left px-4 py-2 hover:bg-slate-50 flex items-center gap-3 text-[#1E1B4B]">
      <span class="w-6 h-6 flex items-center justify-center bg-indigo-50 text-indigo-500 rounded">💡</span> Callout
    </button>
    <button @click="select('reflection')" class="w-full text-left px-4 py-2 hover:bg-slate-50 flex items-center gap-3 text-[#1E1B4B]">
      <span class="w-6 h-6 flex items-center justify-center bg-amber-50 text-amber-500 rounded">🤔</span> Reflection Prompt
    </button>
  </div>
</template>

<script setup>
import { onMounted, onUnmounted } from 'vue';

const emit = defineEmits(['select', 'close']);

function select(type) {
  emit('select', type);
}

function handleClickOutside(e) {
  emit('close');
}

onMounted(() => {
  setTimeout(() => {
    document.addEventListener('click', handleClickOutside);
  }, 10);
});

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside);
});
</script>
