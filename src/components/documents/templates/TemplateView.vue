<script lang="ts" setup>
import { RpcError } from '@protobuf-ts/runtime-rpc';
import { useConfirmDialog } from '@vueuse/core';
import DocumentAccessEntry from '~/components/documents/DocumentAccessEntry.vue';
import TemplatePreviewModal from '~/components/documents/templates/TemplatePreviewModal.vue';
import TemplateRequirementsList from '~/components/documents/templates/TemplateRequirementsList.vue';
import ConfirmDialog from '~/components/partials/ConfirmDialog.vue';
import DataErrorBlock from '~/components/partials/data/DataErrorBlock.vue';
import DataNoDataBlock from '~/components/partials/data/DataNoDataBlock.vue';
import DataPendingBlock from '~/components/partials/data/DataPendingBlock.vue';
import { useNotificatorStore } from '~/store/notificator';
import { AccessLevel } from '~~/gen/ts/resources/documents/access';
import { Template, TemplateRequirements } from '~~/gen/ts/resources/documents/templates';

const props = defineProps<{
    templateId: string;
}>();

const {
    data: template,
    pending,
    refresh,
    error,
} = useLazyAsyncData(`documents-template-${props.templateId}`, () => getTemplate());

const { $grpc } = useNuxtApp();

const notifications = useNotificatorStore();

const { t } = useI18n();

const reqs = ref<undefined | TemplateRequirements>();

async function getTemplate(): Promise<Template | undefined> {
    try {
        const call = $grpc.getDocStoreClient().getTemplate({
            templateId: props.templateId,
            render: false,
        });
        const { response } = await call;

        if (response.template?.schema) {
            reqs.value = response.template?.schema?.requirements;
        }

        return response.template!;
    } catch (e) {
        $grpc.handleError(e as RpcError);
        throw e;
    }
}

async function deleteTemplate(id: string): Promise<void> {
    try {
        await $grpc.getDocStoreClient().deleteTemplate({ id });

        notifications.dispatchNotification({
            title: { key: 'notifications.templates.deleted.title', parameters: {} },
            content: { key: 'notifications.templates.deleted.content', parameters: {} },
            type: 'success',
        });

        await navigateTo({ name: 'documents-templates' });
    } catch (e) {
        $grpc.handleError(e as RpcError);
        throw e;
    }
}

const templateAccessTypes = [{ id: 1, name: t('common.job', 2) }];
const contentAccessTypes = [
    { id: 0, name: t('common.citizen', 2) },
    { id: 1, name: t('common.job', 2) },
];

const templateAccess = ref<
    Map<
        string,
        {
            id: string;
            type: number;
            values: {
                job?: string;
                accessRole?: AccessLevel;
                minimumGrade?: number;
            };
        }
    >
>(new Map());

const contentAccess = ref<
    Map<
        string,
        {
            id: string;
            type: number;
            values: {
                job?: string;
                char?: number;
                accessRole?: AccessLevel;
                minimumGrade?: number;
            };
            required?: boolean;
        }
    >
>(new Map());

watch(template, () => {
    if (!template.value) {
        return;
    }

    const tplAccess = template.value.jobAccess;
    if (tplAccess) {
        let accessId = 0;

        tplAccess.forEach((job) => {
            const id = accessId.toString();
            templateAccess.value.set(id, {
                id,
                type: 1,
                values: {
                    job: job.job,
                    accessRole: job.access,
                    minimumGrade: job.minimumGrade,
                },
            });
            accessId++;
        });
    }

    const docAccess = template.value.contentAccess;
    if (docAccess) {
        let accessId = 0;

        docAccess.users.forEach((access) => {
            const id = accessId.toString();
            contentAccess.value.set(id, {
                id,
                type: 0,
                values: { char: access.userId, accessRole: access.access },
                required: access.required,
            });
            accessId++;
        });

        docAccess.jobs.forEach((access) => {
            const id = accessId.toString();
            contentAccess.value.set(id, {
                id,
                type: 1,
                values: {
                    job: access.job,
                    accessRole: access.access,
                    minimumGrade: access.minimumGrade,
                },
                required: access.required,
            });
            accessId++;
        });
    }
});

const openPreview = ref(false);

const { isRevealed, reveal, confirm, cancel, onConfirm } = useConfirmDialog();

onConfirm(async (id) => deleteTemplate(id));
</script>

<template>
    <ConfirmDialog :open="isRevealed" :cancel="cancel" :confirm="() => confirm(templateId)" />

    <TemplatePreviewModal v-if="openPreview" :id="templateId" :open="openPreview" @close="openPreview = false" />

    <div class="m-2">
        <DataPendingBlock v-if="pending" :message="$t('common.loading', [$t('common.template', 2)])" />
        <DataErrorBlock v-else-if="error" :title="$t('common.unable_to_load', [$t('common.template', 2)])" :retry="refresh" />
        <div v-else-if="template" class="px-1 sm:px-2 lg:px-4">
            <div class="sm:flex sm:items-center">
                <div class="inline-flex sm:flex-auto">
                    <NuxtLink
                        v-if="can('DocStoreService.CreateTemplate')"
                        :to="{ name: 'documents-templates-edit-id', params: { id: templateId.toString() } }"
                        class="flex w-full justify-center rounded-md bg-primary-500 px-3 py-2 text-sm font-semibold text-neutral transition-colors hover:bg-primary-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-base-300"
                    >
                        {{ $t('common.edit') }}
                    </NuxtLink>
                    <button
                        v-if="can('DocStoreService.CreateTemplate')"
                        type="button"
                        class="ml-4 flex w-full justify-center rounded-md bg-accent-600 px-3 py-2 text-sm font-semibold text-neutral transition-colors hover:bg-accent-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-base-300"
                        @click="openPreview = true"
                    >
                        {{ $t('common.preview') }}
                    </button>
                    <button
                        v-if="can('DocStoreService.DeleteTemplate')"
                        type="submit"
                        class="ml-4 flex w-full justify-center rounded-md bg-error-600 px-3 py-2 text-sm font-semibold text-neutral transition-colors hover:bg-error-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-base-300"
                        @click="reveal()"
                    >
                        {{ $t('common.delete') }}
                    </button>
                </div>
            </div>
            <div class="sm:flex sm:items-center">
                <div>
                    <h2 class="text-2xl text-neutral">
                        {{ template.title }}
                    </h2>
                    <p class="text-base text-neutral">
                        <span class="font-semibold">{{ $t('common.description') }}:</span> {{ template.description }}
                    </p>
                </div>
            </div>
            <div class="mb-6 mt-4 flow-root">
                <div class="-my-2 mx-0 overflow-x-auto">
                    <div class="my-2">
                        <h3 class="block text-base font-medium leading-6 text-gray-100">
                            {{ $t('common.template', 2) }} {{ $t('common.weight') }}
                        </h3>
                        <div class="my-2">
                            <input
                                type="text"
                                name="weight"
                                class="block w-full rounded-md border-0 bg-base-700 py-1.5 text-neutral placeholder:text-accent-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                disabled
                                :value="template.weight"
                                @focusin="focusTablet(true)"
                                @focusout="focusTablet(false)"
                            />
                        </div>
                    </div>
                    <div v-if="template.jobAccess" class="my-2">
                        <h3 class="block text-base font-medium leading-6 text-gray-100">
                            {{ $t('common.template', 2) }} {{ $t('common.access') }}
                        </h3>
                        <div class="my-2">
                            <DocumentAccessEntry
                                v-for="entry in templateAccess.values()"
                                :key="entry.id"
                                :init="entry"
                                :access-types="templateAccessTypes"
                                :read-only="true"
                            />
                        </div>
                    </div>
                    <div class="my-2">
                        <h3 class="block text-base font-medium leading-6 text-gray-100">
                            {{ $t('common.content') }} {{ $t('common.title') }}
                        </h3>
                        <div class="my-2">
                            <textarea
                                rows="4"
                                name="contentTitle"
                                class="block w-full whitespace-pre-wrap rounded-md border-0 bg-base-700 py-1.5 text-neutral placeholder:text-accent-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                disabled
                                :value="template.contentTitle"
                                @focusin="focusTablet(true)"
                                @focusout="focusTablet(false)"
                            />
                        </div>
                    </div>
                    <div v-if="template.state">
                        <h3 class="block text-base font-medium leading-6 text-gray-100">
                            {{ $t('common.content') }} {{ $t('common.state') }}
                        </h3>
                        <div class="my-2">
                            <input
                                type="text"
                                name="state"
                                class="block w-full whitespace-pre-wrap rounded-md border-0 bg-base-700 py-1.5 text-neutral placeholder:text-accent-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                disabled
                                :value="template.state"
                                @focusin="focusTablet(true)"
                                @focusout="focusTablet(false)"
                            />
                        </div>
                    </div>
                    <div v-if="template.category">
                        <h3 class="block text-base font-medium leading-6 text-gray-100">
                            {{ $t('common.category') }}
                        </h3>
                        <div class="my-2">
                            <p class="text-sm leading-6 text-gray-100">
                                {{ template.category?.name }} ({{ $t('common.description') }}:
                                {{ template.category?.description }})
                            </p>
                        </div>
                    </div>
                    <div class="my-2">
                        <h3 class="block text-base font-medium leading-6 text-gray-100">
                            {{ $t('common.content') }}
                        </h3>
                        <div class="my-2">
                            <textarea
                                rows="4"
                                name="content"
                                class="block w-full rounded-md border-0 bg-base-700 py-1.5 text-neutral placeholder:text-accent-200 focus:ring-2 focus:ring-inset focus:ring-base-300 sm:text-sm sm:leading-6"
                                disabled
                                :value="template.content"
                                @focusin="focusTablet(true)"
                                @focusout="focusTablet(false)"
                            />
                        </div>
                    </div>
                    <div v-if="reqs">
                        <h3 class="block text-base font-medium leading-6 text-gray-100">
                            {{ $t('common.schema') }}
                        </h3>
                        <div class="my-2">
                            <ul
                                class="mb-2 max-w-md list-inside list-disc space-y-1 text-sm font-medium text-gray-100 dark:text-gray-300"
                            >
                                <li v-if="reqs.users">
                                    <TemplateRequirementsList name="User" :specs="reqs.users!" />
                                </li>
                                <li v-if="reqs.vehicles">
                                    <TemplateRequirementsList name="Vehicle" :specs="reqs.vehicles!" />
                                </li>
                                <li v-if="reqs.documents">
                                    <TemplateRequirementsList name="User" :specs="reqs.documents!" />
                                </li>
                            </ul>
                        </div>
                    </div>
                    <div v-if="template.contentAccess" class="my-2">
                        <h3 class="block text-base font-medium leading-6 text-gray-100">
                            {{ $t('common.access') }}
                        </h3>
                        <div class="my-2">
                            <DocumentAccessEntry
                                v-for="entry in contentAccess.values()"
                                :key="entry.id"
                                :init="entry"
                                :access-types="contentAccessTypes"
                                :read-only="true"
                                :show-required="true"
                            />
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <DataNoDataBlock v-else :type="$t('common.template', 2)" />
    </div>
</template>
