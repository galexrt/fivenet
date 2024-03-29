<script lang="ts" setup>
import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue';
import { RpcError } from '@protobuf-ts/runtime-rpc';
import { max, min, required } from '@vee-validate/rules';
import { useThrottleFn } from '@vueuse/core';
import { CloseIcon, HoopHouseIcon, LoadingIcon } from 'mdi-vue3';
import { defineRule } from 'vee-validate';
import { dispatchStatusToBGColor, dispatchStatuses } from '~/components/centrum/helpers';
import IDCopyBadge from '~/components/partials/IDCopyBadge.vue';
import { useCentrumStore } from '~/store/centrum';
import { StatusDispatch } from '~~/gen/ts/resources/centrum/dispatches';

const props = defineProps<{
    open: boolean;
    dispatchId: string;
    status?: StatusDispatch;
}>();

const emit = defineEmits<{
    (e: 'close'): void;
}>();

const { $grpc } = useNuxtApp();

const centrumStore = useCentrumStore();
const { settings } = storeToRefs(centrumStore);

const status: number = props.status ?? StatusDispatch.NEW;

interface FormData {
    status: number;
    code?: string;
    reason?: string;
}

async function updateDispatchStatus(dispatchId: string, values: FormData): Promise<void> {
    try {
        const call = $grpc.getCentrumClient().updateDispatchStatus({
            dispatchId,
            status: values.status,
            code: values.code,
            reason: values.reason,
        });
        await call;

        emit('close');

        setFieldValue('status', values.status.valueOf());
    } catch (e) {
        $grpc.handleError(e as RpcError);
        throw e;
    }
}

defineRule('required', required);
defineRule('min', min);
defineRule('max', max);

const { handleSubmit, meta, setFieldValue } = useForm<FormData>({
    validationSchema: {
        status: { required: true },
        code: { required: false },
        reason: { required: false, min: 3, max: 255 },
    },
    initialValues: {
        status,
    },
    validateOnMount: true,
});

const canSubmit = ref(true);
const onSubmit = handleSubmit(
    async (values): Promise<void> =>
        await updateDispatchStatus(props.dispatchId, values).finally(() => setTimeout(() => (canSubmit.value = true), 400)),
);
const onSubmitThrottle = useThrottleFn(async (e) => {
    canSubmit.value = false;
    await onSubmit(e);
}, 1000);

watch(props, () => {
    if (props.status) {
        setFieldValue('status', props.status.valueOf());
    }
});

function updateReasonField(value: string): void {
    setFieldValue('reason', value);
}
</script>

<template>
    <TransitionRoot as="template" :show="open">
        <Dialog as="div" class="relative z-30" @close="$emit('close')">
            <div class="fixed inset-0" />

            <div class="fixed inset-0 overflow-hidden">
                <div class="absolute inset-0 overflow-hidden">
                    <div class="pointer-events-none fixed inset-y-0 right-0 flex max-w-2xl pl-10 sm:pl-16">
                        <TransitionChild
                            as="template"
                            enter="transform transition ease-in-out duration-100 sm:duration-200"
                            enter-from="translate-x-full"
                            enter-to="translate-x-0"
                            leave="transform transition ease-in-out duration-100 sm:duration-200"
                            leave-from="translate-x-0"
                            leave-to="translate-x-full"
                        >
                            <DialogPanel class="pointer-events-auto w-screen max-w-3xl">
                                <form
                                    class="flex h-full flex-col divide-y divide-gray-200 bg-primary-900 shadow-xl"
                                    @submit.prevent="onSubmitThrottle"
                                >
                                    <div class="h-0 flex-1 overflow-y-auto">
                                        <div class="bg-primary-700 px-4 py-6 sm:px-6">
                                            <div class="flex items-center justify-between">
                                                <DialogTitle class="inline-flex text-base font-semibold leading-6 text-neutral">
                                                    {{ $t('components.centrum.update_dispatch_status.title') }}:
                                                    <IDCopyBadge :id="dispatchId" class="ml-2" prefix="DSP" />
                                                </DialogTitle>
                                                <div class="ml-3 flex h-7 items-center">
                                                    <button
                                                        type="button"
                                                        class="rounded-md bg-gray-100 text-gray-500 hover:text-gray-400 focus:outline-none focus:ring-2 focus:ring-neutral"
                                                        @click="$emit('close')"
                                                    >
                                                        <span class="sr-only">{{ $t('common.close') }}</span>
                                                        <CloseIcon class="h-5 w-5" aria-hidden="true" />
                                                    </button>
                                                </div>
                                            </div>
                                        </div>
                                        <div class="flex flex-1 flex-col justify-between">
                                            <div class="divide-y divide-gray-200 px-4 sm:px-6">
                                                <div class="mt-1">
                                                    <dl class="divide-y divide-neutral/10 border-b border-neutral/10">
                                                        <div class="px-4 py-3 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
                                                            <dt class="text-sm font-medium leading-6 text-neutral">
                                                                <label
                                                                    for="status"
                                                                    class="block text-sm font-medium leading-6 text-neutral"
                                                                >
                                                                    {{ $t('common.status') }}
                                                                </label>
                                                            </dt>
                                                            <dd
                                                                class="mt-1 text-sm leading-6 text-gray-300 sm:col-span-2 sm:mt-0"
                                                            >
                                                                <VeeField
                                                                    v-slot="{ field }"
                                                                    name="status"
                                                                    as="div"
                                                                    class="grid w-full grid-cols-2 gap-0.5"
                                                                    :placeholder="$t('common.status')"
                                                                    :label="$t('common.status')"
                                                                >
                                                                    <button
                                                                        v-for="(item, idx) in dispatchStatuses"
                                                                        :key="item.name"
                                                                        type="button"
                                                                        class="bg-primary group my-0.5 flex w-full flex-col items-center rounded-md p-1.5 text-xs font-medium text-neutral hover:bg-primary-100/10 hover:text-neutral hover:transition-all"
                                                                        :class="[
                                                                            idx >= dispatchStatuses.length - 1
                                                                                ? 'col-span-2'
                                                                                : '',
                                                                            item.class,
                                                                            field.value == item.status
                                                                                ? 'disabled bg-base-500 hover:bg-base-400'
                                                                                : item.status
                                                                                  ? dispatchStatusToBGColor(item.status)
                                                                                  : item.class,
                                                                            ,
                                                                        ]"
                                                                        :disabled="field.value == item.status"
                                                                        @click="
                                                                            setFieldValue('status', item.status?.valueOf() ?? 0)
                                                                        "
                                                                    >
                                                                        <component
                                                                            :is="item.icon ?? HoopHouseIcon"
                                                                            class="h-5 w-5 shrink-0 text-base-100 group-hover:text-neutral"
                                                                            aria-hidden="true"
                                                                        />
                                                                        <span class="mt-1">
                                                                            {{
                                                                                item.status
                                                                                    ? $t(
                                                                                          `enums.centrum.StatusDispatch.${
                                                                                              StatusDispatch[item.status ?? 0]
                                                                                          }`,
                                                                                      )
                                                                                    : $t(item.name)
                                                                            }}
                                                                        </span>
                                                                    </button>
                                                                </VeeField>
                                                                <VeeErrorMessage
                                                                    name="status"
                                                                    as="p"
                                                                    class="mt-2 text-sm text-error-400"
                                                                />
                                                            </dd>
                                                        </div>
                                                        <div class="px-4 py-3 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
                                                            <dt class="text-sm font-medium leading-6 text-neutral">
                                                                <label
                                                                    for="code"
                                                                    class="block text-sm font-medium leading-6 text-neutral"
                                                                >
                                                                    {{ $t('common.code') }}
                                                                </label>
                                                            </dt>
                                                            <dd
                                                                class="mt-1 text-sm leading-6 text-gray-300 sm:col-span-2 sm:mt-0"
                                                            >
                                                                <VeeField
                                                                    type="text"
                                                                    name="code"
                                                                    class="block w-full rounded-md border-0 bg-base-700 py-1.5 text-neutral placeholder:text-accent-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                                                    :placeholder="$t('common.code')"
                                                                    :label="$t('common.code')"
                                                                    @focusin="focusTablet(true)"
                                                                    @focusout="focusTablet(false)"
                                                                />
                                                                <VeeErrorMessage
                                                                    name="code"
                                                                    as="p"
                                                                    class="mt-2 text-sm text-error-400"
                                                                />
                                                            </dd>
                                                        </div>
                                                        <div class="px-4 py-3 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
                                                            <dt class="text-sm font-medium leading-6 text-neutral">
                                                                <label
                                                                    for="reason"
                                                                    class="block text-sm font-medium leading-6 text-neutral"
                                                                >
                                                                    {{ $t('common.reason') }}
                                                                </label>
                                                            </dt>
                                                            <dd
                                                                class="mt-1 text-sm leading-6 text-gray-300 sm:col-span-2 sm:mt-0"
                                                            >
                                                                <VeeField
                                                                    type="text"
                                                                    name="reason"
                                                                    class="block w-full rounded-md border-0 bg-base-700 py-1.5 text-neutral placeholder:text-accent-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                                                    :placeholder="$t('common.reason')"
                                                                    :label="$t('common.reason')"
                                                                    @focusin="focusTablet(true)"
                                                                    @focusout="focusTablet(false)"
                                                                />
                                                                <VeeErrorMessage
                                                                    name="reason"
                                                                    as="p"
                                                                    class="mt-2 text-sm text-error-400"
                                                                />
                                                            </dd>
                                                        </div>

                                                        <div
                                                            v-if="
                                                                settings?.predefinedStatus &&
                                                                settings?.predefinedStatus.dispatchStatus.length > 0
                                                            "
                                                            class="px-4 py-3 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0"
                                                        >
                                                            <dt class="text-sm font-medium leading-6 text-neutral">
                                                                <label
                                                                    for="dispatchStatus"
                                                                    class="block text-sm font-medium leading-6 text-neutral"
                                                                >
                                                                    {{ $t('common.predefined', 2) }}
                                                                    {{ $t('common.reason', 2) }}
                                                                </label>
                                                            </dt>
                                                            <dd
                                                                class="mt-1 text-sm leading-6 text-gray-300 sm:col-span-2 sm:mt-0"
                                                            >
                                                                <select
                                                                    name="dispatchStatus"
                                                                    class="mt-1 block w-full rounded-md border-0 bg-base-700 py-1.5 text-neutral placeholder:text-accent-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                                                    @focusin="focusTablet(true)"
                                                                    @focusout="focusTablet(false)"
                                                                    @change="
                                                                        updateReasonField(
                                                                            ($event.target as HTMLSelectElement).value,
                                                                        )
                                                                    "
                                                                >
                                                                    <option value=""></option>
                                                                    <option
                                                                        v-for="(preStatus, idx) in settings?.predefinedStatus
                                                                            .dispatchStatus"
                                                                        :key="idx"
                                                                        :value="preStatus"
                                                                    >
                                                                        {{ preStatus }}
                                                                    </option>
                                                                </select>
                                                            </dd>
                                                        </div>
                                                    </dl>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="flex flex-shrink-0 justify-end px-4 py-4">
                                        <span class="isolate inline-flex w-full rounded-md pr-4 shadow-sm">
                                            <button
                                                type="submit"
                                                class="relative flex w-full items-center rounded-l-md px-3.5 py-2.5 text-sm font-semibold text-neutral"
                                                :disabled="!meta.valid || !canSubmit"
                                                :class="[
                                                    !meta.valid || !canSubmit
                                                        ? 'disabled bg-base-500 hover:bg-base-400 focus-visible:outline-base-500'
                                                        : 'bg-primary-500 hover:bg-primary-400 focus-visible:outline-primary-500',
                                                ]"
                                            >
                                                <template v-if="!canSubmit">
                                                    <LoadingIcon class="mr-2 h-5 w-5 animate-spin" aria-hidden="true" />
                                                </template>
                                                {{ $t('common.update') }}
                                            </button>
                                            <button
                                                type="button"
                                                class="relative -ml-px inline-flex w-full items-center rounded-r-md bg-neutral px-3 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-200 hover:text-gray-900"
                                                @click="$emit('close')"
                                            >
                                                {{ $t('common.close', 1) }}
                                            </button>
                                        </span>
                                    </div>
                                </form>
                            </DialogPanel>
                        </TransitionChild>
                    </div>
                </div>
            </div>
        </Dialog>
    </TransitionRoot>
</template>
