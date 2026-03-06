<script setup>
import { onMounted, ref } from 'vue';
import AdminLayout from '@/layouts/AdminLayout.vue';
import { Bus, Ticket, Users } from 'lucide-vue-next';
import { Chart } from 'chart.js/auto';

const monthlyChart = ref(null)
const statusChart = ref(null)

onMounted(() => {
    new Chart(monthlyChart.value, {
        type: 'line',
        data: {
            labels: ['Jan', 'Feb', 'Mar', 'Apr', 'Mei', 'Jun', 'Jul', 'Agu', 'Sep', 'Okt', 'Nov', 'Des'],
            datasets: [{
                label: 'Jumlah Pemesanan',
                data: [320, 450, 480, 520, 600, 750, 820, 910, 870, 940, 1020, 1150],
                borderColor: '#f97316',
                backgroundColor: 'rgba(249, 115, 22, 0.1)',
                tension: 0.3,
                fill: true,
                pointBackgroundColor: '#f97316',
                pointBorderColor: '#fff',
                pointRadius: 4,
                pointHoverRadius: 6
            }]
        },
        options: {
            responsive: true,
            maintainAspectRatio: true, // Pertahankan rasio dalam container
            plugins: {
                legend: { display: false }
            },
            scales: {
                y: { beginAtZero: true }
            }
        }
    });

    new Chart(statusChart.value, {
        type: 'doughnut',
        data: {
            labels: ['Selesai', 'Diproses', 'Dijadwalkan', 'Dibatalkan'],
            datasets: [{
                data: [540, 210, 135, 45],
                backgroundColor: ['#10b981', '#f59e0b', '#3b82f6', '#ef4444'],
                borderWidth: 0,
                hoverOffset: 5
            }]
        },
        options: {
            responsive: true,
            maintainAspectRatio: true,
            cutout: '70%',
            plugins: {
                legend: {
                    position: 'bottom',
                    labels: { boxWidth: 12, padding: 15 }
                }
            }
        }
    });


})


</script>

<template>
    <AdminLayout>
        <div class="mb-8">
            <h2 class="text-2xl lg:text-3xl font-bold text-gray-800">Dashboard Admin</h2>
            <p class="text-gray-600">Ringkasan data dan aktivitas sistem NgeBus.</p>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 mb-8">
            <div class="bg-white p-5 rounded-2xl shadow-sm border border-gray-100 flex items-center justify-between hover:shadow-md transition">
                <div>
                    <p class="text-sm text-gray-500">Total Bus</p>
                    <p class="text-2xl font-bold text-gray-800">124</p>
                </div>
                <div class="bg-blue-100 w-12 h-12 rounded-full flex items-center justify-center text-blue-600">
                    <Bus class="w-7 h-7" />
                </div>
            </div>
            <div class="bg-white p-5 rounded-2xl shadow-sm border border-gray-100 flex items-center justify-between hover:shadow-md transition">
                <div>
                    <p class="text-sm text-gray-500">Total Pemesanan Tiket</p>
                    <p class="text-2xl font-bold text-gray-800">342</p>
                </div>
                <div class="bg-green-100 w-12 h-12 rounded-full flex items-center justify-center text-green-600">
                    <Ticket class="w-7 h-7" />
                </div>
            </div>
            <div class="bg-white p-5 rounded-2xl shadow-sm border border-gray-100 flex items-center justify-between hover:shadow-md transition">
                <div>
                    <p class="text-sm text-gray-500">Total Pengguna</p>
                    <p class="text-2xl font-bold text-gray-800">350</p>
                </div>
                <div class="bg-orange-100 w-12 h-12 rounded-full flex items-center justify-center text-orange-600">
                    <Users class="w-7 h-7" />
                </div>
            </div>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
            <div class="lg:col-span-1 bg-white rounded-2xl shadow-sm border border-gray-100 p-5">
                <h3 class="text-lg font-semibold text-gray-800 mb-4">Pemesanan per Bulan</h3>
                <div class=" relative">
                    <!-- Chart here! -->
                    <canvas ref="monthlyChart"></canvas>
                </div>
            </div>

            <div class="bg-white rounded-2xl shadow-sm border border-gray-100 p-5">
                <h3 class="text-lg font-semibold text-gray-800 mb-4">Status Pemesanan</h3>
                <div class="relative flex items-center justify-center">
                    <canvas ref="statusChart" class="max-h-full max-w-full"></canvas>
                </div>
            </div>
        </div>
    </AdminLayout>
</template>