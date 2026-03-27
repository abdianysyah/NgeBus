<script setup>
import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue';

defineProps({
  modelValue: Boolean,
  title: String
})

defineEmits(['update:modelValue', 'confirm'])
</script>

<template>
  <TransitionRoot as="template" :show="modelValue">
    <Dialog class="relative z-10" @close="$emit('update:modelValue', false)">
      <TransitionChild
        as="template"
        enter="ease-out duration-300"
        enter-from="opacity-0"
        enter-to="opacity-100"
        leave="ease-in duration-200"
        leave-from="opacity-100"
        leave-to="opacity-0"
      >
        <div class="fixed inset-0 bg-gray-900/50 transition-opacity"></div>
      </TransitionChild>
      <div class="fixed inset-0 z-10 w-screen overflow-y-auto">
        <div class="flex min-h-full items-center justify-center p-4">
          <TransitionChild
            as="template"
            enter="ease-out duration-300"
            enter-from="opacity-0 scale-95"
            enter-to="opacity-100 scale-100"
            leave="ease-in duration-200"
            leave-from="opacity-100 scale-100"
            leave-to="opacity-0 scale-95"
          >
            <DialogPanel class="w-full max-w-lg rounded-xl bg-white p-6 shadow-xl">

              <!-- Header -->
              <DialogTitle class="text-lg font-semibold text-black">
                {{ title }}
              </DialogTitle>

              <div class="mt-4 text-gray-300">
                <slot />
              </div>

              <div class="mt-6 flex justify-end gap-2">
                <button class="px-3 py-2 bg-gray-500 text-white rounded" @click="$emit('update:modelValue', false)">
                  Batal
                </button>

                <button type="submit" class="px-3 py-2 bg-blue-500 text-white rounded" @click="$emit('confirm')">
                  Lakukan
                </button>
              </div>
            </DialogPanel>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>