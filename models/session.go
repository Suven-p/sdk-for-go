package models

// Session Model
type Session struct {
	// Session ID.
	Id string `json:"$id"`
	// Session creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// User ID.
	UserId string `json:"userId"`
	// Session expiration date in ISO 8601 format.
	Expire string `json:"expire"`
	// Session Provider.
	Provider string `json:"provider"`
	// Session Provider User ID.
	ProviderUid string `json:"providerUid"`
	// Session Provider Access Token.
	ProviderAccessToken string `json:"providerAccessToken"`
	// The date of when the access token expires in ISO 8601 format.
	ProviderAccessTokenExpiry string `json:"providerAccessTokenExpiry"`
	// Session Provider Refresh Token.
	ProviderRefreshToken string `json:"providerRefreshToken"`
	// IP in use when the session was created.
	Ip string `json:"ip"`
	// Operating system code name. View list of [available
	// options](https://github.com/appwrite/appwrite/blob/master/docs/lists/os.json).
	OsCode string `json:"osCode"`
	// Operating system name.
	OsName string `json:"osName"`
	// Operating system version.
	OsVersion string `json:"osVersion"`
	// Client type.
	ClientType string `json:"clientType"`
	// Client code name. View list of [available
	// options](https://github.com/appwrite/appwrite/blob/master/docs/lists/clients.json).
	ClientCode string `json:"clientCode"`
	// Client name.
	ClientName string `json:"clientName"`
	// Client version.
	ClientVersion string `json:"clientVersion"`
	// Client engine name.
	ClientEngine string `json:"clientEngine"`
	// Client engine name.
	ClientEngineVersion string `json:"clientEngineVersion"`
	// Device name.
	DeviceName string `json:"deviceName"`
	// Device brand name.
	DeviceBrand string `json:"deviceBrand"`
	// Device model name.
	DeviceModel string `json:"deviceModel"`
	// Country two-character ISO 3166-1 alpha code.
	CountryCode string `json:"countryCode"`
	// Country name.
	CountryName string `json:"countryName"`
	// Returns true if this the current user session.
	Current bool `json:"current"`
}
