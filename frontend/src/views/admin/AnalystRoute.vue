<script setup>
import AdminLayout from '@/layouts/AdminLayout.vue';
import { 
    Database,
    ArrowUp,
    User,
    ChartLine,
    DollarSignIcon,
    RouteIcon
} from 'lucide-vue-next';
</script>


<template>
    <AdminLayout>
        <div class="mb-8 p-4 bg-red-50 border border-red-200 rounded-lg flex items-start gap-3">
            <svg class="w-5 h-5 text-red-600 mt-0.5 shrink-0" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
            </svg>
            <div>
                <h3 class="font-semibold text-red-800">Masih Maintenance</h3>
                <p class="text-sm text-red-700 mt-1">Fitur ini sedang dalam tahap pemeliharaan dan belum tersedia saat ini.</p>
            </div>
        </div>
        <div class="mb-8">
            <h2 class="text-2xl lg:text-3xl font-bold text-gray-800">Analisis Rute</h2>
            <p class="text-gray-600">Analisis performa rute, okupansi, dan pendapatan berbasis data warehouse.</p>
        </div>

        <div class="bg-white p-4 rounded-xl shadow-sm border border-gray-100 mb-8 flex flex-wrap gap-3 items-center justify-between">
            <div class="flex gap-2">
                <button class="px-4 py-2 text-sm rounded-lg bg-orange-500 text-white">Hari Ini</button>
                <button class="px-4 py-2 text-sm rounded-lg bg-gray-100 text-gray-700 hover:bg-gray-200">Minggu Ini</button>
                <button class="px-4 py-2 text-sm rounded-lg bg-gray-100 text-gray-700 hover:bg-gray-200">Bulan Ini</button>
                <button class="px-4 py-2 text-sm rounded-lg bg-gray-100 text-gray-700 hover:bg-gray-200">Kustom</button>
            </div>
            <div class="flex text-sm text-gray-500">
                <Database class="w-5 h-5 mr-1"/> Data diupdate setiap 6 jam
            </div>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
            <div class="bg-white p-5 rounded-2xl shadow-sm border border-gray-100">
                <div class="flex items-center justify-between">
                    <div>
                        <p class="text-sm text-gray-500">Total Penumpang (Bulan Ini)</p>
                        <p class="text-2xl font-bold text-gray-800">12,482</p>
                        <p class="flex text-xs text-green-500 mt-1"> <ArrowUp class="w-4 h-4" /> 8.2% </p>
                    </div>
                    <div class="bg-blue-100 w-12 h-12 rounded-full flex items-center justify-center text-blue-600">
                        <User class="w-8 h-8" />
                    </div>
                </div>
            </div>
            <div class="bg-white p-5 rounded-2xl shadow-sm border border-gray-100">
                <div class="flex items-center justify-between">
                    <div>
                        <p class="text-sm text-gray-500">Rata-rata Okupansi</p>
                        <p class="text-2xl font-bold text-gray-800">72.5%</p>
                        <p class="flex text-xs text-green-500 mt-1"><ArrowUp class="w-4 h-4" /> +3.1%</p>
                    </div>
                    <div class="bg-green-100 w-12 h-12 rounded-full flex items-center justify-center text-green-600">
                        <ChartLine class="w-8 h-8" />
                    </div>
                </div>
            </div>
            <div class="bg-white p-5 rounded-2xl shadow-sm border border-gray-100">
                <div class="flex items-center justify-between">
                    <div>
                        <p class="text-sm text-gray-500">Pendapatan (Bulan ini)</p>
                        <p class="text-2xl font-bold text-gray-800">Rp 1.28M</p>
                        <p class="flex text-xs text-green-500 mt-1"><ArrowUp class="w-4 h-4" /> +12.3%</p>
                    </div>
                    <div class="bg-orange-100 w-12 h-12 rounded-full flex items-center justify-center text-orange-600">
                        <DollarSignIcon class="w-8 h-8" />
                    </div>
                </div>
            </div>
            <div class="bg-white p-5 rounded-2xl shadow-sm border border-gray-100">
                <div class="flex items-center justify-between">
                    <div>
                        <p class="text-sm text-gray-500">Total Rute Aktif</p>
                        <p class="text-2xl font-bold text-gray-800">24</p>
                        <p class="flex text-xs text-green-500 mt-1"><ArrowUp class="w-4 h-4" /> +2 dari bulan lalu</p>
                    </div>
                    <div class="bg-purple-100 w-12 h-12 rounded-full flex items-center justify-center text-purple-600">
                        <RouteIcon class="w-8 h-8" />
                    </div>
                </div>
            </div>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
            <div class="bg-white rounded-2xl shadow-sm border border-gray-100 p-5">
                <h3 class="text-lg font-semibold text-gray-800 mb-4">10 Rute Terpopuler (Jumlah Penumpang)</h3>
            </div>
        </div>
    </AdminLayout>
</template>