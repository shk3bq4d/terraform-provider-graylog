package extractor

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/one-2-one/terraform-provider-graylog/graylog/client"
)

func create(d *schema.ResourceData, m interface{}) error {
	log.Printf("shk3bq4d create 0")
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}

	log.Printf("shk3bq4d create 1")
	data, err := getDataFromResourceData(d)
	log.Printf("shk3bq4d create 2")
	if err != nil {
		log.Printf("shk3bq4d create 3")
		return err
	}

	log.Printf("shk3bq4d create 4")
	inputID := data[keyInputID].(string)
	delete(data, keyInputID)
	delete(data, keyID)

	log.Printf("shk3bq4d create 5")
	ac, _, err := cl.Extractor.Create(ctx, inputID, data)
	log.Printf("shk3bq4d create 6")
	if err != nil {
		log.Printf("shk3bq4d create 7")
		return fmt.Errorf("failed to create a extractor (input id: %s): %w", inputID, err)
	}
	log.Printf("shk3bq4d create 8")
	acID := ac[keyExtractorID].(string)
	log.Printf("shk3bq4d create 9")
	if err := d.Set(keyExtractorID, acID); err != nil {
		log.Printf("shk3bq4d create 10")
		return err
	}
	log.Printf("shk3bq4d create 11")
	d.SetId(inputID + "/" + acID)
	log.Printf("shk3bq4d create 12")
	return nil
}
