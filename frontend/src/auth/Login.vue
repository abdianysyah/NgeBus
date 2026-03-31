<script setup>
import { ref } from "vue"
import { login } from "@/services/auth"
import AuthLayout from "@/layouts/AuthLayout.vue"
import { RouterLink, useRouter } from "vue-router"
import { Eye } from "lucide-vue-next"
import { toast } from "vue3-toastify"

const router = useRouter()

const email = ref("")
const password = ref("")
const showPassword = ref(false)


const togglePassword = () => {
    showPassword.value = !showPassword.value
}

const loading = ref(false)

const handleLogin = async () => {
  loading.value = true
  try {
    const res = await login({
      email: email.value,
      password: password.value
    })

    localStorage.setItem("token", res.data.token)
    localStorage.setItem("role", res.data.role)

    await toast("Berhasil Login!",{
      "type":"success",
      "autoClose": 1000,
    })

    if (res.data.role === "admin") {
      router.push("/admin/dashboard")
    } else {
      router.push("/dashboard")
    }

  } catch (err) {
    toast("Gagal Login", {
      "type" : "error",
      "autoClose" : 1000
    })
  } finally {
    loading.value = false
  }
}
</script>

<template>
    <AuthLayout>
        <form @submit.prevent="handleLogin">
            <!-- Email -->
            <div class="mb-5">
                <label for="email" class="block text-sm font-semibold text-gray-700 mb-1">Alamat Email</label>
                <input v-model="email" type="email" id="email" placeholder="nama@email.com" 
                       class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition">
            </div>
            
            <!-- Password -->
            <div class="mb-3">
                <label for="password" class="block text-sm font-semibold text-gray-700 mb-1">Kata Sandi</label>
                <div class="relative">
                    <input v-model="password" :type="showPassword ? 'text' : 'password' " id="password" placeholder="••••••••" 
                           class="w-full border border-gray-300 rounded-xl px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent transition">
                    <button @click="togglePassword" type="button" onclick="togglePassword()" class="absolute right-3 top-3 text-gray-500 hover:text-gray-700">
                        <Eye id="toggleIcon" />
                    </button>
                </div>
            </div>
            
            <!-- Lupa Password & Remember me (opsional) -->
            <div class="flex items-center justify-between mb-6">
                <!-- <label class="flex items-center text-sm text-gray-600">
                    <input type="checkbox" class="rounded border-gray-300 text-orange-500 focus:ring-orange-500 mr-2">
                    Ingat saya
                </label> -->
                <a href="#" class="text-sm text-orange-600 hover:text-orange-700 font-medium">Lupa kata sandi?</a>
            </div>
            
            <!-- Tombol Login -->
            <button :disabled="loading" type="submit" class="w-full bg-orange-500 hover:bg-orange-600 text-white font-semibold py-3 px-4 rounded-xl flex justify-center transition shadow-md">
              <!-- Spinner -->
              <svg
                  v-if="loading"
                  class="animate-spin h-5 w-5 mr-2 text-white"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
              >
                  <circle
                      class="opacity-25"
                      cx="12"
                      cy="12"
                      r="10"
                      stroke="currentColor"
                      stroke-width="4"
                  ></circle>
                  <path
                      class="opacity-75"
                      fill="currentColor"
                      d="M4 12a8 8 0 018-8v8H4z"
                  ></path>
              </svg>
              <span>{{ loading ? 'Loading...' : 'Login' }}</span>
            </button>
            
            <!-- Link Daftar -->
            <p class="text-center text-gray-600 mt-6">
                Belum punya akun?
                <RouterLink to="/register" class="text-orange-600 hover:text-orange-700 font-semibold">Daftar Sekarang</RouterLink>
            </p>
            <div class="text-center">
              <RouterLink to="/" class="text-orange-600 hover:text-orange-700 font-semibold">Kembali ke beranda</RouterLink>
            </div>
        </form>
    </AuthLayout>
</template>