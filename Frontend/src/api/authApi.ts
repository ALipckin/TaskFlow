import axios from '@/api/axiosInstance.ts';
import { API_URLS } from './apiUrls';

interface GetUsersParams {
  ids?: string[];
  search?: string;
  page?: number;
  limit?: number;
}

export interface SelectUser {
   id: bigint;
   name: string;
   email: string;
   displayName: string
}


export const getUserData = () => {
  return axios.get(API_URLS.VALIDATE);
};

export const getUsers = async ({ ids, search, page, limit }: GetUsersParams = {}) => {
  const params: Record<string, any> = {};

  if (ids?.length) params.ids = ids.join(',');
  if (search) params.search = search;
  if (page) params.page = page;
  if (limit) params.limit = limit;

  return await axios.get(`${API_URLS.USER}`, {
    params,
    withCredentials: true,
  });
};

export const login = (credentials: { email: string; password: string }) => {
  return axios.post(API_URLS.LOGIN, credentials);
};
