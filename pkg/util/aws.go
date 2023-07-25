package util

import (
	"msk-glue-demo/pkg/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/linkedin/goavro/v2"
)

// NewGlueSvc returns a new glue client.
func NewGlueSvc() *glue.Glue {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(GetEnvFallback(config.AwsRegionEnv, config.AwsDefaultRegion)),
			Credentials: credentials.NewStaticCredentials(
				GetEnv(config.AwsAccessKeyIdEnv),
				GetEnv(config.AwsSecretAccessKeyEnv),
				GetEnv(config.AwsSessionTokenEnv),
			),
		},
	)
	if err != nil {
		panic("Error creating session: " + err.Error())
	}
	return glue.New(sess)
}

// CreateSchemaRegistry creates a new schema registry or panic when failed.
func CreateSchemaRegistry(svc *glue.Glue, name string) {
	createRegistryInput := &glue.CreateRegistryInput{
		RegistryName: aws.String(name),
	}

	_, err := svc.CreateRegistry(createRegistryInput)
	if err != nil {
		panic("Error creating new registry: " + err.Error())
	}
}

// RegisterSchema registers a new schema or panic when failed.
func RegisterSchema(svc *glue.Glue, registry, schemaName, schemaDefinition string) {
	createSchemaInput := &glue.CreateSchemaInput{
		RegistryId: &glue.RegistryId{
			RegistryName: &registry,
		},
		SchemaName:       aws.String(schemaName),
		DataFormat:       aws.String("AVRO"),
		SchemaDefinition: aws.String(schemaDefinition),
		Compatibility:    aws.String("FULL"),
	}

	_, err := svc.CreateSchema(createSchemaInput)
	if err != nil {
		panic("Error creating new schema: " + err.Error())
	}
}

// GetSchema returns the latest version of the given schema.
func GetSchema(glueSvc *glue.Glue, registry, schema string) string {
	result, err := glueSvc.GetSchemaVersion(
		&glue.GetSchemaVersionInput{
			SchemaId: &glue.SchemaId{
				RegistryName: aws.String(registry),
				SchemaName:   aws.String(schema),
			},
			SchemaVersionNumber: &glue.SchemaVersionNumber{
				LatestVersion: aws.Bool(true),
			},
		})

	if err != nil {
		panic("Error getting schema: " + err.Error())
	}

	return *result.SchemaDefinition
}

// EncodeMessage encodes a message using the given schema definition.
func EncodeMessage(msg map[string]interface{}, schemaDefinition string) []byte {
	codec, err := goavro.NewCodec(schemaDefinition)
	if err != nil {
		panic("Error creating codec: " + err.Error())
	}

	binary, err := codec.BinaryFromNative(nil, msg)
	if err != nil {
		panic("Error encoding message: " + err.Error())
	}

	return binary
}
