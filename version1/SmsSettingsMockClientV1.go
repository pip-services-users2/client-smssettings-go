package version1

import (
	"context"
)

type SmsSettingsMockClientV1 struct {
	settings []*SmsSettingsV1
}

func NewSmsSettingsMockClientV1() *SmsSettingsMockClientV1 {
	return &SmsSettingsMockClientV1{
		settings: make([]*SmsSettingsV1, 0),
	}
}

func (c *SmsSettingsMockClientV1) GetSettingsByIds(ctx context.Context, correlationId string, recipientIds []string) ([]*SmsSettingsV1, error) {
	settings := make([]*SmsSettingsV1, 0)

	for _, setting := range c.settings {
		for _, id := range recipientIds {
			if id == setting.Id {
				settings = append(settings, setting)
				break
			}
		}
	}

	return settings, nil
}

func (c *SmsSettingsMockClientV1) GetSettingsById(ctx context.Context, correlationId string, recipientId string) (*SmsSettingsV1, error) {
	var settings *SmsSettingsV1

	for _, s := range c.settings {
		if s.Id == recipientId {
			settings = s
			break
		}
	}

	return settings, nil
}

func (c *SmsSettingsMockClientV1) GetSettingsByPhoneSettings(ctx context.Context, correlationId string, phone string) (*SmsSettingsV1, error) {
	var settings *SmsSettingsV1

	for _, s := range c.settings {
		if s.Phone == phone {
			settings = s
			break
		}
	}

	return settings, nil
}

func (c *SmsSettingsMockClientV1) SetSettings(ctx context.Context, correlationId string, settings *SmsSettingsV1) (*SmsSettingsV1, error) {
	settings.Verified = false

	for i, s := range c.settings {
		if s.Id == settings.Id {
			c.settings = append(c.settings[:i], c.settings[i+1:]...)
			break
		}
	}

	c.settings = append(c.settings, settings)

	return settings, nil
}

func (c *SmsSettingsMockClientV1) SetVerifiedSettings(ctx context.Context, correlationId string, settings *SmsSettingsV1) (*SmsSettingsV1, error) {
	settings.Verified = true

	for i, s := range c.settings {
		if s.Id == settings.Id {
			c.settings = append(c.settings[:i], c.settings[i+1:]...)
			break
		}
	}

	c.settings = append(c.settings, settings)

	return settings, nil
}

func (c *SmsSettingsMockClientV1) SetRecipient(ctx context.Context, correlationId string, recipientId string, name string, phone string, language string) (*SmsSettingsV1, error) {
	var settings *SmsSettingsV1

	for _, s := range c.settings {
		if s.Id == recipientId {
			settings = s
			break
		}
	}

	if settings != nil {
		settings.Name = name
		settings.Phone = phone
		settings.Language = language
	} else {
		settings = &SmsSettingsV1{
			Id:       recipientId,
			Name:     name,
			Phone:    phone,
			Language: language,
			Verified: false,
		}
		c.settings = append(c.settings, settings)
	}

	return settings, nil
}

func (c *SmsSettingsMockClientV1) SetSubscriptions(ctx context.Context, correlationId string, recipientId string, subscriptions any) (*SmsSettingsV1, error) {
	var settings *SmsSettingsV1

	for _, s := range c.settings {
		if s.Id == recipientId {
			settings = s
			break
		}
	}

	if settings != nil {
		settings.Subscriptions = subscriptions
	} else {
		settings = &SmsSettingsV1{
			Id:            recipientId,
			Name:          "",
			Phone:         "",
			Language:      "",
			Subscriptions: subscriptions,
		}
		c.settings = append(c.settings, settings)
	}

	return settings, nil
}

func (c *SmsSettingsMockClientV1) DeleteSettingsById(ctx context.Context, correlationId string, recipientId string) error {
	for index, s := range c.settings {
		if s.Id == recipientId {
			c.settings = append(c.settings[:index], c.settings[index+1:]...)
			break
		}
	}

	return nil
}

func (c *SmsSettingsMockClientV1) ResendVerification(ctx context.Context, correlationId string, recipientId string) error {
	return nil
}

func (c *SmsSettingsMockClientV1) VerifyPhone(ctx context.Context, correlationId string, recipientId string, code string) error {
	var settings *SmsSettingsV1
	for _, s := range c.settings {
		if s.Id == recipientId {
			settings = s
			break
		}
	}

	if settings != nil && settings.VerCode == code {
		settings.Verified = true
		settings.VerCode = ""
	}

	return nil
}
