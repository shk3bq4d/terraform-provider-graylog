package extractor

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/one-2-one/terraform-provider-graylog/graylog/convert"
)

const schemaVersion = 1

var stateUpgraders = []schema.StateUpgrader{
	stateUpgraderV1,
}

func extractorResourceV0() *schema.Resource {
	return &schema.Resource{}
}

var stateUpgraderV1 = schema.StateUpgrader{
	Version: 0,
	Type:    extractorResourceV0().CoreConfigSchema().ImpliedType(),
	Upgrade: func(_ context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
		log.Printf("shk3bq4d state_upgrader 0")
		inputID := rawState[keyInputID].(string)
		extractorID := rawState[keyID].(string)

		rawState[keyExtractorID] = extractorID
		rawState[keyID] = inputID + "/" + extractorID

		attrTypes := []string{
			"grok_type_extractor_config",
			"json_type_extractor_config",
			"regex_type_extractor_config",
		}

		generalAttrTypes := []string{
			"general_int_extractor_config",
			"general_bool_extractor_config",
			"general_float_extractor_config",
			"general_string_extractor_config",
		}

		attrs := map[string]interface{}{}

		for _, a := range attrTypes {
			v, ok := rawState[a]
			if !ok || v == nil {
				continue
			}
			arr := v.([]interface{})
			if len(arr) == 0 {
				continue
			}
			for k, attr := range arr[0].(map[string]interface{}) {
				attrs[k] = attr
			}
			delete(rawState, a)
		}

		for _, a := range generalAttrTypes {
			v, ok := rawState[a]
			if !ok || v == nil {
				continue
			}
			for k, attr := range v.(map[string]interface{}) {
				attrs[k] = attr
			}
			delete(rawState, a)
		}

		b, err := json.Marshal(attrs)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal attributes '%s' as JSON: %w", keyConfig, err)
		}
		rawState[keyExtractorConfig] = string(b)

		log.Printf("shk3bq4d state_upgrader 1")
		if a, ok := rawState[keyConverters]; ok {
			log.Printf("shk3bq4d state_upgrader 2")
			list := a.([]interface{})
			log.Printf("shk3bq4d state_upgrader 3")
			for i, e := range list {
				log.Printf("shk3bq4d state_upgrader 4")
				if err := convert.OneSizeListToJSON(e.(map[string]interface{}), keyConfig); err != nil {
					log.Printf("shk3bq4d state_upgrader 5")
					return nil, err
				}
				log.Printf("shk3bq4d state_upgrader 6")
				list[i] = e
				log.Printf("shk3bq4d state_upgrader 7")
			}
			log.Printf("shk3bq4d state_upgrader 8")
			rawState[keyConverters] = list
			log.Printf("shk3bq4d state_upgrader 9")
		}

		log.Printf("shk3bq4d state_upgrader 10")
		return rawState, nil
	},
}
