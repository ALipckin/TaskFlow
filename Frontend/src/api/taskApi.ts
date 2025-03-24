import axios from 'axios';
import { API_URLS } from './apiUrls';
import type { Task } from '@/types/task';
import type { ApiResponse } from '@/types/common';

export const getTasksData = async (): Promise<Task[]> => {
  const response = await axios.get<ApiResponse<Task[]>>(API_URLS.TASK, { withCredentials: true });
  return response.data.data;
};

export const getTaskData = async (id: bigint): Promise<Task> => {
  const response = await axios.post<ApiResponse<Task>>(`${API_URLS.TASK}/${id}`, { withCredentials: true });
  return response.data.data;
};
