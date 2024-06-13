package session

import (
	"errors"
	"github.com/extraktor/dobermann-sdk/settings"
	"github.com/monaco-io/request"
)

func Session(config settings.Setup, method settings.Method, endpoint string, body interface{}) (interface{}, error) {
	c := request.Client{
		URL:    config.Environment.String() + endpoint,
		Method: method.String(),
		Header: map[string]string{"Content-Type": "application/json", "access_token": config.Token},
	}

	if method.String() != settings.GET.String() && body != nil {
		c.JSON = body
	}

	var result interface{}
	resp := c.Send().Scan(&result)

	switch resp.Code() {
	case 200:
		return result, nil
	case 400:
		return nil, errors.New("Bad Request")
	case 401:
		return nil, errors.New("Unauthorized")
	case 403:
		return nil, errors.New("Forbidden")
	case 404:
		return nil, errors.New("Not Found")
	case 429:
		return nil, errors.New("Too Many Requests")
	case 500:
		return nil, errors.New("Internal Server Error")
	default:
		return nil, nil
	}
}
