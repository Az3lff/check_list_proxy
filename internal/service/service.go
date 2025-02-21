package service

import (
	"context"
	"fmt"

	"github.com/Az3lff/check_list_proxy/internal/delivery/grpc"
	"github.com/Az3lff/check_list_proxy/internal/delivery/grpc/task"
	"github.com/Az3lff/check_list_proxy/internal/models"
)

type Service struct {
	taskClient *grpc.Client
}

func NewService(cli *grpc.Client) *Service {
	return &Service{
		taskClient: cli,
	}
}

func (s *Service) CreateTask(ctx context.Context, userID int64, reqData models.CreateTaskRequest) (*models.TaskIDResponse, error) {
	resp, err := s.taskClient.CreateTask(ctx, &task.CreateTaskRequest{
		UserID:      userID,
		Title:       reqData.Title,
		Description: reqData.Description,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to createtask: %w", err)
	}

	if resp == nil || resp.TaskID == 0 {
		return nil, fmt.Errorf("empty response from gRPC server")
	}

	return &models.TaskIDResponse{
		TaskID: resp.TaskID,
	}, nil
}

func (s *Service) GetList(ctx context.Context, userID int64) (*models.GetListResponse, error) {
	resp, err := s.taskClient.GetList(ctx, &task.GetListRequest{
		UserID: userID,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get task list: %w", err)
	}

	if resp == nil {
		return nil, fmt.Errorf("empty response from gRPC server")
	}

	var list models.GetListResponse
	if resp.Tasks != nil {
		list.List = make([]models.Task, 0, len(resp.Tasks))

		for i := range resp.Tasks {
			list.List = append(list.List, models.Task{
				TaskID:      resp.Tasks[i].TaskID,
				Title:       resp.Tasks[i].Title,
				Description: resp.Tasks[i].Description,
			})
		}
	}

	return &list, nil
}

func (s *Service) DeleteTask(ctx context.Context, userID int64, reqData models.DeleteTaskRequest) (*models.TaskIDResponse, error) {
	resp, err := s.taskClient.DeleteTask(ctx, &task.DeleteTaskRequest{
		UserID: userID,
		TaskID: reqData.TaskID,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to delete task: %w", err)
	}

	if resp == nil || resp.TaskID == 0 {
		return nil, fmt.Errorf("empty response from gRPC server")
	}

	return &models.TaskIDResponse{
		TaskID: resp.TaskID,
	}, nil
}

func (s *Service) DoneTask(ctx context.Context, userID int64, reqData models.DoneTaskRequest) (*models.TaskIDResponse, error) {
	resp, err := s.taskClient.DoneTask(ctx, &task.DoneTaskRequest{
		UserID: userID,
		TaskID: reqData.TaskID,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to done task: %w", err)
	}

	if resp == nil || resp.TaskID == 0 {
		return nil, fmt.Errorf("empty response from gRPC server")
	}

	return &models.TaskIDResponse{
		TaskID: resp.TaskID,
	}, nil
}
