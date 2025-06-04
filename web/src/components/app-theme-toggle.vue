<script setup lang="ts">
import { useTheme } from "@/composables/useTheme"
import { Sun, Moon } from "lucide-vue-next"
import { Button } from "@/components/ui/button"

const { isDark, toggleDark } = useTheme()
</script>

<template>
    <Button
        variant="ghost"
        size="icon"
        @click="toggleDark()"
        :aria-label="buttonLabel"
    >
        <Sun
            class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
        />
        <Moon
            class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
        />
    </Button>
</template>
