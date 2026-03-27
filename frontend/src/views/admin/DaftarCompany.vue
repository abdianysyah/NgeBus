<script setup>
import AdminLayout from '@/layouts/AdminLayout.vue';
import { PlusCircle, Search, Edit2Icon, Trash, Eye } from 'lucide-vue-next';
import Modal from '@/components/ui/Modal.vue';
import { toast } from 'vue3-toastify';
import { ref } from 'vue';

const createCompany = async (data) => data
const updateCompany = async (id, data) => data

const openModal = ref(false)
const isEdit = ref(false)

const form = ref({
    id: null,
    name_company: '',
    total_bus: 0,
})

const openCreate = () => {
    isEdit.value = false
    form.value = {
        id: null,
        name_company: '',
        total_bus: 0
    }
    openModal.value = true
}

const openEdit = (data) => {
    isEdit.value = true
    form.value = { ...data }
    openModal.value = true
}

const handleSubmit = async () => {
    try {
        if (!form.value.name_company) {
            toast.error("Nama PO Wajib diisi!", {
                autoClose: 1000,
            })
        } else {
            if (isEdit.value) {
                await updateCompany(form.value.id, form.value)
    
                toast.success("Data Berhasil diupdate!", {
                    autoClose: 1000
                })
            } else {
                await createCompany(form.value)
    
                toast.success("Data Berhasil ditambahkan!", {
                    autoClose: 1000
                })
            }
        }
        openModal.value = false
    } catch (error) {
        toast.error("Gagal menyimpan data!", {
            autoClose: 1000
        })
    }
}
</script>

<template>
    <AdminLayout>
        <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-6">
            <div>
                <h2 class="text-2xl lg:text-3xl font-bold text-gray-800">Kelola PO Bus</h2>
                <p class="text-gray-600">Daftar semua armada bus yang terdaftar.</p>
            </div>
            <div class="mt-4 sm:mt-0">
                <button @click="openCreate" class="bg-orange-500 hover:bg-orange-600 hover:shadow-none text-white font-semibold px-5 py-3 rounded-xl shadow-md transition inline-flex items-center">
                    <PlusCircle class="mr-2" /> Tambah PO Bus
                </button>
            </div>
        </div>

        <div class="bg-white p-4 rounded-xl shadow-sm border-gray-100 mb-6 flex flex-col sm:flex-row gap-3 items-center justify-between">
            <div class="relative w-full sm:w-64">
                <Search class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" />
                <input v-model="search" type="text" name="" id="" class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-orange-500" @input="getAllDataBus" placeholder="PO Bus">
            </div>
        </div>

        <div class="bg-white rounded-2xl shadow-sm border-gray-100 overflow-hidden">
            <div class="overflow-x-auto">
                <table class="min-w-full divide-y divide-gray-200">
                    <thead class="bg-gray-50">
                        <tr>
                            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">No</th>
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Nama PO Bus</th>
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Total Armada</th>
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Aksi</th>
                        </tr>
                    </thead>
                    <!-- <tbody class="bg-white divide-y divide-gray-200">
                        <template v-if="bus.length > 0">
                            <tr v-for="(item, index) in bus" :key="item.id" class="hover:bg-gray-50">
                                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ index + 1 }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-center text-gray-700">{{ item.bus_number }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-center text-gray-700">{{ item.bus_name }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-center text-gray-700">{{ item.total_seats }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-center">
                                    <span class="px-2 py-1 text-xs rounded-full" 
                                        :class="{
                                            'bg-green-100 text-green-800' : item.status === 'active',
                                            'bg-red-100 text-red-800' : item.status === 'nonactive',
                                            'bg-yellow-100 text-yellow-800' : item.status === 'maintenance'

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
                    </tbody> -->
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

        <!-- Modal -->
        <Modal v-model="openModal" :title="isEdit ? 'Edit Data' : 'Tambah Data'" @confirm="handleSubmit">
            <div class="space-y-3 text-black">
                <form @submit.prevent="handleSubmit">
                    <div class="mb-4">
                        <label for="name" class="block text-sm font-semibold text-gray-700 mb-1">Nama PO Bus</label>
                        <input id="name" type="text" v-model="form.name_company" placeholder="Masukkan Nama PO Bus" class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition" required>
                    </div>
                </form>
            </div>
        </Modal>
    </AdminLayout>
</template>