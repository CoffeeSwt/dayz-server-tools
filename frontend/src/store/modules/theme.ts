import { defineStore } from 'pinia'
import { useDark, useToggle, useMouse } from '@vueuse/core'
// 第一个参数是应用程序中 store 的唯一 id
const useThemeStore = defineStore('theme', () => {
    const isDark = useDark()
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
    return { toggleDark, isDark }
})

export { useThemeStore } 
