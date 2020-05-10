package mongoOp

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"project/libraryManagementServer/db"
	"project/libraryManagementServer/util"
)

var (
	client *mongo.Client
	err    error
)

// dbConnect : starts db connection
func dbConnect() (context.Context, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(db.MongoDbURI)

	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		cancel()
		util.Log("Failed connection, error : ", err)
		return nil, err
	}

	return ctx, nil
}

func RegisterUser(newUser db.User) error {

	ctx, err := dbConnect()
	if err != nil {
		return err
	}

	collection := client.Database(db.Database).Collection(db.MongoUserCollection)
	ctx, ctxCancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer ctxCancelFunc()

	_, err = collection.InsertOne(ctx, newUser)
	if err != nil {
		util.Logf("Failed to Insert User record, err :", err)
		return err
	}

	return nil
}

func UserExist(user db.User) (bool, error) {

	ctx, err := dbConnect()
	if err != nil {
		return false, err
	}

	collection := client.Database(db.Database).Collection(db.MongoUserCollection)
	ctx, ctxCancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer ctxCancelFunc()

	filter := bson.D{
		primitive.E{Key: "firstName", Value: user.FName},
		primitive.E{Key: "lastName", Value: user.LName},
		primitive.E{Key: "dateOfBirth", Value: user.DOB},
		primitive.E{Key: "address", Value: user.Address},
		primitive.E{Key: "email", Value: user.Email},
		primitive.E{Key: "contactNo", Value: user.ContactNo}}

	var newuser db.User

	err = collection.FindOne(ctx, filter).Decode(&newuser)
	if err != nil {
		util.Logf("Failed to fetch single user, err :", err)
		return false, err
	}

	return true, nil
}

// //InsertBook : Add book to db
// func (bookName *db.NewBook) InsertBook() (interface{}, error) {
// 	bookName.ID = primitive.NewObjectID()
// 	bookName.InsertTime = time.Now()

// 	ctx, err := dbConnect()
// 	if err != nil {
// 		return nil, err
// 	}

// 	collection := client.Database(database).Collection(collection)
// 	ctx, ctxCancelFunc := context.WithTimeout(ctx, 5*time.Second)
// 	defer ctxCancelFunc()

// 	insertOneRslt, err := collection.InsertOne(ctx, bookName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return insertOneRslt.InsertedID, nil
// }

// // FindOneBook : Retrieves one book from db
// func (bookName *db.NewBook) FindOneBook() (*NewBook, error) {
// 	ctx, err := dbConnect()
// 	if err != nil {
// 		return nil, err
// 	}

// 	collection := client.Database(database).Collection(collection)
// 	ctx, ctxCancelFunc := context.WithTimeout(ctx, 5*time.Second)
// 	defer ctxCancelFunc()

// 	filter := bson.D{
// 		primitive.E{Key: "serialNo", Value: bookName.SerialNos},
// 		primitive.E{Key: "bookName", Value: bookName.BookName}}

// 	var book NewBook

// 	err = collection.FindOne(ctx, filter).Decode(&book)
// 	if err != nil {
// 		return nil, NewErrorNotFound()
// 	}

// 	return &book, nil
// }

// // UpdateBook : updates the book into db
// func (bookName *db.NewBook) UpdateBook() (interface{}, error) {
// 	ctx, err := dbConnect()
// 	if err != nil {
// 		return nil, err
// 	}

// 	collection := client.Database(database).Collection(collection)
// 	ctx, ctxCancelFunc := context.WithTimeout(ctx, 10*time.Second)
// 	defer ctxCancelFunc()

// 	bookName.UpdateTime = time.Now()

// 	filter := bson.D{
// 		primitive.E{Key: "serialNo", Value: bookName.SerialNos},
// 		primitive.E{Key: "bookName", Value: bookName.BookName}}

// 	update := bson.M{
// 		"$set": bson.M{
// 			"author":      bookName.Author,
// 			"category":    bookName.Category,
// 			"update_time": bookName.UpdateTime,
// 		},
// 	}

// 	updateResult, err := collection.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return updateResult, nil
// }

// // DeleteBook : deletes one book from db
// func (bookName *db.NewBook) DeleteBook() (int64, error) {
// 	ctx, err := dbConnect()
// 	if err != nil {
// 		return 0, err
// 	}

// 	collection := client.Database(database).Collection(collection)
// 	ctx, ctxCancelFunc := context.WithTimeout(ctx, 5*time.Second)
// 	defer ctxCancelFunc()

// 	filter := bson.D{
// 		primitive.E{Key: "serialNo", Value: bookName.SerialNos},
// 		primitive.E{Key: "bookName", Value: bookName.BookName}}

// 	deleteResult, err := collection.DeleteOne(ctx, filter)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return deleteResult.DeletedCount, nil
// }

// // FindAllBooks retrieves all books from db
// func FindAllBooks(bookName string) ([]db.NewBook, error) {
// 	ctx, err := dbConnect()
// 	if err != nil {
// 		return nil, err
// 	}

// 	collection := client.Database(database).Collection(collection)
// 	ctx, ctxCancelFunc := context.WithTimeout(ctx, 10*time.Second)
// 	defer ctxCancelFunc()

// 	var filter primitive.D

// 	if bookName == "" {
// 		filter = bson.D{}
// 	} else {
// 		filter = bson.D{primitive.E{Key: "bookName", Value: bookName}}
// 	}

// 	cursor, err := collection.Find(ctx, filter)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var books []NewBook

// 	err = cursor.All(ctx, &books)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return books, nil
// }
