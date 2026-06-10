// Tests for the hand-patched UnmarshalJSON in
// model_get_application_by_id_200_response.go. See P81-126016 / OPEN-04.
//
// The fixtures below mirror real server responses observed during the
// 2026-06-05 regression of CheckPoint/checkpoint-sase. Without the patch,
// every fixture would fail to decode with
//   "data matches more than one schema in oneOf(GetApplicationById200Response)"
// because the swagger spec lacks `discriminator: { propertyName: type }`.

package perimeter81sdk

import (
	"encoding/json"
	"strings"
	"testing"
)

const httpsApplicationFixture = `{
  "id":"XmUkDLGkVY",
  "name":"tfCrudApp",
  "type":"https",
  "host":{"source":"fixed","value":"internal.example.com"},
  "port":{"source":"fixed","value":443},
  "enabled":true,
  "alias":{"aliasEnabled":false},
  "attributes":{"checkStatus":200,"sslCertificateVerification":true},
  "auth":{},
  "displayIconAtLogin":true,
  "fqdn":"XmUkDLGkVY.pzero.perimeter81.biz",
  "network":{"id":"O5QGwEUsNL","name":"tfCrudNetUpd7","dns":"oleksandrc-test1-child1-o5qgweusnl.pzero.perimeter81.biz","subnet":"10.200.0.0/16"},
  "users":[{"id":"YjTMdHR2cf","fullName":"","username":"test2@checkpoint.com"}],
  "icon":{"name":"applications/https.svg","url":"https://static.perimeter81.biz/api/files/applications/https.svg"},
  "createdAt":"2026-06-04T13:01:22.139Z",
  "updatedAt":"2026-06-04T13:01:22.140Z"
}`

func TestGetApplicationById200Response_UnmarshalJSON_HttpsApplication(t *testing.T) {
	var resp GetApplicationById200Response
	if err := json.Unmarshal([]byte(httpsApplicationFixture), &resp); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.HttpsApplication == nil {
		t.Fatal("expected HttpsApplication to be populated")
	}
	if resp.HttpApplication != nil || resp.RdpApplication != nil || resp.SshApplication != nil || resp.VncApplication != nil {
		t.Fatalf("expected only HttpsApplication to be populated, got %+v", resp)
	}
	if resp.HttpsApplication.Id != "XmUkDLGkVY" {
		t.Errorf("Id = %q, want %q", resp.HttpsApplication.Id, "XmUkDLGkVY")
	}
	if resp.HttpsApplication.Name != "tfCrudApp" {
		t.Errorf("Name = %q, want %q", resp.HttpsApplication.Name, "tfCrudApp")
	}
	if resp.HttpsApplication.Type != "https" {
		t.Errorf("Type = %q, want %q", resp.HttpsApplication.Type, "https")
	}
	if resp.HttpsApplication.Network.Id != "O5QGwEUsNL" {
		t.Errorf("Network.Id = %q, want %q", resp.HttpsApplication.Network.Id, "O5QGwEUsNL")
	}
}

func TestGetApplicationById200Response_UnmarshalJSON_DispatchesByType(t *testing.T) {
	cases := []struct {
		typeValue string
		check     func(t *testing.T, r GetApplicationById200Response)
	}{
		{"http", func(t *testing.T, r GetApplicationById200Response) {
			if r.HttpApplication == nil {
				t.Fatal("expected HttpApplication to be populated for type=http")
			}
			if r.HttpsApplication != nil || r.RdpApplication != nil || r.SshApplication != nil || r.VncApplication != nil {
				t.Fatal("expected only HttpApplication to be populated")
			}
		}},
		{"https", func(t *testing.T, r GetApplicationById200Response) {
			if r.HttpsApplication == nil {
				t.Fatal("expected HttpsApplication to be populated for type=https")
			}
		}},
		{"rdp", func(t *testing.T, r GetApplicationById200Response) {
			if r.RdpApplication == nil {
				t.Fatal("expected RdpApplication to be populated for type=rdp")
			}
		}},
	}
	for _, tc := range cases {
		t.Run(tc.typeValue, func(t *testing.T) {
			body := strings.Replace(httpsApplicationFixture, `"type":"https"`, `"type":"`+tc.typeValue+`"`, 1)
			var resp GetApplicationById200Response
			if err := json.Unmarshal([]byte(body), &resp); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			tc.check(t, resp)
		})
	}
}

func TestGetApplicationById200Response_UnmarshalJSON_UnknownTypeFails(t *testing.T) {
	body := strings.Replace(httpsApplicationFixture, `"type":"https"`, `"type":"smb"`, 1)
	var resp GetApplicationById200Response
	err := json.Unmarshal([]byte(body), &resp)
	if err == nil {
		t.Fatal("expected error for unknown application type")
	}
	if !strings.Contains(err.Error(), "unknown application type") {
		t.Errorf("error = %q, want it to mention 'unknown application type'", err.Error())
	}
}

func TestGetApplicationById200Response_UnmarshalJSON_MissingTypeFails(t *testing.T) {
	body := `{"id":"x","name":"y"}`
	var resp GetApplicationById200Response
	err := json.Unmarshal([]byte(body), &resp)
	if err == nil {
		t.Fatal("expected error when type discriminator is missing")
	}
}
