package main

import (
	"context"
	"fmt"

	"msk-glue-demo/pkg/config"
	"msk-glue-demo/pkg/util"

	"github.com/linkedin/goavro/v2"
)

func main() {
	// 1. connect to Glue
	glueSvc := util.NewGlueSvc()

	// 2. get the target schema
	schema := util.GetSchema(glueSvc, config.RegistryName, config.SchemaName)
	codec, err := goavro.NewCodec(schema)
	if err != nil {
		panic("Error creating codec: " + err.Error())
	}

	// 3. connect to MSK
	kreader := util.NewKafkaReader()

	// 4. read messages and decode them
	for {
		m, err := kreader.ReadMessage(context.Background())
		if err != nil {
			panic("Error reading message: " + err.Error())
		}
		msg, _, err := codec.NativeFromBinary(m.Value)
		if err != nil {
			panic("Error decoding message: " + err.Error())
		}
		record := msg.(map[string]interface{})
		fmt.Printf(
			"Id: %v, Title: %v, Content: %v\n",
			record["id"], record["title"], record["content"],
		)
	}
}
