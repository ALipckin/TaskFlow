import axios from 'axios';
import { API_URLS } from './apiUrls';

export const getUserData = () => {
  return axios.get(API_URLS.VALIDATE);
};

export const login = (credentials: { email: string; password: string }) => {
  return axios.post(API_URLS.LOGIN, credentials);
};
