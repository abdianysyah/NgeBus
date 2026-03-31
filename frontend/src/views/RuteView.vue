<script setup>
import { ref, onMounted } from 'vue'
import AppLayout from '@/layouts/AppLayout.vue'
import L from 'leaflet'
import 'leaflet/dist/leaflet.css'
import { useSchedule } from '@/composables/useSchedule'


const { schedule, fetchDataSchedule } = useSchedule()
const mapContainer = ref(null)
let map = null

const selectedRoute = ref(null)
const selectedSchedule = ref([])


const groupByRoute = (data) => {
  const map = {}

  data.forEach(item => {
    const key = item.route_id
    if (!map[key]) {
      map[key] = {
        route: item.route,
        schedules: []
      }
    }

    map[key].schedules.push(item)
  })

  return Object.values(map)
}

const drawRoutes = () => {
  if (!schedule.value.length) return
  const grouped = groupByRoute(schedule.value)

  grouped.forEach(group => {
    const r = group.route
    if (!r) return 

      const latlngs = [
        [r.origin_lat, r.origin_lng],
        [r.dest_lat, r.dest_lng]
      ]

      const polyline = L.polyline(latlngs, {
        color: '#f97316',
        weight: 4
      }).addTo(map)

      polyline.on('click', () => {
        selectedRoute.value = r
        selectedSchedule.value = group.schedules
        })
        L.marker(latlngs[0]).addTo(map)
        L.marker(latlngs[1]).addTo(map)
    })
}

const pilihTiket = (item) => {
  alert(`Pilih ${item.bus.bus_name} - ${item.bus.company?.name_company}`)
}

onMounted(async () => {
  map = L.map(mapContainer.value).setView([-2.5, 118], 5)

  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png')
    .addTo(map)

  await fetchDataSchedule()

  drawRoutes()

  const grouped = groupByRoute(schedule.value)

  const allPoints = grouped
  .filter(g => g.route)
  .flatMap(g => [
    [g.route.origin_lat, g.route.origin_lng],
    [g.route.dest_lat, g.route.dest_lng]
  ])

  if (allPoints.length) {
    map.fitBounds(allPoints, { padding: [50, 50] })
  }

  setTimeout(() => {
    map.invalidateSize()
  }, 100)
})
</script>

<template>
  <AppLayout>
    <section class="bg-linear-to-r from-blue-900 to-blue-800 text-white py-10">
      <div class="container mx-auto px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl md:text-4xl font-bold mb-2">Pilih Rute Perjalanan</h1>
        <p class="text-blue-100 max-w-2xl">
          Klik rute pada peta untuk melihat bus dan tiket yang tersedia
        </p>
      </div>
    </section>

    <div class="container mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- Peta -->
      <div class="mb-8">
        <div
          ref="mapContainer"
          class="h-125 w-full shadow-lg rounded-2xl overflow-hidden"
        ></div>
        <p class="text-sm text-gray-500 mt-2 text-center">
          Klik pada rute yang ditampilkan (garis biru) untuk melihat bus dan tiket.
        </p>
      </div>

      <div id="route-result" class="bg-white rounded-2xl overflow-hidden p-6">
        <h2 class="text-2xl font-bold text-gray-800 mb-4">Pilih Rute</h2>
        <p class="text-gray-500 mb-6">
          Klik salah satu rute di peta untuk menampilkan bus dan tiket yang tersedia
        </p>

        <div v-if="selectedRoute" class="space-y-4">
          <div class="mb-4">
            <h3 class="text-xl font-bold">{{ selectedRoute.origin }} -> {{ selectedRoute.destination }}</h3>
          </div>
        </div>

        <div v-for="item in selectedSchedule" :key="item.id" class="border p-4 rounded-xl">
          <h4 class="font-bold">
            {{ item.bus?.bus_name }}
            ({{ item.bus?.company?.name_company || 'Error' }})
          </h4>

          <p class="text-sm text-gray-900">Berangkat: {{ new Date(item.departure_time).toLocaleString('id-ID') }}</p>

          <button @click="pilihTiket(item)" class="bg-orange-500 text-white px-3 py-1 rounded mt-2 hover:bg-orange-700">Pilih Tiket</button>
        </div>
      </div>
    </div>
  </AppLayout>
</template>