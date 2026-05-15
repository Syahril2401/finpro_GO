<template>
  <div class="flex-1 flex flex-col h-full overflow-hidden bg-white relative">
    
    <!-- Editor Header -->
    <div class="h-16 border-b border-slate-100 flex items-center justify-between px-10 flex-shrink-0 bg-white z-20 relative">
      <div class="text-[11px] font-black uppercase tracking-widest flex items-center gap-2" :class="saveStatusColor">
        <span class="w-2 h-2 rounded-full animate-pulse" :class="saveStatusBg"></span>
        {{ saveStatusText }}
      </div>
      <button v-if="saveStatus === 'error' || saveStatus === 'unsaved'" @click="$emit('save')" class="bg-[#3D3ACE] text-white px-4 py-1.5 rounded-lg text-xs font-bold hover:bg-[#322fb0]">
        Save Now
      </button>
    </div>

    <!-- Editor Scroll Area -->
    <div class="flex-1 overflow-y-auto p-10 lg:px-20 xl:px-40 scrollbar-hide pb-32" @click="focusEditorEnd">
      <div class="max-w-3xl mx-auto w-full">
        
        <!-- Title -->
        <input 
          type="text" 
          :value="title"
          @input="$emit('update:title', $event.target.value)"
          placeholder="Untitled"
          class="w-full text-4xl font-black text-[#1E1B4B] mb-8 outline-none bg-transparent placeholder-slate-300"
        >

        <!-- Empty State Hint -->
        <div v-if="blocks.length === 0" class="text-slate-400 font-medium cursor-text" @click.stop="createFirstBlock">
          Start writing your reflection, or type <span class="bg-slate-100 px-2 py-0.5 rounded text-slate-600 font-black">/</span> for commands...
        </div>

        <!-- Blocks -->
        <div v-else class="space-y-1">
          <NoteBlock 
            v-for="(block, index) in blocks" 
            :key="block.id"
            :block="block"
            :isFocused="focusedBlockIndex === index"
            @update="updateBlock(index, $event)"
            @split="splitBlock(index, $event)"
            @delete="deleteBlock(index)"
            @focus="focusedBlockIndex = index"
            @move-up="moveFocusUp(index)"
            @move-down="moveFocusDown(index)"
            @slash-command="openSlashCommand(index)"
          />
        </div>
      </div>
    </div>

    <!-- Slash Command Menu -->
    <SlashCommandMenu 
      v-if="slashMenuOpen" 
      :style="{ top: slashMenuPos.y + 'px', left: slashMenuPos.x + 'px' }"
      @select="applyCommand"
      @close="closeSlashCommand"
    />
  </div>
</template>

<script setup>
import { ref, onMounted, nextTick } from 'vue';
import NoteBlock from './NoteBlock.vue';
import SlashCommandMenu from './SlashCommandMenu.vue';

const props = defineProps({
  title: String,
  blocks: Array,
  saveStatus: {
    type: String, // 'saved', 'saving', 'unsaved', 'error'
    default: 'saved'
  }
});

const emit = defineEmits(['update:title', 'update:blocks', 'save']);

const focusedBlockIndex = ref(-1);
const slashMenuOpen = ref(false);
const slashMenuPos = ref({ x: 0, y: 0 });
const activeSlashBlockIndex = ref(-1);

const saveStatusColor = computed(() => {
  if (props.saveStatus === 'saved') return 'text-slate-400';
  if (props.saveStatus === 'saving') return 'text-amber-500';
  if (props.saveStatus === 'error') return 'text-rose-500';
  return 'text-slate-500';
});

const saveStatusBg = computed(() => {
  if (props.saveStatus === 'saved') return 'bg-slate-300';
  if (props.saveStatus === 'saving') return 'bg-amber-400';
  if (props.saveStatus === 'error') return 'bg-rose-500';
  return 'bg-slate-400';
});

const saveStatusText = computed(() => {
  if (props.saveStatus === 'saved') return 'Saved';
  if (props.saveStatus === 'saving') return 'Saving...';
  if (props.saveStatus === 'error') return 'Failed to save';
  return 'Unsaved changes';
});

function createFirstBlock() {
  const newBlocks = [{ id: Date.now().toString(), type: 'paragraph', content: '' }];
  emit('update:blocks', newBlocks);
  focusedBlockIndex.value = 0;
}

function updateBlock(index, updatedBlock) {
  const newBlocks = [...props.blocks];
  newBlocks[index] = updatedBlock;
  emit('update:blocks', newBlocks);
}

function splitBlock(index, newBlockContent) {
  const newBlocks = [...props.blocks];
  const currentBlockType = newBlocks[index].type;
  
  // If current is bullet/number/todo, continue the list. Otherwise fallback to paragraph.
  let newType = 'paragraph';
  if (['bullet', 'number', 'todo'].includes(currentBlockType)) {
    newType = currentBlockType;
  }
  
  newBlocks.splice(index + 1, 0, {
    id: Date.now().toString(),
    type: newType,
    content: newBlockContent || '',
    checked: false
  });
  
  emit('update:blocks', newBlocks);
  nextTick(() => {
    focusedBlockIndex.value = index + 1;
  });
}

function deleteBlock(index) {
  if (props.blocks.length === 1 && !props.blocks[0].content) {
    emit('update:blocks', []);
    focusedBlockIndex.value = -1;
    return;
  }
  const newBlocks = [...props.blocks];
  newBlocks.splice(index, 1);
  emit('update:blocks', newBlocks);
  if (index > 0) {
    focusedBlockIndex.value = index - 1;
  } else {
    focusedBlockIndex.value = 0;
  }
}

function moveFocusUp(index) {
  if (index > 0) focusedBlockIndex.value = index - 1;
}

function moveFocusDown(index) {
  if (index < props.blocks.length - 1) focusedBlockIndex.value = index + 1;
}

function focusEditorEnd(e) {
  // If clicking empty space at bottom, append a paragraph
  if (e.target === e.currentTarget && props.blocks.length > 0) {
    splitBlock(props.blocks.length - 1, '');
  }
}

function openSlashCommand(index) {
  // Attempt to position menu near selection
  const sel = window.getSelection();
  if (sel.rangeCount > 0) {
    const rect = sel.getRangeAt(0).getBoundingClientRect();
    slashMenuPos.value = { x: rect.left, y: rect.bottom + 10 };
  } else {
    slashMenuPos.value = { x: 300, y: 300 };
  }
  activeSlashBlockIndex.value = index;
  slashMenuOpen.value = true;
}

function closeSlashCommand() {
  slashMenuOpen.value = false;
  activeSlashBlockIndex.value = -1;
}

function applyCommand(type) {
  if (activeSlashBlockIndex.value >= 0) {
    const newBlocks = [...props.blocks];
    newBlocks[activeSlashBlockIndex.value].type = type;
    newBlocks[activeSlashBlockIndex.value].content = newBlocks[activeSlashBlockIndex.value].content.replace('/', '');
    emit('update:blocks', newBlocks);
  }
  closeSlashCommand();
}

import { computed } from 'vue';
</script>
