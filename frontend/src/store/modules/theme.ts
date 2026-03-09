import { defineStore } from 'pinia'
import { useDark, useToggle, useMouse } from '@vueuse/core'
import { ref } from 'vue'
import { WindowToggleMaximise } from '@wails/runtime/runtime.js'
// 第一个参数是应用程序中 store 的唯一 id
const useThemeStore = defineStore('theme', () => {
    const isDark = useDark()
    const isMaximised = ref(false)
    const toggleDarkFunc = useToggle(isDark)
    const { x, y } = useMouse()
    const toggleDark = () => {
        document.documentElement.style.setProperty('--x', x.value + 'px')
        document.documentElement.style.setProperty('--y', y.value + 'px')
        if ((document as any).startViewTransition) {
            (document as any).startViewTransition(() => {
                toggleDarkFunc()
            });
        } else {
            toggleDarkFunc()
        }
    }
    const toggleMaximise = () => {
        WindowToggleMaximise()
        isMaximised.value = !isMaximised.value
        console.log(isMaximised.value)
    }
    return { toggleDark, isDark, isMaximised, toggleMaximise }
})

export { useThemeStore } 
