package db

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

// NewBook struct to contain book info when storing
type NewBook struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	SerialNos  int                `json:"serial_no" bson:"serialNo"`
	BookName   string             `json:"name" bson:"bookName" `
	Author     string             `json:"author" bson:"author"`
	Category   string             `json:"category" bson:"category"`
	InsertTime time.Time          `json:"entry_time" bson:"insertTime"`
	UpdateTime time.Time          `json:"update_time" bson:"updateTime"`
}

type User struct {
	FName     string `json:"first_name" bson:"firstName"`
	LName     string `json:"last_name" bson:"lastName"`
	DOB       string `json:"date_of_birth" bson:"dateOfBirth"`
	Address   string `json:"address" bson:"address"`
	Email     string `json:"email" bson:"email"`
	ContactNo string `json:"contact_no" bson:"contactNo"`
	Password  string `json:"pswd" bson:"password"`
}

type LoginUser struct {
	UserName string `json:"user_name"`
	//Email        string `json:"email"`
	PasswordHash string `json:"password"`
	Role         string `json:"role"`
}

type TokenClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
}

// ErrNotFound : custom error for object not
type ErrNotFound struct {
	message string
}

var (
	client *mongo.Client
	err    error
)

const (
	Database            = "library"
	MongoBookCollection = "booksStore"
	MongoUserCollection = "userRecords"
	MongoDbURI          = "mongodb://localhost:27017"
)

const (
	RefreshTokenValidTime = time.Hour * 72
	AuthTokenValidTime    = time.Minute * 15
)

// NewErrorNotFound : creates a new error type specifying Not Found value
func NewErrorNotFound() *ErrNotFound {
	return &ErrNotFound{
		message: "Not Found",
	}
}

func (err *ErrNotFound) Error() string {
	return err.message
}
