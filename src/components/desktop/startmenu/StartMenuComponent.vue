<script lang="ts" setup>
import { useAuthStore } from '~/store/auth';

defineProps<{
    open: boolean;
}>();

defineEmits<{
    (e: 'close'): void;
}>();

const authStore = useAuthStore();
const { activeChar } = storeToRefs(authStore);

const closePowerModal = ref(false);

const shutdown = () => {
    closePowerModal.value = false;
};
</script>

<template>
    <div
        v-if="open"
        id="startmenu"
        class="w-96 h-[436px] select-none rounded-t-lg bg-white/50 fixed bottom-[30px] left-0 z-[999] flex flex-col items-center justify-between font-segoe-ui"
    >
        <div
            class="windows-blue-gradient shrink-0 rounded-t-lg flex items-center h-16 select-none text-white justify-start gap-2 pl-2 w-full"
        >
            <div class="border border-white w-12 h-12 rounded-md overflow-hidden">
                <img draggable="false" :src="activeChar.avatar.url" alt="start-pp" />
            </div>
            <p class="font-bold font-segoe-ui tracking-wider text-base">{{ activeChar.firstname }} {{ activeChar.lastname }}</p>
        </div>
        <div class="bg-zinc-500 h-full w-full flex items-center justify-center border-x-2 border-[#0831d9] flex-wrap">
            <div class="h-[2px] w-full bg-gradient-to-r from-transparent via-[#e4a668] to-transparent"></div>
            <!-- LEFT AREA -->
            <div class="w-1/2 h-full bg-white text-sm tracking-tighter font-segoe-ui flex flex-col justify-end items-center">
                <ul class="w-full px-px py-1">
                    <li
                        class="text-center w-full flex items-center py-1 justify-start gap-1 px-2 hover:!text-white hover:bg-[#316ac5] text-xs font-bold"
                    >
                        <img class="w-8" draggable="false" src="" alt="internet" />
                        <div class="flex flex-col items-start justify-start">
                            <p>Internet</p>
                            <p class="font-thin opacity-60">Internet Explorer</p>
                        </div>
                    </li>
                </ul>
                <ul class="w-full px-px py-1">
                    <li
                        class="text-center w-full flex items-center py-1 justify-start gap-1 px-2 hover:!text-white hover:bg-[#316ac5] text-xs font-bold"
                    >
                        <img class="w-8" draggable="false" src="" alt="email" />
                        <div class="flex flex-col items-start justify-start">
                            <p>E-mail</p>
                            <p class="font-thin opacity-60">Outlook Express</p>
                        </div>
                    </li>
                </ul>
                <div class="w-full h-px bg-gradient-to-r from-transparent via-zinc-300 to-transparent"></div>
                <ul class="w-full px-px py-1">
                    <li
                        class="text-center w-full flex items-center py-1 justify-start gap-1 px-2 hover:text-white hover:bg-[#316ac5] text-xs"
                    >
                        <img class="w-8" draggable="false" src="" alt="msnexplorer" />MSN Explorer
                    </li>
                </ul>
                <ul class="w-full px-px py-1">
                    <li
                        class="text-center w-full flex items-center py-1 justify-start gap-1 px-2 hover:text-white hover:bg-[#316ac5] text-xs"
                    >
                        <img class="w-8" draggable="false" src="" alt="mediaplayer" />Windows Media Player
                    </li>
                </ul>
                <ul class="w-full px-px py-1">
                    <li
                        class="text-center w-full flex items-center py-1 justify-start gap-1 px-2 hover:text-white hover:bg-[#316ac5] text-xs"
                    >
                        <img class="w-8" draggable="false" src="" alt="tour" />Tour Windows XP
                    </li>
                </ul>
                <ul class="w-full px-px py-1">
                    <li
                        class="text-center w-full flex items-center py-1 justify-start gap-1 px-2 hover:text-white hover:bg-[#316ac5] text-xs"
                    >
                        <img class="w-8" draggable="false" src="" alt="moviemaker" />Windows Movie Maker
                    </li>
                </ul>
                <div class="w-full h-px bg-gradient-to-r from-transparent via-zinc-300 to-transparent"></div>
                <ul class="w-full px-px py-1">
                    <li
                        class="text-center w-full font-semibold flex items-center py-1 justify-center gap-4 hover:text-white hover:bg-[#316ac5]"
                    >
                        All Programs <img draggable="false" src="" alt="greenarrow" />
                    </li>
                </ul>
            </div>
            <!-- RIGHT AREA -->
            <div class="w-1/2 h-full bg-[#d2e5fa] border-l border-l-blue-300">
                <ul class="px-px py-1">
                    <li
                        class="flex items-center hover:text-white hover:bg-[#316ac5] justify-start gap-2 text-xs font-bold py-1 px-2 text-[#0a1835]"
                    >
                        <img draggable="false" class="w-6" src="" alt="documents" />
                        <p>My Documents</p>
                    </li>
                    <li
                        class="flex items-center hover:text-white hover:bg-[#316ac5] justify-start gap-2 text-xs font-bold py-1 px-2 text-[#0a1835]"
                    >
                        <img draggable="false" class="w-6" src="" alt="pictures" />
                        <p>My Pictures</p>
                    </li>
                    <li
                        class="flex items-center hover:text-white hover:bg-[#316ac5] justify-start gap-2 text-xs font-bold py-1 px-2 text-[#0a1835]"
                    >
                        <img draggable="false" class="w-6" src="" alt="musics" />
                        <p>My Music</p>
                    </li>
                    <li
                        class="flex items-center hover:text-white hover:bg-[#316ac5] justify-start gap-2 text-xs font-bold py-1 px-2 text-[#0a1835]"
                    >
                        <img draggable="false" class="w-6" src="" alt="computer" />
                        <p>My Computer</p>
                    </li>
                </ul>
                <div class="w-11/12 h-px bg-gradient-to-r from-transparent via-zinc-400 to-transparent"></div>
                <ul class="px-px py-1">
                    <li
                        class="flex items-center hover:text-white hover:bg-[#316ac5] justify-start gap-2 text-xs font-thin px-2 py-1 text-[#0a1835]"
                    >
                        <img draggable="false" class="w-6" src="" alt="control" />
                        <p>Control Panel</p>
                    </li>
                </ul>
                <div class="w-11/12 h-px bg-gradient-to-r from-transparent via-zinc-400 to-transparent"></div>
                <ul class="px-px py-1">
                    <li
                        class="flex items-center hover:text-white hover:bg-[#316ac5] justify-start gap-2 text-xs font-thin px-2 py-1 text-[#0a1835]"
                    >
                        <img draggable="false" class="w-6" src="" alt="help" />
                        <p>Help and Support</p>
                    </li>
                    <li
                        class="flex items-center hover:text-white hover:bg-[#316ac5] justify-start gap-2 text-xs font-thin px-2 py-1 text-[#0a1835]"
                    >
                        <img draggable="false" class="w-6" src="" alt="search" />
                        <p>Search</p>
                    </li>
                    <li
                        class="flex items-center hover:text-white hover:bg-[#316ac5] justify-start gap-2 text-xs font-thin px-2 py-1 text-[#0a1835]"
                    >
                        <img draggable="false" class="w-6" src="" alt="run" />
                        <p>Run...</p>
                    </li>
                </ul>
            </div>
        </div>
        <div class="windows-blue-gradient-bottom shrink-0 h-10 self-end w-full flex items-center justify-end px-2 gap-2">
            <div
                class="flex text-white font-thin text-xs items-center justify-center group gap-1"
                @click="closePowerModal = true"
            >
                <div class="bg-white rounded-sm">
                    <img draggable="false" class="w-6 h-6 group-active:opacity-70" src="" alt="log-off" />
                </div>
                <p>Log Off</p>
            </div>
            <div
                class="flex text-white font-thin text-xs items-center justify-center group gap-1"
                @click="closePowerModal = true"
            >
                <div class="bg-white rounded-sm">
                    <img draggable="false" class="w-6 h-6 group-active:opacity-70" src="" alt="turn-off" />
                </div>
                <p>Turn Off Computer</p>
            </div>
        </div>
    </div>
    <div
        v-if="closePowerModal"
        class="w-screen select-none h-screen z-[999999] backdrop-saturate-anim fixed top-0 left-0 flex items-center justify-center"
    >
        <div class="border border-black w-[314px] h-[200px] bg-[#003399] shrink-0 flex flex-col">
            <div class="text-white flex items-center justify-between px-2 h-[70px]">
                <p class="font-medium text-lg">Turn off computer</p>
                <img draggable="false" src="/images/logo.png" alt="Logo" class="h-8" />
            </div>
            <div
                class="bg-gradient-to-r from-[#698be4] via-[#95b4f4] to-[#698be4] w-full h-full flex flex-col items-start justify-start"
            >
                <div class="w-full h-[3px] bg-gradient-to-r from-transparent via-white to-transparent"></div>
                <div
                    class="w-full h-full select-none text-white text-xs font-medium font-segoe-ui flex items-center justify-center gap-12"
                >
                    <div class="flex flex-col items-center justify-center gap-2" @click="shutdown">
                        <img draggable="false" class="select-none active:opacity-70 bg-white" src="" alt="stand-by" />
                        <p>Stand By</p>
                    </div>
                    <div class="flex flex-col items-center justify-center gap-2" @click="shutdown">
                        <img draggable="false" class="select-none active:opacity-70 bg-white" src="" alt="turnoff" />
                        <p>Turn Off</p>
                    </div>
                    <div class="flex flex-col items-center justify-center gap-2" @click="shutdown">
                        <img draggable="false" class="select-none active:opacity-70 bg-white" src="" alt="restart" />
                        <p>Restart</p>
                    </div>
                </div>
            </div>
            <div class="flex items-center justify-end px-4 h-[70px]">
                <button class="button" @click="closePowerModal = false">Cancel</button>
            </div>
        </div>
    </div>
</template>

<style scoped>
.windows-blue-gradient {
    background-image: linear-gradient(0deg, #0072f4 0%, #3889e6 11%, #1472e6 42%, #004cb0 92%, #70aafb 96%, #255db5 100%);
    border-top: 1px solid #0831d9;
    border-left: 1px solid #0831d9;
    border-right: 1px solid #001ea0;

    box-shadow: 0px 1px 22px #0832d99f inset;
    text-shadow: 2px 2px 1px #020236;
}
.windows-blue-gradient-bottom {
    background-image: linear-gradient(0deg, #004cb0 0%, #004cb0 11%, #2b7ce3 42%, #418be4 92%, #418fe4 96%);
    border-top: 1px solid #0831d9;
    border-left: 1px solid #0831d9;
    border-right: 1px solid #001ea0;
    box-shadow: 10px -2px 51px #02029267 inset;
}

.backdrop-saturate-anim {
    animation: saturate-zero 4s ease both;
}

@keyframes saturate-zero {
    from {
        backdrop-filter: saturate(100%);
    }
    to {
        backdrop-filter: saturate(0%);
    }
}
</style>
