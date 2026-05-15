<template>
  <div 
    class="p-4 rounded-2xl border transition-all cursor-pointer group"
    :class="isActive ? 'bg-[#3D3ACE] border-[#3D3ACE] text-white shadow-lg shadow-indigo-100' : 'bg-white border-slate-100 hover:border-indigo-200 hover:shadow-md text-[#1E1B4B]'"
    @click="$emit('click')"
  >
    <div class="flex items-start justify-between mb-2">
      <h4 class="font-bold text-sm truncate flex-1 pr-2" :class="isActive ? 'text-white' : 'text-[#1E1B4B]'">
        {{ note.title || 'Untitled' }}
      </h4>
      <div class="flex gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
        <button v-if="note.is_pinned" class="text-amber-400" @click.stop="$emit('toggle-pin')">
          <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20"><path d="M5 4a2 2 0 012-2h6a2 2 0 012 2v14l-5-2.5L5 18V4z"/></svg>
        </button>
        <button v-else class="hover:text-amber-400" :class="isActive ? 'text-indigo-200' : 'text-slate-300'" @click.stop="$emit('toggle-pin')">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z"/></svg>
        </button>
      </div>
    </div>
    
    <p class="text-xs line-clamp-2 mb-3" :class="isActive ? 'text-indigo-100' : 'text-slate-500'">
      {{ previewText }}
    </p>

    <div class="flex items-center justify-between mt-auto">
      <div class="flex items-center gap-2">
        <span class="text-[9px] font-black uppercase tracking-widest px-1.5 py-0.5 rounded-md"
          :class="isActive ? 'bg-white/20 text-white' : 'bg-indigo-50 text-[#3D3ACE]'">
          {{ note.focus_dimension || 'General' }}
        </span>
        <span v-if="note.target_id" title="Linked to Weekly Target">🎯</span>
        <span v-if="note.planner_session_id" title="Linked to Planner">📅</span>
      </div>
      <span class="text-[10px] font-bold" :class="isActive ? 'text-indigo-200' : 'text-slate-400'">
        {{ formattedDate }}
      </span>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
  note: Object,
  isActive: Boolean
});

defineEmits(['click', 'toggle-pin']);

const previewText = computed(() => {
  if (props.note.content_text && props.note.content_text.trim().length > 0) {
    return props.note.content_text;
  }
  return 'No content yet...';
});

const formattedDate = computed(() => {
  if (!props.note.updated_at) return '';
  const date = new Date(props.note.updated_at);
  return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric' });
});
</script>
