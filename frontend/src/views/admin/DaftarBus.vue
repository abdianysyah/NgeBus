<script setup>
import { PlusCircle, Search, Edit2Icon, Trash } from 'lucide-vue-next';
import AdminLayout from '@/layouts/AdminLayout.vue';
import { nextTick, onMounted, ref } from 'vue';
import { getDataBus } from '@/services/auth';
import Modal from '@/components/ui/Modal.vue';

const showModal = ref(false)
const bus = ref([])

const getAllDataBus = async () => {
    try {
        const res = await getDataBus()
        bus.value = res.data.data

        await nextTick()
    } catch {
        console.error(error);
    }
}

const openModal = () => {
    showModal.value = true
}

const closeModal = () => {
    showModal.value = false
}

onMounted(() => {
    getAllDataBus()
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
                <button @click="openModal" class="bg-orange-500 hover:bg-orange-600 hover:shadow-none text-white font-semibold px-5 py-3 rounded-xl shadow-md transition inline-flex items-center">
                    <PlusCircle class="mr-2" /> Tambah Bus
                </button>
            </div>
        </div>

        <div class="bg-white p-4 rounded-xl shadow-sm border-gray-100 mb-6 flex flex-col sm:flex-row gap-3 items-center justify-between">
            <div class="relative w-full sm:w-64">
                <Search class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" />
                <input type="text" name="" id="" class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-orange-500" placeholder="Cari bus...">
            </div>
            <div class="flex gap-2 w-full sm:w-auto">
                <select name="" id="" class="border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-orange-500">
                    <option value="">Semua Kelas</option>
                    <option value="">Eksekutif</option>
                    <option value="">Patas</option>
                    <option value="">Ekonomi</option>
                </select>
                <select name="" id="" class="border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-orange-500">
                    <option value="">Semua Status</option>
                    <option value="">Aktif</option>
                    <option value="">Nonaktif</option>
                    <option value="">Perbaikan</option>
                </select>
            </div>
        </div>

        <div class="bg-white rounded-2xl shadow-sm border-gray-100 overflow-hidden">
            <div class="overflow-x-auto">
                <table class="min-w-full divide-y divide-gray-200">
                    <thead class="bg-gray-50">
                        <tr>
                            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">No</th>
                            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Plat Nomor</th>
                            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Nama Bus</th>
                            <!-- <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Kelas</th> -->
                            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Kapasitas</th>
                            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Status</th>
                            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Aksi</th>
                        </tr>
                    </thead>
                    <tbody class="bg-white divide-y divide-gray-200">
                        <tr v-for="(item, index) in bus" :key="item.id" class="hover:bg-gray-50">
                            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ index + 1 }}</td>
                            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700">{{ item.bus_number }}</td>
                            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700">{{ item.bus_name }}</td>
                            <!-- <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700">Eksklusif</td> -->
                            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700">{{ item.total_seats }}</td>
                            <td class="px-6 py-4 whitespace-nowrap">
                                <span class="px-2 py-1 text-xs rounded-full bg-green-100 text-green-800">{{ item.status }}</span>
                            </td>
                            <td class="px-6 py-4 whitespace-nowrap-text-sm text-gray-700">
                                <button class="text-blue-600 hover:text-blue-800 mr-3" title="Edit">
                                    <Edit2Icon />
                                </button>
                                <button class="text-red-600 hover:text-red-800" title="Hapus">
                                    <Trash />
                                </button>
                            </td>
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

        <!-- Modal -->
        <Modal :show="showModal" @close="closeModal">
<form @submit.prevent="submitBus">

        <div class="space-y-4">

            <!-- Nama Bus -->
            <div>
                <label class="block text-sm font-medium">Nama Bus</label>
                <input 
                    type="text"
                    class="w-full border rounded-lg px-3 py-2"
                    placeholder="Masukkan nama bus"
                />
            </div>

            <!-- Nomor Plat -->
            <div>
                <label class="block text-sm font-medium">Nomor Plat</label>
                <input 
                    type="text"
                    class="w-full border rounded-lg px-3 py-2"
                    placeholder="Contoh: B 1234 ABC"
                />
            </div>

            <!-- Kapasitas -->
            <div>
                <label class="block text-sm font-medium">Kapasitas Kursi</label>
                <input 
                    type="number"
                    class="w-full border rounded-lg px-3 py-2"
                    placeholder="Jumlah kursi"
                />
            </div>

            <!-- Status -->
            <div>
                <label class="block text-sm font-medium">Status</label>
                <select 
                    class="w-full border rounded-lg px-3 py-2"
                >
                    <option value="active">Active</option>
                    <option value="maintenance">Maintenance</option>
                    <option value="inactive">Inactive</option>
                </select>
            </div>

        </div>

        <!-- Button -->
        <div class="flex justify-end gap-2 mt-6">

            <button 
                type="button"
                @click="closeModal"
                class="px-4 py-2 bg-gray-200 rounded"
            >
                Cancel
            </button>

            <button 
                type="submit"
                class="px-4 py-2 bg-orange-500 text-white rounded"
            >
                Simpan
            </button>

        </div>

    </form>
        </Modal>
    </AdminLayout>
</template>