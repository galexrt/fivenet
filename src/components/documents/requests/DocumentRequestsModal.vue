<script lang="ts" setup>
import { RpcError } from '@protobuf-ts/runtime-rpc';
import {
    Dialog,
    DialogPanel,
    DialogTitle,
    Listbox,
    ListboxButton,
    ListboxOption,
    ListboxOptions,
    TransitionChild,
    TransitionRoot,
} from '@headlessui/vue';
import { max, min, required } from '@vee-validate/rules';
import { useThrottleFn } from '@vueuse/core';
import { CheckIcon, ChevronDownIcon, CloseIcon, LoadingIcon } from 'mdi-vue3';
import { defineRule } from 'vee-validate';
import { DocActivityType } from '~~/gen/ts/resources/documents/activity';
import DocumentRequestsList from '~/components/documents/requests/DocumentRequestsList.vue';
import type { DocumentShort } from '~~/gen/ts/resources/documents/documents';
import { useAuthStore } from '~/store/auth';
import { useNotificatorStore } from '~/store/notificator';
import { AccessLevel, type DocumentAccess } from '~~/gen/ts/resources/documents/access';
import { checkDocAccess } from '~/components/documents/helpers';

const props = defineProps<{
    open: boolean;
    access: DocumentAccess;
    doc: DocumentShort;
}>();

defineEmits<{
    (e: 'close'): void;
    (e: 'refresh'): void;
}>();

const { $grpc } = useNuxtApp();

const authStore = useAuthStore();
const { activeChar } = storeToRefs(authStore);

const notifications = useNotificatorStore();

type RequestType = { key: DocActivityType; attrKey: string };
const requestTypes = [
    { key: props.doc.closed ? DocActivityType.REQUESTED_OPENING : DocActivityType.REQUESTED_CLOSURE, attrKey: 'Closure' },
    { key: DocActivityType.REQUESTED_UPDATE, attrKey: 'Update' },
    { key: DocActivityType.REQUESTED_OWNER_CHANGE, attrKey: 'OwnerChange' },
    { key: DocActivityType.REQUESTED_DELETION, attrKey: 'Deletion' },
] as RequestType[];

const availableRequestTypes = computed<RequestType[]>(() =>
    requestTypes.filter((rt) => attr('DocStoreService.CreateDocumentReq', 'Types', rt.attrKey)),
);

const selectedRequestType = ref<RequestType | undefined>(availableRequestTypes.value.at(0));

interface FormData {
    reason?: string;
}

async function createDocumentRequest(values: FormData): Promise<void> {
    if (selectedRequestType.value === undefined) {
        return;
    }

    try {
        const call = $grpc.getDocStoreClient().createDocumentReq({
            documentId: props.doc.id,
            requestType: selectedRequestType.value.key,
            reason: values.reason,
        });
        await call;

        notifications.dispatchNotification({
            title: { key: 'notifications.docstore.requests.created.title' },
            content: { key: 'notifications.docstore.requests.created.content' },
            type: 'success',
        });
    } catch (e) {
        $grpc.handleError(e as RpcError);
        throw e;
    }
}

defineRule('required', required);
defineRule('max', max);
defineRule('min', min);

const { handleSubmit, meta } = useForm<FormData>({
    validationSchema: {
        reason: { required: true, min: 3, max: 255 },
    },
    validateOnMount: true,
});

const canCreate =
    props.doc.creatorId !== activeChar.value?.userId &&
    availableRequestTypes.value.length > 0 &&
    can('DocStoreService.CreateDocumentReq') &&
    checkDocAccess(props.access, props.doc.creator, AccessLevel.VIEW);

const canSubmit = ref(true);
const onSubmit = handleSubmit(
    async (values): Promise<void> =>
        await createDocumentRequest(values).finally(() => setTimeout(() => (canSubmit.value = true), 400)),
);
const onSubmitThrottle = useThrottleFn(async (e) => {
    canSubmit.value = false;
    await onSubmit(e);
}, 1000);
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
                            class="relative w-full transform overflow-hidden rounded-lg bg-base-800 px-4 pb-4 pt-5 text-left text-neutral transition-all sm:my-8 sm:min-h-[28rem] sm:max-w-5xl sm:p-6"
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
                                {{ $t('common.request', 2) }}
                            </DialogTitle>

                            <template v-if="canCreate">
                                <div class="my-2 space-y-24">
                                    <div class="form-control flex-1">
                                        <label for="reason" class="block text-sm font-medium leading-6 text-neutral">
                                            {{ $t('common.reason') }}
                                        </label>
                                        <VeeField
                                            type="text"
                                            name="reason"
                                            class="block w-full rounded-md border-0 bg-base-700 py-1.5 text-neutral placeholder:text-accent-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                            :placeholder="$t('common.reason')"
                                            :label="$t('common.reason')"
                                            @focusin="focusTablet(true)"
                                            @focusout="focusTablet(false)"
                                        />
                                        <VeeErrorMessage name="reason" as="p" class="mt-2 text-sm text-error-400" />
                                    </div>
                                </div>
                                <div class="my-2 space-y-20">
                                    <div class="form-control flex-1">
                                        <label for="requestsType" class="block text-sm font-medium leading-6 text-neutral">
                                            {{ $t('common.type', 2) }}
                                        </label>
                                        <VeeField
                                            type="text"
                                            name="requestsType"
                                            class="block w-full rounded-md border-0 bg-base-700 py-1.5 text-neutral placeholder:text-accent-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                            :placeholder="$t('common.type', 2)"
                                            :label="$t('common.type', 2)"
                                            @focusin="focusTablet(true)"
                                            @focusout="focusTablet(false)"
                                        >
                                            <Listbox v-model="selectedRequestType" as="div">
                                                <div class="relative">
                                                    <ListboxButton
                                                        class="block w-full rounded-md border-0 bg-base-700 py-1.5 pl-3 text-left text-neutral placeholder:text-accent-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                                    >
                                                        <span class="block truncate">
                                                            {{
                                                                $t(
                                                                    `enums.docstore.DocActivityType.${DocActivityType[selectedRequestType?.key ?? 0]}`,
                                                                    2,
                                                                )
                                                            }}
                                                        </span>
                                                        <span
                                                            class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2"
                                                        >
                                                            <ChevronDownIcon class="h-5 w-5 text-gray-400" aria-hidden="true" />
                                                        </span>
                                                    </ListboxButton>

                                                    <transition
                                                        leave-active-class="transition duration-100 ease-in"
                                                        leave-from-class="opacity-100"
                                                        leave-to-class="opacity-0"
                                                    >
                                                        <ListboxOptions
                                                            class="absolute z-10 mt-1 max-h-44 w-full overflow-auto rounded-md bg-base-700 py-1 text-base sm:text-sm"
                                                        >
                                                            <ListboxOption
                                                                v-for="requestType in availableRequestTypes"
                                                                :key="requestType.key"
                                                                v-slot="{ active, selected }"
                                                                as="template"
                                                                :value="requestType"
                                                            >
                                                                <li
                                                                    :class="[
                                                                        active ? 'bg-primary-500' : '',
                                                                        'relative cursor-default select-none py-2 pl-8 pr-4 text-neutral',
                                                                    ]"
                                                                >
                                                                    <span
                                                                        :class="[
                                                                            selected ? 'font-semibold' : 'font-normal',
                                                                            'block truncate',
                                                                        ]"
                                                                    >
                                                                        {{
                                                                            $t(
                                                                                `enums.docstore.DocActivityType.${DocActivityType[requestType.key]}`,
                                                                                2,
                                                                            )
                                                                        }}
                                                                    </span>

                                                                    <span
                                                                        v-if="selected"
                                                                        :class="[
                                                                            active ? 'text-neutral' : 'text-primary-500',
                                                                            'absolute inset-y-0 left-0 flex items-center pl-1.5',
                                                                        ]"
                                                                    >
                                                                        <CheckIcon class="h-5 w-5" aria-hidden="true" />
                                                                    </span>
                                                                </li>
                                                            </ListboxOption>
                                                        </ListboxOptions>
                                                    </transition>
                                                </div>
                                            </Listbox>
                                        </VeeField>
                                        <VeeErrorMessage name="requestsType" as="p" class="mt-2 text-sm text-error-400" />
                                    </div>
                                </div>
                            </template>

                            <div class="absolute bottom-0 left-0 flex w-full">
                                <button
                                    type="button"
                                    class="rounded-bd flex-1 px-3.5 py-2.5 text-sm font-semibold bg-neutral text-gray-900 hover:bg-gray-200"
                                    @click="$emit('close')"
                                >
                                    {{ $t('common.close', 1) }}
                                </button>
                                <button
                                    v-if="canCreate"
                                    type="button"
                                    class="rounded-bd flex flex-1 justify-center px-3.5 py-2.5 text-sm font-semibold text-neutral"
                                    :disabled="!meta.valid || !canSubmit"
                                    :class="[
                                        !meta.valid || !canSubmit
                                            ? 'disabled bg-base-500 hover:bg-base-400 focus-visible:outline-base-500'
                                            : 'bg-primary-500 hover:bg-primary-400 focus-visible:outline-primary-500',
                                    ]"
                                    @click="onSubmitThrottle($event)"
                                >
                                    <template v-if="!canSubmit">
                                        <LoadingIcon class="mr-2 h-5 w-5 animate-spin" aria-hidden="true" />
                                    </template>
                                    {{ $t('common.add') }}
                                </button>
                            </div>

                            <DocumentRequestsList :doc="doc" :access="access" @refresh="$emit('refresh')" />
                        </DialogPanel>
                    </TransitionChild>
                </div>
            </div>
        </Dialog>
    </TransitionRoot>
</template>
