package version1

import "time"

type SmsSettingsV1 struct {
	// Recipient information
	Id       string `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Language string `json:"language"`

	// Phone management
	Subscriptions any       `json:"subscriptions"`
	Verified      bool      `json:"verified"`
	VerCode       string    `json:"ver_code"`
	VerExpireTime time.Time `json:"ver_expire_time"`

	// Custom fields
	CustomHdr any `json:"custom_hdr"`
	CustomDat any `json:"custom_dat"`
}
