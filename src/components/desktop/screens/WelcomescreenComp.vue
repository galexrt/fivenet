<template>
    <div class="fixed top-0 left-0 w-screen h-screen bg-[#003399] select-none py-20">
        <div class="bg top-0 left-0 origin-top-left w-full h-full flex flex-col items-center justify-between">
            <div class="w-full h-[3px] bg-gradient-to-r from-transparent via-white to-transparent"></div>
            <h1 class="text-4xl font-bold text-white font-segoe-ui italic">welcome</h1>
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
