<template>
  <div w-full h-full flex items-center dark:bg-dark-black-1 bg-light-white-2 text-light-white-1>
    <div w-51 flex-center shrink-0 bg-light-white-3 dark:bg-dark-black-2 h-16>
      <img v-show="isDark" src="/static/images/logo-white.png" alt="logo" h-14 px-4 aspect-ratio-1k>
      <img v-show="!isDark" src="/static/images/logo-black.png" alt="logo" h-14 px-4 aspect-ratio-1k>
    </div>
    <div flex-grow-1></div>
    <div shrink-0 flex items-center gap-3 px-8>
      <button @click="toggleDark" i-material-symbols:dark-mode dark:i-material-symbols:light-mode text-2xl
        cursor-pointer>
      </button>
      <button i-codicon:chrome-minimize text-2xl cursor-pointer text-light-white-1></button>
      <button i-material-symbols:chrome-maximize-outline text-2xl cursor-pointer text-light-white-1></button>
      <button i-material-symbols:close-rounded text-2xl cursor-pointer text-light-white-1></button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useThemeStore } from '@/store/modules/theme.ts'
import { storeToRefs } from 'pinia'
const themeStore = useThemeStore()
const { toggleDark } = themeStore
const { isDark } = storeToRefs(themeStore)
</script>
