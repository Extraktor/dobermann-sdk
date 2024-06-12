package dobermann_sdk

import (
	"github.com/extraktor/dobermann-sdk/settings"
)

func Start(token string, environment settings.Environment) settings.Configuration {
	config := settings.Configuration{Token: token, Environment: environment}
	return config
}
