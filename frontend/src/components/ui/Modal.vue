<script setup>
import { defineProps, defineEmits } from 'vue'

const props = defineProps({
  show: Boolean
})

const emit = defineEmits(['close'])

const closeModal = () => {
  emit('close')
}
</script>

<template>
  <Teleport to="body">

    <Transition
      enter-active-class="transition-opacity duration-300"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition-opacity duration-200"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div
        v-if="show"
        class="fixed inset-0 z-50 flex items-center justify-center"
      >

        <!-- backdrop -->
        <div
          class="absolute inset-0 bg-black/50"
          @click="closeModal"
        ></div>

        <!-- modal -->
        <Transition
          enter-active-class="transition duration-300 ease-out"
          enter-from-class="opacity-0 scale-90 translate-y-4"
          enter-to-class="opacity-100 scale-100 translate-y-0"
          leave-active-class="transition duration-200 ease-in"
          leave-from-class="opacity-100 scale-100 translate-y-0"
          leave-to-class="opacity-0 scale-90 translate-y-4"
        >
          <div
            v-if="show"
            class="relative bg-white rounded-xl shadow-xl w-125 p-6"
          >
            <h2 class="text-lg font-semibold mb-4">
              Modal Title
            </h2>

            <slot />

            <div class="flex justify-end mt-4">
              <button
                @click="closeModal"
                class="px-4 py-2 bg-gray-500 text-white rounded-lg"
              >
                Close
              </button>
            </div>
          </div>
        </Transition>

      </div>
    </Transition>

  </Teleport>
</template>