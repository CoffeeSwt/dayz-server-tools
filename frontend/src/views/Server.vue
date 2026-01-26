<template>
    <div size-full box-border overflow-hidden>
        <div size-full overflow-auto p-4 box-border class="custom-scrollbar">
            <div w-full grid grid-cols-4 gap-4>
                <template v-for="item in serverList">
                    <ServerCard aspect-ratio="5/4" bg-red />
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
const serverList = ref<Server[]>([])

onMounted(() => {
    serverList.value = [
        {
            id: 1,
            name: '服务器1',
            ip: '127.0.0.1',
            port: 2302,
            password: '123456',
            status: '运行中',
        },
        {
            id: 1,
            name: '服务器1',
            ip: '127.0.0.1',
            port: 2302,
            password: '123456',
            status: '运行中',
        },
        {
            id: 1,
            name: '服务器1',
            ip: '127.0.0.1',
            port: 2302,
            password: '123456',
            status: '运行中',
        },
        {
            id: 1,
            name: '服务器1',
            ip: '127.0.0.1',
            port: 2302,
            password: '123456',
            status: '运行中',
        },
    ]
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
