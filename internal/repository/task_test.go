package repository_test

import (
	"context"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"example.com/the-boring-to-do-list-1/internal/entity"
	"example.com/the-boring-to-do-list-1/internal/repository"
	gormprovider "example.com/the-boring-to-do-list-1/pkg/provider/gorm"
)

func NewTaskRepository(t *testing.T) repository.TaskRepository {
	_, b, _, _ := runtime.Caller(0)
	folderPath := filepath.Dir(b)
	err := godotenv.Load(folderPath + "/../../.env")
	if err != nil {
		t.Error(err)
	}

	gormProvider := gormprovider.NewTestProvider(t, folderPath+"/../../db/1_schema.sql")
	return repository.NewTaskRepository(gormProvider)
}

func TestTaskRepository_CreateTask(t *testing.T) {
	repo := NewTaskRepository(t)
	title := "Test Task"

	task := entity.Task{Title: title}
	err := repo.CreateTask(context.Background(), &task)
	if assert.NoError(t, err) {
		assert.NotZero(t, task.Id)
		assert.NotZero(t, task.CreatedAt)
	}
}

func TestTaskRepository_ListTasks(t *testing.T) {
	// var pageSize uint = 10

	// type Test struct {
	// 	TestName               string
	// 	Ops                    *service.ListTasksOpts
	// 	ValidateM3OReadRequest func(t *testing.T, test Test, req *db.ReadRequest)
	// }

	// tests := []Test{
	// 	{
	// 		TestName: "Without options",
	// 		ValidateM3OReadRequest: func(t *testing.T, test Test, req *db.ReadRequest) {
	// 			assert.Equal(t, entity.TasksTableName, req.Table, "Table")
	// 			assert.Equal(t, int32(pageSize), req.Limit, "Limit")
	// 			assert.Empty(t, req.OrderBy, "OrderBy")
	// 			assert.Empty(t, req.Query, "Query")
	// 		},
	// 	},
	// 	{
	// 		TestName: "With PageId option",
	// 		Ops:      &service.ListTasksOpts{PageId: ulid.Make().String()},
	// 		ValidateM3OReadRequest: func(t *testing.T, test Test, req *db.ReadRequest) {
	// 			assert.Equal(t, entity.TasksTableName, req.Table, "Table")
	// 			assert.Equal(t, int32(pageSize), req.Limit, "Limit")
	// 			assert.Equal(t, "id", req.OrderBy, "OrderBy")
	// 			assert.Contains(t, req.Query, fmt.Sprintf("'%s'", test.Ops.PageId), "Query")
	// 		},
	// 	},
	// }

	// for _, test := range tests {
	// 	expectedId := ulid.Make().String()

	// 	mockM3ODbClient := mocks.NewM3ODb(t)
	// 	mockM3ODbClient.On("Read", mock.Anything).
	// 		Return(func(req *db.ReadRequest) *db.ReadResponse {
	// 			test.ValidateM3OReadRequest(t, test, req)

	// 			return &db.ReadResponse{Records: []map[string]interface{}{
	// 				entity.TaskToMap(entity.Task{AbstractEntity: entity.AbstractEntity{Id: expectedId}}),
	// 			}}
	// 		}, nil)

	// 	s := service.NewTaskRepository(mockM3ODbClient)
	// 	tasks, err := s.ListTasks(pageSize, test.Ops)
	// 	if assert.NoError(t, err) && assert.Len(t, tasks, 1) {
	// 		assert.Equal(t, expectedId, tasks[0].Id)
	// 	}
	// }
}
