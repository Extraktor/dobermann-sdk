package dobermann_sdk

import (
	"github.com/extraktor/dobermann-sdk/session"
	"github.com/extraktor/dobermann-sdk/settings"
)

func CreateClient(config settings.Configuration, input map[string]interface{}) (interface{}, error) {
	result, err := session.Session(config, settings.GET, "/v3/customers", input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetOneClient(config settings.Configuration, input interface{}) (error, interface{}) {
	return nil, nil
}

func GetListClient(config settings.Configuration, input map[string]interface{}) (error, interface{}) {
	return nil, nil
}

func UpdateClient(config settings.Configuration, input map[string]interface{}) (error, interface{}) {
	return nil, nil
}

func DeleteClient(config settings.Configuration, input map[string]interface{}) error {
	return nil
}
