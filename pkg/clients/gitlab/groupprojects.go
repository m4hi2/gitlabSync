package gitlab

import (
	"encoding/json"
	"fmt"
	"github.com/m4hi2/gitlabSync/pkg/utils/glserrors"
	"io"
	"net/http"
)

type GroupRespPld struct {
	ID       int        `json:"id"`
	Path     string     `json:"path"`
	Projects []*Project `json:"projects"`
}
type Project struct {
	Path          string `json:"path"`
	SSHURLToRepo  string `json:"ssh_url_to_repo"`
	HTTPURLToRepo string `json:"http_url_to_repo"`
}

func (c *Client) buildGroupProjectGetURL(groupID int) string {
	url := fmt.Sprintf("https://%s/%s/groups/%d", c.GitlabHost, BasePath, groupID)
	url = c.addTokenParam(url)

	return url
}

func (c *Client) GetGroupData(groupID int) (*GroupRespPld, error) {
	gpURL := c.buildGroupProjectGetURL(groupID)
	req, err := http.NewRequest("GET", gpURL, nil)
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, glserrors.PassExecutionErr(ErrCallingGitlab, err)
	}

	if resp.StatusCode != http.StatusOK {
		err := httpResponseErrorParse(resp.StatusCode)
		if err != nil {
			return nil, err
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, glserrors.PassExecutionErr(ErrReadingBody, err)
	}

	groupReps := &GroupRespPld{}
	err = json.Unmarshal(body, groupReps)
	if err != nil {
		return nil, glserrors.PassExecutionErr(ErrUnmarshalError, err)
	}

	return groupReps, nil
}
