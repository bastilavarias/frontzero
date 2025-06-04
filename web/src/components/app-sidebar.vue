<script setup lang="ts">
import {
    AudioWaveform,
    BarChart2,
    BookOpen,
    Bot,
    Command,
    Database,
    GalleryVerticalEnd,
    Settings,
    Settings2,
    SquareTerminal,
    Webhook,
} from "lucide-vue-next"
import NavMain from "@/components/nav-main.vue"
import NavProject from "@/components/nav-project.vue"
import NavUser from "@/components/nav-user.vue"
import {
    Sidebar,
    SidebarContent,
    SidebarFooter,
    SidebarHeader,
    SidebarRail,
} from "@/components/ui/sidebar"
import type { HTMLAttributes } from "vue"
import ProjectSwitcher from "@/components/project-switcher.vue"

export interface SidebarProps {
    side?: "left" | "right"
    variant?: "sidebar" | "floating" | "inset"
    collapsible?: "offcanvas" | "icon" | "none"
    class?: HTMLAttributes["class"]
}

const props = withDefaults(defineProps<SidebarProps>(), {
    collapsible: "icon",
})

const data = {
    user: {
        name: "shadcn",
        email: "m@example.com",
        avatar: "/avatars/shadcn.jpg",
    },
    teams: [
        {
            name: "San Miguel Corporation",
            logo: GalleryVerticalEnd,
            plan: "Enterprise",
        },
        {
            name: "Procter & Gamble Philippines",
            logo: AudioWaveform,
            plan: "Startup",
        },
        {
            name: "Ayala Corporation",
            logo: Command,
            plan: "Free",
        },
    ],
    navMain: [
        {
            title: "Playground",
            url: "#",
            icon: SquareTerminal,
            isActive: true,
            items: [
                {
                    title: "History",
                    url: "#",
                },
                {
                    title: "Starred",
                    url: "#",
                },
                {
                    title: "Settings",
                    url: "#",
                },
            ],
        },
        {
            title: "Models",
            url: "#",
            icon: Bot,
            items: [
                {
                    title: "Genesis",
                    url: "#",
                },
                {
                    title: "Explorer",
                    url: "#",
                },
                {
                    title: "Quantum",
                    url: "#",
                },
            ],
        },
        {
            title: "Documentation",
            url: "#",
            icon: BookOpen,
            items: [
                {
                    title: "Introduction",
                    url: "#",
                },
                {
                    title: "Get Started",
                    url: "#",
                },
                {
                    title: "Tutorials",
                    url: "#",
                },
                {
                    title: "Changelog",
                    url: "#",
                },
            ],
        },
        {
            title: "Settings",
            url: "#",
            icon: Settings2,
            items: [
                {
                    title: "General",
                    url: "#",
                },
                {
                    title: "Team",
                    url: "#",
                },
                {
                    title: "Billing",
                    url: "#",
                },
                {
                    title: "Limits",
                    url: "#",
                },
            ],
        },
    ],
    projectActions: [
        {
            name: "Collections",
            url: "#",
            icon: Database,
        },
        {
            name: "Web Hooks",
            url: "#",
            icon: Webhook,
        },
        {
            name: "Analytics",
            url: "#",
            icon: BarChart2,
        },
        {
            name: "Configuration",
            url: "#",
            icon: Settings,
        },
    ],
}
</script>

<template>
    <Sidebar v-bind="props">
        <SidebarHeader>
            <ProjectSwitcher :projects="data.teams" />
        </SidebarHeader>
        <SidebarContent>
            <NavProject :project-actions="data.projectActions" />
            <NavMain :items="data.navMain" />
        </SidebarContent>
        <SidebarFooter>
            <NavUser :user="data.user" />
        </SidebarFooter>
        <SidebarRail />
    </Sidebar>
</template>
