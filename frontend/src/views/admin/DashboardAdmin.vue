<script setup>
import { onMounted, ref, nextTick } from 'vue';
import AdminLayout from '@/layouts/AdminLayout.vue';
import { Bus, Ticket, Users, Building, TriangleAlert, TriangleAlertIcon } from 'lucide-vue-next';
import Modal from '@/components/ui/Modal.vue';
import { Chart } from 'chart.js/auto';
import { countData } from '@/services/auth';
import { toast } from 'vue3-toastify';

const showModal = ref(false)
const monthlyChart = ref(null)
const statusChart = ref(null)
const dashboardData = ref({})
const ChartData = ref([])

const open = ref(false)

const handleConfirm = () => {
    toast("Confirmed!", {
        autoClose: 1000,
    });
}

const getDashboardData = async () => {
    try {
        const res = await countData()
        dashboardData.value = res.data.data
        const currentMonth = new Date().getMonth() + 1
        ChartData.value = res.data.data.monthly_orders.slice(0, currentMonth)

        await nextTick()
        createChart()
    } catch (error) {
        console.error(error);
    }
}

const createChart = () => {
    new Chart(monthlyChart.value, {
        type: 'line',
        data: {
            labels: ['Jan', 'Feb', 'Mar', 'Apr', 'Mei', 'Jun', 'Jul', 'Agu', 'Sep', 'Okt', 'Nov', 'Des'],
            datasets: [{
                label: 'Jumlah Pemesanan',
                data: ChartData.value,
                borderColor: '#f97316',
                backgroundColor: 'rgba(249, 115, 22, 0.1)',
                tension: 0.3,
                fill: true,
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
    })
}

onMounted(() => {

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

    getDashboardData()
})

const openModal = () => {
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
}

</script>

<template>
    <AdminLayout>
        <div class="mb-8">
            <h2 class="text-2xl lg:text-3xl font-bold text-gray-800">Dashboard Admin</h2>
            <p class="text-gray-600">Ringkasan data dan aktivitas sistem NgeBus.</p>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
            <div class="bg-white p-5 rounded-2xl shadow-sm border border-gray-100 flex items-center justify-between hover:shadow-md transition">
                <div>
                    <p class="text-sm text-gray-500">Total Bus</p>
                    <p class="text-2xl font-bold text-gray-800">{{ dashboardData.total_bus }}</p>
                </div>
                <div class="bg-blue-100 w-12 h-12 rounded-full flex items-center justify-center text-blue-600">
                    <Bus class="w-7 h-7" />
                </div>
            </div>
            <div class="bg-white p-5 rounded-2xl shadow-sm border border-gray-100 flex items-center justify-between hover:shadow-md transition">
                <div>
                    <p class="text-sm text-gray-500">Total Pemesanan Tiket</p>
                    <p class="text-2xl font-bold text-gray-800">{{ dashboardData.total_orders }}</p>
                </div>
                <div class="bg-green-100 w-12 h-12 rounded-full flex items-center justify-center text-green-600">
                    <Ticket class="w-7 h-7" />
                </div>
            </div>
            <div class="bg-white p-5 rounded-2xl shadow-sm border border-gray-100 flex items-center justify-between hover:shadow-md transition">
                <div>
                    <p class="text-sm text-gray-500">Total Pengguna</p>
                    <p class="text-2xl font-bold text-gray-800">{{ dashboardData.total_users }}</p>
                </div>
                <div class="bg-orange-100 w-12 h-12 rounded-full flex items-center justify-center text-orange-600">
                    <Users class="w-7 h-7" />
                </div>
            </div>
            <div class="bg-white p-5 rounded-2xl shadow-sm border border-gray-100 flex items-center justify-between hover:shadow-md transition">
                <div>
                    <p class="text-sm text-gray-500">Total PO Bus</p>
                    <p class="text-2xl font-bold text-gray-800">Maintenance Mode</p>
                </div>
                <div class="bg-indigo-100 w-12 h-12 rounded-full flex items-center justify-center text-indigo-600">
                    <Building class="w-7 h-7" />
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
                <p class="text-sm text-gray-500">Maintenance Mode</p>
                <div class="relative flex items-center justify-center">
                    <canvas ref="statusChart" class="max-h-full max-w-full"></canvas>
                </div>
            </div>
        </div>
        <button @click="open = true" class="py-3 px-4 bg-amber-500 rounded-xl shadow">Buka Modal</button>

        <Modal v-model="open" title="Testing Modal" @confirm="handleConfirm">
            <p>Cuman Test doang</p>
        </Modal>
    </AdminLayout>
</template>