<script setup lang="ts">
import {
    ResizableHandle,
    ResizablePanel,
    ResizablePanelGroup,
} from "@/components/ui/resizable"
import CollectionList from "@/components/collection-list.vue"
import CollectionTable from "@/components/collection-table.vue"
</script>

<template>
    <div class="flex h-screen">
        <ResizablePanelGroup
            id="resize-panel-group-1"
            direction="horizontal"
            class="h-full items-stretch"
        >
            <ResizablePanel id="resize-panel-1" :min-size="15" :max-size="20">
                <CollectionList />
            </ResizablePanel>
            <ResizableHandle id="resize-handle-1" with-handle />
            <ResizablePanel id="resize-panel-2">
                <CollectionTable />
            </ResizablePanel>
        </ResizablePanelGroup>
    </div>
</template>
