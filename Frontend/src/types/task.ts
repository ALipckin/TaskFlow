export interface Task {
  id: number;
  title: string;
  description: string;
  status: "pending" | "in_progress" | "completed";
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
