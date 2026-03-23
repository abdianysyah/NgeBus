<script setup>
import { onMounted, ref } from 'vue';
import AdminLayout from '@/layouts/AdminLayout.vue';
import { 
    PlusCircle,
    Search,
    Edit2Icon,
    Trash,
    Eye                         
 } from 'lucide-vue-next';
import { getDataRoute, postDataRoute, editDataRoute, deleteDataRoute  } from '@/services/auth';
import Swal from 'sweetalert2';

const rute = ref([])
const search = ref('')


const getAllRoute = async () => {
    try {
        const res = await getDataRoute({
            search: search.value
        })
        rute.value = res.data.data
    } catch (error) {
        console.error(error);
    }
}

// Tambah dan Edit Data Bus
const FormDataRoute = async (data = null) => {
    const { value } = await Swal.fire({
        title: data ? 'Edit Rute' : 'Tambah Rute',
        html: ` 
            <input type="text" id="origin"
                class="swal2-input"
                placeholder="Kota Asal"
                value="${data?.origin || ''}"
            />

            <input type="text" id="destination"
                class="swal2-input"
                placeholder="Kota Tujuan"
                value="${data?.destination || ''}"
            />

            <input id="distance"
                type="number"
                class="swal2-input"
                placeholder="Waktu Tempuh (menit)"
                value="${data?.distance || ''}"
            />
        `,
        showCancelButton: true,
        confirmButtonText: data ? "Update" : "Simpan",
        focusConfirm: false,
        preConfirm: () => {
            const origin = document.getElementById("origin").value
            const destination = document.getElementById("destination").value
            const distance = Number(document.getElementById("distance").value)

            if (!origin || !destination) {
                Swal.showValidationMessage("Kota Asal & Tujuan Wajib diisi!")
                return false
            }

            return { origin, destination, distance }
        }
    })

    if (!value) return

    try {
        if (data) {
            await editBus(data.id, value)
        } else {
            await postDataRoute(value)
        }

        Swal.fire({
            icon: 'success',
            title: 'Berhasil',
            text: 'Data berhasil disimpan!'
        })

        getAllRoute()

    } catch (error) {
        Swal.fire({
            icon: 'error',
            title: 'Gagal!',
            text: error.response?.data?.error || "Terjadi kesalahan!"
        })
    }
}

const deleteRoute = async (id) => {
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
                await deleteDataRoute(id)
                Swal.fire({
                    icon: "success",
                    title: "Berhasil",
                    text: "Bus Berhasil dihapus!"
                })
                getAllRoute()
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

onMounted(() => {
    getAllRoute()
})
</script>

<template>
    <AdminLayout>
        <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-6">
            <div>
                <h2 class="text-2xl lg:text-3xl font-bold text-gray-800">Kelola Rute</h2>
                <p class="text-gray-600">Daftar semua rute</p>
            </div>
            <div class="mt-4 sm:mt-0">
                <button @click="FormDataRoute()"  class="bg-orange-500 hover:bg-orange-600 hover:shadow-none text-white font-semibold px-5 py-3 rounded-xl shadow-md transition inline-flex items-center">
                    <PlusCircle class="mr-2" /> Tambah Rute
                </button>
            </div>
        </div>
        <div class="bg-white p-4 rounded-xl shadow-sm border-gray-100 mb-6 flex flex-col sm:flex-row gap-3 items-center justify-between">
            <div class="relative w-full sm:w-64">
                <Search class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" />
                <input type="text" name="" id="" class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-orange-500" placeholder="Rute">
            </div>
        </div>

        <div class="bg-white rounded-2xl shadow-sm border-gray-100 overflow-hidden">
            <div class="overflow-x-auto">
                <table class="min-w-full divide-y divide-gray-200">
                    <thead class="bg-gray-50">
                        <tr>
                            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">No</th>
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Kota Asal</th>
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Kota Tujuan</th>
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Waktu Tempuh (Menit)</th>
                            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Aksi</th>
                        </tr>
                    </thead>
                    <tbody class="bg-white divide-y divide-gray-200">
                        <template v-if="rute.length > 0">
                            <tr v-for="(item, index) in rute" :key="item.id" class="hover:bg-gray-50">
                                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ index + 1 }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-center text-gray-700">{{ item.origin }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-center text-gray-700">{{ item.destination }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-center text-gray-700">{{ item.distance }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-gray-700 flex justify-center gap-2">
                                    <button class="px-2 py-2 bg-amber-500 hover:bg-amber-700 text-white rounded-md transition ease-in-out" title="Hapus">
                                        <Eye class="w-5 h-5" />
                                    </button>
                                    <button @click="FormDataRoute(item)" class="px-2 py-2 bg-blue-500 hover:bg-blue-700 text-white rounded-md transition ease-in-out" title="Edit">
                                        <Edit2Icon class="w-5 h-5" />
                                    </button>
                                    <button @click="deleteRoute(item.id)" class="px-2 py-2 bg-red-500 hover:bg-red-700 text-white rounded-md transition ease-in-out" title="Hapus">
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
    </AdminLayout>
</template>