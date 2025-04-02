<template>
  <header>
    <div class="wrapper">
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
        <RouterLink v-if="isAuth" to="/dashboard">Главная</RouterLink>
        <RouterLink v-if="!isAuth" to="/login">Логин</RouterLink>
        <RouterLink v-if="!isAuth" to="/register">Регистрация</RouterLink>
        <div class="user-menu" v-if="isAuth">
          <font-awesome-icon :icon="['fas', 'user']" @click="toggleMenu" />
          <div class="dropdown" v-if="isMenuOpen">
            <button @click="logout">Выйти</button>
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
import { ref, onMounted } from 'vue';
import { useAuthStore } from '@/store/auth.ts';

const authStore = useAuthStore();
const isAuth = ref(false);
const isMenuOpen = ref(false);

onMounted(async () => {
  isAuth.value = await authStore.isAuthenticated();
});

const logout = async () => {
  try {
    await authStore.logout();
    isAuth.value = false;
    // Перенаправление на главную страницу после выхода
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
