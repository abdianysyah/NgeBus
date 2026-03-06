<script setup>
import { ref } from 'vue';
import { XSquare, LayoutDashboard, BusFrontIcon, RouteIcon, UserCircle, LogOut  } from 'lucide-vue-next';
import { RouterLink } from 'vue-router';

const sidebarOpen = ref(false)

const toggleSidebar = () => {
    sidebarOpen.value = !sidebarOpen.value
}

const closeSidebar = () => {
    sidebarOpen.value = false
}
</script>

<template>
    <aside class="fixed inset-y-0 left-0 w-72 bg-white shadow-lg z-30 transform transition-transform duration-200 ease-in-out overflow-y-auto no-scrollbar"
        :class="sidebarOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0'">
        <div class="p-6 relative">
            <button @click="closeSidebar" class="absolute top-4 right-4 text-gray-500 hover:text-gray-800 lg:hidden">
                <XSquare class="text-2xl" />
            </button>
            <RouterLink to="/admin/dashboard" class="text-3xl font-bold text-blue-800 tracking-tight block mb-8">
                Nge<span class="text-orange-500">Bus</span>
            </RouterLink>
            <nav class="space-y-1">
                <RouterLink to="" class="flex items-center space-x-3 px-4 py-3 bg-orange-50 text-orange-600 rounded-xl font-medium">
                    <LayoutDashboard class="w-5" />
                    <span>Dashboard</span>
                </RouterLink>
                <RouterLink to="" class="flex items-center space-x-3 px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-xl transition">
                    <BusFrontIcon class="w-5" />
                    <span>Kelola Bus</span>
                </RouterLink>
                <RouterLink to="" class="flex items-center space-x-3 px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-xl transition">
                    <RouteIcon class="w-5" />
                    <span>Kelola Rute</span>
                </RouterLink>
                <RouterLink to="" class="flex items-center space-x-3 px-4 py-3 text-gray-700 hover:bg-gray-100 rounded-xl transition">
                    <UserCircle class="w-5" />
                    <span>Kelola Pengguna</span>
                </RouterLink>
            </nav>

            <div class="pt-8 mt-8 border-t border-gray-200">
                <button class="flex items-center space-x-3 px-4 py-3 text-red-600 hover:bg-red-50 rounded-xl transition">
                    <LogOut class="w-5" />
                    <span>Log Out</span>
                </button>
            </div>
            <p class="text-xs text-gray-400 mt-8">
                &copy; 2026 NgeBus
            </p>
        </div>
    </aside>
    <slot />
</template>

