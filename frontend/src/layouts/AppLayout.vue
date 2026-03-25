<script setup>
import { onMounted, ref } from "vue";
import { jwtDecode } from "jwt-decode";
import { 
    Menu,
    FacebookIcon,
    TwitterIcon,
    InstagramIcon,
    Phone,
    Send,
    MapIcon
    
} from 'lucide-vue-next';
const isLogin = ref(false)
const role = ref(null)
const mobileMenu = ref(false)
const toggleMenu = () => {
    mobileMenu.value = !mobileMenu.value
}
const closeMenu = () => {
    mobileMenu.value = false
}
onMounted(() => {
    const token = localStorage.getItem('token')
    if (token) {
        const decoded = jwtDecode(token)
        role.value = decoded.role
        isLogin.value = true
    }
})
</script>

<template>
    <nav class="bg-white shadow-sm sticky top-0 z-50">
        <div class="container mx-auto px-4 sm:px-6 lg:px-8">
            <div class="flex justify-between items-center py-4">
                <a href="#" class="text-2xl font-bold text-blue-800 tracking-tight">
                    <img src="@/assets/img/logo_ngebus.webp" alt="logo" class="h-15">
                </a>

                <div class="hidden md:flex space-x-8 text-gray-700 font-medium">
                    <RouterLink to="/" class="hover:text-orange-500 transition">Beranda</RouterLink>
                    <RouterLink to="/bus" class="hover:text-orange-500 transition">Bus</RouterLink>
                    <RouterLink to="/rute" class="hover:text-orange-500 transition">Rute</RouterLink>
                    <RouterLink to="/contact" class="hover:text-orange-500 transition">Kontak</RouterLink>
                </div>

                <div class="hidden md:flex items-center space-x-3">
                    <template v-if="!isLogin">
                        <RouterLink to="/login" class="text-gray-700 hover:text-orange-500 py-1">Login</RouterLink>
                        <RouterLink to="/register" class="bg-orange-500 text-white px-4 py-2 rounded-full hover:bg-orange-600 text-center">Register</RouterLink>
                    </template>
                    <template v-else-if="role === 'admin'">
                        <RouterLink to="/admin/dashboard" class="bg-orange-500 text-white px-4 py-2 rounded-full hover:bg-orange-600 text-center">Dashboard</RouterLink>
                    </template>
                    <template v-else>
                        <RouterLink to="/dashboard" class="bg-orange-500 text-white px-4 py-2 rounded-full hover:bg-orange-600 transition font-medium shadow-sm ease-in-out duration-500">Dashboard</RouterLink>
                    </template>
                </div>

                <div class="md:hidden flex items-center">
                    <button @click="toggleMenu" class="text-gray-600 hover:text-gray-900 focus:outline-none">
                        <Menu class="text-2xl" />
                    </button>
                </div>
            </div>
        </div>

        <div v-show="mobileMenu" class="md:hidden bg-white border-t border-gray-100 px-4 py-3 space-y-2">
            <RouterLink to="/" class="block py-2 text-gray-700 hover:text-orange-500 font-medium">Beranda</RouterLink>
            <RouterLink to="/bus" class="block py-2 text-gray-700 hover:text-orange-500 font-medium">Bus</RouterLink>
            <RouterLink to="/rute" class="block py-2 text-gray-700 hover:text-orange-500 font-medium">Rute</RouterLink>
            <RouterLink to="/contact" class="block py-2 text-gray-700 hover:text-orange-500 font-medium">Contact</RouterLink>
            <div class="flex flex-col space-y-2 pt-3 border-t border-gray-200 mt-2">
                <template v-if="!isLogin">
                    <RouterLink to="/login"></RouterLink>
                    <a href="" class="px-4 py-2 hover:bg-orange-100 rounded-full text-center text-gray-700 hover:text-orange-500 transition">Masuk</a>
                    <a href="" class="bg-orange-500 text-white px-4 py-2 rounded-full hover:bg-orange-600 text-center transition">Daftar</a>
                </template>
                <template v-else-if="role === 'admin'">
                    <RouterLink to="/admin/dashboard" class="bg-orange-500 text-white px-4 py-2 rounded-full hover:bg-orange-600 text-center">Dashboard</RouterLink>
                </template>
                <template v-else>
                    <RouterLink to="/dashboard" class="bg-orange-500 text-white px-4 py-2 rounded-full hover:bg-orange-600 transition font-medium shadow-sm ease-in-out">Dashboard</RouterLink>
                </template>
            </div>
        </div>
    </nav>
    <slot />
    <footer class="bg-gray-900 text-white pt-12 pb-6">
        <div class="container mx-auto px-4 sm:px-6 lg:px-8">
            <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
                <div>
                    <h3 class="text-2xl font-bold text-white mb-4">Nge<span class="text-orange-500">Bus</span></h3>
                    <p class="text-gray-400 text-sm leading-relaxed">Solusi pemesanan tiket bus online terpercaya di Indonesia. Perjalanan Anda, prioritas kami</p>
                    <div class="flex space-x-4 mt-4">
                        <a href="" class="text-gray-400 hover:text-white transition">
                            <FacebookIcon />
                        </a>
                        <a href="" class="text-gray-400 hover:text-white transition">
                            <TwitterIcon />
                        </a>
                        <a href="" class="text-gray-400 hover:text-white transition">
                            <InstagramIcon />
                        </a>
                    </div>
                </div>
                <div>
                    <h4 class="font-semibold text-lg mb-4">Perusahaan</h4>
                    <ul class="space-y-2 text-gray-400">
                        <li><a href="" class="hover:text-white transition">Tentang Kami</a></li>
                        <li><a href="" class="hover:text-white transition">Karir</a></li>
                        <li><a href="" class="hover:text-white transition">Blog</a></li>
                        <li><a href="" class="hover:text-white transition">Mitra</a></li>
                    </ul>
                </div>
                <div>
                    <h4 class="font-semibold text-lg mb-4">Bantuan</h4>
                    <ul class="space-y-2 text-gray-400">
                        <li><a href="" class="hover:text-white transition">Pusat Bantuan</a></li>
                        <li><a href="" class="hover:text-white transition">Syarat & Ketentuan</a></li>
                        <li><a href="" class="hover:text-white transition">Kebijakan Privasi</a></li>
                        <li><a href="" class="hover:text-white transition">FAQ</a></li>
                    </ul>
                </div>
                <div>
                    <h4 class="font-semibold text-lg mb-4">Kontak</h4>
                    <ul class="space-y-2 text-gray-400">
                        <li class="flex items-center"><Phone class="w-5 mr-2"/> 021-1234-5678</li>
                        <li class="flex items-center"><Send class="w-5 mr-2"/> halo@ngebus.id</li>
                        <li class="flex items-center"><MapIcon class="w-5 mr-2" /> Jakarta, Indonesia</li>
                    </ul>
                </div>
            </div>
            <div class="border-t border-gray-800 mt-8 pt-6 text-center text-gray-500 text-sm">
                &copy; 2026 NgeBus. All rights reserved.
            </div>
        </div>
    </footer>
</template>