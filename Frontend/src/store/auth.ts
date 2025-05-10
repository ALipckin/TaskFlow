import { defineStore } from 'pinia';
import axios from '@/api/axiosInstance';
import { API_URLS } from "@/api/apiUrls.ts";
import type { AxiosError } from 'axios';

interface User {
  id: bigint;
  name: string;
  email: string;
}

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null as User | null
  }),
  actions: {
    async login(credentials: { email: string; password: string }) {
      try {
        const response = await axios.post(API_URLS.LOGIN, credentials);
        const { token, ...userData } = response.data;
        this.user = userData;
        localStorage.setItem('token', token);
      } catch (error) {
        console.error("Login error", error);
        throw new Error('Auth error');
      }
    },
    async fetchUser() {
      try {
        if (localStorage.getItem('token')) {
          const response = await axios.get<User & { message?: string }>(API_URLS.VALIDATE);
          if (response.status !== 200) throw new Error(`Error: server response ${response.status}`);
          const {message, ...userData} = response.data;
          this.user = userData;
        }
        else{
          throw new Error(`Error: no token`)
        }
      } catch (error: unknown) {
        this.user = null;

        const axiosError = error as AxiosError;
        if (axiosError?.isAxiosError) {
          console.error("Validation error", axiosError.response?.data || axiosError.message);
        } else {
          console.error("Unknown error", error);
        }

        throw error;
      }
    },

    async register(credentials: { email: string; password: string, name: string }) {
      try {
        await axios.post(API_URLS.REGISTER, credentials);
      } catch (error) {
        console.error("Register error", error);
        throw new Error('Register error');
      }
    },

    async isAuthenticated() {
      try {
        await this.fetchUser();
        return !!this.user;
      } catch {
        return false;
      }
    },

    async logout() {
      localStorage.removeItem('token');
      this.user = null;
    }
  }
});
