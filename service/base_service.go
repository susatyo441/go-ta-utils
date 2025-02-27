package service

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/susatyo441/go-ta-utils/db"
	"github.com/susatyo441/go-ta-utils/entity"
	"github.com/susatyo441/go-ta-utils/parser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Base Service Interface
type Service[T any] interface {
	//
	// Reads
	//

	// Returns the first document that matches the filter.
	FindOne(
		ctx context.Context,
		filter interface{},
		opts ...*options.FindOneOptions,
	) (*T, error)
	// Find all documents that match the filter.
	Find(
		ctx context.Context,
		filter interface{},
		opts ...*options.FindOptions,
	) ([]T, error)
	// Aggregate documents.
	Aggregate(
		v any,
		ctx context.Context,
		pipeline mongo.Pipeline,
		opts ...*options.AggregateOptions,
	) error
	// High-level method to get one document or return an error.
	// If the document is not found, it will return entity.NotFound with the message.
	//
	// The message will use the default message from the BaseServiceOptions
	// during service initialization if not provided.
	GetOneOrFail(
		ctx context.Context,
		filter interface{},
		opts ...*GetOneOrFailOptions,
	) (*T, *entity.HttpError)
	// High-level method to find documents or return an error if one or more documents are not found.
	// If one or more documents are not found, it will return entity.NotFound with the message.
	//
	// The message will use the default message from the BaseServiceOptions
	// during service initialization if not provided.
	FindOrFail(
		ctx context.Context,
		filter interface{},
		expectedLength int,
		opts ...*FindOrFailOptions,
	) ([]T, *entity.HttpError)

	//
	// Inserts
	//

	// Insert and return the insert result.
	Create(
		ctx context.Context,
		createData T,
		opts ...*options.InsertOneOptions,
	) (*T, error)
	// Insert and return the inserted id.
	InsertOne(
		ctx context.Context,
		data T,
		opts ...*options.InsertOneOptions,
	) (*primitive.ObjectID, error)
	// Insert many documents and return the inserted ids.
	InsertMany(
		ctx context.Context,
		data []T,
		opts ...*options.InsertManyOptions,
	) ([]interface{}, error)

	//
	// Updates
	//

	// Update one document.
	UpdateOne(
		ctx context.Context,
		filter interface{},
		updateData bson.M,
		opts ...*options.UpdateOptions,
	) (int, error)
	// Update many documents.
	UpdateMany(
		ctx context.Context,
		filter interface{},
		data bson.M,
		opts ...*options.UpdateOptions,
	) (int, error)
	// Update document that match the filter and return the updated document.
	FindOneAndUpdate(
		ctx context.Context,
		filter interface{},
		updateData bson.M,
		opts ...*options.FindOneAndUpdateOptions,
	) (*T, error)

	// Deletes
	// Delete one document and return the number of documents deleted.
	DeleteOne(
		ctx context.Context,
		filter interface{},
		opts ...*options.DeleteOptions,
	) (int, error)
	// Delete many documents and return the number of documents deleted.
	DeleteMany(
		ctx context.Context,
		filter interface{},
		opts ...*options.DeleteOptions,
	) (int, error)
	// Delete a document and return the document before deletion.
	FindOneAndDelete(
		ctx context.Context,
		filter interface{},
		opts ...*options.FindOneAndDeleteOptions,
	) (*T, error)

	//
	// Utilities
	//

	// Count the number of documents that match the filter.
	CountDocuments(
		ctx context.Context,
		filter interface{},
	) (int, error)
	// Create a unique index on the collection.
	MakeUnique(
		ctx context.Context,
		filter interface{},
	) error
	// Create an index to set data deletion time
	SetDeleteFromDatabaseAttribute(
		ctx context.Context,
		filter interface{},
	) error
	// Create an index on the collection.
	CreateIndex(
		ctx context.Context,
		fields interface{},
		opts ...*options.CreateIndexesOptions,
	) error
	//perform bulk write operation
	BulkWrite(
		ctx context.Context,
		models []mongo.WriteModel,
		opts ...*options.BulkWriteOptions,
	) (*mongo.BulkWriteResult, error)

	//alt for update many without automatically update updatedAt
	UpdateManyOld(ctx context.Context, filter interface{}, data interface{},
		opts ...*options.UpdateOptions) (int, error)
}

type BaseServiceOptions struct {
	GetOneOrFailMessage string
	FindOrFailMessage   string
}

var defaultBaseServiceOptions = BaseServiceOptions{
	GetOneOrFailMessage: "Data not found",
	FindOrFailMessage:   "One or more data not found",
}

func determineOptions(
	defaultOpts BaseServiceOptions,
	opts ...BaseServiceOptions,
) BaseServiceOptions {
	if len(opts) > 0 {
		return opts[0]
	}

	return defaultOpts
}

type BaseService[T any] struct {
	collection *mongo.Collection
	options    BaseServiceOptions
}

// Create a service for a specific company.
//
// NOTE: Use NewBaseService for non-company-specific services.
func NewCustomService[T any](
	dbName string,
	collection string,
	opts ...BaseServiceOptions,
) *BaseService[T] {
	db.ConnectToCustomDb(dbName)
	coll := db.GetCollection(collection)

	return &BaseService[T]{
		collection: coll,
		options:    determineOptions(defaultBaseServiceOptions, opts...),
	}
}

// Create a service for a specific company.
//
// NOTE: Use NewBaseService for non-company-specific services.
func NewCompanyService[T any](
	companyCode string,
	collection string,
	opts ...BaseServiceOptions,
) *BaseService[T] {

	db.ConnectToCompanyDb(companyCode)
	coll := db.GetCollection(collection)

	return &BaseService[T]{
		collection: coll,
		options:    determineOptions(defaultBaseServiceOptions, opts...),
	}
}

// Partner in Admin console will have separate console & db and will be using this
func NewPartnerService[T any](
	companyCode string,
	collection string,
	opts ...BaseServiceOptions,
) *BaseService[T] {
	db.ConnectToPartnerDb(companyCode)
	coll := db.GetCollection(collection)

	return &BaseService[T]{
		collection: coll,
		options:    determineOptions(defaultBaseServiceOptions, opts...),
	}
}

func NewAdminService[T any](collection string, opts ...BaseServiceOptions) *BaseService[T] {
	db.ConnectToAdminDb()
	coll := db.GetCollection(collection)

	return &BaseService[T]{
		collection: coll,
		options:    determineOptions(defaultBaseServiceOptions, opts...),
	}
}

func ShopVisionService[T any](collection string, opts ...BaseServiceOptions) *BaseService[T] {
	db.ConnectToAdminDb()
	coll := db.GetCollection(collection)

	return &BaseService[T]{
		collection: coll,
		options:    determineOptions(defaultBaseServiceOptions, opts...),
	}
}

func NewGlobalService[T any](collection string, opts ...BaseServiceOptions) *BaseService[T] {
	db.ConnectToGlobalDb()
	coll := db.GetCollection(collection)

	return &BaseService[T]{
		collection: coll,
		options:    determineOptions(defaultBaseServiceOptions, opts...),
	}
}

// Use this when you want to create a service for a specific collection.
//
// NOTE: Use NewCompanyService for company-specific services.
func NewBaseService[T any](collection *mongo.Collection) *BaseService[T] {
	return &BaseService[T]{
		collection: collection,
	}
}

//
// Reads
//

// Returns the first document that matches the filter.
func (s *BaseService[T]) FindOne(
	ctx context.Context,
	filter interface{},
	opts ...*options.FindOneOptions,
) (*T, error) {
	var data T

	result := s.collection.FindOne(ctx, filter, opts...)
	if result.Err() != nil {
		return nil, result.Err()
	}

	err := result.Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// Find all documents that match the filter.
func (s *BaseService[T]) Find(
	ctx context.Context,
	filter interface{},
	opts ...*options.FindOptions,
) ([]T, error) {
	var data []T

	cursor, err := s.collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Aggregate documents.
func (s *BaseService[T]) Aggregate(
	v any,
	ctx context.Context,
	pipeline mongo.Pipeline,
	opts ...*options.AggregateOptions,
) error {
	// Default opts if none are provided.
	actualOpts := options.
		Aggregate().
		SetCollation(&options.Collation{Strength: 3, Locale: "en"})
	if len(opts) > 0 {
		actualOpts = opts[0]
	}

	cursor, err := s.collection.Aggregate(ctx, pipeline, actualOpts)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, v)
	if err != nil {
		return err
	}

	return nil
}

type GetOneOrFailOptions struct {
	Message string
	Options *options.FindOneOptions
}

// High-level method to get one document or return an error.
// If the document is not found, it will return entity.NotFound with the message.
//
// The message will use the default message from the BaseServiceOptions
// during service initialization if not provided.
func (s *BaseService[T]) GetOneOrFail(
	ctx context.Context,
	filter interface{},
	opts ...*GetOneOrFailOptions,
) (*T, *entity.HttpError) {
	// Decide which message to use.
	actualOpts := &GetOneOrFailOptions{
		Message: s.options.GetOneOrFailMessage,
		Options: nil,
	}
	if len(opts) > 0 {
		actualOpts = opts[0]
	}

	findOne, findOneErr := s.FindOne(ctx, filter, actualOpts.Options)
	if findOneErr == mongo.ErrNoDocuments {
		return nil, entity.NotFound(actualOpts.Message)
	}
	if findOneErr != nil {
		return nil, entity.InternalServerError(findOneErr.Error())
	}

	return findOne, nil
}

type FindOrFailOptions struct {
	Message string
	Options *options.FindOptions
}

// High-level method to find documents or return an error if one or more documents are not found.
// If one or more documents are not found, it will return entity.NotFound with the message.
//
// The message will use the default message from the BaseServiceOptions
// during service initialization if not provided.
func (s *BaseService[T]) FindOrFail(
	ctx context.Context,
	filter interface{},
	expectedLength int,
	opts ...*FindOrFailOptions,
) ([]T, *entity.HttpError) {
	// Decide which message to use.
	actualOpts := &FindOrFailOptions{
		Message: s.options.FindOrFailMessage,
		Options: nil,
	}
	if len(opts) > 0 {
		actualOpts = opts[0]
	}

	find, findErr := s.Find(ctx, filter, actualOpts.Options)
	if len(find) != expectedLength {
		return nil, entity.NotFound(actualOpts.Message)
	}
	if findErr != nil {
		return nil, entity.InternalServerError(findErr.Error())
	}

	return find, nil
}

//
// Inserts
//

// Insert and return the insert result.
func (s *BaseService[T]) Create(
	ctx context.Context,
	createData T,
	opts ...*options.InsertOneOptions,
) (*T, error) {
	createResult, createErr := s.InsertOne(ctx, createData, opts...)
	if createErr != nil {
		return nil, createErr
	}

	findOneResult, findOneErr := s.FindOne(ctx, bson.M{"_id": createResult})
	if findOneErr != nil {
		return nil, findOneErr
	}

	return findOneResult, nil
}

// Insert and return the inserted id.
func (s *BaseService[T]) InsertOne(
	ctx context.Context,
	data T,
	opts ...*options.InsertOneOptions,
) (*primitive.ObjectID, error) {

	payload, err := s.setTimeStamp(data)
	if err != nil {
		return nil, err
	}

	result, err := s.collection.InsertOne(ctx, payload, opts...)
	if err != nil {
		return nil, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return &oid, nil
	} else {
		// This should never happen since any error should be caught by the InsertOne method.
		return nil, errors.New("somehow... inserted id is not an ObjectID")
	}
}

// Insert many documents and return the inserted ids.
func (s *BaseService[T]) InsertMany(
	ctx context.Context,
	data []T,
	opts ...*options.InsertManyOptions,
) ([]interface{}, error) {
	payloads := []any{}

	for _, d := range data {
		payload, err := s.setTimeStamp(d)
		if err != nil {
			return nil, err
		}

		payloads = append(payloads, payload)
	}

	result, err := s.collection.InsertMany(ctx, payloads, opts...)

	if err != nil {
		return nil, err
	}

	return result.InsertedIDs, nil
}

//
// Updates
//

// Update one document.
func (s *BaseService[T]) UpdateOne(
	ctx context.Context,
	filter interface{},
	updateData bson.M,
	opts ...*options.UpdateOptions,
) (int, error) {
	var data T
	if s.hasTimestamp(data) {
		d, err := parser.StructToMap(updateData["$set"])
		if err != nil {
			return 0, err
		}
		delete(d, "createdAt")
		d["updatedAt"] = time.Now()
		updateData["$set"] = d
	}
	result, err := s.collection.UpdateOne(ctx, filter, updateData, opts...)
	if err != nil {
		return 0, err
	}

	return int(result.ModifiedCount), nil
}

// Update many documents.
func (s *BaseService[T]) UpdateMany(
	ctx context.Context,
	filter interface{},
	updateData bson.M,
	opts ...*options.UpdateOptions,
) (int, error) {
	var data T
	if s.hasTimestamp(data) {
		if updateData["$set"] == nil {
			updateData["$set"] = bson.M{}
		}

		d, err := parser.StructToMap(updateData["$set"])
		if err != nil {
			return 0, err
		}
		delete(d, "createdAt")
		d["updatedAt"] = time.Now()
		updateData["$set"] = d
	}
	result, err := s.collection.UpdateMany(ctx, filter, updateData, opts...)
	if err != nil {
		return 0, err
	}

	return int(result.ModifiedCount), nil
}

// alt for update many
func (s *BaseService[T]) UpdateManyOld(ctx context.Context, filter interface{}, data interface{},
	opts ...*options.UpdateOptions) (int, error) {
	result, err := s.collection.UpdateMany(ctx, filter, data, opts...)
	if err != nil {
		return 0, err
	}

	return int(result.ModifiedCount), nil
}

// Update document that match the filter and return the updated document.
func (s *BaseService[T]) FindOneAndUpdate(
	ctx context.Context,
	filter interface{},
	updateData bson.M,
	opts ...*options.FindOneAndUpdateOptions,
) (*T, error) {
	var data T
	if s.hasTimestamp(data) {
		if _, exists := updateData["$set"]; !exists {
			updateData["$set"] = bson.M{}
		}

		d, err := parser.StructToMap(updateData["$set"])
		if err != nil {
			return nil, err
		}
		delete(d, "createdAt")
		d["updatedAt"] = time.Now()
		updateData["$set"] = d
	}

	// Defaults to returning the updated document.
	actualOpts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	if len(opts) > 0 {
		actualOpts = opts[0]
	}

	result := s.collection.FindOneAndUpdate(ctx, filter, updateData, actualOpts)
	if result.Err() != nil {
		return nil, result.Err()
	}

	err := result.Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

//
// Deletes
//

// Delete one document and return the number of documents deleted.
func (s *BaseService[T]) DeleteOne(
	ctx context.Context,
	filter interface{},
	opts ...*options.DeleteOptions,
) (int, error) {
	result, err := s.collection.DeleteOne(ctx, filter, opts...)
	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), nil
}

// Delete many documents and return the number of documents deleted.
func (s *BaseService[T]) DeleteMany(
	ctx context.Context,
	filter interface{},
	opts ...*options.DeleteOptions,
) (int, error) {
	result, err := s.collection.DeleteMany(ctx, filter, opts...)
	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), nil
}

// Delete a document and return the document before deletion.
func (s *BaseService[T]) FindOneAndDelete(
	ctx context.Context,
	filter interface{},
	opts ...*options.FindOneAndDeleteOptions,
) (*T, error) {
	var data T

	result := s.collection.FindOneAndDelete(ctx, filter, opts...)
	if result.Err() != nil {
		return nil, result.Err()
	}

	err := result.Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

//
// Utilities
//

func (s *BaseService[T]) hasTimestamp(v interface{}) bool {
	t := reflect.TypeOf(v)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return false
	}

	_, hasCreatedAt := t.FieldByName("CreatedAt")
	_, hasUpdatedAt := t.FieldByName("UpdatedAt")

	return hasCreatedAt && hasUpdatedAt
}

func (s *BaseService[T]) setTimeStamp(v T) (bson.M, error) {
	payload, err := parser.StructToMap(v)
	if err != nil {
		return nil, err
	}

	if s.hasTimestamp(v) {
		payload["createdAt"] = time.Now()
		payload["updatedAt"] = time.Now()
	}

	return payload, nil
}

// Count the number of documents that match the filter.
func (s *BaseService[T]) CountDocuments(
	ctx context.Context,
	filter interface{},
) (int, error) {
	documentCounts, err := s.collection.CountDocuments(ctx, filter)

	if err != nil {
		return 0, err
	}

	return int(documentCounts), nil
}

// Create a unique index on the collection.
func (s *BaseService[T]) MakeUnique(
	ctx context.Context,
	filter interface{},
) error {
	indexModel := mongo.IndexModel{
		Keys:    filter,
		Options: options.Index().SetUnique(true),
	}

	_, err := s.collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}

	return nil
}

// Create a unique index on the collection.
func (s *BaseService[T]) SetDeleteFromDatabaseAttribute(
	ctx context.Context,
	filter interface{},
) error {
	indexModel := mongo.IndexModel{
		Keys:    filter,
		Options: options.Index().SetExpireAfterSeconds(0),
	}

	_, err := s.collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}

	return nil
}

// Create an index on the collection.
func (s *BaseService[T]) CreateIndex(
	ctx context.Context,
	fields interface{},
	opts ...*options.CreateIndexesOptions,
) error {
	indexModel := mongo.IndexModel{
		Keys: fields,
	}

	_, err := s.collection.Indexes().CreateOne(ctx, indexModel, opts...)

	return err
}

func (s *BaseService[T]) BulkWrite(
	ctx context.Context,
	models []mongo.WriteModel,
	opts ...*options.BulkWriteOptions,
) (*mongo.BulkWriteResult, error) {
	return s.collection.BulkWrite(ctx, models, opts...)
}
