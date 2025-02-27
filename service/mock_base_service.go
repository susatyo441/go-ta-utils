package service

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/susatyo441/go-ta-utils/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockBaseService[T any] struct {
	mock.Mock
}

//
// Reads
//

func (m *MockBaseService[T]) FindOne(
	ctx context.Context,
	filter interface{},
	opts ...*options.FindOneOptions,
) (*T, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*T), args.Error(1)
}

func (m *MockBaseService[T]) Find(
	ctx context.Context,
	filter interface{},
	opts ...*options.FindOptions,
) ([]T, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).([]T), args.Error(1)
}

func (m *MockBaseService[T]) Aggregate(
	v any,
	ctx context.Context,
	pipeline mongo.Pipeline,
	opts ...*options.AggregateOptions,
) error {
	args := m.Called(v, ctx, pipeline, opts)
	return args.Error(0)
}

func (m *MockBaseService[T]) GetOneOrFail(
	ctx context.Context,
	filter interface{},
	opts ...*GetOneOrFailOptions,
) (*T, *entity.HttpError) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*T), args.Get(1).(*entity.HttpError)
}

func (m *MockBaseService[T]) FindOrFail(
	ctx context.Context,
	filter interface{},
	expectedLength int,
	opts ...*FindOrFailOptions,
) ([]T, *entity.HttpError) {
	args := m.Called(ctx, filter, expectedLength, opts)
	return args.Get(0).([]T), args.Get(1).(*entity.HttpError)
}

//
// Creates
//

func (m *MockBaseService[T]) Create(
	ctx context.Context,
	createData T,
	opts ...*options.InsertOneOptions,
) (*T, error) {
	args := m.Called(ctx, createData, opts)
	return args.Get(0).(*T), args.Error(1)
}

func (m *MockBaseService[T]) InsertOne(
	ctx context.Context,
	data T,
	opts ...*options.InsertOneOptions,
) (*primitive.ObjectID, error) {
	args := m.Called(ctx, data, opts)
	return args.Get(0).(*primitive.ObjectID), args.Error(1)
}

func (m *MockBaseService[T]) InsertMany(
	ctx context.Context,
	data []T,
	opts ...*options.InsertManyOptions,
) ([]interface{}, error) {
	args := m.Called(ctx, data, opts)
	return args.Get(0).([]interface{}), args.Error(1)
}

//
// Updates
//

func (m *MockBaseService[T]) UpdateOne(
	ctx context.Context,
	filter interface{},
	updateData bson.M,
	opts ...*options.UpdateOptions,
) (int, error) {
	args := m.Called(ctx, filter, updateData, opts)
	return args.Get(0).(int), args.Error(1)
}

// REVIEW: Resolve sonarlint identical code issue
func (m *MockBaseService[T]) UpdateMany(
	ctx context.Context,
	filter interface{},
	updateData bson.M,
	opts ...*options.UpdateOptions,
) (int, error) {
	args := m.Called(ctx, filter, updateData, opts)
	return args.Get(0).(int), args.Error(1)
}

func (m *MockBaseService[T]) UpdateManyOld(
	ctx context.Context,
	filter interface{},
	data interface{},
	opts ...*options.UpdateOptions,
) (int, error) {
	args := m.Called(ctx, filter, data, opts)
	return args.Get(0).(int), args.Error(1)
}

func (m *MockBaseService[T]) FindOneAndUpdate(
	ctx context.Context,
	filter interface{},
	updateData bson.M,
	opts ...*options.FindOneAndUpdateOptions,
) (*T, error) {
	args := m.Called(ctx, filter, updateData, opts)
	return args.Get(0).(*T), args.Error(1)
}

//
// Deletes
//

func (m *MockBaseService[T]) DeleteOne(
	ctx context.Context,
	filter interface{},
	opts ...*options.DeleteOptions,
) (int, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(int), args.Error(1)
}

// REVIEW: Resolve sonarlint identical code issue
func (m *MockBaseService[T]) DeleteMany(
	ctx context.Context,
	filter interface{},
	opts ...*options.DeleteOptions,
) (int, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(int), args.Error(1)
}

func (m *MockBaseService[T]) FindOneAndDelete(
	ctx context.Context,
	filter interface{},
	opts ...*options.FindOneAndDeleteOptions,
) (*T, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*T), args.Error(1)
}

func (m *MockBaseService[T]) BulkWrite(
	ctx context.Context,
	models []mongo.WriteModel,
	opts ...*options.BulkWriteOptions,
) (*mongo.BulkWriteResult, error) {
	args := m.Called(ctx, models, opts)
	return args.Get(0).(*mongo.BulkWriteResult), args.Error(1)
}

//
// Utilities
//

func (m *MockBaseService[T]) CountDocuments(
	ctx context.Context,
	filter interface{},
) (int, error) {
	args := m.Mock.Called(ctx, filter)

	if args.Get(0) == nil {
		return 0, nil
	}

	return args.Get(0).(int), args.Error(1)
}

func (m *MockBaseService[T]) MakeUnique(
	ctx context.Context,
	filter interface{},
) error {
	args := m.Called(ctx, filter)
	return args.Error(0)
}

func (m *MockBaseService[T]) SetDeleteFromDatabaseAttribute(
	ctx context.Context,
	filter interface{},
) error {
	args := m.Called(ctx, filter)
	return args.Error(0)
}

func (m *MockBaseService[T]) CreateIndex(
	ctx context.Context,
	fields interface{},
	opts ...*options.CreateIndexesOptions,
) error {
	args := m.Called(ctx, fields, opts)
	return args.Error(0)
}

//
// "On" Aliases
//

//
// Reads
//

// Alias for Mock.On("FindOne", mock.Anything, ...)
func (m *MockBaseService[T]) OnFindOne() *mock.Call {
	return m.Mock.On(
		"FindOne",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	)
}

// Alias for Mock.On("Find", mock.Anything, ...)
func (m *MockBaseService[T]) OnFind() *mock.Call {
	return m.Mock.On(
		"Find",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	)
}

// Alias for Mock.On("Aggregate", mock.Anything, ...)
func (m *MockBaseService[T]) OnAggregate() *mock.Call {
	return m.Mock.On(
		"Aggregate",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	)
}

// Alias for Mock.On("GetOneOrFail", mock.Anything, ...)
func (m *MockBaseService[T]) OnGetOneOrFail() *mock.Call {
	return m.Mock.On(
		"GetOneOrFail",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	)
}

// Alias for Mock.On("FindOrFail", mock.Anything, ...)
func (m *MockBaseService[T]) OnFindOrFail() *mock.Call {
	return m.Mock.On(
		"FindOrFail",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	)
}

//
// Creates
//

// Alias for Mock.On("Create", mock.Anything, ...)
func (m *MockBaseService[T]) OnCreate() *mock.Call {
	return m.Mock.On(
		"Create",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	)
}

// Alias for Mock.On("InsertOne", mock.Anything, ...)
func (m *MockBaseService[T]) OnInsertOne() *mock.Call {
	return m.Mock.On(
		"InsertOne",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	)
}

// Alias for Mock.On("InsertMany", mock.Anything, ...)
func (m *MockBaseService[T]) OnInsertMany() *mock.Call {
	return m.Mock.On(
		"InsertMany",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	)
}

//
// Updates
//

// Alias for Mock.On("UpdateOne", mock.Anything, ...)
func (m *MockBaseService[T]) OnUpdateOne() *mock.Call {
	return m.Mock.On(
		"UpdateOne",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	)
}

// Alias for Mock.On("UpdateMany", mock.Anything, ...)
func (m *MockBaseService[T]) OnUpdateMany() *mock.Call {
	return m.Mock.On(
		"UpdateMany",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	)
}

// Alias for Mock.On("UpdateManyOld", mock.Anything, ...)
func (m *MockBaseService[T]) OnUpdateManyOld() *mock.Call {
	return m.Mock.On(
		"UpdateManyOld",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	)
}

// Alias for Mock.On("FindOneAndUpdate", mock.Anything, ...)
func (m *MockBaseService[T]) OnFindOneAndUpdate() *mock.Call {
	return m.Mock.On(
		"FindOneAndUpdate",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	)
}

//
// Deletes
//

// Alias for Mock.On("DeleteOne", mock.Anything, ...)
func (m *MockBaseService[T]) OnDeleteOne() *mock.Call {
	return m.Mock.On(
		"DeleteOne",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	)
}

// Alias for Mock.On("DeleteMany", mock.Anything, ...)
func (m *MockBaseService[T]) OnDeleteMany() *mock.Call {
	return m.Mock.On(
		"DeleteMany",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	)
}

// Alias for Mock.On("FindOneAndDelete", mock.Anything, ...)
func (m *MockBaseService[T]) OnFindOneAndDelete() *mock.Call {
	return m.Mock.On(
		"FindOneAndDelete",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	)
}

// Alias for Mock.On("BulkWrite", mock.Anything, ...)
func (m *MockBaseService[T]) OnBulkWrite() *mock.Call {
	return m.Mock.On(
		"BulkWrite",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	)
}

//
// Utilities
//

// Alias for Mock.On("CountDocuments", mock.Anything, ...)
func (m *MockBaseService[T]) OnCountDocuments() *mock.Call {
	return m.Mock.On(
		"CountDocuments",
		mock.Anything,
		mock.Anything,
	)
}

// Alias for Mock.On("MakeUnique", mock.Anything, ...)
func (m *MockBaseService[T]) OnMakeUnique() *mock.Call {
	return m.Mock.On(
		"MakeUnique",
		mock.Anything,
		mock.Anything,
	)
}

// Alias for Mock.On("CreateIndex", mock.Anything, ...)
func (m *MockBaseService[T]) OnCreateIndex() *mock.Call {
	return m.Mock.On(
		"CreateIndex",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	)
}
