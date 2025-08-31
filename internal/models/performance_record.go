package models

type PerformanceRecord struct {
	ID          string `bson:"_id" json:"id"`
	Category    string `bson:"category" json:"category"`
	Description string `bson:"description" json:"description"`
	Value       string `bson:"value" json:"value"`
	Unit        string `bson:"unit" json:"unit"`
	CreatedAt   int64  `bson:"createdAt" json:"createdAt"`
	ModifiedAt  int64  `bson:"modifiedAt" json:"modifiedAt"`
}

type PerformanceRecordCreate struct {
	Category    string `bson:"category" json:"category"`
	Description string `bson:"description" json:"description"`
	Value       string `bson:"value" json:"value"`
	Unit        string `bson:"unit" json:"unit"`
}
