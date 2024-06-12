package dobermann_sdk

import (
	"github.com/extraktor/dobermann-sdk/session"
	"github.com/extraktor/dobermann-sdk/settings"
)

func CreateClient(config settings.Setup, input interface{}) (interface{}, error) {
	result, err := session.Session(config, settings.GET, "/v3/customers", input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetOneClient(config settings.Setup, input interface{}) (error, interface{}) {
	return nil, nil
}

func GetListClient(config settings.Setup, input interface{}) (error, interface{}) {
	return nil, nil
}

func UpdateClient(config settings.Setup, input interface{}) (error, interface{}) {
	return nil, nil
}

func DeleteClient(config settings.Setup, input interface{}) error {
	return nil
}
