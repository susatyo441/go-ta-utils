package service

import (
	"context"
	"errors"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/susatyo441/go-ta-utils/db"
	"github.com/susatyo441/go-ta-utils/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IChangelogUseCase interface {
	CreateChangelog(
		ctx context.Context,
		data model.Changelog,
		userId primitive.ObjectID,
	) error
}

type ChangelogUseCase struct {
	ChangelogService Service[model.Changelog]
	UserService      Service[model.User]
}

func NewCompanyChangelogUseCase(companyCode string) IChangelogUseCase {
	return &ChangelogUseCase{
		ChangelogService: NewCompanyService[model.Changelog](companyCode, db.ChangelogModelName),
		UserService:      NewCompanyService[model.User](companyCode, db.UserModelName),
	}
}

func NewAdminChangelogUseCase() IChangelogUseCase {
	return &ChangelogUseCase{
		ChangelogService: NewAdminService[model.Changelog](db.ChangelogModelName),
		UserService:      NewAdminService[model.User](db.UserModelName),
	}
}

func (u *ChangelogUseCase) CreateChangelog(
	ctx context.Context,
	data model.Changelog,
	userId primitive.ObjectID,
) error {
	loggedInUser, loggedInErr := u.UserService.FindOne(ctx, bson.M{"_id": userId})
	if loggedInErr == mongo.ErrNoDocuments {
		return errors.New("user not found while trying to create changelog")
	}
	if loggedInErr != nil {
		return loggedInErr
	}

	changelog := model.Changelog{
		ID:           primitive.NewObjectID(),
		Action:       data.Action,
		Field:        data.Field,
		OldValue:     data.OldValue,
		NewValue:     data.NewValue,
		ModifiedBy:   loggedInUser.Name,
		ModifiedById: loggedInUser.ID,
		Object:       data.Object,
		ObjectId:     data.ObjectId,
		ObjectName:   data.ObjectName,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	_, err := u.ChangelogService.InsertOne(ctx, changelog)
	if err != nil {
		return err
	}

	return nil
}

type MockChangelogUseCase struct {
	mock.Mock
}

func (m *MockChangelogUseCase) CreateChangelog(
	ctx context.Context,
	data model.Changelog,
	userId primitive.ObjectID,
) error {
	args := m.Called(ctx, data, userId)
	return args.Error(0)
}

// Alias for Mock.On("CreateChangelog", mock.Anything, ...)
func (m *MockChangelogUseCase) OnCreateChangelog() *mock.Call {
	return m.Mock.On(
		"CreateChangelog",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	)
}
