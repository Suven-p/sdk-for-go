package health

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/models"
	"strings"
	"github.com/appwrite/sdk-for-go/client"
)

// Health service
type Health struct {
	client client.Client
}

func NewHealth(clt client.Client) *Health {
	return &Health{
		client: clt,
	}
}


// Get check the Appwrite HTTP server is up and responsive.
func (srv *Health) Get()  (*models.HealthStatus, error) {
	path := "/health"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.HealthStatus
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.HealthStatus)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetAntivirus check the Appwrite Antivirus server is up and connection is
// successful.
func (srv *Health) GetAntivirus()  (*models.HealthAntivirus, error) {
	path := "/health/anti-virus"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.HealthAntivirus
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.HealthAntivirus)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetCache check the Appwrite in-memory cache server is up and connection is
// successful.
func (srv *Health) GetCache()  (*models.HealthStatus, error) {
	path := "/health/cache"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.HealthStatus
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.HealthStatus)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetDB check the Appwrite database server is up and connection is
// successful.
func (srv *Health) GetDB()  (*models.HealthStatus, error) {
	path := "/health/db"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.HealthStatus
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.HealthStatus)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetQueueCertificates get the number of certificates that are waiting to be
// issued against [Letsencrypt](https://letsencrypt.org/) in the Appwrite
// internal queue server.
func (srv *Health) GetQueueCertificates()  (*models.HealthQueue, error) {
	path := "/health/queue/certificates"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.HealthQueue
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetQueueFunctions
func (srv *Health) GetQueueFunctions()  (*models.HealthQueue, error) {
	path := "/health/queue/functions"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.HealthQueue
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetQueueLogs get the number of logs that are waiting to be processed in the
// Appwrite internal queue server.
func (srv *Health) GetQueueLogs()  (*models.HealthQueue, error) {
	path := "/health/queue/logs"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.HealthQueue
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetQueueWebhooks get the number of webhooks that are waiting to be
// processed in the Appwrite internal queue server.
func (srv *Health) GetQueueWebhooks()  (*models.HealthQueue, error) {
	path := "/health/queue/webhooks"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.HealthQueue
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetStorageLocal check the Appwrite local storage device is up and
// connection is successful.
func (srv *Health) GetStorageLocal()  (*models.HealthStatus, error) {
	path := "/health/storage/local"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.HealthStatus
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.HealthStatus)
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
func (srv *Health) GetTime()  (*models.HealthTime, error) {
	path := "/health/time"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.HealthTime
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.HealthTime)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
