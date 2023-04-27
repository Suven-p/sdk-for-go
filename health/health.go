package health

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/appwritemodels"
	"strings"
	"github.com/appwrite/sdk-for-go/appwrite"
)

// Health service
type Health struct {
	client appwrite.Client
}

func NewHealth(clt appwrite.Client) *Health {
	return &Health{
		client: clt,
	}
}

// Get check the Appwrite HTTP server is up and responsive.
type GetOptions struct {
}

type GetOption func(*GetOptions)



func (srv *Health) Get()  (*appwriteModel.HealthStatus, error) {
	path := "/health"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.HealthStatus
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.HealthStatus)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetAntivirus check the Appwrite Antivirus server is up and connection is
// successful.
type GetAntivirusOptions struct {
}

type GetAntivirusOption func(*GetAntivirusOptions)



func (srv *Health) GetAntivirus()  (*appwriteModel.HealthAntivirus, error) {
	path := "/health/anti-virus"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.HealthAntivirus
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.HealthAntivirus)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetCache check the Appwrite in-memory cache server is up and connection is
// successful.
type GetCacheOptions struct {
}

type GetCacheOption func(*GetCacheOptions)



func (srv *Health) GetCache()  (*appwriteModel.HealthStatus, error) {
	path := "/health/cache"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.HealthStatus
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.HealthStatus)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetDB check the Appwrite database server is up and connection is
// successful.
type GetDBOptions struct {
}

type GetDBOption func(*GetDBOptions)



func (srv *Health) GetDB()  (*appwriteModel.HealthStatus, error) {
	path := "/health/db"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.HealthStatus
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.HealthStatus)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetQueueCertificates get the number of certificates that are waiting to be
// issued against [Letsencrypt](https://letsencrypt.org/) in the Appwrite
// internal queue server.
type GetQueueCertificatesOptions struct {
}

type GetQueueCertificatesOption func(*GetQueueCertificatesOptions)



func (srv *Health) GetQueueCertificates()  (*appwriteModel.HealthQueue, error) {
	path := "/health/queue/certificates"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.HealthQueue
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetQueueFunctions
type GetQueueFunctionsOptions struct {
}

type GetQueueFunctionsOption func(*GetQueueFunctionsOptions)



func (srv *Health) GetQueueFunctions()  (*appwriteModel.HealthQueue, error) {
	path := "/health/queue/functions"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.HealthQueue
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetQueueLogs get the number of logs that are waiting to be processed in the
// Appwrite internal queue server.
type GetQueueLogsOptions struct {
}

type GetQueueLogsOption func(*GetQueueLogsOptions)



func (srv *Health) GetQueueLogs()  (*appwriteModel.HealthQueue, error) {
	path := "/health/queue/logs"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.HealthQueue
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetQueueWebhooks get the number of webhooks that are waiting to be
// processed in the Appwrite internal queue server.
type GetQueueWebhooksOptions struct {
}

type GetQueueWebhooksOption func(*GetQueueWebhooksOptions)



func (srv *Health) GetQueueWebhooks()  (*appwriteModel.HealthQueue, error) {
	path := "/health/queue/webhooks"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.HealthQueue
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetStorageLocal check the Appwrite local storage device is up and
// connection is successful.
type GetStorageLocalOptions struct {
}

type GetStorageLocalOption func(*GetStorageLocalOptions)



func (srv *Health) GetStorageLocal()  (*appwriteModel.HealthStatus, error) {
	path := "/health/storage/local"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.HealthStatus
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.HealthStatus)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetTime check the Appwrite server time is synced with Google remote NTP
// server. We use this technology to smoothly handle leap seconds with no
// disruptive events. The [Network Time
// Protocol](https://en.wikipedia.org/wiki/Network_Time_Protocol) (NTP) is
// used by hundreds of millions of computers and devices to synchronize their
// clocks over the Internet. If your computer sets its own clock, it likely
// uses NTP.
type GetTimeOptions struct {
}

type GetTimeOption func(*GetTimeOptions)



func (srv *Health) GetTime()  (*appwriteModel.HealthTime, error) {
	path := "/health/time"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.HealthTime
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.HealthTime)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
