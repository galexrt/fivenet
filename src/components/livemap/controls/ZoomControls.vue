<script lang="ts" setup>
import { LControl } from '@vue-leaflet/vue-leaflet';
import { MinusThickIcon, PlusThickIcon } from 'mdi-vue3';
import { useLivemapStore } from '~/store/livemap';

const livemapStore = useLivemapStore();
const { zoom } = storeToRefs(livemapStore);
</script>

<template>
    <LControl position="topleft">
        <div class="inline-flex w-full flex-col">
            <button
                type="button"
                class="inset-0 inline-flex items-center justify-center rounded-md border-2 border-black/20 bg-neutral bg-clip-padding p-1.5 text-black text-primary-500 hover:bg-[#f4f4f4] focus:outline-none"
                @click="if (zoom + 1 <= 7) zoom++;"
            >
                <PlusThickIcon class="h-5 w-5" aria-hidden="true" />
            </button>

            <div
                class="inset-0 inline-flex items-center justify-center rounded-md border-2 border-black/20 bg-neutral bg-clip-padding text-black text-primary-500 focus:outline-none text-base select-none"
            >
                {{ zoom }}
            </div>

            <button
                type="button"
                class="inset-0 inline-flex items-center justify-center rounded-md border-2 border-black/20 bg-neutral bg-clip-padding p-1.5 text-black text-primary-500 hover:bg-[#f4f4f4] focus:outline-none"
                @click="if (zoom - 1 >= 1) zoom--;"
            >
                <MinusThickIcon class="h-5 w-5" aria-hidden="true" />
            </button>
        </div>
    </LControl>
</template>
