<script setup>
import AppLayout from '@/layouts/AppLayout.vue';
import { Clock, MapPin, Tag } from 'lucide-vue-next';
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
                <div class="flex flex-wrap gap-3 mt-4 pt-4 border-t border-gray-100">
                    <span class="text-sm text-gray-500 font-medium">Filter:</span>
                    <button class="px-3 py-1 text-sm border border-gray-300 rounded-full hover:border-orange-500 hover:text-orange-500 transition">Eksekutif</button>
                    <button class="px-3 py-1 text-sm border border-gray-300 rounded-full hover:border-orange-500 hover:text-orange-500 transition">Patas</button>
                    <button class="px-3 py-1 text-sm border border-gray-300 rounded-full hover:border-orange-500 hover:text-orange-500 transition">Ekonomi</button>
                    <button class="px-3 py-1 text-sm border border-gray-300 rounded-full hover:border-orange-500 hover:text-orange-500 transition">AC</button>
                    <button class="px-3 py-1 text-sm border border-gray-300 rounded-full hover:border-orange-500 hover:text-orange-500 transition">Toilet</button>
                    <button class="px-3 py-1 text-sm border border-gray-300 rounded-full hover:border-orange-500 hover:text-orange-500 transition">Wifi</button>
                </div>
            </div>
        </section>

        <section class="container mx-auto px-4 sm:px-6 lg:px-8 py-12">
            <h2 class="text-2xl font-bold text-gray-800 mb-6">Bus Tersedia <span class="text-orange-500 text-lg font-normal ml-12">(12 ditemukan)</span></h2>
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                <div class="bg-white rounded-2xl shadow-md overflow-hidden hover:shadow-xl transition hover:-translate-y-1">
                    <img src="" alt="image" class="w-full h-48 object-cover">
                    <div class="p-5">
                        <div class="flex justify-between items-start">
                            <div>
                                <h3 class="font-bold text-lg">Sanjaya Putra</h3>
                                <p class="text-sm text-gray-500">Eksekutif 48 Kursi</p>
                            </div>
                            <span class="bg-orange-100 text-orange-600 text-xs font-semibold px-2 py-1 rounded-full">AC, Toilet</span>
                        </div>
                        <div class="mt-3 space-y-1">
                            <p class="text-sm flex text-gray-600"><MapPin class="text-orange-500 w-4 mr-1" /> Jakarta -> Bandung</p>
                            <p class="text-sm flex text-gray-600"><Clock class="text-orange-500 w-4 mr-1" /> 08:00 - 11:30 (3j 30m)</p>
                            <p class="text-sm flex text-gray-600"><Tag class="text-orange-500 w-4 mr-1" /> Mulai <span class="font-bold text-orange-600">Rp 100.000</span>/org</p>
                        </div>
                        <div class="mt-4 flex items-center justify-between">
                            <div class="text-yellow-400 text-sm">★★★★★ (124)</div>
                            <a href="#" class="bg-orange-500 hover:bg-orange-600 text-white px-4 py-2 rounded-lg text-sm font-semibold transition">Pilih Kursi</a>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    </AppLayout>
</template>