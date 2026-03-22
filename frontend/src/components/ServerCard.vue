<template>
    <div rounded-lg bg-red overflow-hidden>
        <div v-if="!server.newServer" @click="testServer" bg-black size-full flex-center cursor-pointer></div>
        <div size-full class="group" v-if="server.newServer" flex-center cursor-pointer relative overflow-hidden
            @click="handleNewServer">
            <div size-full group-hover:scale-110 duration-300 ease-in-out overflow-hidden>
                <div size-full overflow-hidden relative>
                    <img :src="newServerBgImagePath" :style="{ transform: transform, transformOrigin: 'left top' }" />
                </div>
            </div>
            <div absolute-center i-oui:ml-create-single-metric-job text-6xl class="text-color"></div>
            <div w-full absolute group-hover:bottom-0 bottom--100% duration-300 ease-in-out left-0 flex-center
                bg-dark-black-2 opacity-60>
                <div text-white text-center text-2xl font-bold>新建服务器</div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { Server } from '@/types/server.ts'
import { useRouter } from 'vue-router'
import { TestServerStart } from '@wails/go/app/Server.js'
const router = useRouter()
const newServerBgImagePath = '/static/images/new_server_bg.png'
const transform = 'translate(-158px, -91px) scale(0.17)'
const props = defineProps({
    server: {
        type: Object as () => Server,
        default: () => ({})
    }
})
const handleNewServer = () => {
    router.push({
        path: '/layout/map',
    })
}

const testServer = () => {
    TestServerStart()
}
</script>


<style scoped>
.text-color {
    color: #b9b9b9;
}
</style>
