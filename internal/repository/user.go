package repository

import (
	"time"

	"github.com/bryanck29/be-test/internal/contract"
	"github.com/bryanck29/be-test/internal/schema/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// userRepository represents user repo object
type userRepository struct {
	col *mongo.Collection
}

// newUserRepository used to intiate auth usecase
func newUserRepository(db *mongo.Database) contract.UserRepository {
	return &userRepository{
		col: db.Collection("user"),
	}
}

// InsertUser inserts user data into db
func (r *userRepository) InsertUser(ctx echo.Context, payload model.User) (err error) {
	_, err = r.col.InsertOne(ctx.Request().Context(), payload)
	return
}

// GetUsers fetches user data from db
func (r *userRepository) GetUsers(ctx echo.Context) (results []model.User, err error) {
	results = []model.User{}
	filter := bson.M{"$or": []bson.M{{"deleted_at": 0}, {"deleted_at": nil}}}
	cursor, err := r.col.Find(ctx.Request().Context(), filter)
	if err != nil {
		return
	}

	if err = cursor.All(ctx.Request().Context(), &results); err != nil {
		return
	}

	return
}

// GetUser fetches a user data from db based on id
func (r *userRepository) GetUser(ctx echo.Context, userId uuid.UUID) (result model.User, err error) {
	filter := bson.M{
		"$or": []bson.M{
			{"deleted_at": 0},
			{"deleted_at": nil},
		},
		"id": bson.M{
			"$eq": userId,
		},
	}
	err = r.col.FindOne(ctx.Request().Context(), filter).Decode(&result)

	if err == mongo.ErrNoDocuments {
		return model.User{}, nil
	}

	return
}

// DeleteUser deletes a user data from db based on id
func (r *userRepository) DeleteUser(ctx echo.Context, userId uuid.UUID) (err error) {
	filter := bson.M{"id": userId}
	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now().UTC().Unix(),
		},
	}
	result := r.col.FindOneAndUpdate(
		ctx.Request().Context(),
		filter,
		update,
	)

	return result.Err()
}

// UpdateUser updates a user data
func (r *userRepository) UpdateUser(ctx echo.Context, payload model.User) (result model.User, err error) {
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := bson.M{"id": payload.Id}
	update := bson.M{
		"$set": bson.M{
			"name":       payload.Name,
			"username":   payload.Username,
			"password":   payload.Password,
			"role":       payload.Role,
			"updated_at": time.Now().UTC().Unix(),
		},
	}

	err = r.col.FindOneAndUpdate(
		ctx.Request().Context(),
		filter,
		update,
		opts,
	).Decode(&result)

	return
}

// GetUserByUsername fetches a user data from db based on username
func (r *userRepository) GetUserByUsername(ctx echo.Context, username string) (result model.User, err error) {
	err = r.col.FindOne(ctx.Request().Context(), bson.D{{
		Key: "username", Value: username,
	}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		return model.User{}, nil
	}

	return
}
