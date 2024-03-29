<script lang="ts" setup>
import { RpcError } from '@protobuf-ts/runtime-rpc';
import { useConfirmDialog } from '@vueuse/core';
import { CloseCircleIcon } from 'mdi-vue3';
import GenericContainerPanelEntry from '~/components/partials/elements/GenericContainerPanelEntry.vue';
import type { OAuth2Account, OAuth2Provider } from '~~/gen/ts/resources/accounts/oauth2';
import OAuth2ConnectButton from '~/components/auth/account/OAuth2ConnectButton.vue';
import ConfirmDialog from '~/components/partials/ConfirmDialog.vue';
const { $grpc } = useNuxtApp();

defineProps<{
    provider: OAuth2Provider;
    account?: OAuth2Account;
}>();

const emit = defineEmits<{
    (e: 'disconnected', provider: string): void;
}>();

async function disconnectOAuth2Connection(provider: OAuth2Provider): Promise<void> {
    try {
        await $grpc.getAuthClient().deleteOAuth2Connection({
            provider: provider.name,
        });

        emit('disconnected', provider.name);
    } catch (e) {
        $grpc.handleError(e as RpcError);
        throw e;
    }
}

const { isRevealed, reveal, confirm, cancel, onConfirm } = useConfirmDialog();

onConfirm(async (provider) => disconnectOAuth2Connection(provider));
</script>

<template>
    <ConfirmDialog :open="isRevealed" :cancel="cancel" :confirm="() => confirm(provider)" />

    <GenericContainerPanelEntry>
        <template #title>
            <NuxtLink :external="true" :to="provider.homepage" target="_blank">
                {{ provider.label }}
            </NuxtLink>
        </template>
        <template #default>
            <div v-if="account !== undefined" class="flex items-center justify-between">
                <img
                    :src="account.avatar"
                    alt="Avatar"
                    class="h-10 w-10 rounded-full bg-base-800 fill-base-300 text-base-300 ring-2 ring-neutral hover:fill-base-100 hover:text-base-100 hover:transition-colors"
                />
                <span class="text-left" :title="`ID: ${account.externalId}`">
                    {{ account.username }}
                </span>

                <button
                    type="button"
                    class="inline-flex items-center gap-1 rounded-md bg-error-600 px-2 py-2 text-sm font-semibold text-neutral hover:bg-error-700"
                    @click="reveal(provider)"
                >
                    <CloseCircleIcon class="h-5 w-5" aria-hidden="true" />
                    <span>{{ $t('common.disconnect') }}</span>
                </button>
            </div>
            <div v-else>
                <OAuth2ConnectButton :provider="provider" />
            </div>
        </template>
    </GenericContainerPanelEntry>
</template>
