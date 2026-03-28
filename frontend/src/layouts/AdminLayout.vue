<script setup>
import { ref } from 'vue';
import { X, LayoutDashboard, BusFrontIcon, RouteIcon, UserCircle, LogOut, MenuIcon, Building, Clock, TicketCheck  } from 'lucide-vue-next';
import { RouterLink, useRouter, useRoute } from 'vue-router';
import Swal from 'sweetalert2';

const route = useRoute()
const router = useRouter()
const sidebarOpen = ref(false)

const isActive = (path) => {
    return route.path === path
}

const toggleSidebar = () => {
    sidebarOpen.value = !sidebarOpen.value
}

const closeSidebar = () => {
    sidebarOpen.value = false
}

const logout = () => {
    Swal.fire({
        icon: "warning",
        title: "logout",
        text: "Anda yakin ingin keluar?",
        showCancelButton: true,
        confirmButtonText: "Logout"
    }).then((result) => {
        if (result.isConfirmed) {
            localStorage.clear()
            router.push("/login")
        }
    })
}

</script>

<template>
    <aside class="fixed inset-y-0 left-0 w-72 bg-white shadow-lg z-30 transform transition-transform duration-600 ease-in-out overflow-y-auto no-scrollbar"
        :class="sidebarOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0'">
        <div class="p-6 relative">
            <button @click="closeSidebar" class="absolute top-4 right-4 text-gray-500 hover:text-gray-800 lg:hidden">
                <X class="text-2xl" />
            </button>
            <RouterLink to="/admin/dashboard" class="text-3xl font-bold text-blue-800 tracking-tight block mb-8">
                Nge<span class="text-orange-500">Bus</span>
                <span class="text-sm bg-orange-100 text-orange-600 px-2 py-1 rounded-full ml-2">Admin</span>
            </RouterLink>
            <nav class="space-y-1">
                <RouterLink to="/admin/dashboard" :class="['flex items-center space-x-3 px-4 py-3 rounded-xl font-medium', isActive('/admin/dashboard') ? 'bg-orange-50 text-orange-600' : 'text-gray-600 hover:bg-gray-100']" >
                    <LayoutDashboard class="w-5" />
                    <span>Dashboard</span>
                </RouterLink>
                <RouterLink to="/admin/data-bus" :class="['flex items-center space-x-3 px-4 py-3 rounded-xl font-medium', isActive('/admin/data-bus') ? 'bg-orange-50 text-orange-600' : 'text-gray-600 hover:bg-gray-100']">
                    <BusFrontIcon class="w-5" />
                    <span>Kelola Bus</span>
                </RouterLink>
                <RouterLink to="/admin/data-route" :class="['flex items-center space-x-3 px-4 py-3 rounded-xl font-medium', isActive('/admin/data-route') ? 'bg-orange-50 text-orange-600' : 'text-gray-600 hover:bg-gray-100']">
                    <RouteIcon class="w-5" />
                    <span>Kelola Rute</span>
                </RouterLink>
                <RouterLink to="/admin/data-company" :class="['flex items-center space-x-3 px-4 py-3 rounded-xl font-medium', isActive('/admin/data-company') ? 'bg-orange-50 text-orange-600' : 'text-gray-600 hover:bg-gray-100']">
                    <Building class="w-5" />
                    <span>Kelola Po Bus</span>
                </RouterLink>
                <RouterLink to="/admin/data-schedule" :class="['flex items-center space-x-3 px-4 py-3 rounded-xl font-medium', isActive('/admin/data-schedule') ? 'bg-orange-50 text-orange-600' : 'text-gray-600 hover:bg-gray-100']">
                    <Clock class="w-5" />
                    <span>Kelola Jadwal</span>
                </RouterLink>
                <RouterLink to="/admin/data-order" :class="['flex items-center space-x-3 px-4 py-3 rounded-xl font-medium', isActive('/admin/data-order') ? 'bg-orange-50 text-orange-600' : 'text-gray-600 hover:bg-gray-100']">
                    <TicketCheck class="w-5" />
                    <span>Kelola Order Tiket</span>
                </RouterLink>
                <RouterLink to="/admin/data-users" :class="['flex items-center space-x-3 px-4 py-3 rounded-xl font-medium', isActive('/admin/data-users') ? 'bg-orange-50 text-orange-600' : 'text-gray-600 hover:bg-gray-100']">
                    <UserCircle class="w-5" />
                    <span>Kelola Pengguna</span>
                </RouterLink>

            </nav>

            <div class="pt-8 mt-8 border-t border-gray-200">
                <button @click="logout" class="flex items-center space-x-3 px-4 py-3 text-red-600 hover:bg-red-50 rounded-xl transition">
                    <LogOut class="w-5" />
                    <span>Log Out</span>
                </button>
            </div>
            <p class="text-xs text-gray-400 mt-8">
                &copy; 2026 NgeBus
            </p>
        </div>
    </aside>

    <div class="fixed inset-0 bg-black bg-opacity-50 z-20 hidden transition-opacity"></div>
    <div class="min-h-screen flex flex-col lg:ml-72">
        <header class="bg-white shadow-sm sticky top-0 z-10">
            <div class="px-4 sm:px-6 lg:px-8 py-3 flex items-center justify-between">
                <button @click="toggleSidebar" class="lg:hidden text-gray-600 hover:text-gray-900 focus:outline-none mr-2">
                    <MenuIcon class="text-2xl" />
                </button>
                <div class="flex items-center space-x-3 ml-auto">
                    <span class="hidden sm:inline text-sm text-gray-700">Admin, Budi</span>
                    <div class="relative">
                        <button class="flex items-center focus:outline-none">
                            <img src="https://randomuser.me/api/portraits/men/2.jpg" alt="Profil" class="w-9 h-9 rounded-full object-cover border-2 border-orange-200">
                            <i class="fas fa-chevron-down ml-1 text-xs text-gray-500"></i>
                        </button>
                    </div>
                </div>
            </div>
        </header>

        <main class="flex-1 p-4 sm:p-6 lg:p-8 bg-gray-50">
            <slot />
        </main>
    </div>
</template>