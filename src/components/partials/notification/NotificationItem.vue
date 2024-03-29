<script lang="ts" setup>
import { AlertCircleIcon, CheckCircleIcon, CloseCircleIcon, CloseIcon, InformationIcon } from 'mdi-vue3';
import { type Notification } from '~/composables/notifications/notifications';
import { useNotificatorStore } from '~/store/notificator';

const notifications = useNotificatorStore();
const { removeNotification } = notifications;

const props = defineProps<{
    notification: Notification;
}>();

async function closeNotification(id: string): Promise<void> {
    removeNotification(id);
}

if (props.notification.callback !== undefined) {
    props.notification.callback();
}
</script>

<template>
    <transition
        enter-active-class="transition duration-300 ease-out transform"
        enter-from-class="translate-y-2 opacity-0 sm:translate-y-0 sm:translate-x-2"
        enter-to-class="translate-y-0 opacity-100 sm:translate-x-0"
        leave-active-class="transition duration-100 ease-in"
        leave-from-class="opacity-100"
        leave-to-class="opacity-0"
    >
        <div
            v-if="notification"
            class="pointer-events-auto z-50 w-full max-w-sm overflow-hidden rounded-lg bg-base-800 text-neutral shadow-float"
        >
            <div class="p-4">
                <div class="flex items-start">
                    <div v-if="notification.type" class="my-auto w-8 flex-shrink-0">
                        <CheckCircleIcon v-if="notification.type === 'success'" class="text-success-400" aria-hidden="true" />
                        <InformationIcon v-else-if="notification.type === 'info'" class="text-info-400" aria-hidden="true" />
                        <AlertCircleIcon v-else-if="notification.type === 'warning'" class="text-warn-400" aria-hidden="true" />
                        <CloseCircleIcon v-else-if="notification.type === 'error'" class="text-error-400" aria-hidden="true" />
                    </div>
                    <div class="ml-3 w-0 flex-1 pt-0.5">
                        <p class="text-sm font-semibold">
                            {{ $t(notification.title.key, notification.title.parameters ?? {}) }}
                        </p>
                        <p class="mt-1 text-sm leading-5">
                            {{ $t(notification.content.key, notification.content.parameters ?? {}) }}
                        </p>
                        <button
                            v-if="notification.onClick !== undefined"
                            type="button"
                            class="mt-1 text-sm leading-5 underline"
                            @click="notification.onClick(notification.data)"
                        >
                            {{
                                notification.onClickText
                                    ? $t(notification.onClickText.key, notification.onClickText.parameters ?? {})
                                    : $t('common.click_here')
                            }}
                        </button>
                    </div>
                    <div class="ml-4 flex flex-shrink-0">
                        <button
                            type="button"
                            class="inline-flex text-neutral hover:text-base-300 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2"
                            @click="() => closeNotification(notification.id)"
                        >
                            <span class="sr-only">{{ $t('common.close') }}</span>
                            <CloseIcon class="h-5 w-5" aria-hidden="true" />
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </transition>
</template>
