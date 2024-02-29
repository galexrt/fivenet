<script lang="ts" setup>
import { DesktopClassicIcon } from 'mdi-vue3';
import ClockComponent from './ClockComponent.vue';
import StartMenuComponent from './StartMenuComponent.vue';
import TaskbarTab from './TaskbarTab.vue';
import NotificationIconPopup from './NotificationIconPopup.vue';

const startMenuOpen = ref(false);
</script>

<template>
    <div class="taskbar flex select-none items-center justify-start z-[9999]">
        <div
            class="taskbar__start flex items-center justify-center gap-1 rounded-r-lg mr-0.5"
            @click="startMenuOpen = !startMenuOpen"
        >
            <img draggable="false" class="taskbar__start--logo !w-5 !h-5" src="/images/logo.png" alt="Windows" />
            <p class="tracking-wider pr-2">start</p>
        </div>

        <!--
        <img draggable="false" class="mx-4 w-5" src="" alt="iexplorer" />
        -->

        <!-- My Computer tab -->
        <TaskbarTab name="My Computer" :active="false" :icon="markRaw(DesktopClassicIcon)" icon-class="fill-gray-300" />

        <!-- Notification Icons Area -->
        <div
            class="taskbar__end ml-auto h-full w-fit text-white border box-border border-t-[#075dca] border-b-[#0a5bc6] border-r-transparent border-l-black flex items-center justify-end px-2 pt-1 text-sm gap-1"
        >
            <NotificationIconPopup title="Better in full screen" message="Press F11 for full screen mode" />
            <ClockComponent />
        </div>
    </div>

    <StartMenuComponent :open="startMenuOpen" @close="startMenuOpen = false" />
</template>

<style scoped>
.taskbar {
    position: fixed;
    bottom: 0;
    background-image: linear-gradient(
        0deg,
        #042b8e 0%,
        #0551f6 6%,
        #0453ff 51%,
        #0551f6 63%,
        #0551f6 81%,
        #3a8be8 90%,
        #0453ff 100%
    );
    width: 100vw;
    height: 30px;
}
.taskbar__start {
    height: 100%;
    width: 100px;
    color: white;
    font-weight: 500;
    font-size: 1.1rem;
    font-style: italic;
    border: none;
    box-shadow: -2px -2px 10px #0000008e inset;
    background-image: linear-gradient(
        0deg,
        #0c450c 0%,
        #308f2f 6%,
        #308f2f 51%,
        #308f2f 63%,
        #308f2f 77%,
        #97c597 85%,
        #97c597 93%,
        #308f2f 97%
    );
    display: flex;
    align-items: center;
    transition: all;
}
.taskbar__start:active {
    background-image: linear-gradient(
        0deg,
        #2f892f 0%,
        #4eb64e 6%,
        #4eb64e 51%,
        #4eb64e 63%,
        #4eb64e 77%,
        #c4ffc4 85%,
        #c4ffc4 93%,
        #2f892f 97%
    );
}
.taskbar__start p {
    filter: drop-shadow(2px 2px 2px rgba(0, 0, 0, 0.685));
}
.taskbar__start--logo {
    filter: drop-shadow(2px 2px 2px rgba(0, 0, 0, 0.705));
}

.taskbar__end {
    background-image: linear-gradient(
        0deg,
        #0a5bc6 0%,
        #1198e9 6%,
        #1198e9 51%,
        #1198e9 63%,
        #1198e9 77%,
        #19b9f3 85%,
        #19b9f3 93%,
        #075dca 97%
    );
    box-shadow: 2px -0px 3px #20e2fc inset;
}

.tab {
    box-shadow: 2px 2px 5px rgba(255, 255, 255, 0.267) inset;
}
</style>
