<script lang="ts" setup>
import { AccountPlusIcon, AccountRemoveIcon, BriefcaseIcon, CoffeeIcon, HelpIcon, PlayIcon, StopIcon } from 'mdi-vue3';
import CitizenInfoPopover from '~/components/partials/citizens/CitizenInfoPopover.vue';
import Time from '~/components/partials/elements/Time.vue';
import { StatusUnit, UnitStatus } from '~~/gen/ts/resources/dispatch/units';
import UnitInfoPopover from '~/components/centrum/units/UnitInfoPopover.vue';

defineProps<{
    activityLength: number;
    item: UnitStatus;
    activityItemIdx: number;
}>();
</script>

<template>
    <li class="relative flex gap-x-2">
        <div
            :class="[
                activityItemIdx === activityLength - 1 ? 'h-6' : '-bottom-6',
                'absolute left-0 top-0 flex w-6 justify-center',
            ]"
        >
            <div class="w-px bg-gray-200" />
        </div>
        <template v-if="item.status === StatusUnit.USER_ADDED">
            <div class="relative flex h-6 w-6 flex-none items-center justify-center bg-gray-300 rounded-lg">
                <AccountPlusIcon class="h-6 w-6 text-primary-600" aria-hidden="true" />
            </div>
            <p class="flex-auto py-0.5 text-xs leading-5 text-gray-200 inline-flex flex-row justify-between">
                <span class="inline-flex items-center gap-1">
                    {{ $t('components.centrum.units.feed.item.USER_ADDED') }}
                    <UnitInfoPopover
                        v-if="item.unit"
                        text-class="font-medium text-gray-400 pl-1"
                        :unit="item.unit"
                        :initials-only="true"
                        :badge="true"
                    />
                </span>

                <CitizenInfoPopover v-if="item.user" text-class="font-medium text-gray-400 pl-1" :user="item.user" />
            </p>
            <span class="flex-none py-0.5 text-xs leading-5 text-gray-200">
                <Time :value="item.createdAt" :type="'compact'" />
            </span>
        </template>
        <template v-else-if="item.status === StatusUnit.USER_REMOVED">
            <div class="relative flex h-6 w-6 flex-none items-center justify-center bg-gray-300 rounded-lg">
                <AccountRemoveIcon class="h-6 w-6 text-primary-600" aria-hidden="true" />
            </div>
            <p class="flex-auto py-0.5 text-xs leading-5 text-gray-200 inline-flex flex-row justify-between">
                <span class="inline-flex items-center gap-1">
                    {{ $t('components.centrum.units.feed.item.USER_REMOVED') }}
                    <UnitInfoPopover
                        v-if="item.unit"
                        text-class="font-medium text-gray-400 pl-1"
                        :unit="item.unit"
                        :initials-only="true"
                        :badge="true"
                    />
                </span>

                <CitizenInfoPopover v-if="item.user" text-class="font-medium text-gray-400 pl-1" :user="item.user" />
            </p>
            <span class="flex-none py-0.5 text-xs leading-5 text-gray-200">
                <Time :value="item.createdAt" :type="'compact'" />
            </span>
        </template>
        <template v-else-if="item.status === StatusUnit.UNAVAILABLE">
            <div class="relative flex h-6 w-6 flex-none items-center justify-center bg-gray-300 rounded-lg">
                <StopIcon class="h-6 w-6 text-primary-600" aria-hidden="true" />
            </div>
            <p class="flex-auto py-0.5 text-xs leading-5 text-gray-200 inline-flex flex-row justify-between">
                <span class="inline-flex items-center gap-1">
                    {{ $t('components.centrum.units.feed.item.UNAVAILABLE') }}
                    <UnitInfoPopover
                        v-if="item.unit"
                        text-class="font-medium text-gray-400 pl-1"
                        :unit="item.unit"
                        :initials-only="true"
                        :badge="true"
                    />
                </span>

                <CitizenInfoPopover v-if="item.user" text-class="font-medium text-gray-400 pl-1" :user="item.user" />
            </p>
            <span class="flex-none py-0.5 text-xs leading-5 text-gray-200">
                <Time :value="item.createdAt" :type="'compact'" />
            </span>
        </template>
        <template v-else-if="item.status === StatusUnit.AVAILABLE">
            <div class="relative flex h-6 w-6 flex-none items-center justify-center bg-gray-300 rounded-lg">
                <PlayIcon class="h-6 w-6 text-primary-600" aria-hidden="true" />
            </div>
            <p class="flex-auto py-0.5 text-xs leading-5 text-gray-200 inline-flex flex-row justify-between">
                <span class="inline-flex items-center gap-1">
                    {{ $t('components.centrum.units.feed.item.AVAILABLE') }}
                    <UnitInfoPopover
                        v-if="item.unit"
                        text-class="font-medium text-gray-400 pl-1"
                        :unit="item.unit"
                        :initials-only="true"
                        :badge="true"
                    />
                </span>

                <CitizenInfoPopover v-if="item.user" text-class="font-medium text-gray-400 pl-1" :user="item.user" />
            </p>
            <span class="flex-none py-0.5 text-xs leading-5 text-gray-200">
                <Time :value="item.createdAt" :type="'compact'" />
            </span>
        </template>
        <template v-else-if="item.status === StatusUnit.ON_BREAK">
            <div class="relative flex h-6 w-6 flex-none items-center justify-center bg-gray-300 rounded-lg">
                <CoffeeIcon class="h-6 w-6 text-primary-600" aria-hidden="true" />
            </div>
            <p class="flex-auto py-0.5 text-xs leading-5 text-gray-200 inline-flex flex-row justify-between">
                <span class="inline-flex items-center gap-1">
                    {{ $t('components.centrum.units.feed.item.ON_BREAK') }}
                    <UnitInfoPopover
                        v-if="item.unit"
                        text-class="font-medium text-gray-400 pl-1"
                        :unit="item.unit"
                        :initials-only="true"
                        :badge="true"
                    />
                </span>

                <CitizenInfoPopover v-if="item.user" text-class="font-medium text-gray-400 pl-1" :user="item.user" />
            </p>
            <span class="flex-none py-0.5 text-xs leading-5 text-gray-200">
                <Time :value="item.createdAt" :type="'compact'" />
            </span>
        </template>
        <template v-else-if="item.status === StatusUnit.BUSY">
            <div class="relative flex h-6 w-6 flex-none items-center justify-center bg-gray-300 rounded-lg">
                <BriefcaseIcon class="h-6 w-6 text-primary-600" aria-hidden="true" />
            </div>
            <p class="flex-auto py-0.5 text-xs leading-5 text-gray-200 inline-flex flex-row justify-between">
                <span class="inline-flex items-center gap-1">
                    {{ $t('components.centrum.units.feed.item.BUSY') }}
                    <UnitInfoPopover
                        v-if="item.unit"
                        text-class="font-medium text-gray-400 pl-1"
                        :unit="item.unit"
                        :initials-only="true"
                        :badge="true"
                    />
                </span>

                <CitizenInfoPopover v-if="item.user" text-class="font-medium text-gray-400 pl-1" :user="item.user" />
            </p>
            <span class="flex-none py-0.5 text-xs leading-5 text-gray-200">
                <Time :value="item.createdAt" :type="'compact'" />
            </span>
        </template>
        <template v-else>
            <div class="relative flex h-6 w-6 flex-none items-center justify-center bg-gray-300 rounded-lg">
                <HelpIcon class="h-6 w-6 text-primary-600" aria-hidden="true" />
            </div>
            <p class="flex-auto py-0.5 text-xs leading-5 text-gray-200 inline-flex flex-row justify-between">
                <span class="inline-flex items-center gap-1">
                    {{ $t('components.centrum.units.feed.item.UNKNOWN') }}
                    <UnitInfoPopover
                        v-if="item.unit"
                        text-class="font-medium text-gray-400 pl-1"
                        :unit="item.unit"
                        :initials-only="true"
                        :badge="true"
                    />
                </span>

                <CitizenInfoPopover v-if="item.user" text-class="font-medium text-gray-400 pl-1" :user="item.user" />
            </p>
            <span class="flex-none py-0.5 text-xs leading-5 text-gray-200">
                <Time :value="item.createdAt" :type="'compact'" />
            </span>
        </template>
    </li>
</template>