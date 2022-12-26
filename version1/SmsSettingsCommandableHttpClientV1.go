package version1

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
	cclients "github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type SmsSettingsCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewSmsSettingsCommandableHttpClientV1() *SmsSettingsCommandableHttpClientV1 {
	c := &SmsSettingsCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/sms_settings"),
	}
	return c
}

func (c *SmsSettingsCommandableHttpClientV1) GetSettingsByIds(ctx context.Context, correlationId string, recipientIds []string) ([]*SmsSettingsV1, error) {
	params := cdata.NewAnyValueMapFromTuples(
		"recipient_ids", recipientIds,
	)

	res, err := c.CallCommand(ctx, "get_settings_by_ids", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[[]*SmsSettingsV1](res, correlationId)
}

func (c *SmsSettingsCommandableHttpClientV1) GetSettingsById(ctx context.Context, correlationId string, recipientId string) (*SmsSettingsV1, error) {
	params := cdata.NewAnyValueMapFromTuples(
		"recipient_id", recipientId,
	)

	res, err := c.CallCommand(ctx, "get_settings_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[*SmsSettingsV1](res, correlationId)
}

func (c *SmsSettingsCommandableHttpClientV1) GetSettingsByPhoneSettings(ctx context.Context, correlationId string, phone string) (*SmsSettingsV1, error) {
	params := cdata.NewAnyValueMapFromTuples(
		"phone", phone,
	)

	res, err := c.CallCommand(ctx, "get_settings_by_phone", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[*SmsSettingsV1](res, correlationId)
}

func (c *SmsSettingsCommandableHttpClientV1) SetSettings(ctx context.Context, correlationId string, settings *SmsSettingsV1) (*SmsSettingsV1, error) {
	params := cdata.NewAnyValueMapFromTuples(
		"settings", settings,
	)

	res, err := c.CallCommand(ctx, "set_settings", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[*SmsSettingsV1](res, correlationId)
}

func (c *SmsSettingsCommandableHttpClientV1) SetVerifiedSettings(ctx context.Context, correlationId string, settings *SmsSettingsV1) (*SmsSettingsV1, error) {
	params := cdata.NewAnyValueMapFromTuples(
		"settings", settings,
	)

	res, err := c.CallCommand(ctx, "set_verified_settings", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[*SmsSettingsV1](res, correlationId)
}

func (c *SmsSettingsCommandableHttpClientV1) SetRecipient(ctx context.Context, correlationId string, recipientId string, name string, phone string, language string) (*SmsSettingsV1, error) {
	params := cdata.NewAnyValueMapFromTuples(
		"recipient_id", recipientId,
		"name", name,
		"phone", phone,
		"language", language,
	)

	res, err := c.CallCommand(ctx, "set_recipient", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[*SmsSettingsV1](res, correlationId)
}

func (c *SmsSettingsCommandableHttpClientV1) SetSubscriptions(ctx context.Context, correlationId string, recipientId string, subscriptions any) (*SmsSettingsV1, error) {
	params := cdata.NewAnyValueMapFromTuples(
		"recipient_id", recipientId,
		"subscriptions", subscriptions,
	)

	res, err := c.CallCommand(ctx, "set_subscriptions", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[*SmsSettingsV1](res, correlationId)
}

func (c *SmsSettingsCommandableHttpClientV1) DeleteSettingsById(ctx context.Context, correlationId string, recipientId string) error {
	params := cdata.NewAnyValueMapFromTuples(
		"recipient_id", recipientId,
	)

	_, err := c.CallCommand(ctx, "delete_settings_by_id", correlationId, params)
	return err
}

func (c *SmsSettingsCommandableHttpClientV1) ResendVerification(ctx context.Context, correlationId string, recipientId string) error {
	params := cdata.NewAnyValueMapFromTuples(
		"recipient_id", recipientId,
	)

	_, err := c.CallCommand(ctx, "resend_verification", correlationId, params)
	return err
}

func (c *SmsSettingsCommandableHttpClientV1) VerifyPhone(ctx context.Context, correlationId string, recipientId string, code string) error {
	params := cdata.NewAnyValueMapFromTuples(
		"recipient_id", recipientId,
		"code", code,
	)

	_, err := c.CallCommand(ctx, "verify_phone", correlationId, params)
	return err
}
