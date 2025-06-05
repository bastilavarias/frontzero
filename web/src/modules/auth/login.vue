<script lang="ts">
export const description =
    "A simple login form with email and password. The submit button says 'Sign in'."
export const iframeHeight = "600px"
export const containerClass =
    "w-full h-screen flex items-center justify-center px-4"
</script>

<script setup lang="ts">
import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
</script>

<template>
    <div class="h-screen w-screen flex justify-center items-center">
        <Card class="w-full max-w-lg border-none shadow-none">
            <CardHeader class="space-y-5">
                <CardTitle class="text-4xl text-center font-bold">
                    frontzero
                </CardTitle>
                <CardDescription class="text-md text-center text-black">
                    Superuser login
                </CardDescription>
            </CardHeader>
            <CardContent class="grid gap-4">
                <div class="grid gap-2">
                    <Label for="email">Email</Label>
                    <Input
                        id="email"
                        type="email"
                        placeholder="m@example.com"
                        required
                    />
                </div>
                <div class="grid gap-2">
                    <Label for="password">Password</Label>
                    <Input id="password" type="password" required />
                </div>
            </CardContent>
            <CardFooter>
                <router-link to="/dashboard" custom v-slot="{ href, navigate }">
                    <Button class="w-full" :href="href" @click="navigate">
                        Sign in
                    </Button>
                </router-link>
            </CardFooter>
        </Card>
    </div>
</template>
