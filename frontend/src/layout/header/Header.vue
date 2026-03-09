<template>
  <div w-full h-full flex items-center dark:bg-dark-black-1 bg-light-white-2 text-light-white-1>
    <div v-show="showLogo" w-51 flex-center shrink-0 bg-light-white-3 dark:bg-dark-black-2 h-16>
      <img v-show="isDark" src="/static/images/logo-white.png" alt="logo" h-14 px-4 aspect-ratio-1k>
      <img v-show="!isDark" src="/static/images/logo-black.png" alt="logo" h-14 px-4 aspect-ratio-1k>
    </div>
    <div v-show="!showLogo" px-8 i-material-symbols:arrow-back text-2xl text-light-white-1 cursor-pointer
      @click="routeBack"></div>
    <div flex-grow-1></div>
    <div shrink-0 flex items-center gap-3 px-8>
      <button @click="toggleDark" i-material-symbols:dark-mode dark:i-material-symbols:light-mode text-2xl
        cursor-pointer>
      </button>
      <button i-codicon:chrome-minimize text-2xl cursor-pointer text-light-white-1 @click="minimise"></button>
      <button v-show="!isMaximised" i-material-symbols:chrome-maximize-outline text-2xl cursor-pointer
        text-light-white-1 @click="toggleMaximise"></button>
      <button v-show="isMaximised" i-fluent:full-screen-minimize-20-filled text-2xl cursor-pointer text-light-white-1
        @click="toggleMaximise"></button>
      <button i-material-symbols:close-rounded text-2xl cursor-pointer text-light-white-1 @click="handleClose"></button>
      <ConfirmDialog v-model="confirmDialogVisible" title="关闭程序" message="确定要关闭程序吗？" @confirm="handleConfirmClose">
      </ConfirmDialog>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useThemeStore } from '@/store/modules/theme.ts'
import { storeToRefs } from 'pinia'
import { WindowMinimise } from '@wails/runtime/runtime.js'
import ConfirmDialog from '@/components/common/ConfirmDialog.vue'
import { ref, computed } from 'vue'
import { CloseApp } from '@wails/go/app/App.js'
import { useLayoutStore } from '@/store/modules/layout.ts'
import { useRouter } from 'vue-router'

const layoutStore = useLayoutStore()
const showLogo = computed(() => layoutStore.showAside)
const confirmDialogVisible = ref(false)
const themeStore = useThemeStore()
const router = useRouter()
const { toggleDark, toggleMaximise } = themeStore
const { isDark, isMaximised } = storeToRefs(themeStore)
const minimise = () => {
  WindowMinimise()
}
const handleClose = () => {
  confirmDialogVisible.value = true
}
const handleConfirmClose = () => {
  CloseApp()
}

const routeBack = () => {
  router.back()
}


</script>
