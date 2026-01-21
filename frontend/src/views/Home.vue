<template>
  <div size-full flex flex-col items-center justify-center>
    <div>HOME</div>
    <div @click="toggleShowAside()">Toggle Aside</div>
    <div @click="toggleDark()">Toggle Dark</div>
    <div @click="testDB()">Test DB</div>
  </div>

</template>

<script setup lang="ts">
import { useLayoutStore } from '@/store/modules/layout.ts'
import { useThemeStore } from '@/store/modules/theme.ts'
import { TestDB } from '@wails/go/app/DB.js'
const layoutStore = useLayoutStore()
const toggleShowAside = () => {
  layoutStore.toggleShowAside()
}
const { toggleDark } = useThemeStore()
const testDB = () => {
  TestDB().then((res) => {
    if (res) {
      console.log("DB is connected")
    } else {
      console.log("DB is not connected")
    }
  })
}
</script>
