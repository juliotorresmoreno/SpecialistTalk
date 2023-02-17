package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	ID        *primitive.ObjectID `json:"id"         bson:"_id"`
	Name      string              `json:"name"       bson:"name"`
	Owner     string              `json:"owner"      bson:"owner"`
	CreatedAt *time.Time          `json:"created_at" bson:"created_at"`
}

func (u *File) TableName() string {
	return "files"
}
