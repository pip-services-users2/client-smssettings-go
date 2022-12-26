package version1

import "context"

type ISmsSettingsClientV1 interface {
	GetSettingsByIds(ctx context.Context, correlationId string, recipientIds []string) ([]*SmsSettingsV1, error)

	GetSettingsById(ctx context.Context, correlationId string, recipientId string) (*SmsSettingsV1, error)

	GetSettingsByPhoneSettings(ctx context.Context, correlationId string, phone string) (*SmsSettingsV1, error)

	SetSettings(ctx context.Context, correlationId string, settings *SmsSettingsV1) (*SmsSettingsV1, error)

	SetVerifiedSettings(ctx context.Context, correlationId string, settings *SmsSettingsV1) (*SmsSettingsV1, error)

	SetRecipient(ctx context.Context, correlationId string, recipientId string,
		name string, phone string, language string) (*SmsSettingsV1, error)

	SetSubscriptions(ctx context.Context, correlationId string, recipientId string, subscriptions any) (*SmsSettingsV1, error)

	DeleteSettingsById(ctx context.Context, correlationId string, recipientId string) error

	ResendVerification(ctx context.Context, correlationId string, recipientId string) error

	VerifyPhone(ctx context.Context, correlationId string, recipientId string, code string) error
}
