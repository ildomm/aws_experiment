package util

import "time"

func SetGlobalTimezoneAsUTC() ( err error ) {
	location, err := time.LoadLocation("UTC")
	if err != nil {
		return
	}
	time.Local = location
	return
}