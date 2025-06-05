<script setup lang="ts">
import { defineProps, defineEmits } from "vue"
import { Sheet, SheetContent } from "@/components/ui/sheet" // Adjust import path as needed
import { Button } from "@/components/ui/button" // Adjust import path as needed
import { Separator } from "@/components/ui/separator" // Adjust import path as needed

// Define the 'open' prop that will be used with v-model
const props = defineProps({
    open: {
        type: Boolean,
        default: false,
    },
})

// Define the 'update:open' emit event for v-model support
const emits = defineEmits(["update:open"])

const documentations = {
    crud: [
        {
            title: "List/Search",
        },
        {
            title: "View",
        },
        {
            title: "Create",
        },
        {
            title: "Update",
        },
        {
            title: "Delete",
        },
        {
            title: "Realtime",
        },
        {
            title: "Batch",
        },
    ],

    auth: [
        {
            title: "List auth methods",
        },
        {
            title: "Auth with password",
        },
        {
            title: "Auth with OAuth2",
        },
        {
            title: "Auth with OTP",
        },
        {
            title: "Auth refresh",
        },
        {
            title: "Verification",
        },
        {
            title: "Password reset",
        },
        {
            title: "Email change",
        },
    ],
}
</script>

<template>
    <Sheet :open="open" @update:open="$emit('update:open', $event)">
        <SheetContent
            class="w-full h-screen sm:max-w-lg md:max-w-xl lg:max-w-2xl xl:max-w-3xl flex flex-col h-full py-0"
        >
            <div class="grid grid-cols-12 gap-4 h-full">
                <div class="col-span-4 px-4 h-full py-4 border-r space-y-4">
                    <p class="text-muted-foreground text-xs">Documentation</p>

                    <nav
                        class="flex space-x-2 lg:flex-col lg:space-x-0 lg:space-y-1"
                    >
                        <Button
                            v-for="item in documentations.crud"
                            :key="item.title"
                            as="a"
                            variant="ghost"
                            class="flex justify-start items-center"
                        >
                            {{ item.title }}
                        </Button>
                    </nav>

                    <Separator />

                    <nav
                        class="flex space-x-2 lg:flex-col lg:space-x-0 lg:space-y-1"
                    >
                        <Button
                            v-for="item in documentations.auth"
                            :key="item.title"
                            as="a"
                            variant="ghost"
                            class="flex justify-start items-center"
                        >
                            {{ item.title }}
                        </Button>
                    </nav>
                </div>

                <div class="col-span-8 py-8">
                    <h4 class="text-xl font-semibold mb-3 text-foreground">
                        List/Search (users)
                    </h4>
                    <p class="text-sm text-muted-foreground mb-4">
                        Fetch a paginated users records list, supporting sorting
                        and filtering.
                    </p>
                    <div
                        class="bg-secondary/5 p-4 rounded-md text-sm font-mono text-foreground"
                    >
                        <pre><code>import PocketBase from 'pocketbase';

const pb = new PocketBase('https://pocketbase.io');

// fetch a paginated records list
const resultList = await pb.collection('users').getList(1, 50, {
  filter: 'someField1 != someField2',
});

// you can also fetch all records at once via getFullList
const records = await pb.collection('users').getFullList({
  sort: '-someField',
});

// or fetch only the first record that matches the specified filter
const record = await pb.collection('users').getFirstListItem('someField="test"', {
  expand: 'relField1,relField2.subRelField',
});</code></pre>
                    </div>
                    <h4 class="text-lg font-semibold mt-6 mb-2 text-foreground">
                        API Details
                    </h4>
                    <div
                        class="bg-muted/50 text-muted-foreground px-4 py-2 rounded-md font-mono inline-block mb-4"
                    >
                        GET /api/collections/users/records
                    </div>
                    <h4 class="text-lg font-semibold mb-2 text-foreground">
                        Query parameters
                    </h4>
                    <table
                        class="w-full text-sm text-left text-muted-foreground border-collapse"
                    >
                        <thead
                            class="text-xs text-foreground uppercase bg-muted"
                        >
                            <tr>
                                <th
                                    scope="col"
                                    class="py-2 px-3 border border-border"
                                >
                                    Param
                                </th>
                                <th
                                    scope="col"
                                    class="py-2 px-3 border border-border"
                                >
                                    Type
                                </th>
                                <th
                                    scope="col"
                                    class="py-2 px-3 border border-border"
                                >
                                    Description
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr class="bg-card border-b border-border">
                                <td class="py-2 px-3 border border-border">
                                    page
                                </td>
                                <td class="py-2 px-3 border border-border">
                                    Number
                                </td>
                                <td class="py-2 px-3 border border-border">
                                    The page (aka. offset) of the paginated list
                                    (default to 1).
                                </td>
                            </tr>
                            <tr class="bg-card border-b border-border">
                                <td class="py-2 px-3 border border-border">
                                    perPage
                                </td>
                                <td class="py-2 px-3 border border-border">
                                    Number
                                </td>
                                <td class="py-2 px-3 border border-border">
                                    Specify the max returned records per page
                                    (default to 30).
                                </td>
                            </tr>
                            <tr class="bg-card">
                                <td class="py-2 px-3 border border-border">
                                    sort
                                </td>
                                <td class="py-2 px-3 border border-border">
                                    String
                                </td>
                                <td class="py-2 px-3 border border-border">
                                    Specify the records order attribute(s). Add
                                    +/- (default) in front of the attribute for
                                    DESC/ASC.
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </SheetContent>
    </Sheet>
</template>
