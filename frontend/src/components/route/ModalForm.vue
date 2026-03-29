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
            origin: '',
            destination: '',
            distance: ''
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

const handleSubmit = () => {
    emits('submit', { ...form.value })
    showModal.value = false
}
</script>

<template>
    <Modal v-model="showModal" :title="mode === 'edit' ? 'Edit Data Rute' : 'Tambah Rute Baru'" @confirm="handleSubmit">
        <div class="space-y-3 text-black">
            <form @submit.prevent="handleSubmit">
                <div class="mb-4">
                    <label for="origin" class="block text-sm font-semibold text-gray-700 mb-1">Kota Asal</label>
                    <input type="text" id="origin" v-model="form.origin" placeholder="Masukkan Kota Asal" class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition" required>
                </div>
                <div class="mb-4">
                    <label for="destination" class="block text-sm font-semibold text-gray-700 mb-1">Kota Tujuan</label>
                    <input type="text" id="destination" v-model="form.destination" placeholder="Masukkan Kota Tujuan" class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition" required>
                </div>
                <div class="mb-4">
                    <label for="distance" class="block text-sm font-semibold text-gray-700 mb-1">Kota Tujuan</label>
                    <input type="number" id="dist" v-model="form.distance" placeholder="Masukkan Jarak Tempuh (KM)" class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition" required>
                </div>
            </form>
        </div>
    </Modal>
</template>