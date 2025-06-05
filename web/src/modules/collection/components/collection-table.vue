<script setup>
import { h, ref, computed, watch } from "vue"
import {
    useVueTable,
    getCoreRowModel,
    getPaginationRowModel,
    getSortedRowModel,
    getFilteredRowModel,
    getFacetedRowModel,
    getFacetedUniqueValues,
    FlexRender,
} from "@tanstack/vue-table"
import { valueUpdater } from "@/lib/utils.js"
import { Button } from "@/components/ui/button/index.js"
import { Checkbox } from "@/components/ui/checkbox/index.js"
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu/index.js"
import { Input } from "@/components/ui/input/index.js"
import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table/index.js"
import { Badge } from "@/components/ui/badge/index.js"
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select/index.js"
import {
    Command,
    CommandEmpty,
    CommandGroup,
    CommandInput,
    CommandItem,
    CommandList,
    CommandSeparator,
} from "@/components/ui/command/index.js"
import {
    ArrowDownIcon,
    ArrowUpIcon,
    SortAscIcon,
    CheckIcon,
    ChevronLeftIcon,
    ChevronRightIcon,
    ChevronsLeftIcon,
    ChevronsRightIcon,
    CircleIcon,
    Search,
    Code2,
    Plus,
} from "lucide-vue-next"
import CollectionFormSheet from "@/modules/collection/components/collection-form-sheet.vue"
import CollectionPreviewSheet from "@/modules/collection/components/collection-preview-sheet.vue"
const cn = (...classes) => classes.filter(Boolean).join(" ")

const statuses = [
    { value: "backlog", label: "Backlog", icon: CircleIcon },
    { value: "todo", label: "Todo", icon: CircleIcon },
    { value: "in progress", label: "In Progress", icon: CircleIcon },
    { value: "done", label: "Done", icon: CircleIcon },
    { value: "canceled", label: "Canceled", icon: CircleIcon },
]

const priorities = [
    { value: "low", label: "Low", icon: ArrowDownIcon },
    { value: "medium", label: "Medium", icon: SortAscIcon },
    { value: "high", label: "High", icon: ArrowUpIcon },
]

const labels = [
    { value: "bug", label: "Bug" },
    { value: "feature", label: "Feature" },
    { value: "enhancement", label: "Enhancement" },
]

const tasks = ref([
    {
        id: "TASK-8782",
        title: "You can't compress the program without quantifying the open-source software.",
        status: "in progress",
        label: "documentation",
        priority: "high",
    },
    {
        id: "TASK-7878",
        title: "Try to input the SAS program, maybe it will index the virtual hard drive!",
        status: "backlog",
        label: "feature",
        priority: "medium",
    },
    {
        id: "TASK-7839",
        title: "We need to bypass the neural system!",
        status: "todo",
        label: "bug",
        priority: "low",
    },
    {
        id: "TASK-7653",
        title: "The AI algorithm is down, parse the system then you can just input the back-end system!",
        status: "done",
        label: "feature",
        priority: "high",
    },
    {
        id: "TASK-7652",
        title: "Use the mobile application, then you can just transmit the CSS application!",
        status: "canceled",
        label: "bug",
        priority: "medium",
    },
    {
        id: "TASK-7432",
        title: "I'll parse the primary XSS circuit, so we can get to the CSS application!",
        status: "backlog",
        label: "documentation",
        priority: "low",
    },
    {
        id: "TASK-7290",
        title: "If we override the interface, we can get to the HTTP interface through the virtual EXE interface!",
        status: "in progress",
        label: "enhancement",
        priority: "high",
    },
    {
        id: "TASK-7153",
        title: "We need to program the primary JSON program!",
        status: "todo",
        label: "feature",
        priority: "medium",
    },
    {
        id: "TASK-7098",
        title: "I'll synthesize the digital TLS circuit, so we can get to the USB protocol!",
        status: "done",
        label: "bug",
        priority: "low",
    },
    {
        id: "TASK-7097",
        title: "You can't generate the program without indexing the virtual UDP program!",
        status: "canceled",
        label: "documentation",
        priority: "high",
    },
    {
        id: "TASK-7096",
        title: "I'll calculate the primary SSL capacitor, so we can get to the FTP protocol!",
        status: "in progress",
        label: "feature",
        priority: "medium",
    },
    {
        id: "TASK-7095",
        title: "We need to connect the digital HTTP interface!",
        status: "todo",
        label: "bug",
        priority: "low",
    },
    {
        id: "TASK-7094",
        title: "If we bypass the interface, we can get to the UDP application through the virtual DNS application!",
        status: "done",
        label: "enhancement",
        priority: "high",
    },
    {
        id: "TASK-7093",
        title: "You can't input the program without transmitting the virtual CSS protocol!",
        status: "canceled",
        label: "documentation",
        priority: "medium",
    },
    {
        id: "TASK-7092",
        title: "I'll quantify the digital PBX circuit, so we can get to the RAM protocol!",
        status: "in progress",
        label: "feature",
        priority: "low",
    },
    {
        id: "TASK-7091",
        title: "We need to program the primary CSS protocol!",
        status: "todo",
        label: "bug",
        priority: "high",
    },
    {
        id: "TASK-7090",
        title: "If we bypass the interface, we can get to the HTTP interface through the virtual DNS application!",
        status: "done",
        label: "enhancement",
        priority: "medium",
    },
    {
        id: "TASK-7089",
        title: "You can't input the program without transmitting the virtual CSS protocol!",
        status: "canceled",
        label: "documentation",
        priority: "low",
    },
])

// --- TanStack Table State Variables ---
const columnFilters = ref([])
const columnVisibility = ref({})
const rowSelection = ref({})
const sorting = ref([])
const globalFilter = ref("")
const isNewRecordFormSheetOpen = ref(false)
const isAPIPreviewSheetOpen = ref(false)

const columns = computed(() => [
    {
        id: "select",
        header: ({ table }) =>
            h(Checkbox, {
                checked:
                    table.getIsAllPageRowsSelected() ||
                    (table.getIsSomePageRowsSelected() && "indeterminate"),
                "onUpdate:checked": (value) =>
                    table.toggleAllPageRowsSelected(!!value),
                ariaLabel: "Select all",
            }),
        cell: ({ row }) =>
            h(Checkbox, {
                checked: row.getIsSelected(),
                "onUpdate:checked": (value) => row.toggleSelected(!!value),
                ariaLabel: "Select row",
                class: "translate-y-[2px]",
            }),
        enableSorting: false,
        enableHiding: false,
    },
    {
        accessorKey: "id",
        header: () => h("div", {}, ["Task ID"]),
        cell: ({ row }) => h("div", { class: "w-[80px]" }, row.getValue("id")),
        enableSorting: false,
        enableHiding: false,
    },
    {
        accessorKey: "title",
        header: ({ column }) =>
            h(
                Button,
                {
                    variant: "ghost",
                    onClick: () =>
                        column.toggleSorting(column.getIsSorted() === "asc"),
                },
                () => [
                    "Title",
                    column.getIsSorted() === "desc"
                        ? h(ArrowUpIcon, { class: "ml-2 h-4 w-4" })
                        : column.getIsSorted() === "asc"
                          ? h(ArrowDownIcon, { class: "ml-2 h-4 w-4" })
                          : h(SortAscIcon, { class: "ml-2 h-4 w-4" }),
                ]
            ),
        cell: ({ row }) => {
            const label = labels.find(
                (label) => label.value === row.original.label
            )
            return h("div", { class: "flex space-x-2" }, [
                label && h(Badge, { variant: "outline" }, () => label.label),
                h(
                    "span",
                    { class: "max-w-[500px] truncate font-medium" },
                    row.getValue("title")
                ),
            ])
        },
    },
    {
        accessorKey: "status",
        header: ({ column }) =>
            h(DataTableColumnHeader, {
                column,
                title: "Status",
                options: statuses,
            }), // Using custom header component
        cell: ({ cell }) => {
            const status = statuses.find(
                (status) => status.value === cell.getValue()
            )
            if (!status) return null
            return h("div", { class: "flex w-[100px] items-center" }, [
                status.icon &&
                    h(status.icon, {
                        class: "mr-2 h-4 w-4 text-muted-foreground",
                    }),
                h("span", {}, status.label),
            ])
        },
        filterFn: (row, id, value) => {
            return value.includes(row.getValue(id))
        },
    },
    {
        accessorKey: "priority",
        header: ({ column }) =>
            h(DataTableColumnHeader, {
                column,
                title: "Priority",
                options: priorities,
            }), // Using custom header component
        cell: ({ cell }) => {
            const priority = priorities.find(
                (priority) => priority.value === cell.getValue()
            )
            if (!priority) return null
            return h("div", { class: "flex items-center" }, [
                priority.icon &&
                    h(priority.icon, {
                        class: "mr-2 h-4 w-4 text-muted-foreground",
                    }),
                h("span", {}, priority.label),
            ])
        },
        filterFn: (row, id, value) => {
            return value.includes(row.getValue(id))
        },
    },
    {
        id: "actions",
        cell: ({ row }) => {
            const task = row.original
            return h(
                DropdownMenu,
                {},
                {
                    trigger: () =>
                        h(
                            Button,
                            {
                                variant: "ghost",
                                class: "h-8 w-8 p-0",
                            },
                            () => h(SortAscIcon, { class: "h-4 w-4" })
                        ),
                    default: () =>
                        h(DropdownMenuContent, { align: "end" }, [
                            h(DropdownMenuLabel, {}, "Actions"),
                            h(
                                DropdownMenuItem,
                                {
                                    onClick: () =>
                                        navigator.clipboard.writeText(task.id),
                                },
                                "Copy task ID"
                            ),
                            h(DropdownMenuSeparator),
                            h(DropdownMenuItem, {}, "View customer"),
                            h(DropdownMenuItem, {}, "View payment details"),
                        ]),
                }
            )
        },
    },
])

const table = useVueTable({
    data: tasks.value,
    columns: columns.value,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    getFacetedRowModel: getFacetedRowModel(),
    getFacetedUniqueValues: getFacetedUniqueValues(),
    state: {
        get columnFilters() {
            return columnFilters.value
        },
        get columnVisibility() {
            return columnVisibility.value
        },
        get rowSelection() {
            return rowSelection.value
        },
        get sorting() {
            return sorting.value
        },
        get globalFilter() {
            return globalFilter.value
        },
    },
    enableRowSelection: true,
    onColumnFiltersChange: (updaterOrValue) =>
        valueUpdater(updaterOrValue, columnFilters),
    onColumnVisibilityChange: (updaterOrValue) =>
        valueUpdater(updaterOrValue, columnVisibility),
    onRowSelectionChange: (updaterOrValue) =>
        valueUpdater(updaterOrValue, rowSelection),
    onSortingChange: (updaterOrValue) => valueUpdater(updaterOrValue, sorting),
    onGlobalFilterChange: (updaterOrValue) =>
        valueUpdater(updaterOrValue, globalFilter),
})

// --- Table Toolbar Logic ---
const isFiltered = computed(
    () => columnFilters.value.length > 0 || globalFilter.value.length > 0
)

const clearFilters = () => {
    table.setColumnFilters([])
    table.setGlobalFilter("")
}

const deleteSelectedRows = () => {
    const selectedRowIds = table
        .getSelectedRowModel()
        .rows.map((row) => row.original.id)
    tasks.value = tasks.value.filter(
        (task) => !selectedRowIds.includes(task.id)
    )
    table.setRowSelection({}) // Clear selection
    // In a real app, you'd send these IDs to an API for deletion
    console.log("Deleted tasks:", selectedRowIds)
}

// --- Custom Column Header Component (for faceted filters) ---
// This component encapsulates the logic for column sorting and faceted filtering
const DataTableColumnHeader = (props) => {
    const { column, title, options } = props

    if (!column.getCanSort() && !column.getCanFilter()) {
        return h(
            "div",
            { class: cn("flex items-center space-x-2", props.class) },
            title
        )
    }

    const sortedUniqueValues = Array.from(
        column.getFacetedUniqueValues().keys()
    )
        .sort()
        .slice(0, 5) // Show top 5 unique values

    const selectedValues = new Set(column.getFilterValue())

    return h("div", { class: cn("flex items-center space-x-2", props.class) }, [
        h(
            DropdownMenu,
            {},
            {
                trigger: () =>
                    h(
                        DropdownMenuTrigger,
                        {
                            asChild: true,
                        },
                        () =>
                            h(
                                Button,
                                {
                                    variant: "ghost",
                                    size: "sm",
                                    class: "-ml-3 h-8 data-[state=open]:bg-accent",
                                },
                                () => [
                                    h("span", {}, title),
                                    column.getCanSort() &&
                                        (column.getIsSorted() === "desc"
                                            ? h(ArrowDownIcon, {
                                                  class: "ml-2 h-4 w-4",
                                              })
                                            : column.getIsSorted() === "asc"
                                              ? h(ArrowUpIcon, {
                                                    class: "ml-2 h-4 w-4",
                                                })
                                              : h(SortAscIcon, {
                                                    class: "ml-2 h-4 w-4",
                                                })),
                                ]
                            )
                    ),
                default: () =>
                    h(
                        DropdownMenuContent,
                        { align: "start", class: "w-[200px]" },
                        [
                            h(DropdownMenuLabel, {}, title),
                            h(DropdownMenuSeparator),
                            h(Command, {}, [
                                h(CommandInput, { placeholder: title }),
                                h(CommandList, {}, [
                                    h(CommandEmpty, {}, "No results found."),
                                    h(CommandGroup, {}, [
                                        options.map((option) => {
                                            const isSelected =
                                                selectedValues.has(option.value)
                                            return h(
                                                CommandItem,
                                                {
                                                    key: option.value,
                                                    onSelect: () => {
                                                        if (isSelected) {
                                                            selectedValues.delete(
                                                                option.value
                                                            )
                                                        } else {
                                                            selectedValues.add(
                                                                option.value
                                                            )
                                                        }
                                                        column.setFilterValue(
                                                            Array.from(
                                                                selectedValues
                                                            )
                                                        )
                                                    },
                                                },
                                                () => [
                                                    h(
                                                        "div",
                                                        {
                                                            class: cn(
                                                                "mr-2 flex h-4 w-4 items-center justify-center rounded-sm border border-primary",
                                                                isSelected
                                                                    ? "bg-primary text-primary-foreground"
                                                                    : "opacity-50 [&_svg]:invisible"
                                                            ),
                                                        },
                                                        h(CheckIcon, {
                                                            class: cn(
                                                                "h-4 w-4"
                                                            ),
                                                        })
                                                    ),
                                                    option.icon &&
                                                        h(option.icon, {
                                                            class: "mr-2 h-4 w-4 text-muted-foreground",
                                                        }),
                                                    h("span", {}, option.label),
                                                ]
                                            )
                                        }),
                                    ]),
                                    sortedUniqueValues.length > 0 &&
                                        h(CommandGroup, { class: "mt-2" }, [
                                            h(
                                                CommandItem,
                                                {
                                                    class: "text-muted-foreground",
                                                },
                                                "Top 5 Values"
                                            ),
                                            sortedUniqueValues.map((value) => {
                                                const option = options.find(
                                                    (opt) => opt.value === value
                                                )
                                                return h(
                                                    CommandItem,
                                                    {
                                                        key: value,
                                                        onSelect: () => {
                                                            selectedValues.clear()
                                                            selectedValues.add(
                                                                value
                                                            )
                                                            column.setFilterValue(
                                                                Array.from(
                                                                    selectedValues
                                                                )
                                                            )
                                                        },
                                                    },
                                                    () => [
                                                        h(
                                                            "div",
                                                            {
                                                                class: "mr-2 flex h-4 w-4 items-center justify-center rounded-sm border border-primary opacity-50",
                                                            },
                                                            h(
                                                                "span",
                                                                {
                                                                    class: "text-[8px] text-primary-foreground",
                                                                },
                                                                value
                                                                    .charAt(0)
                                                                    .toUpperCase()
                                                            )
                                                        ),
                                                        option
                                                            ? h(
                                                                  "span",
                                                                  {},
                                                                  option.label
                                                              )
                                                            : h(
                                                                  "span",
                                                                  {},
                                                                  value
                                                              ),
                                                    ]
                                                )
                                            }),
                                        ]),
                                    selectedValues.size > 0 &&
                                        h(h(CommandSeparator)),
                                    selectedValues.size > 0 &&
                                        h(
                                            h(CommandGroup),
                                            {},
                                            h(
                                                CommandItem,
                                                {
                                                    onSelect: () =>
                                                        column.setFilterValue(
                                                            undefined
                                                        ),
                                                    class: "justify-center text-center",
                                                },
                                                "Clear filters"
                                            )
                                        ),
                                ]),
                            ]),
                        ]
                    ),
            }
        ),
        column.getCanSort() &&
            h(
                Button,
                {
                    variant: "ghost",
                    size: "sm",
                    onClick: () => column.toggleSorting(false), // Toggle sort asc
                    class: "h-8 w-8 p-0",
                },
                () => h(ArrowUpIcon, { class: "h-4 w-4" })
            ), // Ascending sort button
        column.getCanSort() &&
            h(
                Button,
                {
                    variant: "ghost",
                    size: "sm",
                    onClick: () => column.toggleSorting(true), // Toggle sort desc
                    class: "h-8 w-8 p-0",
                },
                () => h(ArrowDownIcon, { class: "h-4 w-4" })
            ), // Descending sort button
        column.getCanHide() &&
            h(
                Button,
                {
                    variant: "ghost",
                    size: "sm",
                    onClick: () => column.toggleVisibility(false),
                    class: "h-8 w-8 p-0",
                },
                () => h(SortAscIcon, { class: "h-4 w-4" })
            ), // Hide column button
    ])
}

// --- Watcher to ensure columnFilters are reset when global filter changes ---
// This is not strictly from shadcn example but good for consistent filtering behavior
watch(globalFilter, () => {
    table.setColumnFilters([])
})
</script>

<template>
    <div>
        <div class="flex flex-col items-center items-start">
            <div class="p-4 flex flex-col w-full gap-4">
                <div class="flex justify-between items-center">
                    <h2 class="text-xl font-bold">Users</h2>
                    <div class="flex gap-2">
                        <Button
                            variant="outline"
                            @click="
                                isAPIPreviewSheetOpen = !isAPIPreviewSheetOpen
                            "
                        >
                            <Code2 /> API Preview</Button
                        >
                        <Button
                            @click="
                                isNewRecordFormSheetOpen =
                                    !isNewRecordFormSheetOpen
                            "
                        >
                            <Plus /> New Record
                        </Button>
                    </div>
                </div>
                <div class="relative items-center">
                    <Input
                        id="search"
                        type="text"
                        placeholder="Search columns..."
                        class="pl-10"
                    />
                    <span
                        class="absolute start-0 inset-y-0 flex items-center justify-center px-2"
                    >
                        <Search class="size-6 text-muted-foreground" />
                    </span>
                </div>
            </div>

            <Table>
                <TableHeader>
                    <TableRow
                        v-for="headerGroup in table.getHeaderGroups()"
                        :key="headerGroup.id"
                    >
                        <TableHead
                            v-for="header in headerGroup.headers"
                            :key="header.id"
                            :colSpan="header.colSpan"
                            :class="
                                cn(
                                    'relative h-12 px-2 text-left align-middle font-medium text-muted-foreground [&:has([role=checkbox])]:pr-0',
                                    header.column.getCanSort() &&
                                        'cursor-pointer',
                                    header.column.id === 'select' && 'w-[3rem]',
                                    header.column.id === 'actions' && 'w-[4rem]' // Fixed width for actions column
                                )
                            "
                        >
                            <FlexRender
                                v-if="!header.isPlaceholder"
                                :render="header.column.columnDef.header"
                                :props="header.getContext()"
                            />
                        </TableHead>
                    </TableRow>
                </TableHeader>
                <TableBody>
                    <template v-if="table.getRowModel().rows?.length">
                        <TableRow
                            v-for="row in table.getRowModel().rows"
                            :key="row.id"
                            :data-state="row.getIsSelected() && 'selected'"
                        >
                            <TableCell
                                v-for="cell in row.getVisibleCells()"
                                :key="cell.id"
                                :class="
                                    cn(
                                        'p-2 align-middle [&:has([role=checkbox])]:pr-0',
                                        cell.column.id === 'select' &&
                                            'text-center', // Center align select checkbox
                                        cell.column.id === 'actions' &&
                                            'text-center', // Center align actions menu
                                        cell.column.id === 'amount' &&
                                            'text-right' // Right align amount
                                    )
                                "
                            >
                                <FlexRender
                                    :render="cell.column.columnDef.cell"
                                    :props="cell.getContext()"
                                />
                            </TableCell>
                        </TableRow>
                    </template>
                    <template v-else>
                        <TableRow>
                            <TableCell
                                :colSpan="columns.length"
                                class="h-24 text-center"
                            >
                                No results.
                            </TableCell>
                        </TableRow>
                    </template>
                </TableBody>
            </Table>

            <div class="flex items-center justify-end space-x-2 p-4">
                <div class="flex-1 text-sm text-muted-foreground">
                    {{ table.getFilteredSelectedRowModel().rows.length }} of
                    {{ table.getFilteredRowModel().rows.length }} row(s)
                    selected.
                </div>
                <div class="flex items-center space-x-2">
                    <p class="text-sm font-medium">Rows per page</p>
                    <Select
                        :model-value="`${table.getState().pagination.pageSize}`"
                        @update:model-value="
                            (value) => table.setPageSize(Number(value))
                        "
                    >
                        <SelectTrigger class="h-8 w-[70px]">
                            <SelectValue
                                :placeholder="
                                    table.getState().pagination.pageSize
                                "
                            />
                        </SelectTrigger>
                        <SelectContent>
                            <SelectItem value="5">5</SelectItem>
                            <SelectItem value="10">10</SelectItem>
                            <SelectItem value="20">20</SelectItem>
                            <SelectItem value="50">50</SelectItem>
                            <SelectItem value="100">100</SelectItem>
                        </SelectContent>
                    </Select>
                </div>
                <div
                    class="flex items-center justify-center text-sm font-medium"
                >
                    Page {{ table.getState().pagination.pageIndex + 1 }} of
                    {{ table.getPageCount() }}
                </div>
                <div class="space-x-2">
                    <Button
                        variant="outline"
                        class="h-8 w-8 p-0"
                        :disabled="!table.getCanPreviousPage()"
                        @click="table.setPageIndex(0)"
                    >
                        <span class="sr-only">Go to first page</span>
                        <ChevronsLeftIcon class="h-4 w-4" />
                    </Button>
                    <Button
                        variant="outline"
                        class="h-8 w-8 p-0"
                        :disabled="!table.getCanPreviousPage()"
                        @click="table.previousPage()"
                    >
                        <span class="sr-only">Go to previous page</span>
                        <ChevronLeftIcon class="h-4 w-4" />
                    </Button>
                    <Button
                        variant="outline"
                        class="h-8 w-8 p-0"
                        :disabled="!table.getCanNextPage()"
                        @click="table.nextPage()"
                    >
                        <span class="sr-only">Go to next page</span>
                        <ChevronRightIcon class="h-4 w-4" />
                    </Button>
                    <Button
                        variant="outline"
                        class="h-8 w-8 p-0"
                        :disabled="!table.getCanNextPage()"
                        @click="table.setPageIndex(table.getPageCount() - 1)"
                    >
                        <span class="sr-only">Go to last page</span>
                        <ChevronsRightIcon class="h-4 w-4" />
                    </Button>
                </div>
            </div>
        </div>

        <CollectionFormSheet v-model:open="isNewRecordFormSheetOpen" />
        <CollectionPreviewSheet v-model:open="isAPIPreviewSheetOpen" />
    </div>
</template>
