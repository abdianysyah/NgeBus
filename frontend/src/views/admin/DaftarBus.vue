<script setup>
import { PlusCircle, Search, Edit2Icon, Trash, Eye } from 'lucide-vue-next';
import AdminLayout from '@/layouts/AdminLayout.vue';
import { nextTick, onMounted, ref } from 'vue';
import { getDataBus, addBus, editBus, deleteBus } from '@/services/auth';
import Modal from '@/components/ui/Modal.vue';
import Swal from 'sweetalert2';

const showModal = ref(false)
const modalEdit = ref(false)
const bus = ref([])
const search = ref('')
const status = ref('')

const form = ref({
    id: null,
    bus_name: '',
    bus_number: '',
    total_seats: '',
    status: '',
})

const editDataBus = (bus) => {
    form.value.id = bus.id
    form.value.bus_name = bus.bus_name
    form.value.bus_number = bus.bus_number
    form.value.total_seats = bus.total_seats
    form.value.status = bus.status

    modalEdit.value = true
}

const submitBus = async () => {
    try {
        if (form.value.id) {
            await editBus(form.value.id, form.value)
            resetForm()
            closemodalEdit()
        } else {
            await addBus(form.value)
            resetForm()
            closeModal()
        }
        getAllDataBus()
        Swal.fire({
            icon: 'success',
            title: 'Berhasil',
            text: 'Data telah berhasil ditambahkan',
            showCloseButton: false
        })
    } catch (error) {
        Swal.fire({
            icon: 'warning',
            title: 'Gagal',
            text: error.response?.data?.error || 'Terjadi Kesalahan',
            showCancelButton: false
        })
    }
}

const getAllDataBus = async () => {
    try {
        const res = await getDataBus({
            search: search.value,
            status: status.value
        })
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

const openmodalEdit = () => {
    modalEdit.value = true
}

const closemodalEdit = () => {
    modalEdit.value = false
}

const deleteDataBus = async (id) => {
    Swal.fire({
        title: 'Yaking ingin menghapus?',
        text: "Data bus akan dihapus permanen!",
        icon: "warning",
        showCancelButton: true,
        confirmButtonColor: '#f97316',
        confirButtonColor: '#6b7280',
        confirmButtonText: "Ya, Hapus!"
    }).then(async (result) => {
        if (result.isConfirmed) {
            try {
                await deleteBus(id)
                Swal.fire({
                    icon: "success",
                    title: "Berhasil",
                    text: "Bus Berhasil dihapus!"
                })
                getAllDataBus()
            } catch (error) {
                Swal.fire({
                    icon: 'error',
                    title: "Gagal",
                    text: error.response?.data?.error || "Terjadi Kesalahan"
                })
            }
        }
    })
}

const resetForm = () => {
    form.value = {
        id: null,
        bus_name: '',
        bus_number: '',
        total_seats: '',
        status: '',
    }
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
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Plat Nomor</th>
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Nama Bus</th>
                            <!-- <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Kelas</th> -->
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Kapasitas</th>
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Status</th>
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Aksi</th>
                        </tr>
                    </thead>
                    <tbody class="bg-white divide-y divide-gray-200">
                        <template v-if="bus.length > 0">
                            <tr v-for="(item, index) in bus" :key="item.id" class="hover:bg-gray-50">
                                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ index + 1 }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-center text-gray-700">{{ item.bus_number }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-center text-gray-700">{{ item.bus_name }}</td>
                                <!-- <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700">Eksklusif</td> -->
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
                                    <button @click="editDataBus(item)" class="px-2 py-2 bg-blue-500 hover:bg-blue-700 text-white rounded-md transition ease-in-out" title="Edit">
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

        <!-- Modal Edit Data -->
        <Modal :show="modalEdit" @close="closemodalEdit">
            <form @submit.prevent="submitBus">
                <input v-model="form.id" type="hidden" >
                <div class="space-y-4">
                    <div>
                        <label for="nama_bus" class="block text-sm font-medium">Nama Bus</label>
                        <input v-model="form.bus_name" type="text" name="nama_bus" id="" class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition" placeholder="Masukkan Nama Bus" value="">
                    </div>

                    <div>
                        <label for="no_plat" class="block text-sm font-medium">Nomor Plat Bus</label>
                        <input v-model="form.bus_number" type="text" name="no_plat" id="" class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition" placeholder="Ex: B 1234 AV">
                    </div>

                    <div>
                        <label for="total_kursi" class="block text-sm font-medium">Kapasitas Kursi</label>
                        <input v-model="form.total_seats" type="number" name="total_kursi" id="" class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition" placeholder="Jumlah Kursi">
                    </div>

                    <div>
                        <label for="status" class="block text-sm font-medium">Status</label>
                        <select v-model="form.status" name="status" id="" class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition">
                            <option value="active">Aktif</option>
                            <option value="maintenance">Maintenance</option>
                            <option value="inactive">Tidak Aktif</option>
                        </select>
                    </div>
                </div>

                <div class="flex justify-end gap-2 mt-6">
                    <button type="submit" class="bg-orange-500 hover:bg-orange-600 hover:shadow-none text-white font-semibold px-5 py-3 rounded-xl shadow-md transition inline-flex items-center">Simpan</button>
                </div>
            </form>
        </Modal>

        <!-- Modal Tambah Data -->
        <Modal :show="showModal" @close="closeModal">
            <form @submit.prevent="submitBus">
                <div class="space-y-4">
                    <div>
                        <label for="nama_bus" class="block text-sm font-medium">Nama Bus</label>
                        <input v-model="form.bus_name" type="text" name="nama_bus" id="" class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition" placeholder="Masukkan Nama Bus">
                    </div>

                    <div>
                        <label for="no_plat" class="block text-sm font-medium">Nomor Plat Bus</label>
                        <input v-model="form.bus_number" type="text" name="no_plat" id="" class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition" placeholder="Ex: B 1234 AV">
                    </div>

                    <div>
                        <label for="total_kursi" class="block text-sm font-medium">Kapasitas Kursi</label>
                        <input v-model="form.total_seats" type="number" name="total_kursi" id="" class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition" placeholder="Jumlah Kursi">
                    </div>

                    <div>
                        <label for="status" class="block text-sm font-medium">Status</label>
                        <select v-model="form.status" name="status" id="" class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition">
                            <option value="active">Aktif</option>
                            <option value="maintenance">Maintenance</option>
                            <option value="inactive">Tidak Aktif</option>
                        </select>
                    </div>
                </div>

                <div class="flex justify-end gap-2 mt-6">
                    <button type="submit" class="bg-orange-500 hover:bg-orange-600 hover:shadow-none text-white font-semibold px-5 py-3 rounded-xl shadow-md transition inline-flex items-center">Simpan</button>
                </div>
            </form>
        </Modal>
    </AdminLayout>
</template>