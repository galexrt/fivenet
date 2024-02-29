<script lang="ts" setup>
const time = ref<string>('00:00');

onMounted(() => {
    setInterval(() => {
        const temptime = new Date();
        const hours = temptime.getHours();
        const minutes = temptime.getMinutes();

        time.value = minutes < 10 ? hours + ':0' + minutes : hours + ':' + minutes;
    }, 100);
});
</script>

<template>
    <div class="text-sm leading-5">
        {{ time }}
    </div>
</template>
