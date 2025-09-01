package models

type PerformanceRecord struct {
	ID          string `bson:"_id" json:"id"`
	Category    string `bson:"category" json:"category"`
	Description string `bson:"description" json:"description"`
	Value       string `bson:"value" json:"value"`
	Unit        string `bson:"unit" json:"unit"`
	CreatedAt   string `bson:"createdAt" json:"createdAt"`
	ModifiedAt  string `bson:"modifiedAt" json:"modifiedAt"`
}

type PerformanceRecordCreate struct {
	Category    string `bson:"category" json:"category"`
	Description string `bson:"description" json:"description"`
	Value       string `bson:"value" json:"value"`
	Unit        string `bson:"unit" json:"unit"`
}
