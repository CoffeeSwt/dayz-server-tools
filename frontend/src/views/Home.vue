<template>
  <div>HOME</div>
  <div @click="toggleShowAside()">Toggle Aside</div>
</template>

<script setup lang="ts">
import { useLayoutStore } from '@/store/modules/layout.ts'
const layoutStore = useLayoutStore()
const toggleShowAside = () => {
    layoutStore.toggleShowAside()
}
</script>
