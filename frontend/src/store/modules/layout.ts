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
    const showLogo = ref(true)
    const toggleShowLogo = () => {
        showLogo.value = !showLogo.value
    }
    const setShowLogo = (newVal: boolean) => {
        showLogo.value = newVal
    }
    return { showAside, toggleShowAside, setShowAside, showLogo, toggleShowLogo, setShowLogo }
})

export { useLayoutStore } 
