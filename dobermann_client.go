package dobermann_sdk

import (
	"fmt"
	"github.com/extraktor/dobermann-sdk/session"
	"github.com/extraktor/dobermann-sdk/settings"
)

func CreateClient(config settings.Setup, input interface{}) (interface{}, error) {
	result, err := session.Session(config, settings.POST, "/v3/customers", input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetOneClient(config settings.Setup, id string) (interface{}, error) {
	endpoint := fmt.Sprintf("/v3/customers/%s", id)
	result, err := session.Session(config, settings.GET, endpoint, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateClient(config settings.Setup, id string, input interface{}) (interface{}, error) {
	endpoint := fmt.Sprintf("/v3/customers/%s", id)
	result, err := session.Session(config, settings.PUT, endpoint, input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteClient(config settings.Setup, id string) error {
	endpoint := fmt.Sprintf("/v3/customers/%s", id)
	_, err := session.Session(config, settings.DELETE, endpoint, nil)
	if err != nil {
		return err
	}
	return nil
}
