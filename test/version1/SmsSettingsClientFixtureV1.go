package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-users2/client-smssettings-go/version1"
	"github.com/stretchr/testify/assert"
)

var SETTINGS = &version1.SmsSettingsV1{
	Id:       "1",
	Name:     "User 1",
	Phone:    "+12345678902",
	Language: "en",
	Verified: false,
}

type SmsSettingsClientFixtureV1 struct {
	Client version1.ISmsSettingsClientV1
}

func NewSessionsClientFixtureV1(client version1.ISmsSettingsClientV1) *SmsSettingsClientFixtureV1 {
	return &SmsSettingsClientFixtureV1{
		Client: client,
	}
}

func (c *SmsSettingsClientFixtureV1) TestCrudOperations(t *testing.T) {

	var settings1 *version1.SmsSettingsV1

	// Create sms settings
	settings, err := c.Client.SetSettings(context.Background(), "123", SETTINGS)

	assert.Nil(t, err)
	assert.NotNil(t, settings)
	assert.Equal(t, SETTINGS.Id, settings.Id)
	assert.Equal(t, SETTINGS.Phone, settings.Phone)
	assert.Equal(t, SETTINGS.Verified, settings.Verified)

	settings1 = settings

	// Create sms settings
	settings, err = c.Client.SetVerifiedSettings(context.Background(), "123", settings1)

	assert.Nil(t, err)
	assert.NotNil(t, settings)
	assert.Equal(t, SETTINGS.Id, settings.Id)
	assert.Equal(t, SETTINGS.Phone, settings.Phone)
	assert.True(t, settings.Verified)

	settings1 = settings

	// Update the sms settings
	settings1.Subscriptions = map[string]any{"engagement": true}

	settings, err = c.Client.SetSettings(context.Background(), "123", settings1)

	assert.Nil(t, err)
	assert.True(t, settings.Subscriptions.(map[string]any)["engagement"].(bool))

	settings1 = settings

	// Get settings
	settingsList, err := c.Client.GetSettingsByIds(context.Background(), "123", []string{settings1.Id})

	assert.Nil(t, err)
	assert.Len(t, settingsList, 1)

	// Delete settings
	err = c.Client.DeleteSettingsById(context.Background(), "123", settings.Id)
	assert.Nil(t, err)

	// Try to get deleted settings
	settings, err = c.Client.GetSettingsById(context.Background(), "123", settings1.Id)
	assert.Nil(t, err)

	assert.Nil(t, settings)
}
