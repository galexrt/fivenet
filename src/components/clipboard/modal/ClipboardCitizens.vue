<script lang="ts" setup>
import { AccountMultipleIcon, TrashCanIcon } from 'mdi-vue3';
import DataNoDataBlock from '~/components/partials/data/DataNoDataBlock.vue';
import { ClipboardUser, useClipboardStore } from '~/store/clipboard';
import { useNotificatorStore } from '~/store/notificator';
import { ObjectSpecs } from '~~/gen/ts/resources/documents/templates';

const clipboardStore = useClipboardStore();
const notifications = useNotificatorStore();

const { users, activeStack } = storeToRefs(clipboardStore);

const emit = defineEmits<{
    (e: 'statisfied', payload: boolean): void;
}>();

const props = withDefaults(
    defineProps<{
        submit?: boolean;
        showSelect?: boolean;
        specs?: ObjectSpecs;
    }>(),
    {
        submit: false,
        showSelect: true,
        specs: undefined,
    },
);

const selected = ref<ClipboardUser[]>([]);

async function select(item: ClipboardUser): Promise<void> {
    const idx = selected.value.indexOf(item);
    if (idx !== undefined && idx > -1) {
        selected.value.splice(idx, 1);
    } else {
        if (props.specs && props.specs.max) {
            selected.value.splice(0, selected.value.length);
        }
        selected.value.push(item);
    }

    const selectedLength = selected.value.length;
    if (props.specs !== undefined) {
        if (props.specs.min !== undefined && selectedLength >= props.specs.min) {
            emit('statisfied', true);
        } else if (props.specs.max && selectedLength === props.specs.max) {
            emit('statisfied', true);
        } else {
            emit('statisfied', false);
        }
    } else {
        emit('statisfied', true);
    }
}

async function remove(item: ClipboardUser, notify: boolean): Promise<void> {
    const idx = selected.value.indexOf(item);
    if (idx !== undefined && idx > -1) {
        selected.value.splice(idx, 1);
    }

    clipboardStore.removeUser(item.userId!);
    if (notify) {
        notifications.dispatchNotification({
            title: { key: 'notifications.clipboard.citizen_removed.title', parameters: {} },
            content: { key: 'notifications.clipboard.citizen_removed.content', parameters: {} },
            duration: 3250,
            type: 'info',
        });
    }
}

async function removeAll(): Promise<void> {
    while (selected.value.length > 0) {
        selected.value.forEach((v) => {
            remove(v, false);
        });
    }

    emit('statisfied', false);
    notifications.dispatchNotification({
        title: { key: 'notifications.clipboard.citizens_removed.title', parameters: {} },
        content: { key: 'notifications.clipboard.citizens_removed.content', parameters: {} },
        duration: 3250,
        type: 'info',
    });
}

watch(props, async (newVal) => {
    if (newVal.submit) {
        if (activeStack.value) {
            activeStack.value.users.length = 0;
            selected.value.forEach((v) => activeStack.value.users.push(v));
        } else if (users.value && users.value.length === 1) {
            selected.value.unshift(users.value[0]);
        }
    }
});
</script>

<template>
    <h3 class="pb-1 pt-1 font-medium">{{ $t('common.citizen', 2) }}</h3>
    <DataNoDataBlock
        v-if="users?.length === 0"
        :icon="AccountMultipleIcon"
        :message="$t('components.clipboard.clipboard_modal.no_data', [$t('common.citizen', 2)])"
    />
    <table v-else class="min-w-full divide-y divide-gray-700">
        <thead>
            <tr>
                <th v-if="showSelect" scope="col" class="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-neutral sm:pl-1">
                    {{ $t('common.select', 1) }}
                </th>
                <th scope="col" class="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-neutral sm:pl-1">
                    {{ $t('common.name') }}
                </th>
                <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-neutral">
                    {{ $t('common.job', 1) }}
                </th>
                <th scope="col" class="relative py-3.5 pl-3 pr-4 sm:pr-0">
                    {{ $t('common.action', 2) }}
                    <button v-if="selected.length > 0" type="button" @click="removeAll()">
                        <TrashCanIcon class="mx-auto h-5 w-5 text-neutral" aria-hidden="true" />
                    </button>
                </th>
            </tr>
        </thead>
        <tbody class="divide-y divide-gray-800">
            <tr v-for="item in users" :key="item.userId">
                <td v-if="showSelect" class="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-neutral sm:pl-1">
                    <template v-if="specs && specs.max === 1">
                        <button
                            class="inline-flex w-full justify-center rounded-md px-3 py-2 text-sm font-semibold text-neutral shadow-sm focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 sm:col-start-2"
                            :class="
                                selected.includes(item)
                                    ? 'bg-accent-500 hover:bg-accent-400 focus-visible:outline-accent-500'
                                    : 'bg-primary-500 hover:bg-primary-400 focus-visible:outline-primary-500'
                            "
                            @click="select(item)"
                        >
                            {{
                                !selected.includes(item)
                                    ? $t('common.select', 1).toUpperCase()
                                    : $t('common.select', 2).toUpperCase()
                            }}
                        </button>
                    </template>
                    <template v-else>
                        <input
                            :key="item.userId"
                            v-model="selected"
                            name="selected"
                            :checked="selected.includes(item)"
                            :value="item"
                            type="checkbox"
                            class="h-4 w-4 rounded border-gray-300 text-primary-600 focus:ring-primary-600"
                            @click="select(item)"
                        />
                    </template>
                </td>
                <td class="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-neutral sm:pl-1">
                    {{ item.firstname }} {{ item.lastname }}
                </td>
                <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-300">
                    {{ item.jobLabel }}
                </td>
                <td class="relative whitespace-nowrap py-4 pl-3 pr-4 text-right text-sm font-medium sm:pr-0">
                    <button type="button" @click="remove(item, true)">
                        <TrashCanIcon class="mx-auto h-5 w-5 text-neutral" aria-hidden="true" />
                    </button>
                </td>
            </tr>
        </tbody>
    </table>
</template>
