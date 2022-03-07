// Code generated by protoc-gen-go-asynq. DO NOT EDIT.
// versions:
// protoc-gen-go-asynq v1.0.0

package example

import (
	context "context"
	asynq "github.com/hibiken/asynq"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the asynq package it is being compiled against.
var _ = new(context.Context)
var _ = new(asynq.Task)
var _ = new(emptypb.Empty)
var _ = new(proto.Message)

type UserJobServer interface {
	CreateUser(context.Context, *CreateUserPayload) error
	UpdateUser(context.Context, *UpdateUserPayload) error
}

func RegisterUserJobServer(mux *asynq.ServeMux, srv UserJobServer) {
	mux.HandleFunc("user:create", _User_CreateUser_Job_Handler(srv))
	mux.HandleFunc("user:update", _User_UpdateUser_Job_Handler(srv))
}

func _User_CreateUser_Job_Handler(srv UserJobServer) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, task *asynq.Task) error {
		var in CreateUserPayload
		if err := proto.Unmarshal(task.Payload(), &in); err != nil {
			return err
		}
		err := srv.CreateUser(ctx, &in)
		return err
	}
}

func _User_UpdateUser_Job_Handler(srv UserJobServer) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, task *asynq.Task) error {
		var in UpdateUserPayload
		if err := proto.Unmarshal(task.Payload(), &in); err != nil {
			return err
		}
		err := srv.UpdateUser(ctx, &in)
		return err
	}
}

type UserSvcJob struct{}

var UserJob UserSvcJob

func (j *UserSvcJob) CreateUser(in *CreateUserPayload, opts ...asynq.Option) (*asynq.Task, error) {
	payload, err := proto.Marshal(in)
	if err != nil {
		return nil, err
	}
	task := asynq.NewTask("user:create", payload, opts...)
	return task, nil
}

func (j *UserSvcJob) UpdateUser(in *UpdateUserPayload, opts ...asynq.Option) (*asynq.Task, error) {
	payload, err := proto.Marshal(in)
	if err != nil {
		return nil, err
	}
	task := asynq.NewTask("user:update", payload, opts...)
	return task, nil
}

type UserJobClient interface {
	CreateUser(ctx context.Context, req *CreateUserPayload, opts ...asynq.Option) (info *asynq.TaskInfo, err error)
	UpdateUser(ctx context.Context, req *UpdateUserPayload, opts ...asynq.Option) (info *asynq.TaskInfo, err error)
}

type UserJobClientImpl struct {
	cc *asynq.Client
}

func NewUserJobClient(client *asynq.Client) UserJobClient {
	return &UserJobClientImpl{client}
}

func (c *UserJobClientImpl) CreateUser(ctx context.Context, in *CreateUserPayload, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	task, err := UserJob.CreateUser(in, opts...)
	if err != nil {
		return nil, err
	}
	info, err := c.cc.Enqueue(task)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (c *UserJobClientImpl) UpdateUser(ctx context.Context, in *UpdateUserPayload, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	task, err := UserJob.UpdateUser(in, opts...)
	if err != nil {
		return nil, err
	}
	info, err := c.cc.Enqueue(task)
	if err != nil {
		return nil, err
	}
	return info, nil
}
