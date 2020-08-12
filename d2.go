package d2

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	version   = "0.1.0"
	bNet      = "https://www.bungie.net/"
	API       = bNet + "Platform/"
	userAgent = "cmacrae/d2/" + version
)

// Possible Membership types collection
var BungieMembershipType = map[string]int{
	"All":           -1,
	"None":          0,
	"TigerXbox":     1,
	"TigerPsn":      2,
	"TigerSteam":    3,
	"TigerBlizzard": 4,
	"TigerStadia":   5,
	"TigerDemon":    10,
	"BungieNext":    254,
}

// Possible Destiny component types collection
type Client struct {
	client      *http.Client
	BungieURL   *url.URL
	PlatformURL *url.URL
	UserAgent   string
	APIKey      string
	Character   *CharacterService
	Player      *PlayerService
	Platform    *PlatformService
}

func NewClient(httpClient *http.Client, apiKey string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	bNetURL, _ := url.Parse(bNet)
	APIURL, _ := url.Parse(API)

	c := &Client{
		client:      httpClient,
		BungieURL:   bNetURL,
		PlatformURL: APIURL,
		UserAgent:   userAgent,
		APIKey:      apiKey,
	}
	c.Platform = &PlatformService{client: c}
	c.Character = &CharacterService{client: c}
	c.Player = &PlayerService{client: c}
	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	u := c.PlatformURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	req.Header.Add("X-API-Key", c.APIKey)
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) ([]byte, error) {
	var body []byte

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	return body, nil
}
