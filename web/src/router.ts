import { createRouter, createWebHistory } from "vue-router"

import DashboardLayout from "@/layout/dashboard.vue"

import HomePage from "@/pages/home.vue"
import CollectionPage from "@/pages/dashboard/collection.vue"

const routes = [
    {
        path: "/",
        name: "home",
        component: () => HomePage,
    },
    {
        path: "/dashboard",
        name: "dashboard",
        component: () => DashboardLayout,
        children: [
            {
                path: "",
                name: "collection",
                component: () => CollectionPage,
            },
        ],
    },
]

const router = createRouter({
    history: createWebHistory(), // Recommended for single-page applications
    routes,
})

export default router
