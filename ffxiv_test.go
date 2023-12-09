/*

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.

*/

package ffxiv_test

import (
	"testing"

	ffxiv "github.com/c032/go-ffxiv"
)

func TestClient_WorldStatus(t *testing.T) {
	var (
		err error
		c   ffxiv.Client
	)

	c, err = ffxiv.NewClient()
	if err != nil {
		t.Fatalf("ffxiv.NewClient() error: %s", err)
	}

	var wss []ffxiv.WorldStatus
	wss, err = c.WorldStatus()
	if err != nil {
		t.Fatalf("c.WorldStatus() error: %s", err)
	}

	if got, notWant := len(wss), 0; got == notWant {
		t.Fatalf("len(c.WorldStatus()) = %d; want non-zero", got)
	}
}
