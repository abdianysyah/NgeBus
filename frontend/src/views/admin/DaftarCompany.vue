<script setup>
import AdminLayout from '@/layouts/AdminLayout.vue';
import { PlusCircle, Search, Edit2Icon, Trash, Eye } from 'lucide-vue-next';
import { toast } from 'vue3-toastify';
import { onMounted, ref, watch } from 'vue';
import ModalForm from '@/components/company/ModalForm.vue';
import ModalDelete from '@/components/company/ModalDelete.vue';
import { useCompany } from '@/composables/useCompany';

const { company, search, fetchCompany, createCompany, updateCompany, removeCompany, detailCompany } = useCompany();

let timeout = null

watch(search, (value) => {
    clearTimeout(timeout)

    timeout = setTimeout(() => {
        fetchCompany()
    }, 500)
})

// Composable

// State modal
const openFormModal = ref(false);
const openDeleteModal = ref(false);

// State form
const form = ref({ id: null, name_company: '', total_bus: 0 });
const modalMode = ref('create');
const selectedId = ref(null);
const selectedName = ref('')

// Event handler modal
const handleOpenCreate = () => {
    modalMode.value = 'create';
    form.value = { id: null, name_company: '', total_bus: 0 };
    openFormModal.value = true;
}

const handleOpenEdit = (data) => {
    console.log('Data Edit', data);
    modalMode.value = 'edit';
    form.value = { ...data };
    openFormModal.value = true;
}

const handleOpenDelete = (item) => {
    selectedId.value = item.id
    selectedName.value = item.name_company
    openDeleteModal.value = true
}

const handleSubmitForm = async (data) => {
    try {
        if (!data.name_company) {
            toast.error("Error: Nama PO Wajib diisi!", { autoClose: 1000 });
            return;
        }

        if (modalMode.value === 'edit') {
            await updateCompany(data.id, data);
            toast.success("Data berhasil diupdate!", { autoClose: 1000 });
        } else {
            await createCompany(data);
            toast.success("Data berhasil ditambahkan!", { autoClose: 1000 });
        }

        await fetchCompany();
    } catch (error) {
        toast.error("Gagal menyimpan data!", { autoClose: 1000 });
    }
}

const handleDelete = async () => {
    try {
        await removeCompany(selectedId.value);
        toast.success("Data berhasil dihapus!", { autoClose: 1000 });
        await fetchCompany();
        openDeleteModal.value = false
    } catch (err) {
        toast.error("Gagal hapus data!", { autoClose: 1000 })
    }
}

onMounted(() => {
    fetchCompany();
})
</script>

<template>
    <AdminLayout>
        <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-6">
            <div>
                <h2 class="text-2xl lg:text-3xl font-bold text-gray-800">Kelola PO Bus</h2>
                <p class="text-gray-600">Daftar semua armada bus yang terdaftar.</p>
            </div>
            <div class="mt-4 sm:mt-0">
                <button @click="handleOpenCreate" class="bg-orange-500 hover:bg-orange-600 hover:shadow-none text-white font-semibold px-5 py-3 rounded-xl shadow-md transition inline-flex items-center">
                    <PlusCircle class="mr-2" /> Tambah PO Bus
                </button>
            </div>
        </div>

        <div class="bg-white p-4 rounded-xl shadow-sm border-gray-100 mb-6 flex flex-col sm:flex-row gap-3 items-center justify-between">
            <div class="relative w-full sm:w-64">
                <Search class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" />
                <input v-model="search" type="text" name="search" id="" class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-orange-500" placeholder="PO Bus">
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
                    <tbody class="bg-white divide-y divide-gray-200">
                        <template v-if="company.length > 0">
                            <tr v-for="(item, index) in company" :key="item.id" class="hover:bg-gray-50">
                                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ index + 1 }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-center text-gray-700">{{ item.name_company }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-center text-gray-700">{{ item.total_bus }}</td>
                                <td class="px-6 py-4 whitespace-nowrap text-gray-700 flex justify-center gap-2">
                                    <!-- <button class="px-2 py-2 bg-amber-500 hover:bg-amber-700 text-white rounded-md transition ease-in-out" title="Hapus">
                                        <Eye class="w-5 h-5" />
                                    </button> -->
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

        <!-- Modal Form -->        
        <ModalForm
            v-model="openFormModal"
            :data="form"
            :mode="modalMode"
            @submit="handleSubmitForm"
        />

        <!-- Modal Delete -->
        <ModalDelete
            v-model="openDeleteModal"
            :name="selectedName"
            @confirm="handleDelete"
        />
    </AdminLayout>
</template>