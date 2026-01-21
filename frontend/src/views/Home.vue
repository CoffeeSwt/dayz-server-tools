<template>
  <div size-full flex flex-col items-center justify-center>
    <div>HOME</div>
    <div @click="toggleShowAside()">Toggle Aside</div>
    <div @click="toggleDark()">Toggle Dark</div>
  </div>

</template>

<script setup lang="ts">
import { useLayoutStore } from '@/store/modules/layout.ts'
import { useThemeStore } from '@/store/modules/theme.ts'
const layoutStore = useLayoutStore()
const toggleShowAside = () => {
  layoutStore.toggleShowAside()
}
const { toggleDark } = useThemeStore()
</script>
