<script setup>
import { useRouter } from 'vue-router';
import Swal from 'sweetalert2';
const router = useRouter()
const logout = () => {
    Swal.fire({
        icon: "warning",
        title: "logout",
        text: "Anda yakin ingin keluar?",
        showCancelButton: true,
        confirmButtonText: "Logout"
    }).then((result) => {
        if (result.isConfirmed) {
            localStorage.clear()
            router.push("/login")
        }
    })
}
</script>

<template>
    <h2>Ini halaman user</h2>
    <button @click="logout" class="bg-orange-500 text-white hover:bg-orange-700 py-2 px-3 rounded-2xl">Log Out</button>
</template>