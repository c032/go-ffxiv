/*

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.

*/

package ffxiv_test

import (
	"os"
	"testing"

	ffxiv "github.com/c032/go-ffxiv"
)

func TestClient_WorldStatus(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

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

func TestParseWorldStatusPage(t *testing.T) {
	var (
		err error
		f   *os.File
	)

	f, err = os.Open("testdata/worldstatus.html")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var wss []ffxiv.WorldStatus
	wss, err = ffxiv.ParseWorldStatusPage(f)
	if err != nil {
		t.Fatal(err)
	}

	expectedAllWorldStatus := []ffxiv.WorldStatus{
		ffxiv.WorldStatus{Group: "Chaos", Name: "Cerberus", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Chaos", Name: "Louisoix", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Chaos", Name: "Moogle", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Chaos", Name: "Omega", Category: ffxiv.CategoryPreferred, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Chaos", Name: "Phantom", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Chaos", Name: "Ragnarok", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Chaos", Name: "Sagittarius", Category: ffxiv.CategoryPreferred, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Chaos", Name: "Spriggan", Category: ffxiv.CategoryPreferred, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: false},
		ffxiv.WorldStatus{Group: "Light", Name: "Alpha", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Light", Name: "Lich", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Light", Name: "Odin", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Light", Name: "Phoenix", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Light", Name: "Raiden", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Light", Name: "Shiva", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Light", Name: "Twintania", Category: ffxiv.CategoryPreferred, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: false},
		ffxiv.WorldStatus{Group: "Light", Name: "Zodiark", Category: ffxiv.CategoryPreferred, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: false},
		ffxiv.WorldStatus{Group: "Materia", Name: "Bismarck", Category: ffxiv.CategoryNew, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Materia", Name: "Ravana", Category: ffxiv.CategoryNew, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Materia", Name: "Sephirot", Category: ffxiv.CategoryNew, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Materia", Name: "Sophia", Category: ffxiv.CategoryNew, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Materia", Name: "Zurvan", Category: ffxiv.CategoryNew, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Aether", Name: "Adamantoise", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Aether", Name: "Cactuar", Category: ffxiv.CategoryCongested, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: false},
		ffxiv.WorldStatus{Group: "Aether", Name: "Faerie", Category: ffxiv.CategoryCongested, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: false},
		ffxiv.WorldStatus{Group: "Aether", Name: "Gilgamesh", Category: ffxiv.CategoryCongested, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: false},
		ffxiv.WorldStatus{Group: "Aether", Name: "Jenova", Category: ffxiv.CategoryCongested, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: false},
		ffxiv.WorldStatus{Group: "Aether", Name: "Midgardsormr", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Aether", Name: "Sargatanas", Category: ffxiv.CategoryCongested, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: false},
		ffxiv.WorldStatus{Group: "Aether", Name: "Siren", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Crystal", Name: "Balmung", Category: ffxiv.CategoryCongested, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: false},
		ffxiv.WorldStatus{Group: "Crystal", Name: "Brynhildr", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Crystal", Name: "Coeurl", Category: ffxiv.CategoryCongested, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: false},
		ffxiv.WorldStatus{Group: "Crystal", Name: "Diabolos", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Crystal", Name: "Goblin", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Crystal", Name: "Malboro", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Crystal", Name: "Mateus", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Crystal", Name: "Zalera", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Dynamis", Name: "Halicarnassus", Category: ffxiv.CategoryNew, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Dynamis", Name: "Maduin", Category: ffxiv.CategoryNew, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Dynamis", Name: "Marilith", Category: ffxiv.CategoryNew, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Dynamis", Name: "Seraph", Category: ffxiv.CategoryNew, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Primal", Name: "Behemoth", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Primal", Name: "Excalibur", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Primal", Name: "Exodus", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Primal", Name: "Famfrit", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Primal", Name: "Hyperion", Category: ffxiv.CategoryCongested, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: false},
		ffxiv.WorldStatus{Group: "Primal", Name: "Lamia", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Primal", Name: "Leviathan", Category: ffxiv.CategoryCongested, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: false},
		ffxiv.WorldStatus{Group: "Primal", Name: "Ultros", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Elemental", Name: "Aegis", Category: ffxiv.CategoryPreferred, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Elemental", Name: "Atomos", Category: ffxiv.CategoryPreferred, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Elemental", Name: "Carbuncle", Category: ffxiv.CategoryPreferred, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Elemental", Name: "Garuda", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Elemental", Name: "Gungnir", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Elemental", Name: "Kujata", Category: ffxiv.CategoryPreferred, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Elemental", Name: "Tonberry", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Elemental", Name: "Typhon", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Gaia", Name: "Alexander", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Gaia", Name: "Bahamut", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Gaia", Name: "Durandal", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Gaia", Name: "Fenrir", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Gaia", Name: "Ifrit", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Gaia", Name: "Ridill", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Gaia", Name: "Tiamat", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Gaia", Name: "Ultima", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Mana", Name: "Anima", Category: ffxiv.CategoryCongested, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: false},
		ffxiv.WorldStatus{Group: "Mana", Name: "Asura", Category: ffxiv.CategoryCongested, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: false},
		ffxiv.WorldStatus{Group: "Mana", Name: "Chocobo", Category: ffxiv.CategoryCongested, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: false},
		ffxiv.WorldStatus{Group: "Mana", Name: "Hades", Category: ffxiv.CategoryCongested, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: false},
		ffxiv.WorldStatus{Group: "Mana", Name: "Ixion", Category: ffxiv.CategoryCongested, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: false},
		ffxiv.WorldStatus{Group: "Mana", Name: "Masamune", Category: ffxiv.CategoryCongested, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: false},
		ffxiv.WorldStatus{Group: "Mana", Name: "Pandaemonium", Category: ffxiv.CategoryCongested, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: false},
		ffxiv.WorldStatus{Group: "Mana", Name: "Titan", Category: ffxiv.CategoryCongested, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: false},
		ffxiv.WorldStatus{Group: "Meteor", Name: "Belias", Category: ffxiv.CategoryPreferred, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Meteor", Name: "Mandragora", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Meteor", Name: "Ramuh", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Meteor", Name: "Shinryu", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Meteor", Name: "Unicorn", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Meteor", Name: "Valefor", Category: ffxiv.CategoryStandard, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Meteor", Name: "Yojimbo", Category: ffxiv.CategoryPreferred, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
		ffxiv.WorldStatus{Group: "Meteor", Name: "Zeromus", Category: ffxiv.CategoryPreferred, ServerStatus: ffxiv.StatusOnline, CanCreateNewCharacters: true},
	}

	if got, want := len(wss), len(expectedAllWorldStatus); got != want {
		t.Fatalf("len(ParseWorldStatus(f)) = %d; want %d", got, want)
	}

	for i, wantWorldStatus := range expectedAllWorldStatus {
		gotWorldStatus := wss[i]

		if got, want := gotWorldStatus.Group, wantWorldStatus.Group; got != want {
			t.Errorf("wss[%d].Group = %#v; want %#v", i, got, want)
		}

		if got, want := gotWorldStatus.Name, wantWorldStatus.Name; got != want {
			t.Errorf("wss[%d].Name = %#v; want %#v", i, got, want)
		}

		if got, want := gotWorldStatus.Category, wantWorldStatus.Category; got != want {
			t.Errorf("wss[%d].Category = %#v; want %#v", i, got, want)
		}

		if got, want := gotWorldStatus.ServerStatus, wantWorldStatus.ServerStatus; got != want {
			t.Errorf("wss[%d].ServerStatus = %#v; want %#v", i, got, want)
		}

		if got, want := gotWorldStatus.CanCreateNewCharacters, wantWorldStatus.CanCreateNewCharacters; got != want {
			t.Errorf("wss[%d].CanCreateNewCharacters = %#v; want %#v", i, got, want)
		}
	}
}
