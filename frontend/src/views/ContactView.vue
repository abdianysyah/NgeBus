<script setup>
import AppLayout from '@/layouts/AppLayout.vue';
import { MapPin, Phone, Clock, MessageCircle, FacebookIcon, InstagramIcon, TwitterIcon, YoutubeIcon, TwitchIcon, Instagram, Youtube } from 'lucide-vue-next';

</script>

<template>
    <AppLayout>
        <section class="bg-linear-to-r from-blue-900 to-blue-800 text-white py-10">
          <div class="container mx-auto px-4 sm:px-6 lg:px-8">
            <h1 class="text-3xl md:text-4xl font-bold mb-2">Hubungi Kami</h1>
            <p class="text-blue-100 max-w-2xl">
              Ada pertanyaan atau butuh bantuan? Tim kami siap membantu Anda 24/7
            </p>
          </div>
        </section>

        <section class="py-16 bg-white">
            <div class="container mx-auto px-4 sm:px-6 lg:px-8">
                <div class="grid grid-cols-1 lg:grid-cols-2 gap-12">
                    <div>
                        <h2 class="text-3xl font-bold text-gray-800 mb-6">Informasi Kontak</h2>
                        <div class="space-y-6">
                            <div class="flex items-start gap-4">
                                <div class="bg-orange-100 w-12 h-12 rounded-xl flex items-center justify-center text-orange-600">
                                    <MapPin />
                                </div>
                                <div>
                                    <h3 class="font-semibold text-gray-800">Kantor Pusat</h3>
                                    <p class="text-gray-600">Jl. Kebon Sirih No. 123, Jakarta Pusat 10110, Indonesia</p>
                                </div>
                            </div>
                            <div class="flex items-start gap-4">
                                <div class="bg-orange-100 w-12 h-12 rounded-xl flex items-center justify-center text-orange-600">
                                    <Phone />
                                </div>
                                <div>
                                    <h3 class="font-semibold text-gray-800">Telepon</h3>
                                    <p class="text-gray-600">(021) 1234-5678</p>
                                    <p class="text-gray-600">WhatsApp: 0812-3456-7890</p>
                                </div>
                            </div>
                            <div class="flex items-start gap-4">
                                <div class="bg-orange-100 w-12 h-12 rounded-xl flex items-center justify-center text-orange-600">
                                    <MessageCircle />
                                </div>
                                <div>
                                    <h3 class="font-semibold text-gray-800">Email</h3>
                                    <p class="text-gray-600">halo@ngebus.id</p>
                                    <p class="text-gray-600">cs@ngebus.id (Customer Support)</p>
                                </div>
                            </div>
                            <div class="flex items-start gap-4">
                                <div class="bg-orange-100 w-12 h-12 rounded-xl flex items-center justify-center text-orange-600">
                                    <Clock />
                                </div>
                                <div>
                                    <h3 class="font-semibold text-gray-800">Jam Operasional</h3>
                                    <p class="text-gray-600">Senin - Jumat: 08:00 - 20:00</p>
                                    <p class="text-gray-600">Sabtu - Minggu: 09:00 - 17:00</p>
                                    <p class="text-gray-600">Layanan Darurat 24/7: (021) 9876-5432</p>
                                </div>
                            </div>
                        </div>
                        <div class="mt-8">
                            <h3 class="font-semibold text-gray-800 mb-3">Ikuti Kami</h3>
                            <div class="flex gap-4">
                                <a href="#" class="w-10 h-10 bg-gray-100 rounded-full flex items-center justify-center text-gray-600 hover:bg-orange-500 hover:text-white transition">
                                    <FacebookIcon />
                                </a>
                                <a href="#" class="w-10 h-10 bg-gray-100 rounded-full flex items-center justify-center text-gray-600 hover:bg-orange-500 hover:text-white transition">
                                    <TwitterIcon />
                                </a>
                                <a href="#" class="w-10 h-10 bg-gray-100 rounded-full flex items-center justify-center text-gray-600 hover:bg-orange-500 hover:text-white transition">
                                    <InstagramIcon />
                                </a>
                                <a href="#" class="w-10 h-10 bg-gray-100 rounded-full flex items-center justify-center text-gray-600 hover:bg-orange-500 hover:text-white transition">
                                    <YoutubeIcon />
                                </a>
                            </div>
                        </div>
                    </div>

                    <div class="bg-gray-50 p-6 rounded-2xl shadow-sm border border-gray-200">
                        <h2 class="text-2xl font-bold text-gray-800 mb-4">Kirim Pesan</h2>
                        <p class="text-gray-600 mb-6">Isi form di bawah ini, tim kami akan merespon dalam 1x24 jam.</p>
                        <form id="contactForm">
                            <div class="mb-4">
                                <label class="block text-sm font-medium text-gray-700 mb-1">Nama Lengkap</label>
                                <input type="text" name="name" placeholder="Nama Anda" class="w-full border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-orange-500">
                            </div>
                            <div class="mb-4">
                                <label class="block text-sm font-medium text-gray-700 mb-1">Email</label>
                                <input type="email" name="email" placeholder="email@example.com" class="w-full border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-orange-500">
                            </div>
                            <div class="mb-4">
                                <label class="block text-sm font-medium text-gray-700 mb-1">Subjek</label>
                                <input type="text" name="subject" placeholder="Subjek pesan" class="w-full border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-orange-500">
                            </div>
                            <div class="mb-6">
                                <label class="block text-sm font-medium text-gray-700 mb-1">Pesan</label>
                                <textarea rows="4" name="message" placeholder="Tulis pesan Anda di sini..." class="w-full border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-orange-500"></textarea>
                            </div>
                            <button type="submit" class="w-full bg-orange-500 hover:bg-orange-600 text-white font-semibold py-3 rounded-xl transition shadow-md">
                                <i class="fas fa-paper-plane mr-2"></i> Kirim Pesan
                            </button>
                        </form>
                        <p id="formMessage" class="text-sm text-center mt-3 text-green-600 hidden"></p>
                    </div>
                </div>
            </div>
        </section>
    </AppLayout>
</template>