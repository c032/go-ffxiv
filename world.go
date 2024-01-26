/*

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.

*/

package ffxiv

type ServerStatus string

const (
	StatusUnknown            ServerStatus = ""
	StatusMaintenance        ServerStatus = "maintenance"
	StatusPartialMaintenance ServerStatus = "partial_maintenance"
	StatusOnline             ServerStatus = "online"
)

type ServerCategory string

const (
	CategoryStandard  ServerCategory = "Standard"
	CategoryPreferred ServerCategory = "Preferred"
	CategoryCongested ServerCategory = "Congested"
	CategoryNew       ServerCategory = "New"
)

type WorldStatus struct {
	Group string

	Name         string
	Category     ServerCategory
	ServerStatus ServerStatus

	CanCreateNewCharacters bool
}
