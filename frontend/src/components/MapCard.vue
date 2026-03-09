<template>
    <div w-66 h-102 rounded-md overflow-hidden cursor-pointer class="group" relative @click="routeToServer" scale-80>
        <div w-full h-full group-hover:scale-110 duration-300 ease-in-out overflow-hidden>
            <div w-full h-full overflow-hidden relative>
                <img :src="image" :style="{ transform: transform, transformOrigin: 'left top' }" />
            </div>
        </div>
        <div w-full h-20 absolute group-hover:bottom-0 bottom--100% duration-300 ease-in-out left-0 flex-center
            bg-dark-black-2 opacity-60>
            <div text-white text-center text-2xl font-bold>{{ chineseName }}</div>
        </div>
        <ConfirmDialogWithForm v-model="showConfirmDialog" :title="'请输入服务器名称'" :message="tips" @confirm="handleConfirm">
            <template #form>
                <div class="flex flex-col items-center justify-center">
                    <div class="text-lg font-bold mb-4">请输入服务器名称</div>
                    <input v-model="serverName" type="text" placeholder="请输入服务器名称" class="w-40 h-8 px-2 mb-4" />
                </div>
            </template>
        </ConfirmDialogWithForm>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import ConfirmDialogWithForm from '@/components/common/ConfirmDialogWithForm.vue'
import { ref } from 'vue'
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
const handleConfirm = () => {
    router.push({
        path: '/layout/newServerConfig',
        query: {
            mapName: props.mapName,
            missionName: props.missionName,
            serverName: serverName.value,
        }
    })
}
const serverName = ref('')
const showConfirmDialog = ref(false)
const routeToServer = () => {
    showConfirmDialog.value = true
    // router.push({
    //     path: '/layout/newServerConfig',
    //     query: {
    //         mapName: props.mapName,
    //         missionName: props.missionName,
    //     }
    // })
}
</script>
