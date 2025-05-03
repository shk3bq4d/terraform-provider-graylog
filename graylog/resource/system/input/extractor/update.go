package extractor

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/one-2-one/terraform-provider-graylog/graylog/client"
)

func update(d *schema.ResourceData, m interface{}) error {
	log.Printf("shk3bq4d update 0")
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}

	data, err := getDataFromResourceData(d)
	if err != nil {
		return err
	}

	delete(data, keyInputID)
	delete(data, keyID)

	inputID := d.Get(keyInputID).(string)
	eID := d.Get(keyExtractorID).(string)
	if _, _, err := cl.Extractor.Update(ctx, inputID, eID, data); err != nil {
		return fmt.Errorf("failed to update a extractor (input id: %s, id: %s): %w", inputID, eID, err)
	}
	return nil
}
