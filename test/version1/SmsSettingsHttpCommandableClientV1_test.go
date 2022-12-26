package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-smssettings-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type smsSettingsHttpCommandableClientV1Test struct {
	client  *version1.SmsSettingsCommandableHttpClientV1
	fixture *SmsSettingsClientFixtureV1
}

func newSmsSettingsHttpCommandableClientV1Test() *smsSettingsHttpCommandableClientV1Test {
	return &smsSettingsHttpCommandableClientV1Test{}
}

func (c *smsSettingsHttpCommandableClientV1Test) setup(t *testing.T) *SmsSettingsClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewSmsSettingsCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewSessionsClientFixtureV1(c.client)

	return c.fixture
}

func (c *smsSettingsHttpCommandableClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestHttpOpenSession(t *testing.T) {
	c := newSmsSettingsHttpCommandableClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
