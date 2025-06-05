import { createRouter, createWebHistory } from "vue-router"

import DashboardLayout from "@/layout/dashboard.vue"

import LoginPage from "@/modules/auth/login.vue"
import CollectionPage from "@/modules/collection/index.vue"

const routes = [
    {
        path: "/",
        name: "Login",
        component: () => LoginPage,
    },
    {
        path: "/dashboard",
        component: () => DashboardLayout,
        children: [
            {
                path: "",
                name: "Collections",
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
