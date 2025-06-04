import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"
import { type Ref } from "vue"

type Updater<T> = T | ((old: T) => T)

/**
 * Merges Tailwind CSS classes, handling conditional classes and conflicts.
 * This is the `cn` utility commonly used in shadcn/ui and shadcn-vue.
 *
 * @param inputs - An array of ClassValue (string, object, array, or boolean).
 * @returns A merged string of Tailwind CSS classes.
 */
export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs))
}

/**
 * Updates a Vue Ref object's value based on an updater.
 * The updater can be either a direct new value or a function that takes the current value
 * and returns the new value. This pattern is common with TanStack Table's state updates.
 *
 * @template T - The type of the Ref's value.
 * @param {Updater<T>} updaterOrValue - The new value or a function to compute the new value.
 * @param {Ref<T>} ref - The Vue Ref object to update.
 */
export function valueUpdater<T>(updaterOrValue: Updater<T>, ref: Ref<T>) {
    // Check if the updater is a function
    if (typeof updaterOrValue === "function") {
        // If it's a function, call it with the current ref value to get the new value
        ref.value = (updaterOrValue as (old: T) => T)(ref.value)
    } else {
        // If it's not a function, it's a direct new value, so assign it
        ref.value = updaterOrValue
    }
}
