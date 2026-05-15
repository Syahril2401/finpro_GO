<template>
  <aside class="w-80 bg-white border-l border-slate-100 flex flex-col h-full z-30">
    <div class="p-6 border-b border-slate-100 flex items-center justify-between">
      <h3 class="text-sm font-black text-[#1E1B4B] uppercase tracking-widest">Properties</h3>
      <div class="flex gap-2">
        <button @click="$emit('toggle-archive')" class="w-8 h-8 rounded-xl flex items-center justify-center transition-colors"
          :class="note.is_archived ? 'bg-amber-50 text-amber-500' : 'bg-slate-50 text-slate-400 hover:bg-slate-100 hover:text-slate-600'" title="Archive">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4"/></svg>
        </button>
        <button @click="$emit('delete')" class="w-8 h-8 bg-rose-50 text-rose-500 rounded-xl flex items-center justify-center hover:bg-rose-500 hover:text-white transition-colors" title="Delete">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/></svg>
        </button>
      </div>
    </div>

    <div class="flex-1 overflow-y-auto p-6 space-y-6 scrollbar-hide">
      
      <!-- Focus Dimension -->
      <div>
        <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">Focus Dimension</label>
        <select 
          :value="note.focus_dimension" 
          @change="update('focus_dimension', $event.target.value)"
          class="w-full bg-slate-50 border border-slate-100 rounded-xl px-3 py-2.5 text-sm font-bold text-[#1E1B4B] focus:bg-white focus:border-[#3D3ACE] outline-none transition-colors"
        >
          <option value="Planning">Planning</option>
          <option value="Time Management">Time Management</option>
          <option value="Cognitive Strategy">Cognitive Strategy</option>
          <option value="Reflection">Reflection</option>
          <option value="General">General</option>
        </select>
      </div>

      <!-- Mood -->
      <div>
        <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">Mood</label>
        <select 
          :value="note.mood" 
          @change="update('mood', $event.target.value)"
          class="w-full bg-slate-50 border border-slate-100 rounded-xl px-3 py-2.5 text-sm font-bold text-[#1E1B4B] focus:bg-white focus:border-[#3D3ACE] outline-none transition-colors"
        >
          <option value="">Select mood...</option>
          <option value="Focused">Focused</option>
          <option value="Neutral">Neutral</option>
          <option value="Tired">Tired</option>
          <option value="Confused">Confused</option>
          <option value="Motivated">Motivated</option>
          <option value="Stressed">Stressed</option>
        </select>
      </div>

      <!-- Confidence Level -->
      <div>
        <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">Confidence Level</label>
        <div class="flex items-center gap-1">
          <button 
            v-for="i in 5" 
            :key="i"
            @click="update('confidence_level', i)"
            class="w-10 h-10 rounded-xl text-sm font-black transition-colors border"
            :class="(note.confidence_level || 0) >= i ? 'bg-emerald-500 border-emerald-500 text-white' : 'bg-slate-50 border-slate-100 text-slate-300 hover:border-emerald-200'"
          >
            {{ i }}
          </button>
        </div>
        <div class="flex justify-between mt-1 text-[9px] font-bold text-slate-400 uppercase tracking-widest">
          <span>Low</span>
          <span>High</span>
        </div>
      </div>

      <!-- Tags -->
      <div>
        <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">Tags</label>
        <input 
          type="text" 
          :value="note.tags"
          @change="update('tags', $event.target.value)"
          placeholder="react, history, essay..."
          class="w-full bg-slate-50 border border-slate-100 rounded-xl px-3 py-2.5 text-sm font-bold text-[#1E1B4B] focus:bg-white focus:border-[#3D3ACE] outline-none transition-colors"
        >
      </div>

      <!-- Linked Target -->
      <div>
        <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">Linked Target</label>
        <input 
          type="text" 
          :value="note.target_id"
          @change="update('target_id', $event.target.value)"
          placeholder="Target ID (Optional)"
          class="w-full bg-slate-50 border border-slate-100 rounded-xl px-3 py-2.5 text-sm font-bold text-[#1E1B4B] focus:bg-white focus:border-[#3D3ACE] outline-none transition-colors"
        >
      </div>

      <!-- Linked Session -->
      <div>
        <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest mb-2">Linked Session</label>
        <input 
          type="text" 
          :value="note.planner_session_id"
          @change="update('planner_session_id', $event.target.value)"
          placeholder="Session ID (Optional)"
          class="w-full bg-slate-50 border border-slate-100 rounded-xl px-3 py-2.5 text-sm font-bold text-[#1E1B4B] focus:bg-white focus:border-[#3D3ACE] outline-none transition-colors"
        >
      </div>

    </div>

    <!-- Info Footer -->
    <div class="p-6 border-t border-slate-100 bg-slate-50">
      <div class="text-[10px] font-bold text-slate-400 space-y-1">
        <div class="flex justify-between">
          <span>Created:</span>
          <span>{{ formatDate(note.created_at) }}</span>
        </div>
        <div class="flex justify-between">
          <span>Updated:</span>
          <span>{{ formatDate(note.updated_at) }}</span>
        </div>
      </div>
    </div>
  </aside>
</template>

<script setup>
const props = defineProps({
  note: {
    type: Object,
    required: true
  }
});

const emit = defineEmits(['update', 'delete', 'toggle-archive']);

function update(field, value) {
  emit('update', field, value);
}

function formatDate(dateStr) {
  if (!dateStr) return 'Unknown';
  const d = new Date(dateStr);
  return d.toLocaleString('en-US', { month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit' });
}
</script>
