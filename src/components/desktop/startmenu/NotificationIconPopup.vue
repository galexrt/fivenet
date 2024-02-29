<script lang="ts" setup>
import { CloseIcon, FullscreenIcon, HelpBoxIcon } from 'mdi-vue3';

defineProps<{
    title: string;
    message: string;
}>();

const open = ref(true);
</script>

<template>
    <div class="relative">
        <div
            v-if="open"
            class="w-52 drop-shadow-lg text-xs bg-white absolute select-none bottom-6 right-2 text-black border border-black px-2 py-1 rounded-lg rounded-br-none"
        >
            <p class="font-bold flex gap-2">
                <HelpBoxIcon class="w-5" aria-hidden="true" />
                {{ title }}
                <CloseIcon
                    class="ml-auto w-4 h-4 cursor-pointer opacity-30 hover:opacity-60"
                    aria-hidden="true"
                    @click="open = false"
                />
            </p>
            <p class="w-fit">{{ message }}</p>
        </div>
        <FullscreenIcon class="w-4 fill-" @click="open = !open" />
    </div>
</template>
