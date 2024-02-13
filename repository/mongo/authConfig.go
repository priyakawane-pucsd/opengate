package mongo

import (
	"context"
	"opengate/models/dao"
	"opengate/utils"

	"github.com/bappaapp/goutils/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	AUTH_CONFIG_COLLECTION = "auth_config"
)

// CreateUpdateConfig creates or updates a service configuration in the MongoDB repository.
func (r *Repository) CreateUpdateAuthConfig(ctx context.Context, cfg *dao.AuthConfigService) error {
	// Specify the MongoDB collection and filter criteria
	collection := r.conn.Database(r.cfg.Database).Collection(AUTH_CONFIG_COLLECTION)

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
