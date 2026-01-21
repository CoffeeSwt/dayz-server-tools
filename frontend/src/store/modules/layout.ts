import { defineStore } from 'pinia'
import { ref } from 'vue'

// 第一个参数是应用程序中 store 的唯一 id
const useLayoutStore = defineStore('layout', () => {
    const showAside = ref(true)
    const toggleShowAside = () => {
        showAside.value = !showAside.value
    }
    const setShowAside = (newVal: boolean) => {
        showAside.value = newVal
    }
    return { showAside, toggleShowAside, setShowAside }
})

export { useLayoutStore } 
