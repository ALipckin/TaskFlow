<template>
  <AuthCard title="Вход">
    <template v-slot:actions>
      <v-form @submit.prevent="handleLogin">
        <v-text-field
          v-model="email"
          label="Email"
          type="email"
          required
        />
        <v-text-field
          v-model="password"
          label="Пароль"
          type="password"
          required
        />
        <v-btn
          type="submit"
          :loading="loading"
          :disabled="loading"
          block
          color="primary"
          class="mt-4"
        >
          Войти
        </v-btn>
        <v-alert
          v-if="errorMessage"
          type="error"
          class="mt-4"
          dense
        >
          {{ errorMessage }}
        </v-alert>
      </v-form>
    </template>
  </AuthCard>
</template>

<script setup lang="ts">
import {ref} from 'vue'
import {useAuthStore} from '../store/auth'
import {useRouter} from 'vue-router'
import AuthCard from "@/components/cards/AuthCard.vue";

const email = ref('')
const password = ref('')
const errorMessage = ref('')
const loading = ref(false)
const authStore = useAuthStore()
const router = useRouter()

const handleLogin = async () => {
  loading.value = true
  errorMessage.value = ''

  try {
    await authStore.login({email: email.value, password: password.value})
    await router.push('/dashboard')
  } catch (error) {
    errorMessage.value = 'Неверные учетные данные'
  } finally {
    loading.value = false
  }
}
</script>
