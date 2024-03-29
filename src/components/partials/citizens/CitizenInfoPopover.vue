<script lang="ts" setup>
import { Float } from '@headlessui-float/vue';
import { Popover, PopoverButton, PopoverPanel } from '@headlessui/vue';
import { AccountIcon } from 'mdi-vue3';
import { type User, type UserShort } from '~~/gen/ts/resources/users/users';
import PhoneNumberBlock from '~/components/partials/citizens/PhoneNumberBlock.vue';
import { ClipboardUser } from '~/store/clipboard';
import ProfilePictureImg from '~/components/partials/citizens/ProfilePictureImg.vue';

withDefaults(
    defineProps<{
        user: ClipboardUser | User | UserShort | undefined;
        noPopover?: boolean;
        textClass?: unknown;
        buttonClass?: unknown;
        showAvatar?: boolean;
    }>(),
    {
        textClass: '' as any,
        buttonClass: '' as any,
        showAvatar: undefined,
    },
);
</script>

<template>
    <template v-if="!user">
        <span class="inline-flex items-center">
            <slot name="before" />
            <span>N/A</span>
            <slot name="after" />
        </span>
    </template>
    <template v-else-if="noPopover">
        <span class="inline-flex items-center">
            <slot name="before" />
            <NuxtLink :to="{ name: 'citizens-id', params: { id: user.userId ?? 0 } }">
                {{ user.firstname }} {{ user.lastname }}
            </NuxtLink>
            <span v-if="user.phoneNumber">
                <PhoneNumberBlock v-if="user.phoneNumber" :number="user.phoneNumber" :hide-number="true" :show-label="false" />
            </span>
            <slot name="after" />
        </span>
    </template>
    <Popover v-else class="relative">
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
                <span :class="textClass"> {{ user.firstname }} {{ user.lastname }} </span>
                <slot name="after" />
            </PopoverButton>

            <PopoverPanel
                class="absolute z-5 w-72 min-w-fit max-w-[18rem] rounded-lg border border-gray-600 bg-gray-800 text-sm text-gray-400 shadow-sm transition-opacity"
            >
                <div class="inline-flex p-3">
                    <div>
                        <div
                            v-if="can('CitizenStoreService.ListCitizens') || user.phoneNumber"
                            class="mb-2 flex items-center gap-2"
                        >
                            <NuxtLink
                                v-if="can('CitizenStoreService.ListCitizens')"
                                :to="{ name: 'citizens-id', params: { id: user.userId ?? 0 } }"
                                class="inline-flex items-center text-primary-500 hover:text-primary-400"
                            >
                                <AccountIcon class="h-5 w-5" aria-hidden="true" />
                                <span class="ml-1">{{ $t('common.profile') }}</span>
                            </NuxtLink>
                            <PhoneNumberBlock
                                v-if="user.phoneNumber"
                                :number="user.phoneNumber"
                                :hide-number="true"
                                :show-label="true"
                            />
                        </div>
                        <div class="inline-flex flex-row gap-2">
                            <div v-if="showAvatar === undefined || showAvatar">
                                <ProfilePictureImg :url="user.avatar?.url" :name="`${user.firstname} ${user.lastname}`" />
                            </div>
                            <div>
                                <p class="text-base font-semibold leading-none text-gray-900 dark:text-neutral">
                                    <NuxtLink :to="{ name: 'citizens-id', params: { id: user.userId ?? 0 } }">
                                        {{ user.firstname }} {{ user.lastname }}
                                    </NuxtLink>
                                </p>
                                <p v-if="user.jobLabel" class="text-sm font-normal">
                                    <span class="font-semibold">{{ $t('common.job') }}:</span>
                                    {{ user.jobLabel }}
                                    <span v-if="(user.jobGrade ?? 0) > 0 && user.jobGradeLabel">
                                        ({{ user.jobGradeLabel }})</span
                                    >
                                </p>
                                <p v-if="user.dateofbirth" class="text-sm font-normal">
                                    <span class="font-semibold">{{ $t('common.date_of_birth') }}:</span>
                                    {{ user.dateofbirth }}
                                </p>
                            </div>
                        </div>
                    </div>
                </div>
            </PopoverPanel>
        </Float>
    </Popover>
</template>
