//go:build darwin
// +build darwin

package sparkle

import "C"

var feedURL = func() string {
	return ""
}

// CGOSetFeedURL is a bridge for a Sparkle, fetching the url from a Go function
// since the Sparkle framework uses a delegate to provide an URL as a feedURLStringForUpdater function result
// to a SPUUpdater, we need to provide a way to call a feedURL function from an Objective-C function
//
//export CGOFeedURL
func CGOFeedURL() *C.char {
	if url := feedURL(); url != "" {
		return C.CString(url)
	}
	return nil
}
