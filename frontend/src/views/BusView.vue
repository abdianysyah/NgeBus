<script setup>
import AppLayout from '@/layouts/AppLayout.vue';
</script>

<template>
    <AppLayout>
        <section class="bg-linear-to-r from-blue-900 to-blue-800 text-white py-12">
            <div class="container mx-auto px-4 sm:px-6 lg:px-8">
                <h1 class="text-3xl md:text-4xl font-bold mb-2">Temukan Bus Favorit Anda</h1>
                <p class="text-blue-100 max-w-2xl">PIlih dari berbagai armada bus terbaik dengan fasilitas lengkap dan harga terjangkau.</p>
            </div>
        </section>

        <section class="container mx-auto px-4 sm:px-6 lg:px-8 -mt-6 relative z-10">
            <div class="bg-white rounded-xl shadow-lg p-5">
                <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
                    <div>
                        <label for="dari" class="block text-gray-700 text-sm font-semibold mb-1">Dari</label>
                        <input type="text" name="" placeholder="Kota Asal" id="" class="w-full border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-orange-500">
                    </div>
                    <div>
                        <label for="dari" class="block text-gray-700 text-sm font-semibold mb-1">Ke</label>
                        <input type="text" name="" placeholder="Kota Tujuan" id="" class="w-full border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-orange-500">
                    </div>
                    <div>
                        <label class="block text-gray-700 text-sm font-semibold mb-1">Tanggal</label>
                        <input type="date" class="w-full border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-orange-500">
                    </div>
                    <div class="flex items-end">
                        <button class="w-full bg-orange-500 hover:bg-orange-600 text-white font-semibold py-2 px-4 rounded-lg transition shadow-md">Cari Bus</button>
                    </div>
                </div>
            </div>
        </section>
    </AppLayout>
</template>