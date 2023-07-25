package config

const (
	RegistryName     = "msk-glue-demo"
	SchemaName       = "demo-schema"
	SchemaDefinition = `{
	"namespace": "example.avro",
	"type": "record",
	"name": "User",
	"fields": [
		{"name": "id", "type": "int"},
		{"name": "title", "type": "string"},
		{"name": "content", "type": "string"}
	]
}`
	TopicName        = "demo-topic"
	AwsDefaultRegion = "us-west-2"

	AwsAccessKeyIdEnv     = "AWS_ACCESS_KEY_ID"
	AwsSecretAccessKeyEnv = "AWS_SECRET_ACCESS_KEY"
	AwsSessionTokenEnv    = "AWS_SESSION_TOKEN"
	AwsRegionEnv          = "AWS_REGION"
	KafkaServersEnv       = "KAFKA_SERVERS"
)
