package gitlab

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"os"
	"testing"
)

type MockHttpClient struct{}

func (m MockHttpClient) Do(r *http.Request) (*http.Response, error) {
	if r.URL.Path == "/api/v4/groups/59" {
		if r.Method != http.MethodGet {
			return nil, fmt.Errorf("unexpected method: %s", r.Method)
		}
		if r.URL.RawQuery != "private_token=youThinkIPutToken" {
			return nil, fmt.Errorf("unexpected query: %s", r.URL.RawQuery)
		}
		if r.Host != "git.iammahir.com" {
			return nil, fmt.Errorf("unexpected host: %s", r.Host)
		}
		// Sorry, can not provide data, since working with real private gitlab instance
		data, err := os.Open("../../../gg/veritas.json")
		if err != nil {
			return nil, err
		}
		return &http.Response{
			Status:     "200 OK",
			StatusCode: http.StatusOK,
			Proto:      "HTTP/1.0",
			ProtoMajor: 1,
			ProtoMinor: 1,
			Body:       data,
		}, nil
	}

	return nil, fmt.Errorf("unexpected path: %s", r.URL.Path)

}

func GetMockClient() *Client {
	return &Client{
		GitlabHost:  "git.iammahir.com",
		GitlabToken: "youThinkIPutToken",
		HttpClient:  &MockHttpClient{},
	}
}

func TestClient_CheckGettingGroupData(t *testing.T) {
	c := GetMockClient()
	data, err := c.GetGroupData(59)
	if err != nil {
		log.Fatal("can not get data from gitlab: ", err)
	}

	assert.Equal(t, data.ID, 59, "Group ID miss match")
	assert.Equal(t, data.Path, "services", "Path miss mismatch")
}
