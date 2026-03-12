import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Login from '@/auth/Login.vue'
import Register from '@/auth/Register.vue'
import DashboardAdmin from '@/views/admin/DashboardAdmin.vue'
import DashboardUser from '@/views/user/DashboardUser.vue'
import NotFound from '@/views/error/NotFound.vue'
import DaftarBus from '@/views/admin/DaftarBus.vue'
import DaftarRoute from '@/views/admin/DaftarRoute.vue'
import DaftarUser from '@/views/admin/DaftarUser.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/register',
      name: 'register',
      component: Register
    },
    {
      path: "/admin/dashboard",
      name: 'Dashboard Admin',
      component: DashboardAdmin
    },
    {
      path: "/admin/data-bus",
      name: 'Daftar Bus',
      component: DaftarBus
    },
    {
      path: "/admin/data-route",
      name: 'Daftar Rute',
      component: DaftarRoute
    },
    {
      path: "/admin/data-users",
      name: 'Daftar Pengguna',
      component: DaftarUser 
    },
    {
      path: "/dashboard",
      name: 'Dashboard User',
      component: DashboardUser
    },
    // {
    //   path: '/about',
    //   name: 'about',
    //   // route level code-splitting
    //   // this generates a separate chunk (About.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: () => import('../views/AboutView.vue'),
    // },
    {
      path: "/:pathMatch(.*)*",
      name: 'Not Found',
      component: NotFound
    },

  ],
  
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem("token")
  const role = localStorage.getItem("role")

  const publicPages = ["/login", "/register", "/"]
  const authRequired = !publicPages.includes(to.path)

  if (authRequired && !token) {
    return next("/login")
  }

  if (token && to.path === "/login") {
    if (role === "admin") {
      return next("/admin/dashboard")
    } else {
      return next("/dashboard")
    }
  }

  if (to.path.startsWith("/admin") && role !== "admin") {
    return next("/dashboard")
  }

  next()
})

export default router
