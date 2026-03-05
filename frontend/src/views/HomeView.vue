<script setup>
import { Bus, Ticket, Shield, ArrowRight } from 'lucide-vue-next';
import AppLayout from '@/layouts/AppLayout.vue';

</script>

<template>
  <AppLayout>
    <section class="relative bg-linear-to-r from-blue-900 to-blue-800 overflow-hidden">
      <div class="absolute inset-0">
        <img src="/src/assets/img/image.webp" alt="ambabus" class="w-full h-full object-cover opacity-20 md:opacity-30">
      </div>
      <div class="relative container mx-auto px-4 sm:px-6 lg:px-8 py-16 md:py-24 text-white">
        <div class="max-w-3xl">
          <h1 class="text-4xl md:text-5xl font-bold leading-tight mb-4">Perjalanan Nyaman <br><span class="text-orange-300">dengan Bus Terpercaya</span></h1>
          <p class="text-lg md:text-xl mb-8 text-gray-200">Pesan tiket bus online, mudah, cepat, dan aman. Banyak pilihan rute dan bus favorit.</p>
        </div>
    
        <!-- Form -->
        <div class="bg-white/95 backdrop-blur-sm p-5 rounded-xl shadow-xl max-w-4xl">
          <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
            <div>
              <label for="dari" class="block text-gray-800 text-sm font-semibold mb-1">Dari</label>
              <input id="dari" type="text" placeholder="Kota Asal" class="w-full border border-gray-300 rounded-lg px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 text-gray-900">
            </div>
            <div>
              <label for="ke" class="block text-gray-800 text-sm font-semibold mb-1">Ke</label>
              <input id="ke" type="text" placeholder="Kota Tujuan" class="w-full border border-gray-300 rounded-lg px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 text-gray-900">
            </div>
            <div>
              <label for="tanggal" class="block text-gray-800 text-sm font-semibold mb-1">Tanggal</label>
              <input id="tanggal" type="date" placeholder="Kota Asal" class="w-full border border-gray-300 rounded-lg px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 text-gray-900">
            </div>
            <div class="flex items-end">
              <button class="w-full bg-orange-500 hover:bg-orange-600 text-white font-bold py-3 px-4 rounded-lg transition shadow-md">Cari Tiket</button>
            </div>
          </div>
        </div>
    
        <div class="flex flex-wrap gap-8 mt-12 text-white">
          <div>
            <span class="block text-3xl font-bold">100+</span>
            <span class="text-sm uppercase tracking-wider opacity-80">Bus Tersedia</span>
          </div>
          <div>
            <span class="block text-3xl font-bold">50+</span>
            <span class="text-sm uppercase tracking-wider opacity-80">Kota Tujuan</span>
          </div>
          <div>
            <span class="block text-3xl font-bold">10k+</span>
            <span class="text-sm uppercase tracking-wider opacity-80">Pelanggan Puas</span>
          </div>
        </div>
      </div>
    </section>
    
    <section class="py-16 bg-white">
      <div class="container mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-12">
          <h2 class="text-2xl md:text-4xl font-bold text-gray-800">Kenapa Pilih NgeBus?</h2>
          <p class="text-gray-600 mt-3 max-w-2xl mx-auto">Kami hadir untuk membuat perjalanan Anda lebih nyaman tanpa ribet.</p>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <!-- Fitur 1 -->
          <div class="group text-center p-6 border border-gray-200 rounded-2xl hover:shadow-xl transition hover:-translate-y-1 bg-white">
            <div class="bg-orange-100 w-16 h-16 rounded-2xl flex items-center justify-center mx-auto mb-4 group-hover:bg-orange-200 transition">
                <Bus class="text-orange-600 text-2xl" />
            </div>
            <h3 class="text-xl font-semibold mb-2">Armada Terbaik</h3>
            <p class="text-gray-600">Bus dengan fasilitas AC, reclining seat, dan hiburan onboard untuk kenyamanan maksimal.</p>
          </div>
          <!-- Fitur 2 -->
          <div class="group text-center p-6 border border-gray-200 rounded-2xl hover:shadow-xl transition hover:-translate-y-1 bg-white">
            <div class="bg-orange-100 w-16 h-16 rounded-2xl flex items-center justify-center mx-auto mb-4 group-hover:bg-orange-200 transition">
                <Ticket class="text-orange-600 text-2xl" />
            </div>
            <h3 class="text-xl font-semibold mb-2">Pemesanan Mudah</h3>
            <p class="text-gray-600">Cukup beberapa klik, tiket bus Anda sudah dipesan. Bayar via berbagai metode.</p>
          </div>
          <!-- Fitur 3 -->
          <div class="group text-center p-6 border border-gray-200 rounded-2xl hover:shadow-xl transition hover:-translate-y-1 bg-white">
            <div class="bg-orange-100 w-16 h-16 rounded-2xl flex items-center justify-center mx-auto mb-4 group-hover:bg-orange-200 transition">
                <Shield class="text-orange-600 text-2xl" />
            </div>
            <h3 class="text-xl font-semibold mb-2">Aman & Terpercaya</h3>
            <p class="text-gray-600">Data Anda aman, pembayaran terenkripsi, dan customer service 24/7 siap membantu.</p>
          </div>
        </div>
      </div>
    </section>
    
    <section class="py-16 bg-gray-100">
      <div class="container mx-auto px-4 sm:px-6 lg:px-8">
        <h2 class="text-3xl md:text-4xl font-bold text-center text-gray-800 mb-12">Rute Populer</h2>
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
          <div class="bg-white rounded-2xl shadow-md overflow-hidden hover:shadow-xl transition hover:-translate-y-1">
            <img src="https://images.unsplash.com/photo-1598965402089-897ce52e8355?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=60" alt="Jakarta" class="w-full h-44 object-cover">
            <div class="p-5">
              <h3 class="font-bold text-lg">Jakarta - Bandung</h3>
              <p class="text-gray-600 text-sm mb-2">Mulai dari <span class="font-semibold text-orange-600">Rp100.000</span></p>
              <a href="#" class="text-orange-500 text-sm font-semibold hover:underline inline-flex items-center">Lihat jadwal <ArrowRight class="text-xs"/></a>
            </div>
          </div>
          <div class="bg-white rounded-2xl shadow-md overflow-hidden hover:shadow-xl transition hover:-translate-y-1">
            <img src="https://images.unsplash.com/photo-1598965402089-897ce52e8355?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=60" alt="Jakarta" class="w-full h-44 object-cover">
            <div class="p-5">
              <h3 class="font-bold text-lg">Surabaya - Malang</h3>
              <p class="text-gray-600 text-sm mb-2">Mulai dari <span class="font-semibold text-orange-600">Rp75.000</span></p>
              <a href="#" class="text-orange-500 text-sm font-semibold hover:underline inline-flex items-center">Lihat jadwal <ArrowRight class="text-xs"/></a>
            </div>
          </div>
          <div class="bg white rounded-2xl shadow-md overflow-hidden hover:shadow-xl transition hover:-translate-y-1">
            <img src="https://images.unsplash.com/photo-1519501025264-65ba15a82390?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=60" alt="Yogyakarta" class="w-full h-44 object-cover">
            <div class="p-5">
              <h3 class="font-bold text-lg">Yogyakarta - Solo</h3>
              <p class="text-gray-600 text-sm mb-2">Mulai dari <span class="font-semibold text-orange-600">Rp50.000</span></p>
              <a href="#" class="text-orange-500 text-sm font-semibold hover:underline inline-flex items-center">Lihat jadwal <ArrowRight class="text-xs"/></a>
            </div>
          </div>
          <div class="bg white rounded-2xl shadow-md overflow-hidden hover:shadow-xl transition hover:-translate-y-1">
            <img src="https://images.unsplash.com/photo-1555899434-94d1368aa7af?ixlib=rb-4.0.3&auto=format&fit=crop&w=500&q=60" alt="Bali" class="w-full h-44 object-cover">
            <div class="p-5">
              <h3 class="font-bold text-lg">Denpasar - Gilimanuk</h3>
              <p class="text-gray-600 text-sm mb-2">Mulai dari <span class="font-semibold text-orange-600">Rp85.000</span></p>
              <a href="#" class="text-orange-500 text-sm font-semibold hover:underline inline-flex items-center">Lihat jadwal <ArrowRight class="text-xs"/></a>
            </div>
          </div>
        </div>
      </div>
    </section>
    
    <section class="py-16 bg-white">
      <div class="container mx-auto px-4 sm:px-6 lg:px-8">
        <h2 class="text-3xl md:text-4xl font-bold text-center text-gray-800 mb-12">Apa Kata Mereka?</h2>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <!-- Testi 1 -->
          <div class="bg-gray-50 p-6 rounded-2xl shadow-sm hover:shadow-md transition">
            <div class="flex items-center mb-4">
              <img src="https://randomuser.me/api/portraits/women/44.jpg" alt="Siti" class="w-12 h-12 rounded-full mr-4 object-cover border-2 border-orange-200">
              <div>
                  <h4 class="font-semibold">Siti Aisyah</h4>
                  <div class="text-yellow-400 text-sm">★★★★★</div>
              </div>
            </div>
            <p class="text-gray-600">"Pesan tiket di NgeBus sangat mudah. Banyak pilihan bus dan harganya bersaing. Recommended!"</p>
          </div>
          <!-- Testi 2 -->
          <div class="bg-gray-50 p-6 rounded-2xl shadow-sm hover:shadow-md transition">
            <div class="flex items-center mb-4">
              <img src="https://randomuser.me/api/portraits/women/44.jpg" alt="Siti" class="w-12 h-12 rounded-full mr-4 object-cover border-2 border-orange-200">
              <div>
                  <h4 class="font-semibold">Amba Rusdy</h4>
                  <div class="text-yellow-400 text-sm">★★★★★</div>
              </div>
            </div>
            <p class="text-gray-600">"Pelayanan cepat, pembayaran aman. Saya sudah beberapa kali pakai dan selalu puas."</p>
          </div>
          <!-- Testi 3 -->
          <div class="bg-gray-50 p-6 rounded-2xl shadow-sm hover:shadow-md transition">
            <div class="flex items-center mb-4">
              <img src="https://randomuser.me/api/portraits/women/44.jpg" alt="Siti" class="w-12 h-12 rounded-full mr-4 object-cover border-2 border-orange-200">
              <div>
                  <h4 class="font-semibold">Amba Rustank</h4>
                  <div class="text-yellow-400 text-sm">★★★★★</div>
              </div>
            </div>
            <p class="text-gray-600">"Aplikasinya keren, fitur lengkap. Suka banget sama pilihan bus eksekutifnya."</p>
          </div>
        </div>
      </div>
    </section>
  </AppLayout>
</template>
