package types

type StratoDNSProviderConfig struct {
	// Change the two fields below according to the format of the configuration
	// to be decoded.
	// These fields will be set by users in the
	// `issuer.spec.acme.dns01.providers.webhook.config` field.
	ApiUrl         string `json:"apiUrl"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	TotpSecret     string `json:"totpSecret"`
	TotpDeviceName string `json:"totpDeviceName"`
	WaitingTime    int    `json:"waitingTime"`
}
