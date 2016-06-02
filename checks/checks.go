package checks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Check holds the info to be used by the functions
type Check struct {
	server   string
	port     int
	vhost    string
	ssl      bool
	user     string
	password string
}

// New returns a Check struct
func New(ssl bool, server string, port int, vhost, user, password string) *Check {
	return &Check{
		server:   server,
		port:     port,
		vhost:    vhost,
		ssl:      ssl,
		user:     user,
		password: password,
	}
}

func (c *Check) request(method, endpoint string, object interface{}) error {

	userInfo := url.UserPassword(c.user, c.password)

	urlObj, err := url.Parse(c.url(endpoint))
	urlObj.User = userInfo

	fmt.Println(urlObj.String())

	req, err := http.NewRequest(method, urlObj.String(), nil)
	if err != nil {
		return err
	}

	httpClient := http.DefaultClient

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	objectByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if object != nil {
		if err = json.Unmarshal(objectByte, &object); err != nil {
			return err
		}
	}

	return nil
}

func (c *Check) url(endpoint string) string {

	var (
		urlString string
	)

	if c.vhost == "" {
		urlString = fmt.Sprintf("://%s:%d/api/%s", c.server, c.port, endpoint)
	} else {
		urlString = fmt.Sprintf("://%s:%d/api/%s/%s", c.server, c.port, c.vhost, endpoint)
	}

	if c.ssl {
		fmt.Printf("https%s\n", urlString)
		return fmt.Sprintf("https%s", urlString)
	}
	fmt.Printf("http%s\n", urlString)
	return fmt.Sprintf("http%s", urlString)

}
