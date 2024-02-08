// pkg/database/mongodb/models/organization.go

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Organization represents the organization model
type Organization struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
}