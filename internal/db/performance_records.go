package db

import (
	"PfMartin/upprove-mcp-server/internal/models"
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (store *MongoDbStore) GetAllPerformanceRecords(ctx context.Context) ([]models.PerformanceRecord, error) {
	var performanceRecords []models.PerformanceRecord

	cursor, err := store.performanceRecordsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Err(err).Msg("failed to find performance record documents")
		return performanceRecords, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &performanceRecords); err != nil {
		log.Err(err).Msg("failed to parse performance records documents")
		return performanceRecords, err
	}

	return performanceRecords, nil
}

func (store *MongoDbStore) CreatePerformanceRecord(ctx context.Context, performanceRecord models.PerformanceRecordCreate) (string, error) {
	timestamp := time.Now().UTC()

	insertData := bson.M{
		"category":    performanceRecord.Category,
		"description": performanceRecord.Description,
		"value":       performanceRecord.Value,
		"unit":        performanceRecord.Unit,
		"createdAt":   timestamp.String(),
		"modifiedAt":  timestamp.String(),
	}

	insertResult, err := store.performanceRecordsCollection.InsertOne(ctx, insertData)
	if err != nil {
		log.Err(err).Msgf("failed to insert performance record with description %s", performanceRecord.Description)
		return "", err
	}

	insertedID := insertResult.InsertedID.(primitive.ObjectID)

	return insertedID.Hex(), nil
}
