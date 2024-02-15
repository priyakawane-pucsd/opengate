package mongo

import (
	"context"
	"errors"
	"net/http"
	"opengate/models/dao"
	"opengate/utils"

	"github.com/bappaapp/goutils/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CONFIG_COLLECTION = "configs"
)

// CreateUpdateConfig creates or updates a service configuration in the MongoDB repository.
func (r *Repository) CreateUpdateConfig(ctx context.Context, cfg *dao.ServiceConfig) error {
	// Specify the MongoDB collection and filter criteria
	collection := r.conn.Database(r.cfg.Database).Collection(CONFIG_COLLECTION)

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
	collection := r.conn.Database(r.cfg.Database).Collection(CONFIG_COLLECTION)
	filter := bson.M{"type": "serviceConfig"}

	cursor, err := collection.Find(ctx, filter)
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
	collection := r.conn.Database(r.cfg.Database).Collection(CONFIG_COLLECTION)

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
	collection := r.conn.Database(r.cfg.Database).Collection(CONFIG_COLLECTION)

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

// CreateUpdateConfig creates or updates a service configuration in the MongoDB repository.
func (r *Repository) CreateUpdateAuthConfig(ctx context.Context, cfg *dao.AuthConfig) error {
	// Specify the MongoDB collection and filter criteria
	collection := r.conn.Database(r.cfg.Database).Collection(CONFIG_COLLECTION)
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

func (r *Repository) CreateUpdateAuthConfig2(ctx context.Context, cfg *dao.AuthConfig) error {
	// Create a filter to match on 'Type'.
	//if cfg
	filter := bson.M{"type": cfg.Type}

	// Create an update document with the provided AuthConfig.
	update := bson.M{
		"$set": cfg,
	}

	// Set upsert option to insert if document not found.
	opts := options.Update().SetUpsert(true)

	// Perform the upsert operation.
	_, err := r.conn.Database(r.cfg.Database).Collection(CONFIG_COLLECTION).UpdateOne(ctx, filter, update, opts)
	if err != nil {
		logger.Error(ctx, "Error inserting new service configuration: %v", err)
		return utils.NewInternalServerError("Failed to insert new config into the database")

	}

	return nil
}

func (r *Repository) GetAuthConfig(ctx context.Context) (*dao.AuthConfig, error) {
	// Create a filter to match on 'Type'.
	filter := bson.M{"type": "authConfig"}

	// Perform the find operation.
	result := r.conn.Database(r.cfg.Database).Collection(CONFIG_COLLECTION).FindOne(ctx, filter)
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			// Handle the case where no document is found (not an error).
			return nil, utils.NewCustomError(http.StatusNotFound, "document not found")
		}
		// Handle other errors, e.g., log it or return an error.
		return nil, result.Err()
	}

	// Decode the result into an AuthConfig struct.
	var authConfig dao.AuthConfig
	if err := result.Decode(&authConfig); err != nil {
		// Handle decoding error, e.g., log it or return an error.
		return nil, err
	}
	return &authConfig, nil
}
