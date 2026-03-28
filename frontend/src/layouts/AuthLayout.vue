<script setup>
import { RouterLink, useRoute } from 'vue-router';
import { 
    Bus,
    Ticket,
    MapPin 
} from 'lucide-vue-next';

const route = useRoute ()

const titles = {
  Login: {
    title: "Selamat datang kembali",
    subtitle: "Silakan masuk ke akun Anda"
  },
  Register: {
    title: "Buat akun baru",
    subtitle: "Daftar untuk mulai menggunakan sistem"
  },
  forgot: {
    title: "Lupa Password",
    subtitle: "Masukkan email untuk reset password"
  }
}
</script>

<template>
    <div class="min-h-screen grid grid-cols-1 md:grid-cols-2">
        <div class="flex items-center justify-center p-6 bg-white">
            <div class="w-full max-w-md">
                <!-- Logo Kecil -->
                <RouterLink to="/" class="inline-block mb-8">
                    <span class="text-3xl font-bold text-blue-800 tracking-tight">Nge<span class="text-orange-500">Bus</span></span>
                </RouterLink>

                <h2 class="text-3xl font-bold text-gray-800 mb-2">{{ titles[route.name]?.title }}</h2>
                <p class="text-gray-500 mb-8">{{ titles[route.name]?.subtitle }}</p>
                <!-- Form -->
                <slot/>
            </div>
        </div>

        <div class="hidden md:flex flex-col items-center justify-center p-12 bg-linear-to-br from-blue-900 to-blue-800 text-white relative overflow-hidden">
            <!-- Background pattern halus (opsional) -->
            <div class="absolute inset-0 opacity-10">
                <Bus class="text-9xl absolute bottom-0 left-0 -ml-10 -mb-10 rotate-12"/>
                <Ticket class="text-8xl absolute top-10 right-10 opacity-20" />
                <MapPin class="text-7xl absolute bottom-20 right-20 opacity-20" />
            </div>
            
            <div class="relative z-10 max-w-md text-center">
                <!-- Logo besar versi putih -->
                <span class="text-5xl font-bold tracking-tight block mb-6">Nge<span class="text-orange-300">Bus</span></span>
                <h2 class="text-3xl font-bold mb-4">Perjalanan Anda, prioritas kami</h2>
                <p class="text-blue-100 leading-relaxed mb-8">
                    Nikmati kemudahan pemesanan tiket bus dengan ribuan rute di seluruh Indonesia. 
                    Dapatkan pengalaman perjalanan yang nyaman dan aman.
                </p>
                
                <!-- Testimonial singkat atau fitur unggulan -->
                <div class="grid grid-cols-2 gap-4 text-left bg-white/10 backdrop-blur-sm p-5 rounded-2xl">
                    <div>
                        <div class="text-orange-300 text-2xl font-bold">100+</div>
                        <div class="text-xs uppercase tracking-wider">Armada Bus</div>
                    </div>
                    <div>
                        <div class="text-orange-300 text-2xl font-bold">50+</div>
                        <div class="text-xs uppercase tracking-wider">Kota Tujuan</div>
                    </div>
                    <div>
                        <div class="text-orange-300 text-2xl font-bold">10k+</div>
                        <div class="text-xs uppercase tracking-wider">Pelanggan Puas</div>
                    </div>
                    <div>
                        <div class="text-orange-300 text-2xl font-bold">24/7</div>
                        <div class="text-xs uppercase tracking-wider">Layanan</div>
                    </div>
                </div>
                
                <!-- Gambar ilustrasi (opsional, bisa tambah) -->
                <div class="mt-8 flex justify-center">
                    <img src="https://images.unsplash.com/placeholder-illustration" alt="" class="w-48 opacity-70">
                </div>
            </div>
        </div>
    </div>
</template>