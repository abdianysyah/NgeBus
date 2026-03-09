import axios from "axios"
import Swal from "sweetalert2"

const api = axios.create({
    baseURL: "http://localhost:8080/api"
})

api.interceptors.request.use((config) => {
    const token = localStorage.getItem("token")

    console.log("TOKEN:", token)

    if (token) {
        config.headers.Authorization = `Bearer ${token}`
    }

    return config
})

api.interceptors.response.use(
    (response) => response,
    (error) => {
        if (error.response && error.response.status === 401) {
            localStorage.removeItem("token")
            window.location.href = "/login"
            Swal.fire({
                icon: "error",
                title: 'Invalid Token',
                text: 'Silahkan Login'
            })
        }
        return Promise.reject(error)
    }
)

export default api