<template>
    <div class="fixed top-0 left-0 w-screen h-screen bg-[#003399] select-none py-20">
        <div class="bg top-0 left-0 origin-top-left w-full h-full flex flex-col items-center justify-between">
            <div class="w-full h-[3px] bg-gradient-to-r from-transparent via-white to-transparent"></div>
            <div class="-mt-32">
                <div class="logotüm scale-[.60] origin-bottom-right -mr-40 mb-3">
                    <div id="line1" class="flex items-end justify-center gap-3 -ml-[240px]">
                        <p
                            class="!leading-4 text-base after:content-['\00a9'] after:text-[10px] after:relative after:-top-2 after:ml-1"
                        >
                            Microsoft
                        </p>
                        <img draggable="false" src="/images/logo.png" alt="xplogo" class="w-32" />
                    </div>
                    <p class="!text-5xl relative">
                        Windows<span class="text-[#FF6821] text-2xl absolute top-0 font-bold">xp</span>
                    </p>
                </div>
                <h1 class="text-2xl font-base text-white font-segoe-ui">Windows is shutting down...</h1>
            </div>
            <div class="h-[3px] w-full bg-gradient-to-r from-transparent via-[#e4a668] to-transparent"></div>
        </div>
    </div>
</template>

<script setup></script>

<style scoped>
h1 {
    text-shadow: 1px 1px #3151b5;
}
.bg {
    background-image: radial-gradient(circle at top left, #8eaeee 16%, #5a7edc 29%);
}
</style>
