<script lang="ts" setup>
import { Float } from '@headlessui-float/vue';
import { Popover, PopoverButton, PopoverPanel } from '@headlessui/vue';
import { DispatchStatus, StatusDispatch } from '~~/gen/ts/resources/centrum/dispatches';
import DispatchDetailsByID from '~/components/centrum/dispatches//DispatchDetailsByID.vue';
import { CarEmergencyIcon } from 'mdi-vue3';

defineProps<{
    status: DispatchStatus | undefined;
    textClass?: unknown;
    buttonClass?: unknown;
}>();

const open = ref(false);
</script>

<template>
    <template v-if="!status">
        <span class="inline-flex items-center">
            <slot name="before" />
            <span>N/A</span>
            <slot name="after" />
        </span>
    </template>
    <template v-else>
        <DispatchDetailsByID :open="open" :dispatch-id="status.dispatchId" @close="open = false" />

        <Popover class="relative">
            <Float
                portal
                auto-placement
                :offset="35"
                enter="transition duration-150 ease-out"
                enter-from="scale-95 opacity-0"
                enter-to="scale-100 opacity-100"
                leave="transition duration-100 ease-in"
                leave-from="scale-100 opacity-100"
                leave-to="scale-95 opacity-0"
            >
                <PopoverButton class="inline-flex items-center" :class="buttonClass">
                    <slot name="before" />
                    <span :class="textClass"> DSP-{{ status.dispatchId }} </span>
                    <slot name="after" />
                </PopoverButton>

                <PopoverPanel
                    class="absolute z-5 w-72 min-w-fit max-w-[18rem] rounded-lg border border-gray-600 bg-gray-800 text-sm text-gray-400 shadow-sm transition-opacity"
                >
                    <div class="inline-flex p-3 gap-2">
                        <div class="mb-2 flex items-center gap-2">
                            <button
                                type="button"
                                :title="$t('common.detail', 2)"
                                class="inline-flex items-center text-primary-500 hover:text-primary-400"
                                @click="open = true"
                            >
                                <CarEmergencyIcon class="h-5 w-5" aria-hidden="true" />
                                <span class="ml-1">{{ $t('common.detail', 2) }}</span>
                            </button>
                        </div>
                        <div>
                            <p class="text-base font-semibold leading-none text-gray-900 dark:text-neutral">
                                DSP-{{ status.dispatchId }}
                            </p>
                            <p class="inline-flex items-center justify-center text-sm font-normal">
                                <span class="font-semibold"> {{ $t('common.status') }} </span>:
                                {{ $t(`enums.centrum.StatusDispatch.${StatusDispatch[status.status ?? 0]}`) }}
                            </p>
                        </div>
                    </div>
                </PopoverPanel>
            </Float>
        </Popover>
    </template>
</template>
