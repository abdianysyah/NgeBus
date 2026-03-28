<script setup>
import Modal from '../ui/Modal.vue';
import { ref, watch, computed, defineProps, defineEmits } from 'vue';
import { 
    Listbox,
    ListboxButton,
    ListboxOptions,
    ListboxOption,

 } from '@headlessui/vue';

const props = defineProps({
    modelValue: Boolean,
    data: {
        type: Object,
        default: () => ({
            id: null,
            company_id: '',
            bus_name: '',
            bus_number: '',
            bus_class: '',
            status: '',
            total_seats: 0,
        })
    },
    mode: { type: String, default: 'create' },
    companies: { type: Array, default: () => [] }
})

const emits = defineEmits(['update:modelValue', 'submit'])
const showModal = computed({
    get: () => props.modelValue,
    set: (val) => emits('update:modelValue', val)
})

const form = ref({ ...props.data })

watch(() => props.data, (val) => {
    form.value = { ...val }
}, { immediate: true })

// Handler Company
const selectedCompany = computed({
    get() {
        return props.companies.find(c => c.id === form.value.company_id) || null
    },
    set(val) {
        form.value.company_id = val ? Number(val.id) : null
    }
})

// Handler Class
const busClasses = [
    'Ekonomi',
    'Bisnis',
    'Executive'
]

const selectedClass = computed({
    get: () => form.value.bus_class,
    set: (val) => form.value.bus_class = val
})

// Handler Status
const busStatus = [
    'Aktif',
    'Nonaktif',
    'Maintenance',
]

const selectedStatus = computed({
    get: () => form.value.status,
    set: (val) => form.value.status = val
})

// Handler submit
const handleSubmit = () => {
    emits('submit', { ...form.value })
    showModal.value = false
}
</script>

<template>
    <Modal v-model="showModal" :title="mode === 'edit' ? 'Edit Data Bus' : 'Tambah Data Bus'" @confirm="handleSubmit">
        <div class="space-y-4 text-black">
            <form @submit.prevent="handleSubmit">
                <!-- Company -->
                <div class="mb-2">
                    <label for="pobus" class="block text-sm font-semibold mb-1">PO Bus</label>
                    <Listbox v-model="selectedCompany">
                        <div class="relative mt-1">
                            <ListboxButton class="relative w-full cursor-default rounded-lg border-gray-300 bg-white py-2 pl-3 pr-10 text-left border focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition">
                                <span class="block truncate">
                                    {{ selectedCompany?.name_company || 'Pilih PO Bus' }}
                                </span>

                            </ListboxButton>
                            <ListboxOptions class="absolute mt-1 max-h-60 w-full overflow-auto rounded-md bg-white shadow-lg z-50 border focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition">
                                <ListboxOption v-for="c in companies" :key="c.id" :value="c" v-slot="{ active, selected }">
                                    <li :class="[
                                        active ? 'bg-orange-100 text-orange-600' : '',
                                        'cursor-pointer select-none py-2 pl-4 pr-4'
                                    ]">
                                        <span :class="selected ? 'font-semibold' : ''">
                                            {{ c.name_company }}
                                        </span>
                                    </li>
                                </ListboxOption>
                            </ListboxOptions>
                        </div>
                    </Listbox>
                </div>
                <div class="mb-4">
                    <label for="bus_name" class="block text-sm font-semibold mb-1">Nama Bus</label>
                    <input v-model="form.bus_name" type="text" name="bus_name" id="bus_name" class="w-full border border-gray-300 rounded-md px-4 py-2 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition" placeholder="Masukkan Nama Bus" required>
                </div>
                <div class="mb-4">
                    <label for="bus_number" class="block text-sm font-semibold mb-1">Plat Nomor</label>
                    <input v-model="form.bus_number" type="text" name="bus_number" id="bus_number" class="w-full border border-gray-300 rounded-md px-4 py-2 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition" placeholder="Masukkan Nomor Plat Bus" required>
                </div>
                <div class="mb-4">
                    <label for="bus_class" class="block text-sm font-semibold mb-1">Kelas</label>
                    <Listbox v-model="selectedClass">
                        <div class="relative mt-1">
                            <ListboxButton class="relative w-full cursor-default rounded-lg border-gray-300 bg-white py-2 pl-3 pr-10 text-left border focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition">
                                <span class="block truncate">
                                    {{ selectedClass || 'Pilih Kelas Bus' }}
                                </span>
                            </ListboxButton>

                            <ListboxOptions class="absolute mt-1 max-h-60 w-full overflow-auto rounded-md bg-white shadow-lg z-50 border focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition">
                                <ListboxOption 
                                    v-for="cls in busClasses"
                                    :key="cls"
                                    :value="cls"
                                    v-slot="{ active, selected }"
                                >
                                <li
                                    :class="[
                                        active ? 'bg-orange-100 text-orange-600' : '',
                                        'cursor-pointer select-none py-2 pl-4 pr-4'
                                    ]"
                                >
                                    <span :class="selected ? 'font-semibold text' : ''">{{ cls }}</span>
                                </li>
                                </ListboxOption>
                            </ListboxOptions>
                        </div>
                    </Listbox>
                </div>
                <div class="mb-4">
                    <label for="total_seats" class="block text-sm font-semibold mb-1">Kapasitas Penumpang</label>
                    <input v-model="form.total_seats" type="number" name="total_seats" id="total_seats" class="w-full border border-gray-300 rounded-md px-4 py-2 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition" placeholder="Masukkan Jumlah Kapasitas" required min="0">
                </div>
                <div class="mb-4">
                    <label for="status" class="block text-sm font-semibold mb-1">Status Bus</label>
                    <Listbox v-model="selectedStatus">
                        <div class="relative mt-1">
                            <ListboxButton class="relative w-full cursor-default rounded-lg border-gray-300 bg-white py-2 pl-3 pr-10 text-left border focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition">
                                <span class="block truncate">
                                    {{ selectedStatus || 'Pilih Status' }}
                                </span>
                            </ListboxButton>

                            <ListboxOptions class="absolute mt-1 max-h-60 w-full overflow-auto rounded-md bg-white shadow-lg z-50 border focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition">
                                <ListboxOption 
                                    v-for="status in busStatus"
                                    :key="status"
                                    :value="status"
                                    v-slot="{ active, selected }"
                                >
                                <li
                                    :class="[
                                        active ? 'bg-orange-100 text-orange-600' : '',
                                        'cursor-pointer select-none py-2 pl-4 pr-4'
                                    ]"
                                >
                                    <span :class="selected ? 'font-semibold text' : ''">{{ status }}</span>
                                </li>
                                </ListboxOption>
                            </ListboxOptions>
                        </div>
                    </Listbox>
                </div>
            </form>
        </div>
    </Modal>
</template>