package avatars

import (
	"encoding/json"
	"errors"
	"strings"
	"github.com/appwrite/sdk-for-go/appwrite"
)

// Avatars service
type Avatars struct {
	client appwrite.Client
}

func NewAvatars(clt appwrite.Client) *Avatars {
	return &Avatars{
		client: clt,
	}
}

// GetBrowser you can use this endpoint to show different browser icons to
// your users. The code argument receives the browser code as it appears in
// your user [GET /account/sessions](/docs/client/account#accountGetSessions)
// endpoint. Use width, height and quality arguments to change the output
// settings.
// 
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 100x100px.
type GetBrowserOptions struct {
	Width int
	Height int
	Quality int
}

type GetBrowserOption func(*GetBrowserOptions)

func WithGetBrowserWidth(v int) GetBrowserOption {
	return func(o *GetBrowserOptions) {
		o.Width = v
	}
}
func WithGetBrowserHeight(v int) GetBrowserOption {
	return func(o *GetBrowserOptions) {
		o.Height = v
	}
}
func WithGetBrowserQuality(v int) GetBrowserOption {
	return func(o *GetBrowserOptions) {
		o.Quality = v
	}
}

			
func (srv *Avatars) GetBrowser(Code string, optionalSetters ...GetBrowserOption)  (*[]byte, error) {
	r := strings.NewReplacer("{code}", Code)
	path := r.Replace("/avatars/browsers/{code}")

	options := &GetBrowserOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["width"] = options.Width
	params["height"] = options.Height
	params["quality"] = options.Quality
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed []byte
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetCreditCard the credit card endpoint will return you the icon of the
// credit card provider you need. Use width, height and quality arguments to
// change the output settings.
// 
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 100x100px.
// 
type GetCreditCardOptions struct {
	Width int
	Height int
	Quality int
}

type GetCreditCardOption func(*GetCreditCardOptions)

func WithGetCreditCardWidth(v int) GetCreditCardOption {
	return func(o *GetCreditCardOptions) {
		o.Width = v
	}
}
func WithGetCreditCardHeight(v int) GetCreditCardOption {
	return func(o *GetCreditCardOptions) {
		o.Height = v
	}
}
func WithGetCreditCardQuality(v int) GetCreditCardOption {
	return func(o *GetCreditCardOptions) {
		o.Quality = v
	}
}

			
func (srv *Avatars) GetCreditCard(Code string, optionalSetters ...GetCreditCardOption)  (*[]byte, error) {
	r := strings.NewReplacer("{code}", Code)
	path := r.Replace("/avatars/credit-cards/{code}")

	options := &GetCreditCardOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["width"] = options.Width
	params["height"] = options.Height
	params["quality"] = options.Quality
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed []byte
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetFavicon use this endpoint to fetch the favorite icon (AKA favicon) of
// any remote website URL.
// 
type GetFaviconOptions struct {
}

type GetFaviconOption func(*GetFaviconOptions)


	
func (srv *Avatars) GetFavicon(Url string)  (*[]byte, error) {
	path := "/avatars/favicon"

	params := map[string]interface{}{}
	params["url"] = Url
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed []byte
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetFlag you can use this endpoint to show different country flags icons to
// your users. The code argument receives the 2 letter country code. Use
// width, height and quality arguments to change the output settings. Country
// codes follow the [ISO 3166-1](http://en.wikipedia.org/wiki/ISO_3166-1)
// standard.
// 
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 100x100px.
// 
type GetFlagOptions struct {
	Width int
	Height int
	Quality int
}

type GetFlagOption func(*GetFlagOptions)

func WithGetFlagWidth(v int) GetFlagOption {
	return func(o *GetFlagOptions) {
		o.Width = v
	}
}
func WithGetFlagHeight(v int) GetFlagOption {
	return func(o *GetFlagOptions) {
		o.Height = v
	}
}
func WithGetFlagQuality(v int) GetFlagOption {
	return func(o *GetFlagOptions) {
		o.Quality = v
	}
}

			
func (srv *Avatars) GetFlag(Code string, optionalSetters ...GetFlagOption)  (*[]byte, error) {
	r := strings.NewReplacer("{code}", Code)
	path := r.Replace("/avatars/flags/{code}")

	options := &GetFlagOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["width"] = options.Width
	params["height"] = options.Height
	params["quality"] = options.Quality
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed []byte
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetImage use this endpoint to fetch a remote image URL and crop it to any
// image size you want. This endpoint is very useful if you need to crop and
// display remote images in your app or in case you want to make sure a 3rd
// party image is properly served using a TLS protocol.
// 
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 400x400px.
// 
type GetImageOptions struct {
	Width int
	Height int
}

type GetImageOption func(*GetImageOptions)

func WithGetImageWidth(v int) GetImageOption {
	return func(o *GetImageOptions) {
		o.Width = v
	}
}
func WithGetImageHeight(v int) GetImageOption {
	return func(o *GetImageOptions) {
		o.Height = v
	}
}

			
func (srv *Avatars) GetImage(Url string, optionalSetters ...GetImageOption)  (*[]byte, error) {
	path := "/avatars/image"

	options := &GetImageOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["url"] = Url
	params["width"] = options.Width
	params["height"] = options.Height
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed []byte
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetInitials use this endpoint to show your user initials avatar icon on
// your website or app. By default, this route will try to print your
// logged-in user name or email initials. You can also overwrite the user name
// if you pass the 'name' parameter. If no name is given and no user is
// logged, an empty avatar will be returned.
// 
// You can use the color and background params to change the avatar colors. By
// default, a random theme will be selected. The random theme will persist for
// the user's initials when reloading the same theme will always return for
// the same initials.
// 
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 100x100px.
// 
type GetInitialsOptions struct {
	Name string
	Width int
	Height int
	Background string
}

type GetInitialsOption func(*GetInitialsOptions)

func WithGetInitialsName(v string) GetInitialsOption {
	return func(o *GetInitialsOptions) {
		o.Name = v
	}
}
func WithGetInitialsWidth(v int) GetInitialsOption {
	return func(o *GetInitialsOptions) {
		o.Width = v
	}
}
func WithGetInitialsHeight(v int) GetInitialsOption {
	return func(o *GetInitialsOptions) {
		o.Height = v
	}
}
func WithGetInitialsBackground(v string) GetInitialsOption {
	return func(o *GetInitialsOptions) {
		o.Background = v
	}
}

	
func (srv *Avatars) GetInitials(optionalSetters ...GetInitialsOption)  (*[]byte, error) {
	path := "/avatars/initials"

	options := &GetInitialsOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = options.Name
	params["width"] = options.Width
	params["height"] = options.Height
	params["background"] = options.Background
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed []byte
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetQR converts a given plain text to a QR code image. You can use the query
// parameters to change the size and style of the resulting image.
// 
type GetQROptions struct {
	Size int
	Margin int
	Download bool
}

type GetQROption func(*GetQROptions)

func WithGetQRSize(v int) GetQROption {
	return func(o *GetQROptions) {
		o.Size = v
	}
}
func WithGetQRMargin(v int) GetQROption {
	return func(o *GetQROptions) {
		o.Margin = v
	}
}
func WithGetQRDownload(v bool) GetQROption {
	return func(o *GetQROptions) {
		o.Download = v
	}
}

			
func (srv *Avatars) GetQR(Text string, optionalSetters ...GetQROption)  (*[]byte, error) {
	path := "/avatars/qr"

	options := &GetQROptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["text"] = Text
	params["size"] = options.Size
	params["margin"] = options.Margin
	params["download"] = options.Download
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed []byte
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
