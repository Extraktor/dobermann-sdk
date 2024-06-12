package dobermann_sdk

import (
	"github.com/extraktor/dobermann-sdk/settings"
)

func Start(token string, environment settings.Environment) settings.Setup {
	config := settings.Setup{Token: token, Environment: environment}
	return config
}
