package build

import (
	clients1 "github.com/pip-services-users2/client-smssettings-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type SmsSettingsClientFactory struct {
	*cbuild.Factory
}

func NewSmsSettingsClientFactory() *SmsSettingsClientFactory {
	c := &SmsSettingsClientFactory{
		Factory: cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("service-smssettings", "client", "null", "*", "1.0")
	// directClientDescriptor := cref.NewDescriptor("service-smssettings", "client", "direct", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-smssettings", "client", "commandable-http", "*", "1.0")
	mockClientDescriptor := cref.NewDescriptor("service-smssettings", "client", "mock", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewSmsSettingsNullClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewSmsSettingsCommandableHttpClientV1)
	c.RegisterType(mockClientDescriptor, clients1.NewSmsSettingsMockClientV1)

	return c
}
