<template>
    <div w-66 h-102 rounded-md overflow-hidden cursor-pointer class="group" relative @click="routeToServer">
        <div w-full h-full group-hover:scale-110 duration-300 ease-in-out overflow-hidden>
            <div w-full h-full overflow-hidden relative>
                <img :src="image" :style="{ transform: transform, transformOrigin: 'left top' }" />
            </div>
        </div>
        <div w-full h-20 absolute group-hover:bottom-0 bottom--100% duration-300 ease-in-out left-0 flex-center
            bg-dark-black-2 opacity-60>
            <div text-white text-center text-2xl font-bold>{{ chineseName }}</div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
const router = useRouter()
const props = defineProps({
    id: {
        type: Number,
        default: 0,
    },
    mapName: {
        type: String,
        default: '',
    },
    chineseName: {
        type: String,
        default: '',
    },
    image: {
        type: String,
        default: '',
    },
    tips: {
        type: String,
        default: '',
    },
    missionName: {
        type: String,
        default: '',
    },
    isWorkShopMap: {
        type: Boolean,
        default: false,
    },
    transform: {
        type: String,
        default: '',
    },
})
const routeToServer = () => {
    router.push({
        path: '/layout',
    })
}
</script>
