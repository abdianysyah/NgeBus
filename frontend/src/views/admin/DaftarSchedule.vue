<script setup>
import AdminLayout from '@/layouts/AdminLayout.vue';
import { PlusCircle, Search, Edit2Icon, Trash, Eye } from 'lucide-vue-next';

</script>

<template>
    <AdminLayout>
        <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-6">
            <div>
                <h2 class="text-2xl lg:text-3xl font-bold text-gray-800">Kelola Jadwal Keberangkatan</h2>
                <p class="text-gray-600">Daftar semua jadwal keberangkatan yang tersedia</p>
            </div>
            <div class="mt-4 sm:mt-0">
                <button class="bg-orange-500 hover:bg-orange-600 hover:shadow-none text-white font-semibold px-5 py-3 rounded-xl shadow-md transition inline-flex items-center">
                    <PlusCircle class="mr-2" /> Tambah Jadwal
                </button>
            </div>
        </div>
    </AdminLayout>
</template>