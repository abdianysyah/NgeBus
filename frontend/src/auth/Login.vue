<script setup>
import AuthLayout from '@/layouts/AuthLayout.vue';
import { RouterLink } from 'vue-router';
import { Eye } from "lucide-vue-next";

</script>

<template>
    <AuthLayout>
        <form action="" method="post">
            <!-- Email -->
            <div class="mb-5">
                <label for="email" class="block text-sm font-semibold text-gray-700 mb-1">Alamat Email</label>
                <input type="email" id="email" placeholder="nama@email.com" 
                       class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition">
            </div>
            
            <!-- Password -->
            <div class="mb-3">
                <label for="password" class="block text-sm font-semibold text-gray-700 mb-1">Kata Sandi</label>
                <div class="relative">
                    <input type="password" id="password" placeholder="••••••••" 
                           class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition">
                    <button type="button" onclick="togglePassword()" class="absolute right-3 top-3 text-gray-500 hover:text-gray-700">
                        <Eye id="toggleIcon" />
                    </button>
                </div>
            </div>
            
            <!-- Lupa Password & Remember me (opsional) -->
            <div class="flex items-center justify-between mb-6">
                <label class="flex items-center text-sm text-gray-600">
                    <input type="checkbox" class="rounded border-gray-300 text-orange-500 focus:ring-orange-500 mr-2">
                    Ingat saya
                </label>
                <a href="#" class="text-sm text-orange-600 hover:text-orange-700 font-medium">Lupa kata sandi?</a>
            </div>
            
            <!-- Tombol Login -->
            <button type="submit" class="w-full bg-orange-500 hover:bg-orange-600 text-white font-semibold py-3 px-4 rounded-xl transition shadow-md">Masuk</button>
            
            <!-- Link Daftar -->
            <p class="text-center text-gray-600 mt-6">
                Belum punya akun?
                <RouterLink to="/register" class="text-orange-600 hover:text-orange-700 font-semibold">Daftar Sekarang</RouterLink>
            </p>
        </form>
    </AuthLayout>
</template>