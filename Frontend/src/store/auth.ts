import { defineStore } from 'pinia';
import axios from 'axios';

interface User {
  id: number;
  email: string;
  name: string;
}

// Получаем API хост из .env
const API_HOST = import.meta.env.BACKEND_API_HOST || 'http://localhost:5437';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null as User | null
  }),
  actions: {
    async login(credentials: { email: string; password: string }) {
      try {
        const { data } = await axios.post<{ user: User }>(
          `${API_HOST}/auth/login`,
          credentials,
          { withCredentials: true } // Куки передаются автоматически
        );
        this.user = data.user;
      } catch (error) {
        throw new Error('Ошибка авторизации');
      }
    },
    async fetchUser() {
      try {
        const { data } = await axios.get<{ user: User }>(`${API_HOST}/auth/me`, {
          withCredentials: true
        });
        this.user = data.user;
      } catch (error) {
        this.user = null;
      }
    },
    async logout() {
      try {
        await axios.post(`${API_HOST}/auth/logout`, {}, { withCredentials: true });
      } catch (error) {
        console.warn('Ошибка при выходе');
      }
      this.user = null;
    }
  }
});
