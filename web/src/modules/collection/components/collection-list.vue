<script setup lang="ts">
import { ChevronRight, Folder, Search, Plus } from "lucide-vue-next"
import {
    SidebarGroup,
    SidebarMenu,
    SidebarMenuButton,
    SidebarMenuItem,
    SidebarMenuSub,
    SidebarMenuSubButton,
    SidebarMenuSubItem,
} from "@/components/ui/sidebar"
import {
    Collapsible,
    CollapsibleContent,
    CollapsibleTrigger,
} from "@/components/ui/collapsible"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { ScrollArea } from "@/components/ui/scroll-area"

const collections = [
    {
        title: "Users",
        isActive: true,
        items: [
            { title: "id" },
            { title: "name" },
            { title: "email" },
            { title: "address" },
            { title: "created_at" },
        ],
    },
    {
        title: "Products",
        isActive: false,
        items: [
            { title: "id" },
            { title: "name" },
            { title: "description" },
            { title: "price" },
            { title: "stock_quantity" },
        ],
    },
    {
        title: "Orders",
        isActive: false,
        items: [
            { title: "id" },
            { title: "user_id" },
            { title: "total_amount" },
            { title: "status" },
            { title: "created_at" },
        ],
    },
    {
        title: "Payments",
        isActive: false,
        items: [
            { title: "id" },
            { title: "order_id" },
            { title: "payment_method" },
            { title: "payment_date" },
            { title: "amount" },
        ],
    },
    {
        title: "Reviews",
        isActive: false,
        items: [
            { title: "id" },
            { title: "user_id" },
            { title: "product_id" },
            { title: "rating" },
            { title: "comment" },
        ],
    },
]
</script>

<template>
    <ScrollArea class="h-full w-full">
        <div class="flex flex-col items-center items-start p-4 gap-4">
            <div class="relative w-full items-center">
                <Input
                    id="search"
                    type="text"
                    placeholder="Search collections..."
                    class="pl-10"
                />
                <span
                    class="absolute start-0 inset-y-0 flex items-center justify-center px-2"
                >
                    <Search class="size-6 text-muted-foreground" />
                </span>
            </div>
            <SidebarGroup class="group-data-[collapsible=icon]:hidden p-0 m-0">
                <SidebarMenu>
                    <Collapsible
                        v-for="item in collections"
                        :key="item.title"
                        as-child
                        :default-open="item.isActive"
                        class="group/collapsible"
                    >
                        <SidebarMenuItem>
                            <CollapsibleTrigger as-child>
                                <SidebarMenuButton
                                    :tooltip="item.title"
                                    :is-active="item.isActive"
                                >
                                    <Folder />
                                    <span>{{ item.title }}</span>
                                    <ChevronRight
                                        class="ml-auto transition-transform duration-200 group-data-[state=open]/collapsible:rotate-90"
                                    />
                                </SidebarMenuButton>
                            </CollapsibleTrigger>
                            <CollapsibleContent>
                                <SidebarMenuSub>
                                    <SidebarMenuSubItem
                                        v-for="subItem in item.items"
                                        :key="subItem.title"
                                    >
                                        <SidebarMenuSubButton as-child>
                                            <span>{{ subItem.title }}</span>
                                        </SidebarMenuSubButton>
                                    </SidebarMenuSubItem>
                                </SidebarMenuSub>
                            </CollapsibleContent>
                        </SidebarMenuItem>
                    </Collapsible>
                </SidebarMenu>
            </SidebarGroup>

            <Button variant="outline" class="w-full">
                <Plus />
                New collection
            </Button>
        </div>
    </ScrollArea>
</template>
