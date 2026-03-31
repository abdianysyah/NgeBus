<script setup>
import { ref, onMounted, watch } from 'vue';
import AdminLayout from '@/layouts/AdminLayout.vue';
import { PlusCircle, Search, Edit2Icon, Trash } from 'lucide-vue-next';
import { useSchedule } from '@/composables/useSchedule';
import { useBus } from '@/composables/useBus';
import { useRouteBus } from '@/composables/useRouteBus';
import { toast } from 'vue3-toastify';
import ModalForm from '@/components/schedule/ModalForm.vue';
import ModalDelete from '@/components/schedule/ModalDelete.vue';

// Composable
const { schedule, fetchDataSchedule, createSchedule, updateSchedule, removeSchedule } = useSchedule()
const { bus, fetchDataBus } = useBus()
const { rute, fetchDataRoute } = useRouteBus()

// State
const openScheduleModal = ref(false)
const openDeleteModal = ref(false)
const modalMode = ref('create')
const selectedId = ref(null)
const selectedName = ref('')
const scheduleForm = ref({})

// Filter
const filters = ref({
    search: '',
    route: '',
    company: '',
    status: ''
})

// Debounce
let timeout = null

watch(filters, (val) => {
    if (timeout) clearTimeout(timeout)

    timeout = setTimeout(() => {
        fetchDataSchedule(val)
    }, 500)
}, { deep: true })

// =======================
// MODAL HANDLER
// =======================
const handleOpenCreate = () => {
    modalMode.value = 'create'
    scheduleForm.value = {
        id: null,
        bus_id: '',
        route_id: '',
        departure_time: '',
    }
    openScheduleModal.value = true
}

const handleOpenEdit = (item) => {
    modalMode.value = 'edit'
    selectedId.value = item.id
    scheduleForm.value = { ...item }
    openScheduleModal.value = true
}

const handleOpenDelete = (item) => {
    selectedId.value = item.id
    selectedName.value = item.bus?.bus_name || 'Data ini'
    openDeleteModal.value = true
}

// =======================
// SUBMIT
// =======================
const handleSubmitForm = async (data) => {
    try {
        if (modalMode.value === 'create') {
            await createSchedule(data)
            toast.success('Jadwal berhasil ditambahkan')
        } else {
            await updateSchedule(selectedId.value, data)
            toast.success('Jadwal berhasil diperbarui')
        }

        fetchDataSchedule(filters.value)
        openScheduleModal.value = false

    } catch (error) {
        toast.error(error.response?.data?.message || 'Terjadi kesalahan')
    }
}

const handleConfirmDelete = async () => {
    try {
        await removeSchedule(selectedId.value)
        toast.success('Jadwal berhasil dihapus')
        fetchDataSchedule(filters.value)
        openDeleteModal.value = false
    } catch (error) {
        toast.error(error.response?.data?.message || 'Terjadi kesalahan')
    }
}

const formatDate = (date) => {
  if (!date) return '-'

  const d = new Date(date)

  return d.toLocaleString('id-ID', {
    day: '2-digit',
    month: 'short',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

// INIT
onMounted(() => {
    fetchDataSchedule()
    fetchDataRoute()
    fetchDataBus()
})
</script>

<template>
    <AdminLayout>
        <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-6">
            <div>
                <h2 class="text-2xl lg:text-3xl font-bold text-gray-800">Kelola Jadwal Keberangkatan</h2>
                <p class="text-gray-600">Daftar semua jadwal keberangkatan yang tersedia</p>
            </div>
            <div class="mt-4 sm:mt-0">
                <button @click="handleOpenCreate" class="bg-orange-500 hover:bg-orange-600 hover:shadow-none text-white font-semibold px-5 py-3 rounded-xl shadow-md transition inline-flex items-center">
                    <PlusCircle class="mr-2" /> Tambah Jadwal
                </button>
            </div>
        </div>

        <div class="bg-white p-4 rounded-xl shadow-sm border-gray-100 mb-6 flex flex-col sm:flex-row gap-3 items-center justify-between">
            <div class="relative w-full sm:w-64">
                <Search class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" />
                <input type="text" name="" id="" class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-orange-500" placeholder="Cari Rute, PO Bus, dll">
            </div>
            <div class="flex gap-2 w-full sm:w-auto">
                <select v-model="statusBus" @change="getAllDataBus" name="" id="" class="border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-orange-500">
                    <option value="">Semua Rute</option>
                    <option value="Pending">Pending</option>
                    <option value="Berhasil">Berhasil</option>
                    <option value="Gagal">Gagal</option>
                </select>
                <select v-model="statusBus" @change="getAllDataBus" name="" id="" class="border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-orange-500">
                    <option value="">Semua PO</option>
                    <option value="Pending">Pending</option>
                    <option value="Berhasil">Berhasil</option>
                    <option value="Gagal">Gagal</option>
                </select>
                <select v-model="statusBus" @change="getAllDataBus" name="" id="" class="border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-orange-500">
                    <option value="">Semua Status</option>
                    <option value="Pending">Pending</option>
                    <option value="Berhasil">Berhasil</option>
                    <option value="Gagal">Gagal</option>
                </select>
            </div>
        </div>

        <div class="bg-white rounded-2xl shadow-sm border-gray-100 overflow-hidden">
            <div class="overflow-x-auto">
                <table class="min-w-full divide-y divide-gray-200">
                    <thead class="bg-gray-50">
                        <tr>
                            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">No</th>
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">PO Bus</th>
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Nama Bus</th>
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Rute</th>
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Jam Keberangkatan</th>
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Aksi</th>
                        </tr>
                    </thead>
                    <tbody class="bg-white divide-y divide-gray-200">
                        <template v-if="schedule.length > 0">
                            <tr v-for="(item, index) in schedule" :key="item.id" class="hover:bg-gray-50">
                                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ index + 1 }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-center text-gray-700">{{ item.Bus?.company?.name_company }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-center text-gray-700">{{ item.Bus?.bus_name }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-center text-gray-700">{{ item.Route?.origin }} -> {{ item.Route?.destination }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-center text-gray-700">{{ formatDate(item.departure_time) }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-gray-700 flex justify-center gap-2">
                                    <button @click="handleOpenEdit(item)" class="px-2 py-2 bg-blue-500 hover:bg-blue-700 text-white rounded-md transition ease-in-out" title="Edit">
                                        <Edit2Icon class="w-5 h-5" />
                                    </button>
                                    <button @click="handleOpenDelete(item)" class="px-2 py-2 bg-red-500 hover:bg-red-700 text-white rounded-md transition ease-in-out" title="Hapus">
                                        <Trash class="w-5 h-5" />
                                    </button>
                                </td>
                            </tr>
                        </template>

                        <tr v-else >
                            <td colspan="6" class="text-center py-4 text-gray-500">Tidak ada data</td>
                        </tr>
                    </tbody>
                </table>
            </div>
            <div class="px-6 py-4 bg-gray-50 border-t border-gray-200 flex items-center justify-between">
                <div class="text-sm text-gray-500">
                    Menampilkan <span class="font-medium">1</span> - <span class="font-medium">1</span> dari <span class="font-medium">12</span> Bus
                </div>
                <div class="flex space-x-2">
                    <button class="px-3 py-1 border border-gray-300 rounded-md text-sm text-gray-600 hover:bg-gray-100 disabled:opacity-50" disabled>Sebelumnya</button>
                    <button class="px-3 py-1 border border-gray-300 rounded-md text-sm text-gray-600 hover:bg-gray-100">Selanjutnya</button>
                </div>
            </div>
        </div>

        <ModalForm 
            v-model="openScheduleModal"
            :data="scheduleForm"
            :mode="modalMode"
            :buses="bus"
            :routes="rute"
            @submit="handleSubmitForm"
        />
    </AdminLayout>
</template>