/* ex: set filetype=go fenc=utf-8 noexpandtab ts=4 sw=4 : */
package extractor

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/shk3bq4d/terraform-provider-graylog/graylog/convert"
	"github.com/shk3bq4d/terraform-provider-graylog/graylog/util"
)

const (
	keyInputID         = "input_id"
	keyExtractorID     = "extractor_id"
	keyExtractorConfig = "extractor_config"
	keyID              = "id"
	keyConfig          = "config"
	keyConverters      = "converters"
	keyType            = "type"
)

func getDataFromResourceData(d *schema.ResourceData) (map[string]interface{}, error) {
	log.Printf("shk3bq4d getDataFromResourceData 0")
	data, err := convert.GetFromResourceData(d, Resource())
	if err != nil {
		return nil, err
	}
	//util.RenameKey(data, "cursor_strategy", "cut_or_copy")
	util.RenameKey(data, "type", "extractor_type")
	util.SetDefaultValue(data, "target_field", "")
	util.SetDefaultValue(data, "condition_value", "")

	if err := convert.JSONToData(data, keyExtractorConfig); err != nil {
		return nil, err
	}
	log.Printf("shk3bq4d getDataFromResourceData 15")
	util.RenameKey(data, keyExtractorID, keyID)

	util.MrDebug(data)
	/*
		converters := convert.ListToMap(data[keyConverters].([]interface{}), keyType)
		for k, v := range converters {
			converters[k] = v.(map[string]interface{})[keyConfig]
		}
		if err := convert.JSONToData(converters); err != nil {
			return nil, err
		}
		data[keyConverters] = converters
		util.MrDebug(data)
	// */

	/*
		for _, v := range data[keyConverters].([]interface{}) {
			if err := convert.JSONToData(v.(map[string]interface{}), keyConfig); err != nil {
				log.Printf("shk3bq4d getDataFromResourceData 16")
				return nil, err
			}
		}
		util.MrDebug(data)
	//	*/
	/*
		converters := data[keyConverters].([]interface{})
		for k, v := range converters {
			converters[k] = v.(map[string]interface{})[keyConfig]
		}
		if err := convert.JSONToData(converters.(map[string]interface{})); err != nil {
			return nil, err
		}
		data[keyConverters] = converters
		util.MrDebug(data)
	//	*/
	log.Printf("shk3bq4d getDataFromResourceData 17")

	return data, nil
}

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	log.Printf("shk3bq4d setDataToResourceData")
	if err := convert.DataToJSON(data, keyExtractorConfig); err != nil {
		log.Printf("shk3bq4d setDataToResourceData 2")
		return err
	}
	log.Printf("shk3bq4d setDataToResourceData 3")
	util.RenameKey(data, keyID, keyExtractorID)
	log.Printf("shk3bq4d setDataToResourceData 4")

	converters := data[keyConverters].([]interface{})
	log.Printf("shk3bq4d setDataToResourceData 5")
	/*
		for i, a := range converters {
			log.Printf("shk3bq4d setDataToResourceData 6")
			elem := a.(map[string]interface{})
			log.Printf("shk3bq4d setDataToResourceData 7")
			b, err := json.Marshal(elem[keyConfig])
			if err != nil {
				log.Printf("shk3bq4d setDataToResourceData 8")
				return err
			}
			log.Printf("shk3bq4d setDataToResourceData 9")
			elem[keyConfig] = string(b)
			log.Printf("shk3bq4d setDataToResourceData 10")
			converters[i] = elem
			log.Printf("shk3bq4d setDataToResourceData 11")
		}
	*/

	log.Printf("shk3bq4d setDataToResourceData 12")
	data[keyConverters] = converters
	log.Printf("shk3bq4d setDataToResourceData 13")

	if err := convert.SetResourceData(d, Resource(), data); err != nil {
		log.Printf("shk3bq4d setDataToResourceData 14")
		return err
	}

	log.Printf("shk3bq4d setDataToResourceData 15")
	d.SetId(d.Get(keyInputID).(string) + "/" + d.Get(keyExtractorID).(string))
	log.Printf("shk3bq4d setDataToResourceData 16")
	return nil
}
