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
import BusView from '@/views/BusView.vue'
import RuteView from '@/views/RuteView.vue'
import ContactView from '@/views/ContactView.vue'
import DaftarCompany from '@/views/admin/DaftarCompany.vue'

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
      name: 'Login',
      component: Login
    },
    {
      path: '/register',
      name: 'Register',
      component: Register
    },
    {
      path: '/bus',
      name: 'Bus',
      component: BusView
    },
    {
      path: "/rute",
      name: 'Rute',
      component: RuteView
    },
    {
      path: '/contact',
      name: 'Kontak Kami',
      component: ContactView
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
      path: "/admin/data-company",
      name: 'Daftar PO Bus',
      component: DaftarCompany
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

router.beforeEach((to) => {
  document.title = to.name || "NgeBus"
  const token = localStorage.getItem("token")
  const role = localStorage.getItem("role")

  const publicPages = ["/login", "/register", "/", "/bus", "/rute", "/contact"]
  const authRequired = !publicPages.includes(to.path)

  if (authRequired && !token) {
    return "/login"
  }

  if (token && to.path === "/login") {
    if (role === "admin") {
      return "/admin/dashboard"
    } else {
      return "/dashboard"
    }
  }

  if (to.path.startsWith("/admin") && role !== "admin") {
    return "/dashboard"
  }

  return true
})

export default router
