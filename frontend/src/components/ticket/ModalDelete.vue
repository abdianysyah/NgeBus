<script setup>
import Modal from '../ui/Modal.vue';
import { defineProps, defineEmits, computed } from 'vue';

const props = defineProps({
    modelValue: Boolean,
    name: String
})

const emits = defineEmits(['update:modelValue', 'confirm'])
const showModal = computed({
    get: () => props.modelValue,
    set: (val) => emits('update:modelValue', val)
})

const handleConfirm = () => {
    emits('confirm')
    showModal.value = false
}
</script>

<template>
    <Modal v-model="showModal" title="Hapus Data">
        <div class="text-black">
            <p> Yakin ingin menghapus PO <span class="font-semibold">{{ name }}</span> </p>
        </div>

        <template #footer>
            <div class="flex justify-end gap-2 mt-4">
                <button class="px-3 py-2 bg-gray-500 text-white rounded" @click="showModal = false">
                    Batal
                </button>

                <button class="px-3 py-2 bg-red-500 text-white rounded" @click="handleConfirm">
                    Hapus
                </button>
            </div>
        </template>
    </Modal>
</template>