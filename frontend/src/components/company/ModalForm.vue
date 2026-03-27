<script setup>
import Modal from '../ui/Modal.vue';
import { ref, watch, computed, defineProps, defineEmits } from 'vue';

const props = defineProps({
  modelValue: Boolean,
  data: { type: Object, default: () => ({ id: null, name_company: '', total_bus: 0 }) },
  mode: { type: String, default: 'create' } // create | edit
})

const emits = defineEmits(['update:modelValue', 'submit'])

const form = ref({ ...props.data })

// Sync form jika props.data berubah
watch(() => props.data, (newVal) => {
  form.value = { ...newVal }
})

// computed untuk v-model
const showModal = computed({
  get: () => props.modelValue,
  set: (val) => emits('update:modelValue', val)
})

const handleSubmit = () => {
  emits('submit', { ...form.value }) // Kirim data ke parent
  showModal.value = false            // Tutup modal
}
</script>

<template>
  <Modal v-model="showModal" :title="mode==='edit'?'Edit PO Bus':'Tambah PO Bus'" @confirm="handleSubmit">
    <div class="space-y-3 text-black">
      <form @submit.prevent="handleSubmit">
        <div class="mb-4">
          <label for="name" class="block text-sm font-semibold text-gray-700 mb-1">Nama PO Bus</label>
          <input 
            id="name" 
            type="text" 
            v-model="form.name_company" 
            placeholder="Masukkan Nama PO Bus" 
            class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition" 
            required
          >
        </div>
      </form>
    </div>
  </Modal>
</template>