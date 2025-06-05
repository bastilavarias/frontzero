<script setup lang="ts">
import {
    ResizableHandle,
    ResizablePanel,
    ResizablePanelGroup,
} from "@/components/ui/resizable"
import CollectionTable from "@/modules/collection/components/collection-table.vue"
import CollectionList from "@/modules/collection/components/collection-list.vue"
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
