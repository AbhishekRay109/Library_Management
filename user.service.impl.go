package main

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *User) error { // Calling the create function to create data
	_, err := u.usercollection.InsertOne(u.ctx, user) // Insertion of Data serially
	return err
}

/*func (u *UserServiceImpl) GetAllByName(bookname *string) (*User, error) {
	var user *User
	query := bson.D{bson.E{Key: "bookname", Value: bookname}} // Finding the data based on the user_name
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}*/

func (u *UserServiceImpl) GetAllByName(name *string) ([]*User, error) {
	var users []*User
	cursor, err := u.usercollection.Find(u.ctx, bson.D{bson.E{Key: "bookname", Value: name}}) // retrieving all the data in bson
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user User
		err := cursor.Decode(&user)

		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}
	return users, nil
}

func (u *UserServiceImpl) GetAllByRange(rent *string) ([]*User, error) {
	var users []*User
	cursor, err := u.usercollection.Find(u.ctx, bson.D{bson.E{Key: "rentperday", Value: rent}}) // retrieving all the data in bson
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user User
		err := cursor.Decode(&user)

		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}
	return users, nil
}

func (u *UserServiceImpl) GetAllByCat(cat *string) ([]*User, error) {
	var users []*User
	cursor, err := u.usercollection.Find(u.ctx, bson.D{bson.E{Key: "category", Value: cat}}) // retrieving all the data in bson
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user User
		err := cursor.Decode(&user)

		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}
	return users, nil
}

func (u *UserServiceImpl) UpdateIssue(user *User) error {
	filter := bson.D{primitive.E{Key: "bookname", Value: user.BookName}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "bookname", Value: user.BookName}, primitive.E{Key: "bookdetail", Value: user.BookDetail}}}}
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update) // replacing the old data with the new data
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *UserServiceImpl) UpdateReturn(user *User) error {
	filter := bson.D{primitive.E{Key: "bookname", Value: user.BookName}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "bookname", Value: user.BookName}, primitive.E{Key: "bookdetail", Value: user.BookDetail}}}}
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update) // replacing the old data with the new data
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}
