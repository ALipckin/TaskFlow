import { defineStore } from 'pinia';
import axios from 'axios';
import {API_URLS} from "@/api/apiUrls.ts";

interface User {
  id: bigint;
  email: string;
}

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null as User | null
  }),
  actions: {
    async login(credentials: { email: string; password: string }) {
      try {
        await axios.post(API_URLS.LOGIN, credentials, { withCredentials: true });
        await this.fetchUser();
      } catch (error) {
        throw new Error('Auth error');
      }
    },
    async fetchUser() {
      try {
        const response = await axios.get<User & { message?: string }>(API_URLS.VALIDATE, {
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
    },
    async register(credentials: { email: string; password: string }) {
      try {
        await axios.post(API_URLS.REGISTER, credentials, { withCredentials: true });
      } catch (error) {
        throw new Error('Register error');
      }
    },
    async isAuthenticated(){
      try {
        await this.fetchUser();
        return !!this.user;
      } catch {
        return false;
      }
    },
    async logout() {
      try {
        await axios.post(API_URLS.LOGOUT, {}, {withCredentials: true});
        this.user = null;
      } catch (error) {
        console.error("Logout error", error);
        throw error;
      }
    }
  }
});
