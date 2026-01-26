<template>
  <div size-full flex flex-col items-center justify-center>
    <div flex gap-6>
      <template v-for="item in maplist" :key="item.id">
        <MapCard :id="item.id" :mapName="item.mapName" :chineseName="item.chineseName" :image="item.image"
          :tips="item.tips" :missionName="item.missionName" :isWorkShopMap="item.isWorkShopMap"
          :transform="item.transform" />
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useLayoutStore } from '@/store/modules/layout.ts'
const { setShowAside } = useLayoutStore()
onMounted(() => {
  setShowAside(false)
})
import MapCard from '@/components/MapCard.vue'
import { onMounted } from 'vue'
const maplist = ref([
  {
    id: 1,
    mapName: 'DayZ',
    chineseName: '车臣尼亚',
    image: '/static/images/chernarusplus.jpg',
    tips: 'DayZ是一个开放世界的沙盒游戏，玩家可以在游戏中自由探索、建造和战斗。',
    missionName: 'DayZ',
    isWorkShopMap: false,
    transform: "translate(-296px, -7px) scale(0.40)",
  },
  {
    id: 2,
    mapName: 'ChernarusPlus',
    chineseName: '利沃尼亚',
    image: '/static/images/enoch.jpg',
    tips: 'ChernarusPlus是一个基于Chernarus地图的增强版，添加了更多的内容和功能。',
    missionName: 'ChernarusPlus',
    isWorkShopMap: false,
    transform: "translate(-285px, 0px) scale(0.38)",
  },
  {
    id: 3,
    mapName: 'Sakhal',
    chineseName: '萨哈尔',
    image: '/static/images/sakhal.jpg',
    tips: 'Sakhal是一个基于Chernarus地图的增强版，添加了更多的内容和功能。',
    missionName: 'Sakhal',
    isWorkShopMap: false,
    transform: "translate(-199px, -2px) scale(0.38)",
  },
])
</script>
