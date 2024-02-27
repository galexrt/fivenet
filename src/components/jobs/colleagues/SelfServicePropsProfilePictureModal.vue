<script lang="ts" setup>
import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue';
import { RpcError } from '@protobuf-ts/runtime-rpc';
import { useThrottleFn } from '@vueuse/core';
import { CloseIcon, LoadingIcon } from 'mdi-vue3';
import ProfilePictureImg from '~/components/partials/citizens/ProfilePictureImg.vue';
import { useAuthStore } from '~/store/auth';
import { useNotificatorStore } from '~/store/notificator';
import type { SetProfilePictureRequest } from '~~/gen/ts/services/citizenstore/citizenstore';

defineProps<{
    open: boolean;
}>();

const emit = defineEmits<{
    (e: 'close'): void;
}>();

const { $grpc } = useNuxtApp();

const authStore = useAuthStore();
const { activeChar } = storeToRefs(authStore);

const notifications = useNotificatorStore();

interface FormData {
    reset?: boolean;
}

async function setProfilePicture(values: FormData): Promise<void> {
    const req = {} as SetProfilePictureRequest;
    if (!values.reset) {
        if (!fileUploadRef.value) {
            return;
        }
        if (!fileUploadRef.value.files || fileUploadRef.value.files.length <= 0) {
            return;
        }
        // File too big
        if (fileUploadRef.value.files[0].size > 2097152) {
            return;
        }

        req.avatar = { data: new Uint8Array(await fileUploadRef.value.files[0].arrayBuffer()) };
    } else {
        req.avatar = { data: new Uint8Array(), delete: true };
    }

    try {
        const call = $grpc.getCitizenStoreClient().setProfilePicture(req);
        const { response } = await call;

        if (activeChar.value) {
            activeChar.value.avatar = response.avatar;
        }

        notifications.dispatchNotification({
            title: { key: 'notifications.action_successfull.title', parameters: {} },
            content: { key: 'notifications.action_successfull.content', parameters: {} },
            type: 'success',
        });

        emit('close');
    } catch (e) {
        $grpc.handleError(e as RpcError);
        throw e;
    }
}

const { handleSubmit, meta, setFieldValue } = useForm<FormData>({
    validationSchema: {
        avatar: {},
    },
    validateOnMount: true,
});

const canSubmit = ref(true);
const onSubmit = handleSubmit(
    async (values): Promise<void> =>
        await setProfilePicture(values).finally(() => setTimeout(() => (canSubmit.value = true), 400)),
);
const onSubmitThrottle = useThrottleFn(async (e) => {
    canSubmit.value = false;
    await onSubmit(e);
    setFieldValue('reset', false);
}, 1000);

const fileUploadRef = ref<HTMLInputElement | null>(null);
</script>

<template>
    <TransitionRoot as="template" :show="open">
        <Dialog as="div" class="relative z-30" @close="$emit('close')">
            <TransitionChild
                as="template"
                enter="ease-out duration-300"
                enter-from="opacity-0"
                enter-to="opacity-100"
                leave="ease-in duration-200"
                leave-from="opacity-100"
                leave-to="opacity-0"
            >
                <div class="fixed inset-0 bg-base-900 bg-opacity-75 transition-opacity" />
            </TransitionChild>

            <div class="fixed inset-0 z-30 overflow-y-auto">
                <div class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
                    <TransitionChild
                        as="template"
                        enter="ease-out duration-300"
                        enter-from="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                        enter-to="opacity-100 translate-y-0 sm:scale-100"
                        leave="ease-in duration-200"
                        leave-from="opacity-100 translate-y-0 sm:scale-100"
                        leave-to="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                    >
                        <DialogPanel
                            class="relative h-96 w-full transform overflow-hidden rounded-lg bg-base-800 px-4 pb-4 pt-5 text-left text-neutral transition-all sm:my-8 sm:max-w-2xl sm:p-6"
                        >
                            <div class="absolute right-0 top-0 block pr-4 pt-4">
                                <button
                                    type="button"
                                    class="rounded-md bg-neutral text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2"
                                    @click="$emit('close')"
                                >
                                    <span class="sr-only">{{ $t('common.close') }}</span>
                                    <CloseIcon class="h-5 w-5" aria-hidden="true" />
                                </button>
                            </div>
                            <DialogTitle as="h3" class="text-base font-semibold leading-6">
                                {{ $t('components.jobs.self_service.set_profile_picture') }}
                            </DialogTitle>
                            <form @submit.prevent="onSubmitThrottle">
                                <div class="my-2 space-y-24">
                                    <div class="form-control flex-1">
                                        <label for="job" class="block text-sm font-medium leading-6 text-neutral">
                                            {{ $t('common.avatar') }}
                                        </label>
                                        <VeeField
                                            v-slot="{ handleChange, handleBlur }"
                                            type="file"
                                            name="avatar"
                                            class="block w-full rounded-md border-0 bg-base-700 py-1.5 text-neutral placeholder:text-accent-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                            :placeholder="$t('common.image')"
                                            :label="$t('common.image')"
                                            @focusin="focusTablet(true)"
                                            @focusout="focusTablet(false)"
                                        >
                                            <input
                                                ref="fileUploadRef"
                                                type="file"
                                                accept="image/png,image/jpeg"
                                                class="block w-full rounded-md border-0 bg-base-700 py-1.5 text-neutral placeholder:text-accent-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                                @change="handleChange"
                                                @blur="handleBlur"
                                            />
                                        </VeeField>
                                        <VeeErrorMessage name="avatar" as="p" class="mt-2 text-sm text-error-400" />
                                    </div>
                                </div>
                                <div class="flex-1 flex items-center">
                                    <ProfilePictureImg
                                        class="mx-auto my-auto"
                                        :url="activeChar?.avatar?.url"
                                        :name="`${activeChar?.firstname} ${activeChar?.lastname}`"
                                        size="huge"
                                    />
                                </div>
                                <div class="absolute bottom-0 left-0 flex w-full">
                                    <button
                                        type="button"
                                        class="rounded-bd flex-1 bg-neutral text-gray-900 hover:bg-gray-200 px-3.5 py-2.5 text-sm font-semibold"
                                        @click="$emit('close')"
                                    >
                                        {{ $t('common.close', 1) }}
                                    </button>
                                    <button
                                        type="button"
                                        class="rounded-bd flex flex-1 justify-center px-3.5 py-2.5 text-sm font-semibold text-neutral"
                                        :disabled="!meta.valid || !canSubmit"
                                        :class="[
                                            !meta.valid || !canSubmit
                                                ? 'disabled bg-base-500 hover:bg-base-400 focus-visible:outline-base-500'
                                                : 'bg-error-500 hover:bg-error-400 focus-visible:outline-error-500',
                                        ]"
                                        @click="
                                            setFieldValue('reset', true);
                                            onSubmitThrottle($event);
                                        "
                                    >
                                        <template v-if="!canSubmit">
                                            <LoadingIcon class="mr-2 h-5 w-5 animate-spin" />
                                        </template>
                                        {{ $t('common.reset') }}
                                    </button>
                                    <button
                                        type="submit"
                                        class="rounded-bd flex flex-1 justify-center px-3.5 py-2.5 text-sm font-semibold text-neutral"
                                        :disabled="!meta.valid || !canSubmit"
                                        :class="[
                                            !meta.valid || !canSubmit
                                                ? 'disabled bg-base-500 hover:bg-base-400 focus-visible:outline-base-500'
                                                : 'bg-primary-500 hover:bg-primary-400 focus-visible:outline-primary-500',
                                        ]"
                                    >
                                        <template v-if="!canSubmit">
                                            <LoadingIcon class="mr-2 h-5 w-5 animate-spin" />
                                        </template>
                                        {{ $t('common.save') }}
                                    </button>
                                </div>
                            </form>
                        </DialogPanel>
                    </TransitionChild>
                </div>
            </div>
        </Dialog>
    </TransitionRoot>
</template>