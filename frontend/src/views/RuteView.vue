<script setup>
import { ref, onMounted } from 'vue'
import AppLayout from '@/layouts/AppLayout.vue'
import L from 'leaflet'
import 'leaflet/dist/leaflet.css'

const mapContainer = ref(null)
let map = null

const busList = ref('')

const routes = [
  {
    id: 1,
    name: "Jakarta → Bandung",
    points: [
      [-6.2088, 106.8456],
      [-6.9175, 107.6191]
    ],
    buses: [
      {
        name: "Sanjaya Putra",
        class: "Eksekutif",
        price: 100000,
        facilities: ["AC", "Toilet", "TV"],
        schedule: "08:00, 14:00, 20:00"
      }
    ]
  }
]

const showRouteResult = (route) => {
  let html = `
    <div class="mb-4">
      <h3 class="text-xl font-bold">${route.name}</h3>
    </div>
  `

  route.buses.forEach(bus => {
    html += `
      <div class="border p-4 rounded-xl mb-3">
        <h4 class="font-bold">${bus.name}</h4>
        <p>Rp ${bus.price.toLocaleString()}</p>
        <button onclick="window.pilihTiket('${bus.name}', '${bus.class}', ${bus.price})"
          class="bg-orange-500 text-white px-3 py-1 rounded">
          Pilih Tiket
        </button>
      </div>
    `
  })

  busList.value = html
}

window.pilihTiket = (busName, kelas, harga) => {
  alert(`Pilih ${busName} (${kelas}) - Rp ${harga}`)
}

delete L.Icon.Default.prototype._getIconUrl

L.Icon.Default.mergeOptions({
  iconRetinaUrl: new URL('leaflet/dist/images/marker-icon-2x.png', import.meta.url).href,
  iconUrl: new URL('leaflet/dist/images/marker-icon.png', import.meta.url).href,
  shadowUrl: new URL('leaflet/dist/images/marker-shadow.png', import.meta.url).href,
})

const drawRoutes = () => {
  routes.forEach(route => {
    const latlngs = route.points.map(p => L.latLng(p[0], p[1]))
    const polyline = L.polyline(latlngs, {
      color: '#f97316',
      weight: 4
    }).addTo(map)

    polyline.on('click', () => {
      showRouteResult(route)

      polyline.setStyle({ color: '#3b82f6', weight: 6 })
      setTimeout(() => {
        polyline.setStyle({ color: '#f97316', weight: 4 })
      }, 500)
    })

    L.marker(route.points[0]).addTo(map)
    L.marker(route.points[route.points.length - 1]).addTo(map)
  })
}

onMounted(() => {
  map = L.map(mapContainer.value).setView([-2.5, 118], 5)

  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; OpenStreetMap'
  }).addTo(map)

  drawRoutes()

  const allPoints = routes.flatMap(r => r.points)
  const bounds = L.latLngBounds(allPoints)
  map.fitBounds(bounds, { padding: [50, 50] })

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

        <div class="space-y-4">
          <div v-html="busList"></div>

          <div v-if="!busList" class="text-center text-gray-400 py-8">
            Belum ada rute dipilih
          </div>
        </div>
      </div>
    </div>
  </AppLayout>
</template>