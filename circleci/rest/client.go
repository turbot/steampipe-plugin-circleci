package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	baseURL     *url.URL
	circleToken string
	client      *http.Client
}

type Config struct {
	URL   string
	Token string
}

func New(config Config) *Client {
	u, _ := url.Parse(config.URL)

	// Ensure endpoint ends with a slash
	endpoint := u.Path
	if !strings.HasSuffix(endpoint, "/") {
		endpoint += "/"
	}

	return &Client{
		baseURL:     u.ResolveReference(&url.URL{Path: endpoint}),
		circleToken: config.Token,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) NewRequest(method string, u *url.URL, payload interface{}) (req *http.Request, err error) {
	var r io.Reader
	if payload != nil {
		buf := &bytes.Buffer{}
		r = buf
		err = json.NewEncoder(buf).Encode(payload)
		if err != nil {
			return nil, err
		}
	}

	req, err = http.NewRequest(method, c.baseURL.ResolveReference(u).String(), r)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Circle-Token", c.circleToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "mrolla/terraform-provider-circleci")
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (c *Client) DoRequest(req *http.Request, resp interface{}) (statusCode int, err error) {
	httpResp, err := c.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode >= 300 {
		httpError := struct {
			Message string `json:"message"`
		}{}
		err = json.NewDecoder(httpResp.Body).Decode(&httpError)
		if err != nil {
			return httpResp.StatusCode, err
		}
		return httpResp.StatusCode, &HTTPError{Code: httpResp.StatusCode, Message: httpError.Message}
	}

	if resp != nil {
		// REST API v2 does not currently return Content-Type

		// if !strings.Contains(httpResp.Header.Get("Content-Type"), "application/json") {
		// 	return httpResp.StatusCode, errors.New("wrong content type received")
		// }

		err = json.NewDecoder(httpResp.Body).Decode(resp)
		if err != nil {
			return httpResp.StatusCode, err
		}
	}
	return httpResp.StatusCode, nil
}

type HTTPError struct {
	Code    int
	Message string
}

func (e *HTTPError) Error() string {
	if e.Code == 0 {
		e.Code = http.StatusInternalServerError
	}
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprintf("response %d (%s)", e.Code, http.StatusText(e.Code))
}

func (c *Client) HasProjectEnvironmentVariable(vcs, org, project, name string) (bool, error) {
	slug := Slug(vcs, org, project)
	u := &url.URL{
		Path: fmt.Sprintf("project/%s/envvar/%s", slug, name),
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return false, err
	}

	_, err = c.DoRequest(req, nil)
	if err != nil {
		var httpError *HTTPError
		if errors.As(err, &httpError) && httpError.Code == 404 {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (c *Client) ListPipelines(vcs, org string) (*PipelineResponse, error) {
	slug := OrgSlug(vcs, org)
	fmt.Println(slug)

	u := &url.URL{
		Path: "pipeline",
	}
	values := u.Query()
	values.Add("org-slug", slug)
	values.Add("mine", "true")
	u.RawQuery = values.Encode()

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	pipelineResp := &PipelineResponse{}

	_, err = c.DoRequest(req, pipelineResp)
	if err != nil {
		return nil, err
	}

	return pipelineResp, nil
}

// Slug returns a project slug, including the VCS, organization, and project names
func Slug(vcs, org, project string) string {
	return fmt.Sprintf("%s/%s/%s", vcs, org, project)
}

func OrgSlug(vcs, org string) string {
	return fmt.Sprintf("%s/%s", vcs, org)
}
