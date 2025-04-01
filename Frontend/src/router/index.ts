import { createRouter, createWebHistory } from 'vue-router';
import type { RouteRecordRaw } from 'vue-router';
import AboutView from "@/views/AboutView.vue";
import RegisterView from "@/views/RegisterView.vue";
import LoginView from "@/views/LoginView.vue";
import HomeView from "@/views/HomeView.vue";
import DashboardView from "@/views/DashboardView.vue";
import TaskView from "@/views/TaskView.vue"; // Импортируем компонент задачи
import { useAuthStore } from "@/store/auth.ts";

const routes: Array<RouteRecordRaw> = [
  { path: '/', component: HomeView },
  { path: '/login', component: LoginView },
  { path: '/register', component: RegisterView },
  { path: '/about', component: AboutView },
  { path: '/dashboard', component: DashboardView, meta: { requiresAuth: true } },
  { path: '/task/:id', component: TaskView, props: true },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach(async (to, _, next) => {
  const authStore = useAuthStore();

  if (to.meta.requiresAuth) {
    try {
      await authStore.fetchUser();
      next();
    } catch {
      next('/login');
    }
  } else {
    next();
  }
});

export default router;
