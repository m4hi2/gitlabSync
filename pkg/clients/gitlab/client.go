package gitlab

import (
	"fmt"
	"github.com/m4hi2/gitlabSync/pkg/utils/httpcaller"
)

type Client struct {
	GitlabHost  string
	GitlabToken string
	HttpClient  httpcaller.HTTPCaller
}

const BasePath = "api/v4"

func (c *Client) addTokenParam(url string) string {
	return fmt.Sprintf("%s?private_token=%s", url, c.GitlabToken)
}
