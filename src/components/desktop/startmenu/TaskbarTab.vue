<script lang="ts" setup>
import { type DefineComponent } from 'vue';

defineProps<{
    active: boolean;
    name: string;
    alt?: string;
    icon?: DefineComponent;
    iconClass?: string;
}>();
</script>

<template>
    <div
        :class="{ '!bg-[#1658dd] !border-[#082875]': active }"
        class="w-40 h-[80%] tab mt-px mr-1 px-2 gap-1 border-t border-y border-[#255be1] rounded-sm bg-[#3980f4] hover:bg-[#1b50b8] hover:border-[#082875] flex items-center justify-start"
    >
        <component :is="icon" draggable="false" :alt="alt ?? name" class="w-4" :class="iconClass" />
        <p class="text-xs text-white">{{ name }}</p>
    </div>
</template>
