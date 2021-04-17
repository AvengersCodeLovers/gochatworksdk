package gochatworksdk

import Response "github.com/AvengersCodeLovers/gochatworksdk/responses"

// Api.Go is setup functional SDK

// Profile is struct type defined from response
type Profile = Response.Profile

// Me is get profile own information
func Me() {}

// GetMyStatus is get profile own statics information
func GetMyStatus() {}

// GetRooms is get user room list
func GetRooms() {}
