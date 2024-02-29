<script lang="ts" setup>
import { useTimeoutFn } from '@vueuse/core';
import NotificationProvider from '~/components/partials/notification/NotificationProvider.vue';
import NotificatorProvider from '~/components/partials/notification/NotificatorProvider.vue';
import DesktopStartMenu from '~/components/desktop/startmenu/DesktopStartMenu.vue';
import LoadingscreenComp from '~/components/desktop/screens/LoadingscreenComp.vue';
import WelcomescreenComp from '~/components/desktop/screens/WelcomescreenComp.vue';

const loading = ref(true);
const welcome = ref(true);

useTimeoutFn(() => {
    loading.value = false;

    useTimeoutFn(() => (welcome.value = false), 1250);
}, 2250);
</script>

<template>
    <NotificationProvider>
        <LoadingscreenComp v-if="loading" />
        <WelcomescreenComp v-else-if="welcome" />

        <section
            aria-labelledby="primary-heading"
            class="h-full min-w-0 lg:order-last"
            :class="!loading && !welcome ? '' : 'hidden'"
        >
            <slot />

            <DesktopStartMenu />
        </section>

        <NotificatorProvider />
    </NotificationProvider>
</template>
