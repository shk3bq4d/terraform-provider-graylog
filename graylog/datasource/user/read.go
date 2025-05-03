package user

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/shk3bq4d/terraform-provider-graylog/graylog/client"
)

func read(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}

	if username, ok := d.GetOk(keyUsername); ok {
		return readByUsername(ctx, d, cl, username.(string))
	}

	return errors.New("username must not exist")
}

func readByUsername(ctx context.Context, d *schema.ResourceData, cl client.Client, username string) error {
	data, resp, err := cl.UserOutput.GetUserByUsername(ctx, username)
	if err != nil {
		if resp != nil {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			fmt.Printf("Error response: %s\n", bodyBytes)
		}
		return err
	}
	return setDataToResourceData(d, data)
}
