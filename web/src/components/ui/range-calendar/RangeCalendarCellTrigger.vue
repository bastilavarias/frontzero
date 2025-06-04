<script lang="ts" setup>
import type { HTMLAttributes } from 'vue'
import { reactiveOmit } from '@vueuse/core'
import { RangeCalendarCellTrigger, type RangeCalendarCellTriggerProps, useForwardProps } from 'reka-ui'
import { cn } from '@/lib/utils'
import { buttonVariants } from '@/components/ui/button'

const props = withDefaults(defineProps<RangeCalendarCellTriggerProps & { class?: HTMLAttributes['class'] }>(), {
  as: 'button',
})

const delegatedProps = reactiveOmit(props, 'class')

const forwardedProps = useForwardProps(delegatedProps)
</script>

<template>
  <RangeCalendarCellTrigger
    data-slot="range-calendar-trigger"
    :class="cn(
      buttonVariants({ variant: 'ghost' }),
      'h-8 w-8 p-0 font-normal data-[selected]:opacity-100',
      '[&[data-today]:not([data-selected])]:bg-neutral-100 [&[data-today]:not([data-selected])]:text-neutral-900 dark:[&[data-today]:not([data-selected])]:bg-neutral-800 dark:[&[data-today]:not([data-selected])]:text-neutral-50',
      // Selection Start
      'data-[selection-start]:bg-neutral-900 data-[selection-start]:text-neutral-50 data-[selection-start]:hover:bg-neutral-900 data-[selection-start]:hover:text-neutral-50 data-[selection-start]:focus:bg-neutral-900 data-[selection-start]:focus:text-neutral-50 dark:data-[selection-start]:bg-neutral-50 dark:data-[selection-start]:text-neutral-900 dark:data-[selection-start]:hover:bg-neutral-50 dark:data-[selection-start]:hover:text-neutral-900 dark:data-[selection-start]:focus:bg-neutral-50 dark:data-[selection-start]:focus:text-neutral-900',
      // Selection End
      'data-[selection-end]:bg-neutral-900 data-[selection-end]:text-neutral-50 data-[selection-end]:hover:bg-neutral-900 data-[selection-end]:hover:text-neutral-50 data-[selection-end]:focus:bg-neutral-900 data-[selection-end]:focus:text-neutral-50 dark:data-[selection-end]:bg-neutral-50 dark:data-[selection-end]:text-neutral-900 dark:data-[selection-end]:hover:bg-neutral-50 dark:data-[selection-end]:hover:text-neutral-900 dark:data-[selection-end]:focus:bg-neutral-50 dark:data-[selection-end]:focus:text-neutral-900',
      // Outside months
      'data-[outside-view]:text-neutral-500 dark:data-[outside-view]:text-neutral-400',
      // Disabled
      'data-[disabled]:text-neutral-500 data-[disabled]:opacity-50 dark:data-[disabled]:text-neutral-400',
      // Unavailable
      'data-[unavailable]:text-neutral-50 data-[unavailable]:line-through dark:data-[unavailable]:text-neutral-50',
      props.class,
    )"
    v-bind="forwardedProps"
  >
    <slot />
  </RangeCalendarCellTrigger>
</template>
