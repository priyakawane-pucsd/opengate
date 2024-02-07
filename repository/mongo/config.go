package mongo

import (
	"context"
	"errors"
	"opengate/models/dao"
	"opengate/utils"

	"github.com/bappaapp/goutils/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	SERVICE_CONFIG_COLLECTION = "service_config"
)

// CreateUpdateConfig creates or updates a service configuration in the MongoDB repository.
func (r *Repository) CreateUpdateConfig(ctx context.Context, cfg *dao.ServiceConfig) error {
	// Specify the MongoDB collection and filter criteria
	collection := r.conn.Database(r.cfg.Database).Collection(SERVICE_CONFIG_COLLECTION)

	// If the ID is provided, check if a document with the given ID already exists
	if cfg.Id != "" {
		existingFilter := bson.M{"_id": cfg.Id}
		existingDocCount, err := collection.CountDocuments(ctx, existingFilter)
		if err != nil {
			logger.Error(ctx, "Error checking existing document: %v", err)
			return utils.NewInternalServerError("Failed to check existing document in the database")
		}

		// If a document with the ID already exists, perform an update
		if existingDocCount > 0 {
			update := bson.M{"$set": cfg}
			_, err := collection.UpdateOne(ctx, existingFilter, update)
			if err != nil {
				logger.Error(ctx, "Error updating existing service configuration: %v", err)
				return utils.NewInternalServerError("Failed to update existing config in the database")
			}
			return nil
		}
	}

	// If the function reaches this point, the document with the given ID doesn't exist, so create a new record
	newObjectID := primitive.NewObjectID()
	cfg.Id = newObjectID.Hex()
	_, err := collection.InsertOne(ctx, cfg)
	if err != nil {
		logger.Error(ctx, "Error inserting new service configuration: %v", err)
		return utils.NewInternalServerError("Failed to insert new config into the database")
	}
	return nil
}

func (r *Repository) GetAllConfigs(ctx context.Context) ([]*dao.ServiceConfig, error) {
	collection := r.conn.Database(r.cfg.Database).Collection(SERVICE_CONFIG_COLLECTION)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		logger.Error(ctx, "Error finding all service configurations: %v", err)
		return nil, utils.NewInternalServerError("Failed to retrieve all configs from the database")
	}
	defer cursor.Close(ctx)

	var configs []*dao.ServiceConfig
	if err := cursor.All(ctx, &configs); err != nil {
		logger.Error(ctx, "Error decoding service configurations: %v", err)
		return nil, utils.NewInternalServerError("Failed to decode service configs")
	}
	return configs, nil
}

func (r *Repository) GetConfigById(ctx context.Context, id string) (*dao.ServiceConfig, error) {
	// Specify the MongoDB collection and filter criteria
	collection := r.conn.Database(r.cfg.Database).Collection(SERVICE_CONFIG_COLLECTION)

	// Perform a query to find the document by ID
	filter := bson.M{"_id": id}
	var result *dao.ServiceConfig
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, utils.NewCustomError(404, "Config not found")
		}
		logger.Error(ctx, "Error finding config by ID: %v", err)
		return nil, utils.NewInternalServerError("Failed to retrieve config by ID")
	}
	return result, nil
}

// Repository method to delete a configuration by ID
func (r *Repository) DeleteConfigById(ctx context.Context, id string) error {
	collection := r.conn.Database(r.cfg.Database).Collection(SERVICE_CONFIG_COLLECTION)

	// Specify the filter based on the ID
	filter := bson.M{"_id": id}

	// Perform the deletion operation
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		logger.Error(ctx, "Error deleting configuration by ID: %v", err)
		return utils.NewInternalServerError("Failed to delete configuration")
	}

	// Check if no documents were matched and deleted
	if result.DeletedCount == 0 {
		return utils.NewCustomError(404, "Configuration not found with this Id")
	}

	return nil
}
