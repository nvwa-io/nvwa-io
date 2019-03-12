package libs

import "net/url"

func GenerateUrl(host string, pathString string, urlValueMap map[string]string) (string, error) {
	Url, err := url.Parse(host)
	if err != nil {
	return "", err
	}
	Url.Path += pathString

	parameters := url.Values{}
	for k, v := range urlValueMap {
		parameters.Add(k, v)
	}

	Url.RawQuery = parameters.Encode()
	return Url.String(), nil
}
