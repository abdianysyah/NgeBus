<script setup>
import Modal from '../ui/Modal.vue'
import { ref, watch, computed } from 'vue'
import {
  Listbox,
  ListboxButton,
  ListboxOptions,
  ListboxOption,
} from '@headlessui/vue'

const props = defineProps({
  modelValue: Boolean,
  data: {
    type: Object,
    default: () => ({
      id: null,
      bus_id: '',
      route_id: '',
      departure_time: '',
    }),
  },
  mode: {
    type: String,
    default: 'create',
  },
  buses: {
    type: Array,
    default: () => [],
  },
  routes: {
    type: Array,
    default: () => [],
  },
})

const emits = defineEmits(['update:modelValue', 'submit'])

const showModal = computed({
  get: () => props.modelValue,
  set: (val) => emits('update:modelValue', val),
})

const form = ref({
  id: null,
  bus_id: '',
  route_id: '',
  departure_time: '',
})

watch(
  () => props.data,
  (val) => {
    form.value = {
      id: val?.id ?? null,
      bus_id: val?.bus_id ?? '',
      route_id: val?.route_id ?? '',
      departure_time: val?.departure_time 
        ? val.departure_time.slice(0, 16) // Format untuk input datetime-local
        : '',
    }
  },
  { immediate: true, deep: true }
)

const selectedBus = computed({
  get() {
    return (
      props.buses.find((b) => Number(b.id) === Number(form.value.bus_id)) || null
    )
  },
  set(val) {
    form.value.bus_id = val ? Number(val.id) : ''
  },
})

const selectedRoute = computed({
  get() {
    return (
      props.routes.find((r) => Number(r.id) === Number(form.value.route_id)) || null
    )
  },
  set(val) {
    form.value.route_id = val ? Number(val.id) : ''
  },
})

const handleSubmit = () => {
    const payload = {
        ...form.value,
        departure_time: form.value.departure_time ? new Date(form.value.departure_time).toISOString() : null,
    }
  emits('submit', payload)
  showModal.value = false
}
</script>

<template>
  <Modal
    v-model="showModal"
    :title="mode === 'edit' ? 'Edit Jadwal' : 'Tambah Jadwal Baru'"
    @confirm="handleSubmit"
  >
    <div class="space-y-4 text-black">
      <form @submit.prevent="handleSubmit" class="space-y-4">
        <div>
          <label class="block mb-1 font-medium">Bus</label>
          <Listbox v-model="selectedBus">
            <div class="relative">
              <ListboxButton
                class="w-full py-2 pl-3 pr-10 text-left bg-white border rounded-lg shadow-sm cursor-pointer focus:outline-none focus:ring-2 focus:ring-orange-500"
              >
                <span class="block truncate">
                  {{ selectedBus ? selectedBus.bus_name : 'Pilih Bus' }}
                </span>
                <span class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
                  <i class="fas fa-chevron-down text-gray-400"></i>
                </span>
              </ListboxButton>

              <ListboxOptions
                class="absolute z-10 w-full py-1 mt-1 overflow-auto text-base bg-white rounded-md shadow-lg max-h-60 ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm"
              >
                <ListboxOption
                  v-for="bus in buses"
                  :key="bus.id"
                  :value="bus"
                  v-slot="{ active, selected }"
                >
                  <li
                    :class="[
                      'cursor-pointer select-none relative py-2 pl-10 pr-4',
                      active ? 'bg-orange-600 text-white' : 'text-gray-900',
                    ]"
                  >
                    <span :class="['block truncate', selected ? 'font-semibold' : 'font-normal']">
                      {{ bus.bus_name }}
                    </span>
                  </li>
                </ListboxOption>
              </ListboxOptions>
            </div>
          </Listbox>
        </div>

        <div>
          <label class="block mb-1 font-medium">Rute</label>
          <Listbox v-model="selectedRoute">
            <div class="relative">
              <ListboxButton
                class="w-full py-2 pl-3 pr-10 text-left bg-white border rounded-lg shadow-sm cursor-pointer focus:outline-none focus:ring-2 focus:ring-orange-500"
              >
                <span class="block truncate">
                  {{
                    selectedRoute
                      ? `${selectedRoute.origin} - ${selectedRoute.destination}`
                      : 'Pilih Rute'
                  }}
                </span>
                <span class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
                  <i class="fas fa-chevron-down text-gray-400"></i>
                </span>
              </ListboxButton>

              <ListboxOptions
                class="absolute z-10 w-full py-1 mt-1 overflow-auto text-base bg-white rounded-md shadow-lg max-h-60 ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm"
              >
                <ListboxOption
                  v-for="route in routes"
                  :key="route.id"
                  :value="route"
                  v-slot="{ active, selected }"
                >
                  <li
                    :class="[
                      'cursor-pointer select-none relative py-2 pl-10 pr-4',
                      active ? 'bg-orange-600 text-white' : 'text-gray-900',
                    ]"
                  >
                    <span :class="['block truncate', selected ? 'font-semibold' : 'font-normal']">
                      {{ route.origin }} - {{ route.destination }}
                    </span>
                  </li>
                </ListboxOption>
              </ListboxOptions>
            </div>
          </Listbox>
        </div>

        <div>
          <label for="departure_time" class="block mb-1 font-medium">Waktu Keberangkatan</label>
          <input
            id="departure_time"
            v-model="form.departure_time"
            type="datetime-local"
            class="w-full px-3 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition"
          >
        </div>
      </form>
    </div>
  </Modal>
</template>