/* ex: set filetype=go fenc=utf-8 noexpandtab ts=4 sw=4 : */
package extractor

import (
	"context"
	"github.com/shk3bq4d/terraform-provider-graylog/graylog/util"
	"github.com/suzuki-shunsuke/go-httpclient/httpclient"
	"log"
	"net/http"
)

type Client struct {
	Client httpclient.Client
}

func (cl Client) Get(ctx context.Context, inputID, id string) (map[string]interface{}, *http.Response, error) {
	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "GET",
		Path:         "/system/inputs/" + inputID + "/extractors/" + id,
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Create(
	ctx context.Context, inputID string, data map[string]interface{},
) (map[string]interface{}, *http.Response, error) {
	log.Printf("shk3bq4d client create 0")
	body := map[string]interface{}{}
	url := "/system/inputs/" + inputID + "/extractors"
	log.Printf("shk3bq4d Data for POST %s", url)
	util.MrDebug(data)
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "POST",
		Path:         url,
		RequestBody:  data,
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Update(
	ctx context.Context, inputID, id string, data map[string]interface{},
) (map[string]interface{}, *http.Response, error) {
	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "PUT",
		Path:         "/system/inputs/" + inputID + "/extractors/" + id,
		RequestBody:  data,
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Delete(ctx context.Context, inputID, id string) (*http.Response, error) {
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method: "DELETE",
		Path:   "/system/inputs/" + inputID + "/extractors/" + id,
	})
	return resp, err
}
