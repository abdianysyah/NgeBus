<script setup>
import { ref } from 'vue';
import AuthLayout from '@/layouts/AuthLayout.vue';
import { RouterLink } from 'vue-router';
import { Eye } from 'lucide-vue-next';
import { register } from '@/services/auth';
import Swal from 'sweetalert2';

const name = ref("")
const email = ref("")
const password = ref("")
const phone = ref("")
const confirmPassword = ref("")
const loading = ref(false)
const showPassword = ref(false)

const handleRegister = async () => {
    if (password.value !== confirmPassword.value) {
        Swal.fire({
            icon: "error",
            title: "Password tidak sama!"
        })
        return
    }

    try {
        loading.value = true

        await register({
            name: name.value,
            email: email.value,
            no_tel: phone.value,
            password: password.value
        })

        Swal.fire({
            icon: "success",
            title: "Register berhasil"
        })
    } catch (error) {
        Swal.fire({
            icon: "error",
            title: "Register gagal"
        })
    } finally {
        loading.value = false
    }
}

const togglePassword = () => {
    showPassword.value = !showPassword.value
}
</script>

<template>
    <AuthLayout>
        <form @submit.prevent="handleRegister">
            <!-- Name -->
            <div class="mb-5">
                <label for="username" class="bloc text-sm font-semibold text-gray-700 mb-1">Nama</label>
                <input v-model="name" id="username" type="text" class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition" placeholder="Masukkan Nama Lengkap!">
            </div>
            <!-- Email -->
            <div class="mb-5">
                <label for="email" class="block text-sm font-semibold text-gray-700 mb-1">Email</label>
                <input v-model="email" id="email" type="email" class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition" placeholder="testing@example.com">
            </div>
            <!-- No Hp -->
            <div class="mb-5">
                <label for="no_tel" class="block text-sm font-semibold text-gray-700 mb-1">Nomor Telpon</label>
                <input v-model="phone" id="no_tel" type="text" class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition" placeholder="Nomor Telpon">
                <span class="text-xs text-gray-500">Pastikan nomor yang di isi terdaftar di WhatsApp</span>
            </div>
            <!-- Password -->
            <div class="mb-5">
                <label for="password" class="block text-sm font-semibold text-gray-700 mb-1">Kata Sandi</label>
                <div class="relative">
                    <input v-model="password" :type="showPassword ? 'text' : 'password' " id="password" placeholder="Masukkan Password" 
                           class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition">
                    <button @click="togglePassword" type="button" class="absolute right-3 top-3 text-gray-500 hover:text-gray-700">
                        <Eye id="toggleIcon" />
                    </button>
                </div>
            </div>
            <div class="mb-5">
                <label for="confirmPassword" class="block text-sm font-semibold text-gray-700 mb-1">Konfirmasi Password</label>
                <input v-model="confirmPassword" type="password" class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition" placeholder="Ulang Password">
            </div>
            
            <button type="submit" class="w-full bg-orange-500 hover:bg-orange-600 text-white font-semibold py-3 px-4 rounded-xl transition shadow-md" :disabled="loading">{{ loading ? "Loading..." : "Register" }}</button>

            <!-- Link Daftar -->
            <p class="text-center text-gray-600 mt-6">
                Sudah punya akun?
                <RouterLink to="/login" class="text-orange-600 hover:text-orange-700 font-semibold">Masuk disini</RouterLink>
            </p>
        </form>
    </AuthLayout>
</template>