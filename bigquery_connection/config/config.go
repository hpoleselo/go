package config

type BigQueryConfig struct {
	// BigQuery settings that are available for all packages for the application
	ProjectID string "mlopsplatform"
	DatasetID string "model_catalog"
	TableID   string "model_catalog_table_t"
}
