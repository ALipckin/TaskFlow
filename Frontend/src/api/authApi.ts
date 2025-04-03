import axios from 'axios';
import { API_URLS } from './apiUrls';

export const getUserData = () => {
  return axios.get(API_URLS.VALIDATE);
};

export const getUsers = async ({ ids, search, page, limit } = {}) => {
  const params = {};

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
