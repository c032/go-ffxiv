/*

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.

*/

package ffxiv

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	goquery "github.com/PuerkitoBio/goquery"
)

// websiteURLStr is the base URL for all requests.
//
// For the current features the region doesn't matter so I just chose EU, but
// if any region-specific feature is added then this should be updated.
const websiteURLStr = "https://eu.finalfantasyxiv.com"

var (
	websiteURL *url.URL

	worldStatusURL *url.URL
)

var _ Client = (*client)(nil)

func init() {
	var err error

	websiteURL, err = url.Parse(websiteURLStr)
	if err != nil {
		panic(err)
	}

	worldStatusURL, err = websiteURL.Parse("/lodestone/worldstatus/")
	if err != nil {
		panic(err)
	}
}

type Client interface {
	WorldStatus() ([]WorldStatus, error)
}

type client struct {
	httpClient *http.Client
}

func (c *client) WorldStatus() ([]WorldStatus, error) {
	var (
		err error

		req *http.Request
	)

	req, err = http.NewRequest(http.MethodGet, worldStatusURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("could not create HTTP request: %w", err)
	}

	hc := c.httpClient

	var resp *http.Response

	resp, err = hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not get world status: %w", err)
	}
	defer resp.Body.Close()

	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not parse response body: %w", err)
	}

	var wss []WorldStatus
	doc.Find(".world-dcgroup__item").Each(func(_ int, selGroup *goquery.Selection) {
		groupName := strings.TrimSpace(
			selGroup.Find(".world-dcgroup__header").First().Text(),
		)

		selGroup.Find(".world-list__item").Each(func(_ int, selWorld *goquery.Selection) {
			worldName := strings.TrimSpace(
				selWorld.Find(".world-list__world_name").First().Text(),
			)
			worldCategory := strings.TrimSpace(
				selWorld.Find(".world-list__world_category").First().Text(),
			)
			canCreateNewCharacters := selWorld.Find(".world-ic__available").Length() == 1

			var serverStatus ServerStatus
			if selWorld.Find(".world-ic__3").Length() == 1 {
				serverStatus = StatusMaintenance
			} else if selWorld.Find(".world-ic__2").Length() == 1 {
				serverStatus = StatusPartialMaintenance
			} else if selWorld.Find(".world-ic__1").Length() == 1 {
				serverStatus = StatusOnline
			}

			ws := WorldStatus{
				Group: groupName,

				Name:         worldName,
				Category:     ServerCategory(worldCategory),
				ServerStatus: serverStatus,

				CanCreateNewCharacters: canCreateNewCharacters,
			}

			wss = append(wss, ws)
		})
	})

	return wss, nil
}

func NewClient() (Client, error) {
	const timeout = 15 * time.Second

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("could not create cookie jar: %w", err)
	}

	httpClient := &http.Client{
		Jar:     jar,
		Timeout: timeout,
	}

	c := &client{
		httpClient: httpClient,
	}

	return c, nil
}
