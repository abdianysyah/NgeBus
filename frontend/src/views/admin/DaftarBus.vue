<script setup>
import { PlusCircle, Search, Edit2Icon, Trash, Eye } from 'lucide-vue-next';
import AdminLayout from '@/layouts/AdminLayout.vue';
import { onMounted, ref } from 'vue';
import ModalForm from '@/components/bus/ModalForm.vue';
import { useBus } from '@/composables/useBus';
import { useCompany } from '@/composables/useCompany';
import { toast } from 'vue3-toastify';

const { bus, search, fetchDataBus, createBus, updateBus, removeBus } = useBus()
const { company, fetchCompany } = useCompany()

const openBusModal = ref(false)
const busForm = ref({})
const modalMode = ref('create')

const handleOpenCreate = () => {
    modalMode.value = 'create'
    busForm.value = {
        company_id: '',
        bus_name: '',
        bus_number: '',
        bus_class: '',
        status: '',
        total_seats: 0
    }
    openBusModal.value = true
}

const handleOpenEdit = (data) => {
    modalMode.value = 'edit'
    busForm.value = { ...data }
    openBusModal.value = true
}

const handleSubmitForm = async (data) => {
    try {
        if (!data.company_id) {
            toast.error("PO Bus wajib dipilih!", { autoClose: 1000 })
            return
        }
        if (!data.bus_name) {
            toast.error("Nama Bus wajib diisi", { autoClose: 1000 })
            return
        }

        if (modalMode.value === 'edit') {
            await updateBus(data.id, data)
            toast.success("Data berhasil diupdate!", { autoClose: 1000 })
        } else {
            await createBus(data)
            toast.success("Bus berhasil ditambahkan!", { autoClose: 1000 })
        }
        await fetchDataBus()
    } catch (error) {
        toast.error("Gagal simpan data!", { autoClose: 1000 })
    }
}

onMounted(() => {
    fetchDataBus()
    fetchCompany()
})

</script>

<template>
    <AdminLayout>
        <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-6">
            <div>
                <h2 class="text-2xl lg:text-3xl font-bold text-gray-800">Kelola Bus</h2>
                <p class="text-gray-600">Daftar semua armada bus yang terdaftar.</p>
            </div>
            <div class="mt-4 sm:mt-0">
                <button @click="handleOpenCreate" class="bg-orange-500 hover:bg-orange-600 hover:shadow-none text-white font-semibold px-5 py-3 rounded-xl shadow-md transition inline-flex items-center">
                    <PlusCircle class="mr-2" /> Tambah Bus
                </button>
            </div>
        </div>

        <div class="bg-white p-4 rounded-xl shadow-sm border-gray-100 mb-6 flex flex-col sm:flex-row gap-3 items-center justify-between">
            <div class="relative w-full sm:w-64">
                <Search class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" />
                <input v-model="search" type="text" name="" id="" class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-orange-500" @input="getAllDataBus" placeholder="Nama Bus / Plat No Bus">
            </div>
            <div class="flex gap-2 w-full sm:w-auto">
                <select name="" id="" class="border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-orange-500">
                    <option value="">Semua Kelas</option>
                    <option value="">Eksekutif</option>
                    <option value="">Patas</option>
                    <option value="">Ekonomi</option>
                </select>
                <select v-model="status" @change="getAllDataBus" name="" id="" class="border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-orange-500">
                    <option value="">Semua Status</option>
                    <option value="active">Aktif</option>
                    <option value="nonaktif">Nonaktif</option>
                    <option value="maintenance">Perbaikan</option>
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
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Plat Nomor</th>
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Nama Bus</th>
                            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Kelas</th>
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Kapasitas</th>
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Status</th>
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Aksi</th>
                        </tr>
                    </thead>
                    <tbody class="bg-white divide-y divide-gray-200">
                        <template v-if="bus.length > 0">
                            <tr v-for="(item, index) in bus" :key="item.id" class="hover:bg-gray-50">
                                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ index + 1 }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-center text-gray-700">{{ item.company_id }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-center text-gray-700">{{ item.bus_number }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-center text-gray-700">{{ item.bus_name }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700">{{ item.bus_class }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-center text-gray-700">{{ item.total_seats }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-center">
                                    <span class="px-2 py-1 text-xs rounded-full" 
                                        :class="{
                                            'bg-green-100 text-green-800' : item.status === 'Aktif',
                                            'bg-red-100 text-red-800' : item.status === 'Nonaktif',
                                            'bg-yellow-100 text-yellow-800' : item.status === 'Maintenance'

                                        }">{{ item.status }}</span>
                                </td>
                                <td class="px-6 py-4 whitespace-nowrap text-gray-700 flex justify-center gap-2">
                                    <button class="px-2 py-2 bg-amber-500 hover:bg-amber-700 text-white rounded-md transition ease-in-out" title="Hapus">
                                        <Eye class="w-5 h-5" />
                                    </button>
                                    <button @click="openBusForm(item)" class="px-2 py-2 bg-blue-500 hover:bg-blue-700 text-white rounded-md transition ease-in-out" title="Edit">
                                        <Edit2Icon class="w-5 h-5" />
                                    </button>
                                    <button @click="deleteDataBus(item.id)" class="px-2 py-2 bg-red-500 hover:bg-red-700 text-white rounded-md transition ease-in-out" title="Hapus">
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

        <!-- Modal Form -->
        <ModalForm 
            v-model="openBusModal"
            :data="busForm"
            :mode="modalMode"
            :companies="company"
            @submit="handleSubmitForm"
        />
    </AdminLayout>
</template>