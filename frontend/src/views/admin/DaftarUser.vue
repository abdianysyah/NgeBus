<script setup>
import AdminLayout from '@/layouts/AdminLayout.vue';
import { 
    PlusCircle,
    Search,
 } from 'lucide-vue-next';
</script>

<template>
    <AdminLayout>
        <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-6">
            <div>
                <h2 class="text-2xl lg:text-3xl font-bold text-gray-800">Kelola Pengguna</h2>
                <p class="text-gray-600">Daftar semua pengguna</p>
            </div>
            <div class="mt-4 sm:mt-0">
                <button class="bg-orange-500 hover:bg-orange-600 hover:shadow-none text-white font-semibold px-5 py-3 rounded-xl shadow-md transition inline-flex items-center">
                    <PlusCircle class="mr-2" /> Tambah Bus
                </button>
            </div>
        </div>
        <div class="bg-white p-4 rounded-xl shadow-sm border-gray-100 mb-6 flex flex-col sm:flex-row gap-3 items-center justify-between">
            <div class="relative w-full sm:w-64">
                <Search class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" />
                <input type="text" name="" id="" class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-orange-500" placeholder="Username / Email">
            </div>
        </div>
    </AdminLayout>
</template>