<script setup lang="ts">
import type { HTMLAttributes } from 'vue'
import { reactiveOmit } from '@vueuse/core'
import {
  DropdownMenuSeparator,
  type DropdownMenuSeparatorProps,
} from 'reka-ui'
import { cn } from '@/lib/utils'

const props = defineProps<DropdownMenuSeparatorProps & {
  class?: HTMLAttributes['class']
}>()

const delegatedProps = reactiveOmit(props, 'class')
</script>

<template>
  <DropdownMenuSeparator
    data-slot="dropdown-menu-separator"
    v-bind="delegatedProps"
    :class="cn('bg-neutral-200 -mx-1 my-1 h-px dark:bg-neutral-800', props.class)"
  />
</template>
