import axios from 'axios';
import { API_URLS } from './apiUrls';
import type {NewTask, Task} from '@/types/task';
import type { ApiResponse } from '@/types/common';

export const getTasksData = async (
  title?: string|null,
  performerId?: bigint|null,
  creatorId?: bigint|null
):
  Promise<Task[]> => {

  const params: Record<string, any> = {};

  if (title?.length) params.title = title;
  if (performerId) params.performer_id = performerId;
  if (creatorId) params.creator_id = creatorId;

  const response =
    await axios.get<ApiResponse<Task[]>>(API_URLS.TASK,
    { params, withCredentials: true});
  return response.data.data;
};

export const getTaskData = async (id: bigint): Promise<Task> => {
  const response =
    await axios.get<ApiResponse<Task>>(
      `${API_URLS.TASK}/${id}`,
      { withCredentials: true }
    );
  return response.data.data;
};

export const createTask = async (data: NewTask): Promise<Task> => {
  const response = await axios.post(`${API_URLS.TASK}`, data, { withCredentials: true });
  return response.data.data;
};

export const deleteTask = async (id: bigint):Promise<any> => {
  return await axios.get<ApiResponse<any>>(`${API_URLS.TASK}/${id}`, { withCredentials: true });
};

export const updateTask = async (id: bigint, data: NewTask): Promise<Task> => {
  const response = await axios.put(`${API_URLS.TASK}/${id}`, data, { withCredentials: true });
  return response.data.data;
};
