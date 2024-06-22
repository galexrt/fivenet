<script lang="ts" setup>
import type { MDCParserResult } from '@nuxtjs/mdc';
import { parseMarkdown } from '@nuxtjs/mdc/runtime';
import DashboardPanel from '~/components/partials/dashboard/DashboardPanel.vue';

const md = ref(`
# Just a Vue app

This is markdown content rendered via the \`<MDCRenderer>\` component, including MDC below.

:UAlert{title="Hello MDC"}

\`\`\`ts
const a = 1;
\`\`\`

## Test1

Citizen Info:

:PartialsCitizensCitizenInfoPopover{userId=26061} Hast du die Sachen ins Lager gebracht?

## Test123 Hello World

TEST

### ABC XYZ 123 456

HELLO WORLD!

`);

const ast = ref<MDCParserResult | null>(null);

onBeforeMount(async () => {
    ast.value = await parseMarkdown(md.value);
});
</script>

<template>
    <UDashboardPage class="h-full">
        <UDashboardPanel
            class="h-full flex-shrink-0 border-b border-gray-200 lg:border-b-0 lg:border-r dark:border-gray-800"
            grow
        >
            <UDashboardNavbar :title="ast?.data.title" />

            <UDashboardPanelContent class="p-0">
                <div class="mx-auto w-full max-w-screen-xl break-words rounded-lg bg-base-900">
                    <div class="prose dark:prose-invert min-w-full px-8 py-2">
                        <MDCRenderer v-if="ast?.body" :body="ast.body" :data="ast.data" />
                    </div>
                </div>
            </UDashboardPanelContent>
        </UDashboardPanel>

        <DashboardPanel collapsible side="right" class="min-w-72 max-w-72">
            <UDashboardNavbar :title="$t('common.toc')" />

            <UDashboardPanelContent>
                <UContentTocLinks v-if="ast?.toc?.links?.length" :links="ast.toc.links" />
            </UDashboardPanelContent>
        </DashboardPanel>
    </UDashboardPage>
</template>
