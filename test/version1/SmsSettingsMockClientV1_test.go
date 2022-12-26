package test_version1

import (
	"testing"

	"github.com/pip-services-users2/client-smssettings-go/version1"
)

type smsSettingsMockClientV1Test struct {
	client  *version1.SmsSettingsMockClientV1
	fixture *SmsSettingsClientFixtureV1
}

func NewSmsSettingsMockClientV1Test() *smsSettingsMockClientV1Test {
	return &smsSettingsMockClientV1Test{}
}

func (c *smsSettingsMockClientV1Test) setup(t *testing.T) *SmsSettingsClientFixtureV1 {

	c.client = version1.NewSmsSettingsMockClientV1()
	c.fixture = NewSessionsClientFixtureV1(c.client)
	return c.fixture
}

func (c *smsSettingsMockClientV1Test) teardown(t *testing.T) {
	c.client = nil
	c.fixture = nil
}

func TestMockOpenSession(t *testing.T) {
	c := NewSmsSettingsMockClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
