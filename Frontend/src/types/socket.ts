export interface SocketMessage {
  type: string
  message: string
  [key: string]: any
}

export interface TaskNotifyMessage {
  event: string;
  title?: string;
  description?: string;
  user_id?: number;
  [key: string]: any;
}
