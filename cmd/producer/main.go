package main

import (
	"context"
	"msk-glue-demo/pkg/config"
	"msk-glue-demo/pkg/util"

	"github.com/segmentio/kafka-go"
)

func main() {
	// 1. connect to the glue schema registry
	glueSvc := util.NewGlueSvc()

	// 2. create a schema registry
	util.CreateSchemaRegistry(glueSvc, config.RegistryName)

	// 3. register a schema to the schema registry
	util.RegisterSchema(glueSvc, config.RegistryName, config.SchemaName, config.SchemaDefinition)

	// 4. connect to the MSK
	kwriter := util.NewKafkaWriter()
	defer kwriter.Close()

	// 5. create a topic
	util.CreateTopic(config.TopicName)

	// 6. encode sample messages and write to kafka
	msgs := []map[string]interface{}{
		{
			"id":      1,
			"title":   "CEO",
			"content": "Chief Executive Officer",
		},
		{
			"id":      2,
			"title":   "CTO",
			"content": "Chief Technology Officer",
		},
		{
			"id":      3,
			"title":   "CFO",
			"content": "Chief Financial Officer",
		},
	}

	for _, msg := range msgs {
		bytes := util.EncodeMessage(msg, config.SchemaDefinition)
		err := kwriter.WriteMessages(
			context.TODO(),
			kafka.Message{
				Value: bytes,
			},
		)
		if err != nil {
			panic("Error writing message: " + err.Error())
		}
	}
}
