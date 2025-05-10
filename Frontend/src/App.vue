<template>
  <Notification/>
  <header>
    <div class="wrapper">
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
        <RouterLink v-if="isAuth" to="/dashboard">Задачи</RouterLink>
        <RouterLink v-if="isAuth" to="/tasks/my">Мои задачи</RouterLink>
        <RouterLink v-if="!isAuth" to="/login">Логин</RouterLink>
        <RouterLink v-if="!isAuth" to="/register">Регистрация</RouterLink>
        <div class="user-menu" v-if="isAuth">
          <font-awesome-icon :icon="['fas', 'user']" @click="toggleMenu" />
          <div class="dropdown" v-if="isMenuOpen">
            <div class="name" v-if="authStore.user">{{authStore.user.name}}</div>
            <v-col cols="auto">
              <v-btn density="comfortable" variant="elevated" @click="logout">
                Выйти
              </v-btn>
            </v-col>
          </div>
        </div>
      </nav>
    </div>
  </header>
  <main>
    <RouterView />
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import { useAuthStore } from '@/store/auth.ts';
const authStore = useAuthStore();
const isAuth = ref(false);
const isMenuOpen = ref(false);
import {API_URLS} from "@/api/apiUrls.ts";
import { initSocket, onMessage } from '@/services/socket'
import Notification from "@/components/Notification.vue";

onMounted(async () => {
  isAuth.value = await authStore.isAuthenticated();
  initSocket(API_URLS.TASK_NOTIFICATIONS)

  onMessage((data) => {
    if (data.type === 'alert') {
      notify(data.message, 'warning')
    } else {
      notify(data.message, 'info')
    }
  })
});

watch(
  () => authStore.user,
  (newUser) => {
    isAuth.value = !!newUser;
  }
);

const logout = async () => {
  try {
    await authStore.logout();
    isAuth.value = false;
    window.location.href = '/';
  } catch (error) {
    console.error("Ошибка при выходе", error);
  }
};

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value;
};
</script>

<style scoped>
.name{
    font-weight: bold;
}

header {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  background: white;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  z-index: 1000;
  padding: 1rem 0;
}

nav {
  width: 100%;
  text-align: center;
}

nav a {
  display: inline-block;
  padding: 0.5rem 1rem;
  border-left: 1px solid var(--color-border);
  text-decoration: none;
  color: var(--color-text);
}

nav a:first-of-type {
  border: none;
}

nav a.router-link-exact-active {
  font-weight: bold;
  color: var(--color-primary);
}

main {
  margin-top: 4rem;
  padding: 1rem;
}

.user-menu {
  position: relative;
  display: inline-block;
}

.dropdown {
  position: absolute;
  right: 0;
  background-color: white;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  padding: 0.5rem;
  border-radius: 4px;
  z-index: 1001;
}

.dropdown button {
  display: block;
  width: 100%;
  padding: 0.5rem;
  text-align: left;
  background: none;
  border: none;
  color: var(--color-text);
  cursor: pointer;
}

.dropdown button:hover {
  background-color: var(--color-border);
}

:root {
  --color-primary: #007bff;
  --color-text: #333;
  --color-border: #ddd;
}
</style>
