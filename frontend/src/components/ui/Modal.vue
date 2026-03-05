<script setup>
import { XIcon } from 'lucide-vue-next'

defineProps({
    show: Boolean
})
const emit = defineEmits(['close'])
const closeModal = () => {
    emit('close')
}
</script>

<template>
    <Transition
        enter-active-class="transition duration-300 ease-out"
        enter-from-class="opacity-0 scale-95"
        enter-to-class="opacity-100 scale-100"
        leave-active-class="transition duration-200 ease-in"
        leave-from-class="opacity-100 scale-100"
        leave-to-class="opacity-0 scale-95"
    >
    <div v-if="show" class="fixed inset-0 backdrop-blur-sm flex items-center justify-center z-50">
        <div class="bg-white rounded-xl p-6 w-96 relative">
            <button class="absolute top-2 right-2 text-gray-500" @click="closeModal">
                <XIcon />
            </button>
            <slot />
        </div>
    </div>
    </Transition>
</template>