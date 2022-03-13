// Code generated by protoc-gen-go-asynq. DO NOT EDIT.
// versions:
// protoc-gen-go-asynq v1.0.2

package example

import (
	context "context"
	asynqx "github.com/amzapi/protoc-gen-go-asynq/asynqx"
	asynq "github.com/hibiken/asynq"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	time "time"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the asynq package it is being compiled against.
var _ = new(time.Time)
var _ = new(context.Context)
var _ = new(asynq.Task)
var _ = new(emptypb.Empty)
var _ = new(proto.Message)
var _ = new(asynqx.Server)

const QueueName = "user"

type UserTaskServer interface {
	CreateUser(context.Context, *CreateUserPayload) error
	UpdateUser(context.Context, *UpdateUserPayload) error
}

func RegisterUserTaskServer(s *asynqx.Server, srv UserTaskServer) {
	s.HandleFunc("user:create", _User_CreateUser_Task_Handler(srv))
	s.HandleFunc("user:update", _User_UpdateUser_Task_Handler(srv))
}

func _User_CreateUser_Task_Handler(srv UserTaskServer) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, task *asynq.Task) error {
		var in CreateUserPayload
		if err := proto.Unmarshal(task.Payload(), &in); err != nil {
			return err
		}
		err := srv.CreateUser(ctx, &in)
		return err
	}
}

func _User_UpdateUser_Task_Handler(srv UserTaskServer) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, task *asynq.Task) error {
		var in UpdateUserPayload
		if err := proto.Unmarshal(task.Payload(), &in); err != nil {
			return err
		}
		err := srv.UpdateUser(ctx, &in)
		return err
	}
}

type UserSvcTask struct{}

var UserTask UserSvcTask

func (j *UserSvcTask) CreateUser(in *CreateUserPayload, opts ...asynq.Option) (*asynq.Task, error) {
	payload, err := proto.Marshal(in)
	if err != nil {
		return nil, err
	}
	opts = append(opts, asynq.Timeout(30*time.Second))
	opts = append(opts, asynq.MaxRetry(10))
	opts = append(opts, asynq.Timeout(60*time.Second))
	opts = append(opts, asynq.Timeout(3600*time.Second))
	task := asynq.NewTask("user:create", payload, opts...)
	return task, nil
}

func (j *UserSvcTask) UpdateUser(in *UpdateUserPayload, opts ...asynq.Option) (*asynq.Task, error) {
	payload, err := proto.Marshal(in)
	if err != nil {
		return nil, err
	}
	opts = append(opts, asynq.Timeout(60*time.Second))
	opts = append(opts, asynq.MaxRetry(10))
	opts = append(opts, asynq.Timeout(60*time.Second))
	opts = append(opts, asynq.Timeout(3600*time.Second))
	task := asynq.NewTask("user:update", payload, opts...)
	return task, nil
}

type UserTaskClient interface {
	CreateUser(ctx context.Context, req *CreateUserPayload, opts ...asynq.Option) (info *asynq.TaskInfo, err error)
	UpdateUser(ctx context.Context, req *UpdateUserPayload, opts ...asynq.Option) (info *asynq.TaskInfo, err error)
}

type UserTaskClientImpl struct {
	cc *asynq.Client
}

func NewUserTaskClient(client *asynq.Client) UserTaskClient {
	return &UserTaskClientImpl{client}
}

func (c *UserTaskClientImpl) CreateUser(ctx context.Context, in *CreateUserPayload, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	task, err := UserTask.CreateUser(in, opts...)
	if err != nil {
		return nil, err
	}
	info, err := c.cc.Enqueue(task)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (c *UserTaskClientImpl) UpdateUser(ctx context.Context, in *UpdateUserPayload, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	task, err := UserTask.UpdateUser(in, opts...)
	if err != nil {
		return nil, err
	}
	info, err := c.cc.Enqueue(task)
	if err != nil {
		return nil, err
	}
	return info, nil
}
