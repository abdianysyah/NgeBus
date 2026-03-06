<script setup>
import { BusFrontIcon, ArrowLeft } from 'lucide-vue-next';
import { RouterLink } from 'vue-router';
</script>

<template>
    <div class="bg-gray-50 antialiased flex flex-col min-h-screen">
        <main class="flex-1 flex items-center justify-center p-6">
            <div class="max-w-2xl w-full text-center">
                <div class="mb-6 text-blue-800 flex justify-center">
                    <BusFrontIcon class="w-24 h-24 md:w-32 md:h-32 opacity-20" />
                </div>
                <h1 class="text-8xl md:text-9xl font-extrabold text-transparent bg-clip-text bg-linear-to-r from-blue-800 to-orange-500 leading-none">
                    404
                </h1>
                <h2 class="text-2xl md:text-3xl font-bold text-gray-800 mt-4 mb-2">
                    Halaman Tidak Ditemukan
                </h2>
                <p class="text-gray-600 mb-8 max-w-md mx-auto">
                    Sepertinya bus yang Anda cari sudah berangkat atau rutenya tidak tersedia. 
                    Yuk, kembali ke halaman utama untuk melanjutkan perjalanan.
                </p>
    
                <div class="mt-4">
                    <RouterLink to="/" class="inline-flex items-center justify-center px-4 py-2 bg-orange-500 text-white rounded-md shadow hover:bg-orange-600 transition">
                        <ArrowLeft class="mr-2" /> Kembali ke Beranda
                    </RouterLink>
                </div>
    
                <p class="text-sm text-gray-400 mt-8">
                    Butuh bantuan? <RouterLink to="/contact" class="text-orange-500 hover:underline">Hubungi Kami!</RouterLink>
                </p>
            </div>
        </main>
    </div>
</template>