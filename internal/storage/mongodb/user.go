package mongodb

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	user "goresizer.com/m/internal/storage"
	"goresizer.com/m/pkg/logging"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

// Create implements user.Storage.
func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	d.logger.Debug("create user")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user due to error: %v", err)
	}

	d.logger.Debug("cinver InsertedID to ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(user)
	return "", fmt.Errorf("failed to convert objectid to hex. oid:%s", oid)
}

// Delete implements user.Storage.
func (d *db) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert user ID to objectID. ID:%s", id)
	}
	filter := bson.M{"_id": oid}

	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute querry. error:%v", err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("not found")
	}

	d.logger.Tracef("Deleted %d documents", result.DeletedCount)
	return nil

}

// FindOne implements user.Storage.
func (d *db) FindOne(ctx context.Context, customFilter user.FindUserByFilter) (u user.User, err error) {
	filter := bson.M{}
	if customFilter.Email != "" {
		filter["email"] = customFilter.Email
	}
	if customFilter.UserID != "" {
		filter["id"] = customFilter.UserID
	}

	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return u, fmt.Errorf("ErrEntityNotFound")
		}
		return u, fmt.Errorf("failed to find one user by id: %s due to error: %v", customFilter, err)
	}
	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user(id:%s) from DB due to error: %v", customFilter, err)
	}

	return u, nil
}

// Update implements user.Storage.
func (d *db) Update(ctx context.Context, user user.User) error {
	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return fmt.Errorf("failed to convert user ID to ObjectID, ID=%s", user.ID)

	}

	filter := bson.M{"_id": objectID}

	userBytes, err := bson.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshall user, error:%v", err)
	}

	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal user bytes. error:%v", err)
	}

	delete(updateUserObj, "_id")

	update := bson.M{
		"$set": updateUserObj,
	}
	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute update user query. error: %v", err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("not found")
	}

	d.logger.Tracef("Matched %d documents and Modified %d documents", result.MatchedCount, result.ModifiedCount)
	return nil
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}

}
