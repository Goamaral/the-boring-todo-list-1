// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "example.com/fiber-m3o-validator/entity"
	mock "github.com/stretchr/testify/mock"
)

// TaskService is an autogenerated mock type for the TaskService type
type TaskService struct {
	mock.Mock
}

// CreateTask provides a mock function with given fields: task
func (_m *TaskService) CreateTask(task entity.Task) (entity.Task, error) {
	ret := _m.Called(task)

	var r0 entity.Task
	if rf, ok := ret.Get(0).(func(entity.Task) entity.Task); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Get(0).(entity.Task)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Task) error); ok {
		r1 = rf(task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTasks provides a mock function with given fields: pageId, pageSize
func (_m *TaskService) ListTasks(pageId string, pageSize uint) ([]entity.Task, error) {
	ret := _m.Called(pageId, pageSize)

	var r0 []entity.Task
	if rf, ok := ret.Get(0).(func(string, uint) []entity.Task); ok {
		r0 = rf(pageId, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, uint) error); ok {
		r1 = rf(pageId, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTaskService interface {
	mock.TestingT
	Cleanup(func())
}

// NewTaskService creates a new instance of TaskService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTaskService(t mockConstructorTestingTNewTaskService) *TaskService {
	mock := &TaskService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
