package server

import (
	"TaskStorageService/initializers"
	"TaskStorageService/models"
	"TaskStorageService/proto/taskpb"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"time"
)

type TaskServer struct {
	taskpb.UnimplementedTaskServiceServer
	DB *gorm.DB
}

// CreateTask создает новую задачу и добавляет в Redis
func (s *TaskServer) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.TaskResponse, error) {
	task := models.Task{
		Title:       req.Title,
		Description: req.Description,
		PerformerId: uint(req.PerformerId),
		CreatorId:   uint(req.CreatorId),
		Observers:   models.ObserversFromIDs(req.ObserverIds),
		Status:      req.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.DB.Create(&task).Error; err != nil {
		return nil, err
	}

	// Сохраняем в Redis
	redisKey := fmt.Sprintf("task:%d", task.ID)
	taskJSON, _ := json.Marshal(task)
	initializers.RedisClient.Set(ctx, redisKey, taskJSON, 10*time.Minute)

	return &taskpb.TaskResponse{Task: convertToProto(task)}, nil
}

// GetTask получает задачу сначала из Redis, затем из PostgreSQL
func (s *TaskServer) GetTask(ctx context.Context, req *taskpb.GetTaskRequest) (*taskpb.TaskResponse, error) {
	redisKey := fmt.Sprintf("task:%d", req.Id)

	// Проверяем в Redis
	taskJSON, err := initializers.RedisClient.Get(ctx, redisKey).Result()
	if err == nil {
		var task models.Task
		json.Unmarshal([]byte(taskJSON), &task)
		return &taskpb.TaskResponse{Task: convertToProto(task)}, nil
	}

	// Если нет в Redis, загружаем из PostgreSQL
	var task models.Task
	if err := s.DB.First(&task, req.Id).Error; err != nil {
		return nil, err
	}

	taskJSONBytes, _ := json.Marshal(task)
	taskJSON = string(taskJSONBytes)
	initializers.RedisClient.Set(ctx, redisKey, taskJSON, 10*time.Minute)

	return &taskpb.TaskResponse{Task: convertToProto(task)}, nil
}

// GetTasks возвращает список всех задач (без кэширования)
func (s *TaskServer) GetTasks(ctx context.Context, req *taskpb.GetTasksRequest) (*taskpb.GetTasksResponse, error) {
	var tasks []models.Task
	s.DB.Find(&tasks)

	var protoTasks []*taskpb.Task
	for _, t := range tasks {
		protoTasks = append(protoTasks, convertToProto(t))
	}

	return &taskpb.GetTasksResponse{Tasks: protoTasks}, nil
}

// UpdateTask обновляет задачу и сбрасывает кэш Redis
func (s *TaskServer) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.TaskResponse, error) {
	var task models.Task
	if err := s.DB.First(&task, req.Id).Error; err != nil {
		return nil, err
	}

	task.Title = req.Title
	task.Description = req.Description
	task.Status = req.Status
	task.UpdatedAt = time.Now()

	s.DB.Save(&task)

	// Удаляем из Redis, чтобы при следующем запросе получить актуальные данные
	redisKey := fmt.Sprintf("task:%d", task.ID)
	initializers.RedisClient.Del(ctx, redisKey)

	return &taskpb.TaskResponse{Task: convertToProto(task)}, nil
}

// DeleteTask удаляет задачу и очищает Redis
func (s *TaskServer) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {
	if err := s.DB.Delete(&models.Task{}, req.Id).Error; err != nil {
		return nil, err
	}

	// Удаляем кэш Redis
	redisKey := fmt.Sprintf("task:%d", req.Id)
	initializers.RedisClient.Del(ctx, redisKey)

	return &taskpb.DeleteTaskResponse{Message: "Task deleted"}, nil
}

// Конвертация модели в gRPC-структуру
func convertToProto(task models.Task) *taskpb.Task {
	return &taskpb.Task{
		Id:          uint64(task.ID),
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		PerformerId: uint64(task.PerformerId),
		CreatorId:   uint64(task.CreatorId),
		ObserverIds: task.ObserverIDs(),
		CreatedAt:   timestamppb.New(task.CreatedAt),
		UpdatedAt:   timestamppb.New(task.UpdatedAt),
	}
}
