import { defineStore } from 'pinia';
import axios from 'axios';

interface User {
  id: number;
  email: string;
}

const API_HOST = import.meta.env.VITE_BACKEND_API_HOST;

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null as User | null
  }),
  actions: {
    async login(credentials: { email: string; password: string }) {
      try {
        await axios.post(`${API_HOST}/auth/login`, credentials, { withCredentials: true });
        await this.fetchUser();
      } catch (error) {
        throw new Error('Auth error');
      }
    },
    async fetchUser() {
      try {
        const response = await axios.get<User & { message?: string }>(`${API_HOST}/auth/validate`, {
          withCredentials: true
        });
        if (response.status !== 200) {
          throw new Error(`Error: server response ${response.status}`);
        }
        const { message, ...userData } = response.data;
        this.user = userData;
      } catch (error) {
        this.user = null;
        if (axios.isAxiosError(error)) {
          console.error("Error validate", error.response?.data || error.message);
        } else {
          console.error("Unknown error", error);
        }
        throw error;
      }
    }
  }
});
