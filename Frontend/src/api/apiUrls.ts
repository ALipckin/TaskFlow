const API_PROTOCOL = import.meta.env.VITE_BACKEND_API_PROTOCOL;
const API_HOST = import.meta.env.VITE_BACKEND_API_HOST;
const API_URL = API_PROTOCOL + API_HOST;

export const API_URLS = {
  BASE_URL: API_URL,
  TASK: `${API_URL}/tasks`,
  TASK_NOTIFICATIONS: `ws://${API_HOST}/tasks/notifications`,
  LOGIN: `${API_URL}/auth/login`,
  REGISTER: `${API_URL}/auth/register`,
  VALIDATE: `${API_URL}/auth/validate`,
  LOGOUT: `${API_URL}/auth/logout`,
  USER: `${API_URL}/auth/users`,

} as const;

export type ApiUrls = typeof API_URLS;
