// pkg/database/mongodb/repository/organization_repository.go

package repository

import (
	"context"
	"log"

	"github.com/mina-droid/Ideanest-Software-Engineering-Task/pkg/database/mongodb"
	"github.com/mina-droid/Ideanest-Software-Engineering-Task/pkg/database/mongodb/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var organizationCollection *mongo.Collection

func init() {
	// Get the MongoDB client
	client := mongodb.GetMongoDBClient()

	// Get the organization collection
	organizationCollection = client.Database("yourdbname").Collection("organizations")
}

// CreateOrganization inserts a new organization into the database
func CreateOrganization(name, description string) (string, error) {
	// Create a new organization document
	organization := models.Organization{
		Name:        name,
		Description: description,
	}

	// Insert the organization into the database
	result, err := organizationCollection.InsertOne(context.Background(), organization)
	if err != nil {
		log.Println("Error creating organization:", err)
		return "", err
	}

	// Retrieve the inserted organization ID
	organizationID := result.InsertedID.(primitive.ObjectID).Hex()

	return organizationID, nil
}
