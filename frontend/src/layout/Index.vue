<template>
    <div w-full h-full flex flex-col>
        <div w-full h-16 shrink-0 text-white draggable>
            <Header></Header>
        </div>
        <div w-full flex flex-grow-1 overflow-hidden>
            <div h-full bg-gray-200 flex-shrink-0 overflow-hidden :class="{ 'w-0': !showAside, 'w-51': showAside }"
                duration="300ms">
                <Aside></Aside>
            </div>
            <div w-full>
                <Main></Main>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import Header from './header/Header.vue'
import Main from './main/Main.vue'
import Aside from './aside/Aside.vue'
import { computed } from 'vue'
import { useLayoutStore } from '@/store/modules/layout.ts'
const layoutStore = useLayoutStore()
const showAside = computed(() => layoutStore.showAside)
</script>
