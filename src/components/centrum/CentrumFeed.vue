<script lang="ts" setup>
import { DispatchStatus } from '~~/gen/ts/resources/centrum/dispatches';
import { UnitStatus } from '~~/gen/ts/resources/centrum/units';
import DispatchFeedItem from '~/components/centrum/dispatches/DispatchFeedItem.vue';
import UnitFeedItem from '~/components/centrum/units/UnitFeedItem.vue';

defineProps<{
    items: (DispatchStatus | UnitStatus)[];
}>();

defineEmits<{
    (e: 'goto', loc: Coordinate): void;
}>();
</script>

<template>
    <div class="h-full overflow-y-auto px-4 sm:px-6 lg:px-8">
        <div class="sm:flex sm:items-center">
            <div class="inline-flex items-center sm:flex-auto">
                <h2 class="text-base font-semibold leading-6 text-gray-100">
                    {{ $t('common.activity', 2) }}
                </h2>
            </div>
        </div>
        <div class="mt-0.5 flow-root">
            <div class="-mx-2 sm:-mx-6 lg:-mx-8">
                <div class="inline-block min-w-full py-2 align-middle sm:px-2 lg:px-2">
                    <ul role="list" class="space-y-2">
                        <template v-for="(activityItem, activityItemIdx) in items">
                            <DispatchFeedItem
                                v-if="'dispatchId' in activityItem"
                                :key="'dsp' + activityItem.id"
                                :activity-length="items?.length ?? 0"
                                :item="activityItem"
                                :activity-item-idx="activityItemIdx"
                                :show-id="true"
                                @goto="$emit('goto', $event)"
                            />
                            <UnitFeedItem
                                v-else
                                :key="'unit' + activityItem.id"
                                :activity-length="items?.length ?? 0"
                                :item="activityItem"
                                :activity-item-idx="activityItemIdx"
                                @goto="$emit('goto', $event)"
                            />
                        </template>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</template>
