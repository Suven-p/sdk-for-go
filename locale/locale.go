package locale

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/appwritemodels"
	"strings"
	"github.com/appwrite/sdk-for-go/appwrite"
)

// Locale service
type Locale struct {
	client appwrite.Client
}

func NewLocale(clt appwrite.Client) *Locale {
	return &Locale{
		client: clt,
	}
}

// Get get the current user location based on IP. Returns an object with user
// country code, country name, continent name, continent code, ip address and
// suggested currency. You can use the locale header to get the data in a
// supported language.
// 
// ([IP Geolocation by DB-IP](https://db-ip.com))
type GetOptions struct {
}

type GetOption func(*GetOptions)



func (srv *Locale) Get()  (*appwriteModel.Locale, error) {
	path := "/locale"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Locale
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Locale)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// ListContinents list of all continents. You can use the locale header to get
// the data in a supported language.
type ListContinentsOptions struct {
}

type ListContinentsOption func(*ListContinentsOptions)



func (srv *Locale) ListContinents()  (*appwriteModel.ContinentList, error) {
	path := "/locale/continents"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.ContinentList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.ContinentList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// ListCountries list of all countries. You can use the locale header to get
// the data in a supported language.
type ListCountriesOptions struct {
}

type ListCountriesOption func(*ListCountriesOptions)



func (srv *Locale) ListCountries()  (*appwriteModel.CountryList, error) {
	path := "/locale/countries"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.CountryList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.CountryList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// ListCountriesEU list of all countries that are currently members of the EU.
// You can use the locale header to get the data in a supported language.
type ListCountriesEUOptions struct {
}

type ListCountriesEUOption func(*ListCountriesEUOptions)



func (srv *Locale) ListCountriesEU()  (*appwriteModel.CountryList, error) {
	path := "/locale/countries/eu"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.CountryList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.CountryList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// ListCountriesPhones list of all countries phone codes. You can use the
// locale header to get the data in a supported language.
type ListCountriesPhonesOptions struct {
}

type ListCountriesPhonesOption func(*ListCountriesPhonesOptions)



func (srv *Locale) ListCountriesPhones()  (*appwriteModel.PhoneList, error) {
	path := "/locale/countries/phones"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.PhoneList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.PhoneList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// ListCurrencies list of all currencies, including currency symbol, name,
// plural, and decimal digits for all major and minor currencies. You can use
// the locale header to get the data in a supported language.
type ListCurrenciesOptions struct {
}

type ListCurrenciesOption func(*ListCurrenciesOptions)



func (srv *Locale) ListCurrencies()  (*appwriteModel.CurrencyList, error) {
	path := "/locale/currencies"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.CurrencyList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.CurrencyList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// ListLanguages list of all languages classified by ISO 639-1 including
// 2-letter code, name in English, and name in the respective language.
type ListLanguagesOptions struct {
}

type ListLanguagesOption func(*ListLanguagesOptions)



func (srv *Locale) ListLanguages()  (*appwriteModel.LanguageList, error) {
	path := "/locale/languages"

	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.LanguageList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.LanguageList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
