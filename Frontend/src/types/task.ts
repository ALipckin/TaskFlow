export type TaskStatus = "pending" | "in_progress" | "completed";

export interface Task {
  id: number;
  title: string;
  description: string;
  status: TaskStatus;
  performer: User;
  creator: User;
  observers: User[];
  createdAt: {
    seconds: number;
    nanos: number;
  };
  updatedAt: {
    seconds: number;
    nanos: number;
  };
}

export interface NewTask {
  title: string;
  description: string;
  status: TaskStatus;
  performer: bigint;
  creator: bigint;
  observers: bigint[];
}

export interface User {
  id: bigint;
  email: string;
  name: string;
}

export interface TaskResponse {
  data: Task;
  status: string;
}

export interface TasksResponse {
  data: Task[];
  status: string;
}
