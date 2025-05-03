package user

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/suzuki-shunsuke/flute/v2/flute"

	"github.com/shk3bq4d/terraform-provider-graylog/graylog/testutil"
)

func TestAccUser(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	getUserRoute := flute.Route{
		Name: "get user",
		Matcher: flute.Matcher{
			Method: "GET",
		},
		Tester: flute.Tester{
			Path:         "/api/users/devops",
			PartOfHeader: testutil.Header(),
		},
		Response: flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(strings.NewReader(`{
  "id": "11f3f1792383646ecd354f8a",
  "username": "devops",
  "email": "devops@test.com",
  "first_name": null,
  "last_name": null,
  "full_name": "DevOps",
  "permissions": [ "users:edit:devops", "streams:read:000000000000000000000001" ],
  "grn_permissions": [ "entity:own:grn::::stream:11cf4cbacadd9c316802221b" ],
  "preferences": { "updateUnfocussed": false, "enableSmartSearch": true },
  "timezone": null,
  "session_timeout_ms": 28800000,
  "external": true,
  "roles": [ "Admin", "Reader", "Views Manager" ],
  "read_only": false,
  "session_active": true,
  "last_activity": "2020-01-16T09:00:00.631+0000",
  "client_address": "10.1.1.3",
  "account_status": "enabled",
  "auth_service_enabled": true,
  "service_account": false
}`)),
				}, nil
			},
		},
	}

	readStep := resource.TestStep{
		ResourceName: "data.graylog_user.test",
		PreConfig: func() {
			testutil.SetHTTPClient(t, getUserRoute)
		},
		Config: `
data "graylog_user" "test" {
  username = "devops"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.graylog_user.test", "id", "11f3f1792383646ecd354f8a"),
			resource.TestCheckResourceAttr("data.graylog_user.test", "username", "devops"),
			resource.TestCheckResourceAttr("data.graylog_user.test", "email", "devops@test.com"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleDataSourceProviders("graylog_user", DataSource()),
		Steps: []resource.TestStep{
			readStep,
		},
	})
}
