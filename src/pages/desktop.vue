<script lang="ts" setup>
import { VueWinBox } from 'vue-winbox';
import { helpers, type RoutesNamedLocations, type RoutesNamesList, type TypedRouteLocationFromName } from '@typed-router';
import DesktopGrid from '~/components/desktop/DesktopGrid.vue';
import { useSettingsStore } from '~/store/settings';

definePageMeta({
    title: 'Desktop',
    requiresAuth: true,
    layout: 'desktop',
});

const settingsStore = useSettingsStore();
if (!settingsStore.toggles.desktop) {
    navigateTo({ name: 'overview' });
}

const router = useRouter();

const windows = ref<{ name: string; route: TypedRouteLocationFromName<RoutesNamesList> }[]>([]);

function newWindow(href: RoutesNamedLocations): void {
    const r = helpers.route(href);
    const route = router.resolve(r);
    windows.value.push({ name: route.name, route });
}

watch(windows, () => console.log(windows));
</script>

<template>
    <div class="min-h-dvh h-full w-full bg-center bg-repeat" style="background-image: url('/images/bg.svg')">
        <DesktopGrid @open-window="newWindow($event)" />

        <VueWinBox
            v-for="route in windows"
            :key="route.name"
            :options="{
                title: $t(route.route.meta.title ?? 'Unknown'),
                class: 'custom',
                background: 'defaultTheme:bg-accent-600 bg-secondary-600',
            }"
        >
            <NuxtLayout name="blank">
                <NuxtPage />
            </NuxtLayout>
        </VueWinBox>
    </div>
</template>

<style>
.winbox.custom {
    .wb-body {
        background-color: rgb(var(--colors-secondary-600));
    }

    .wb-title {
        color: #fff;
    }

    .wb-full {
        display: none;
    }
}
</style>
