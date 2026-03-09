<template>
    <div size-full box-border overflow-hidden>
        <div size-full overflow-auto p-4 box-border class="custom-scrollbar">
            <div w-full grid grid-cols-3 gap-4>
                <template v-for="item in serverList">
                    <ServerCard aspect-ratio="16/11" :server="item"></ServerCard>
                </template>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import ServerCard from '@/components/ServerCard.vue'
import { ref } from 'vue'
import { Server } from '@/types/server.ts'
import { useLayoutStore } from '@/store/modules/layout.ts'

const { setShowAside } = useLayoutStore()
onMounted(() => {
    setShowAside(true)
})
const serverList = ref<Server[]>([])
onMounted(() => {
    // 初始化服务器列表
    serverList.value = [
        {
            id: 1,
            name: '服务器1',
            port: 2302,
        }
    ]
    // 添加新服务器
    serverList.value.push({
        newServer: true,
    })
})


</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
    width: 6px;
    height: 6px;
}

.custom-scrollbar::-webkit-scrollbar-track {
    background: transparent;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
    background-color: rgba(0, 0, 0, 0.4);
    border-radius: 3px;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
    background-color: rgba(0, 0, 0, 0.5);
}

:global(html.dark) .custom-scrollbar::-webkit-scrollbar-thumb {
    background-color: rgba(255, 255, 255, 0.4);
}

:global(html.dark) .custom-scrollbar::-webkit-scrollbar-thumb:hover {
    background-color: rgba(255, 255, 255, 0.5);
}
</style>
