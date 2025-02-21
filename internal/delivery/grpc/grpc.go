package grpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"

	"github.com/Az3lff/check_list_proxy/internal/config"
	"github.com/Az3lff/check_list_proxy/internal/delivery/grpc/task"
)

type Client struct {
	api task.TaskClient
}

func New(cfg config.GRPCClient) (*Client, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			grpcretry.UnaryClientInterceptor(
				grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
				grpcretry.WithMax(uint(cfg.RetryCount)),
				grpcretry.WithPerRetryTimeout(cfg.Timeout),
			),
		),
	}

	cliConn, err := grpc.NewClient(cfg.Address, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect service: %w", err)
	}

	return &Client{
		api: task.NewTaskClient(cliConn),
	}, nil
}

func (c *Client) CreateTask(ctx context.Context, req *task.CreateTaskRequest) (*task.CreateTaskResponse, error) {
	return c.api.CreateTask(ctx, req)
}

func (c *Client) GetList(ctx context.Context, req *task.GetListRequest) (*task.GetListResponse, error) {
	return c.api.GetList(ctx, req)
}

func (c *Client) DeleteTask(ctx context.Context, req *task.DeleteTaskRequest) (*task.DeleteTaskResponse, error) {
	return c.api.DeleteTask(ctx, req)
}

func (c *Client) DoneTask(ctx context.Context, req *task.DoneTaskRequest) (*task.DoneTaskResponse, error) {
	return c.api.DoneTask(ctx, req)
}
