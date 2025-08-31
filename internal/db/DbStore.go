package db

import (
	"PfMartin/upprove-mcp-server/internal/models"
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DbStore interface {
	CreatePerformanceRecord(ctx context.Context, performanceRecord models.PerformanceRecordCreate) (string, error)
	GetAllPerformanceRecords(ctx context.Context) ([]models.PerformanceRecord, error)
}

type MongoDbStore struct {
	performanceRecordsCollection *mongo.Collection
}

func newDatabaseClient(authSource string, username string, password string, uri string) (*mongo.Client, context.CancelFunc) {
	credentials := options.Credential{
		AuthSource: authSource,
		Username:   username,
		Password:   password,
	}

	clientOptions := options.Client().ApplyURI(uri).SetAuth(credentials)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	dbClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		cancel()
		log.Fatal().Msgf("failed to connect to database: %s", err)
	}

	if err = dbClient.Ping(ctx, readpref.Primary()); err != nil {
		cancel()
		log.Fatal().Msgf("failed to ping database: %s", err)
	}

	log.Info().Msg("Connected to database")

	return dbClient, cancel
}

func NewMongoDbStore(dbName, dbUser, dbPassword, dbURI string) *MongoDbStore {
	client, _ := newDatabaseClient(dbName, dbUser, dbPassword, dbURI)

	database := client.Database(dbName)

	return &MongoDbStore{
		performanceRecordsCollection: database.Collection("performanceRecords"),
	}
}
