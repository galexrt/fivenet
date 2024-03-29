<script lang="ts" setup>
import { RpcError } from '@protobuf-ts/runtime-rpc';
import CardsList from '~/components/partials/CardsList.vue';
import DataErrorBlock from '~/components/partials/data/DataErrorBlock.vue';
import DataNoDataBlock from '~/components/partials/data/DataNoDataBlock.vue';
import DataPendingBlock from '~/components/partials/data/DataPendingBlock.vue';
import { type CardElements } from '~/utils/types';
import { TemplateShort } from '~~/gen/ts/resources/documents/templates';

const { $grpc } = useNuxtApp();

defineEmits<{
    (e: 'selected', t: TemplateShort): void;
}>();

const { data: templates, pending, refresh, error } = useLazyAsyncData(`documents-templates`, () => listTemplates());

async function listTemplates(): Promise<TemplateShort[]> {
    try {
        const call = $grpc.getDocStoreClient().listTemplates({});
        const { response } = await call;

        return response.templates;
    } catch (e) {
        $grpc.handleError(e as RpcError);
        throw e;
    }
}

const items = ref<CardElements>([]);
watch(templates, () =>
    templates.value?.forEach((v) => {
        items.value.push({ title: v?.title, description: v?.description });
    }),
);

function selected(idx: number): TemplateShort {
    return templates.value![idx];
}
</script>

<template>
    <div>
        <DataPendingBlock v-if="pending" :message="$t('common.loading', [$t('common.template', 2)])" />
        <DataErrorBlock v-else-if="error" :title="$t('common.unable_to_load', [$t('common.template', 2)])" :retry="refresh" />
        <DataNoDataBlock v-else-if="templates && templates.length === 0" :type="$t('common.template', 2)" />

        <div v-else class="flex justify-center">
            <CardsList :items="items" :show-icon="false" @selected="$emit('selected', selected($event))" />
        </div>
    </div>
</template>
