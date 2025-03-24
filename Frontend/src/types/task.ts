export interface Task {
  id: number;
  title: string;
  description: string;
  status: "pending" | "in_progress" | "completed";
  performerId: number;
  creatorId: number;
  observerIds: number[];
  createdAt: {
    seconds: number;
    nanos: number;
  };
  updatedAt: {
    seconds: number;
    nanos: number;
  };
}

export interface TaskResponse {
  data: Task;
  status: string;
}

export interface TasksResponse {
  data: Task[];
  status: string;
}
