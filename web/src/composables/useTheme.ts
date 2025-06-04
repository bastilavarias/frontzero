import { watchEffect, onMounted } from "vue"
import { useDark, useToggle } from "@vueuse/core"

export function useTheme() {
    const isDark = useDark() // Automatically checks localStorage and prefers-color-scheme
    const toggleDark = useToggle(isDark)

    onMounted(() => {
        document.documentElement.classList.toggle("dark", isDark.value)
    })

    // Watch for changes in isDark and update the class on the html element
    watchEffect(() => {
        document.documentElement.classList.toggle("dark", isDark.value)
    })

    return {
        isDark,
        toggleDark,
    }
}
