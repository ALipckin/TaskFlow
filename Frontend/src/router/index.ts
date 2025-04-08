import { createRouter, createWebHistory } from 'vue-router';
import type { RouteRecordRaw } from 'vue-router';
import AboutView from "@/views/AboutView.vue";
import RegisterView from "@/views/RegisterView.vue";
import LoginView from "@/views/LoginView.vue";
import HomeView from "@/views/HomeView.vue";
import DashboardView from "@/views/DashboardView.vue";
import TaskEdit from "@/views/TaskEditView.vue";
import CreateTaskView from "@/views/CreateTaskView.vue";
import { useAuthStore } from "@/store/auth.ts";
import MyTasksView from "@/views/MyTasksView.vue";

const routes: Array<RouteRecordRaw> = [
  { path: '/', component: HomeView },
  { path: '/login', component: LoginView, meta: { onlyGuest: true }},
  { path: '/register', component: RegisterView, meta: { onlyGuest: true }},
  { path: '/about', component: AboutView },
  { path: '/dashboard', component: DashboardView, meta: { requiresAuth: true }},
  { path: '/tasks/my/', component: MyTasksView, meta: { requiresAuth: true }},
  { path: '/task/edit/:id', component: TaskEdit, props: true, meta: { requiresAuth: true } },
  { path: '/task/create', component: CreateTaskView, meta: { requiresAuth: true }},
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach(async (to, _, next) => {
  const authStore = useAuthStore();

  if (to.meta.requiresAuth) {
      if (await authStore.isAuthenticated()) {
        next();
      } else {
        next('/login');
      }
  }
  else {
    if (to.meta.onlyGuest) {
      if (await authStore.isAuthenticated()) {
        next('/dashboard');
      } else {
        next();
      }
    } else {
      next();
    }
  }
});

export default router;
