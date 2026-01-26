<template>
    <div rounded-lg></div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { Server } from '@/types/server.ts'
const props = defineProps({
    server: {
        type: Object as () => Server,
        default: () => ({})
    }
})
</script>
