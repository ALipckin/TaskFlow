const API_HOST = import.meta.env.VITE_BACKEND_API_HOST;

export const API_URLS = {
  BASE_URL: API_HOST,
  TASK: `${API_HOST}/tasks`,
  LOGIN: `${API_HOST}/auth/login`,
  REGISTER: `${API_HOST}/auth/register`,
  VALIDATE: `${API_HOST}/auth/validate`,
  USER: `${API_HOST}/auth/users`,

} as const;

export type ApiUrls = typeof API_URLS;
