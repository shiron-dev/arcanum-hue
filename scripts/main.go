package main

import (
	"encoding/json"
	"os"

	"github.com/invopop/jsonschema"
	"github.com/shiron-dev/arcanum-hue/internal/converter/model"
)

func main() {
	file, err := os.Create("./schema.json")
	if err != nil {
		panic(err)
	}

	//nolint:exhaustruct
	ref := &jsonschema.Reflector{
		FieldNameTag:               "yaml",
		RequiredFromJSONSchemaTags: true,
	}

	//nolint:exhaustruct
	schema := ref.Reflect(&model.Config{})

	schemaStr, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		panic(err)
	}

	_, err = file.Write(schemaStr)
	if err != nil {
		panic(err)
	}
}
