<template>
  <aside class="group" relative w-full h-full flex items-center justify-start bg-light-white-3 text-light-black-1
    dark:bg-dark-black-2 dark:text-dark-gray-1 flex-col>
    <template v-for="item in menuList" :key="item.path">
      <div @click="routerTo(item.path)" w-85% h-12 flex-center cursor-pointer hover:bg-light-black-2 dark:hover:bg-dark-black-1 rounded-2xl px-4
        box-border duration-150 text-xl>
        {{ item.name }}
      </div>
    </template>
  </aside>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
const menuList = [
  {
    name: '服务器',
    path: '/layout',
  },
  // {
  //   name: '地图选择',
  //   path: '/layout/server',
  // },
]

const router = useRouter()
const routerTo = (path: string) => {
  router.push(path)
}
</script>
