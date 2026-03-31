<script setup>
import { onMounted, ref } from 'vue';
import { X, LayoutDashboard, BusFrontIcon, RouteIcon, UserCircle, LogOut, MenuIcon, Building, Clock, TicketCheck, UserCircleIcon  } from 'lucide-vue-next';
import { RouterLink, useRouter, useRoute } from 'vue-router';
import { Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/vue';

const route = useRoute()
const router = useRouter()
const sidebarOpen = ref(false)
const user = ref(null)

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
    localStorage.clear()
    router.push("/login")
}

onMounted(() => {
    const data = localStorage.getItem('user')
    if (data) {
        user.value = JSON.parse(data)
    }
})

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
        </div>
    </aside>

    <div class="fixed inset-0 bg-black bg-opacity-50 z-20 hidden transition-opacity"></div>
    <div class="min-h-screen flex flex-col lg:ml-72">
        <header class="bg-white shadow-sm sticky top-0 z-10">
            <div class="px-4 sm:px-6 lg:px-8 py-3 flex items-center justify-between">
                <button @click="toggleSidebar" class="lg:hidden text-gray-600 hover:text-gray-900 focus:outline-none mr-2">
                    <MenuIcon class="text-2xl" />
                </button>
                <Menu as="div" class="relative flex items-center ml-auto">
                  <MenuButton class="flex items-center focus:outline-none hover:bg-gray-100 rounded-full px-2 py-1 transition">
                    <span class="hidden sm:inline text-sm text-gray-700 mr-2">
                      {{ user ? `${user.name}` : 'Maintenance Mode' }}
                    </span>
                    <img
                      src="https://randomuser.me/api/portraits/men/2.jpg"
                      class="w-9 h-9 rounded-full border-2 border-orange-200"
                    />
                  </MenuButton>
              
                  <transition
                    enter-active-class="transition duration-100 ease-out"
                    enter-from-class="transform scale-95 opacity-0"
                    enter-to-class="transform scale-100 opacity-100"
                    leave-active-class="transition duration-75 ease-in"
                    leave-from-class="transform scale-100 opacity-100"
                    leave-to-class="transform scale-95 opacity-0"
                  >
                    <MenuItems class="absolute right-0 top-full mt-2 w-56 origin-top-right rounded-md bg-white shadow-lg ring-1 ring-black/5 z-50">
                      <div class="px-1 py-1">
                        <MenuItem v-slot="{ active }">
                          <button
                            :class="[
                              active ? 'bg-orange-500 text-white' : 'text-gray-500',
                              'group flex w-full items-center rounded-md px-2 py-2 text-sm',
                            ]"
                          >
                            <UserCircleIcon class="mr-2 h-5 w-5" />
                            Profil Saya
                          </button>
                        </MenuItem>
                        <MenuItem v-slot="{ active }">
                          <button @click="logout"
                            :class="[
                              active ? 'bg-orange-500 text-white' : 'text-gray-500',
                              'group flex w-full items-center rounded-md px-2 py-2 text-sm',
                            ]"
                          >
                            <LogOut class="mr-2 h-5 w-5" />
                            Keluar
                          </button>
                        </MenuItem>
                      </div>
                    </MenuItems>
                  </transition>
                </Menu>
            </div>
        </header>

        <main class="flex-1 p-4 sm:p-6 lg:p-8 bg-gray-50">
            <slot />
        </main>
    </div>
</template>