package version1

import "context"

type SmsSettingsNullClientV1 struct {
}

func NewSmsSettingsNullClientV1() *SmsSettingsNullClientV1 {
	return &SmsSettingsNullClientV1{}
}

func (c *SmsSettingsNullClientV1) GetSettingsByIds(ctx context.Context, correlationId string, recipientIds []string) ([]*SmsSettingsV1, error) {
	return make([]*SmsSettingsV1, 0), nil
}

func (c *SmsSettingsNullClientV1) GetSettingsById(ctx context.Context, correlationId string, recipientId string) (*SmsSettingsV1, error) {
	return nil, nil
}

func (c *SmsSettingsNullClientV1) GetSettingsByPhoneSettings(ctx context.Context, correlationId string, phone string) (*SmsSettingsV1, error) {
	return nil, nil
}

func (c *SmsSettingsNullClientV1) SetSettings(ctx context.Context, correlationId string, settings *SmsSettingsV1) (*SmsSettingsV1, error) {
	return settings, nil
}

func (c *SmsSettingsNullClientV1) SetVerifiedSettings(ctx context.Context, correlationId string, settings *SmsSettingsV1) (*SmsSettingsV1, error) {
	return settings, nil
}

func (c *SmsSettingsNullClientV1) SetRecipient(ctx context.Context, correlationId string, recipientId string, name string, phone string, language string) (*SmsSettingsV1, error) {
	return &SmsSettingsV1{
		Id:       recipientId,
		Name:     name,
		Phone:    phone,
		Language: language,
		Verified: false,
	}, nil
}

func (c *SmsSettingsNullClientV1) SetSubscriptions(ctx context.Context, correlationId string, recipientId string, subscriptions any) (*SmsSettingsV1, error) {
	return &SmsSettingsV1{
		Id:       recipientId,
		Name:     "",
		Phone:    "",
		Language: "",
		Verified: false,
	}, nil
}

func (c *SmsSettingsNullClientV1) DeleteSettingsById(ctx context.Context, correlationId string, recipientId string) error {
	return nil
}

func (c *SmsSettingsNullClientV1) ResendVerification(ctx context.Context, correlationId string, recipientId string) error {
	return nil
}

func (c *SmsSettingsNullClientV1) VerifyPhone(ctx context.Context, correlationId string, recipientId string, code string) error {
	return nil
}
