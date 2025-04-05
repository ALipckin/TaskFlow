export type TaskStatus = "pending" | "in_progress" | "completed";

export interface Task {
  id: number;
  title: string;
  description: string;
  status: TaskStatus;
  performer_id: bigint;
  creator_id: bigint;
  observer_ids: bigint[];
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
  performer_id: bigint;
  creator_id: bigint;
  observer_ids: bigint[];
}

export interface TaskResponse {
  data: Task;
  status: string;
}

export interface TasksResponse {
  data: Task[];
  status: string;
}
