<script setup lang="ts">
import { defineEmits, defineProps } from "vue"
import {
    Sheet,
    SheetContent,
    SheetFooter,
    SheetHeader,
    SheetTitle,
} from "@/components/ui/sheet"
import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import { Switch } from "@/components/ui/switch"
import { Input } from "@/components/ui/input"

defineProps({
    open: {
        type: Boolean,
        default: false,
    },
})

const emits = defineEmits(["update:open"])
</script>

<template>
    <Sheet :open="open" @update:open="$emit('update:open', $event)">
        <SheetContent
            class="w-full sm:max-w-sm md:max-w-md lg:max-w-lg xl:max-w-xl flex flex-col h-full"
        >
            <SheetHeader class="sticky top-0 z-10 p-4 border-b">
                <SheetTitle>New users record</SheetTitle>
            </SheetHeader>

            <div class="flex-grow overflow-y-auto px-4">
                <div class="grid gap-4 py-4">
                    <div class="space-y-2">
                        <Label for="id">id</Label>
                        <Input
                            id="id"
                            placeholder="Leave empty to auto generate..."
                        />
                    </div>

                    <div class="space-y-2">
                        <Label for="email">Email *</Label>
                        <div class="flex items-center space-x-2">
                            <Input id="email" type="email" />
                            <span class="text-sm text-gray-500"
                                >Public: Off</span
                            >
                        </div>
                    </div>

                    <div class="grid grid-cols-2 gap-4">
                        <div class="space-y-2">
                            <Label for="password">Password *</Label>
                            <Input id="password" type="password" />
                        </div>
                        <div class="space-y-2">
                            <Label for="password-confirm"
                                >Password confirm *</Label
                            >
                            <Input id="password-confirm" type="password" />
                        </div>
                    </div>

                    <div class="flex items-center space-x-2">
                        <Switch id="verified" />
                        <Label for="verified">Verified</Label>
                    </div>

                    <div class="space-y-2">
                        <Label for="username">Username</Label>
                        <Input
                            id="username"
                            placeholder="Leave empty to autogenerate..."
                        />
                    </div>

                    <div class="space-y-2">
                        <Label for="name">Name</Label>
                        <Input id="name" />
                    </div>

                    <div class="space-y-2">
                        <Label for="avatar">Avatar</Label>
                        <Input id="avatar" type="file" class="hidden" />
                        <label
                            for="avatar"
                            class="flex items-center justify-center p-4 border rounded-md cursor-pointer hover:bg-gray-50"
                        >
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                class="h-5 w-5 mr-2"
                                viewBox="0 0 20 20"
                                fill="currentColor"
                            >
                                <path
                                    fill-rule="evenodd"
                                    d="M3 17a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM6.293 6.707a1 1 0 010-1.414l3-3a1 1 0 011.414 0l3 3a1 1 0 01-1.414 1.414L11 5.414V13a1 1 0 11-2 0V5.414L6.293 6.707a1 1 0 01-1.414 0z"
                                    clip-rule="evenodd"
                                />
                            </svg>
                            Upload new file
                        </label>
                    </div>

                    <div class="space-y-2">
                        <Label for="website">Website</Label>
                        <Input id="website" />
                    </div>

                    <div class="space-y-2">
                        <Label for="id">id</Label>
                        <Input
                            id="id"
                            placeholder="Leave empty to auto generate..."
                        />
                    </div>

                    <div class="space-y-2">
                        <Label for="email">Email *</Label>
                        <div class="flex items-center space-x-2">
                            <Input id="email" type="email" />
                            <span class="text-sm text-gray-500"
                                >Public: Off</span
                            >
                        </div>
                    </div>

                    <div class="grid grid-cols-2 gap-4">
                        <div class="space-y-2">
                            <Label for="password">Password *</Label>
                            <Input id="password" type="password" />
                        </div>
                        <div class="space-y-2">
                            <Label for="password-confirm"
                                >Password confirm *</Label
                            >
                            <Input id="password-confirm" type="password" />
                        </div>
                    </div>

                    <div class="flex items-center space-x-2">
                        <Switch id="verified" />
                        <Label for="verified">Verified</Label>
                    </div>

                    <div class="space-y-2">
                        <Label for="username">Username</Label>
                        <Input
                            id="username"
                            placeholder="Leave empty to autogenerate..."
                        />
                    </div>

                    <div class="space-y-2">
                        <Label for="name">Name</Label>
                        <Input id="name" />
                    </div>

                    <div class="space-y-2">
                        <Label for="avatar">Avatar</Label>
                        <Input id="avatar" type="file" class="hidden" />
                        <label
                            for="avatar"
                            class="flex items-center justify-center p-4 border rounded-md cursor-pointer"
                        >
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                class="h-5 w-5 mr-2"
                                viewBox="0 0 20 20"
                                fill="currentColor"
                            >
                                <path
                                    fill-rule="evenodd"
                                    d="M3 17a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM6.293 6.707a1 1 0 010-1.414l3-3a1 1 0 011.414 0l3 3a1 1 0 01-1.414 1.414L11 5.414V13a1 1 0 11-2 0V5.414L6.293 6.707a1 1 0 01-1.414 0z"
                                    clip-rule="evenodd"
                                />
                            </svg>
                            Upload new file
                        </label>
                    </div>

                    <div class="space-y-2">
                        <Label for="website">Website</Label>
                        <Input id="website" />
                    </div>
                </div>
            </div>

            <SheetFooter
                class="items-end sticky bottom-0 z-10 px-4 pb-4 pt-2 border-t flex justify-between"
            >
                <div></div>
                <div class="flex gap-2">
                    <Button
                        variant="ghost"
                        @click="$emit('update:open', false)"
                    >
                        Cancel
                    </Button>
                    <Button>Create</Button>
                </div>
            </SheetFooter>
        </SheetContent>
    </Sheet>
</template>
