//go:build darwin
// +build darwin

// Package sparkle provides go bindings for Sparkle.
//
// Sparkle is a secure and reliable software update framework for Cocoa developers.
// See https://github.com/sparkle-project/Sparkle for more info.
package sparkle

/*
#cgo CFLAGS: -I ${SRCDIR}/Sparkle.framework
#cgo LDFLAGS: -F ${SRCDIR} -framework Sparkle -framework Foundation

#include <stdlib.h>

void initialize();

void checkForUpdates();

void checkForUpdatesInBackground();

void setAutomaticallyChecksForUpdates(int);
int automaticallyChecksForUpdates();

void setAutomaticallyDownloadsUpdates(int);
int automaticallyDownloadsUpdates();

void setUpdateCheckInterval(int);
double updateCheckInterval();

void checkForUpdateInformation();

void setFeedURL(const char*);
const char* feedURL();

void setUserAgentString(const char*);
const char* userAgentString();

void setSendsSystemProfile(int);
int sendsSystemProfile();

void setDecryptionPassword(const char*);
const char* decryptionPassword();

double lastUpdateCheckDate();

void resetUpdateCycle();

int updateInProgress();

*/
import "C"

import (
	"math"
	"runtime"
	"time"
	"unsafe"

	_ "github.com/abemedia/go-sparkle/Sparkle.framework"
)

func init() {
	runtime.LockOSThread()
	C.initialize()
}

// Explicitly checks for updates and displays a progress dialog while doing so.
//
// This method is meant for a main menu item.
// Connect any menu item to this action in Interface Builder,
// and Sparkle will check for updates and report back its findings verbosely
// when it is invoked.
//
// This will find updates that the user has opted into skipping.
func CheckForUpdates() {
	C.checkForUpdates()
}

// Checks for updates, but does not display any UI unless an update is found.
//
// This is meant for programmatically initating a check for updates. That is,
// it will display no UI unless it actually finds an update, in which case it
// proceeds as usual.
//
// If automatic downloading of updates it turned on and allowed, however,
// this will invoke that behavior, and if an update is found, it will be downloaded
// in the background silently and will be prepped for installation.
//
// This will not find updates that the user has opted into skipping.
func CheckForUpdatesInBackground() {
	C.checkForUpdatesInBackground()
}

// Sets whether or not to check for updates automatically.
func SetAutomaticallyChecksForUpdates(check bool) {
	C.setAutomaticallyChecksForUpdates(bool2int(check))
}

// Returns whether or not to check for updates automatically.
//
// Setting this property will persist in the host bundle's user defaults.
// The update schedule cycle will be reset in a short delay after the property's new value is set.
// This is to allow reverting this property without kicking off a schedule change immediately
func AutomaticallyChecksForUpdates() bool {
	return C.automaticallyChecksForUpdates() != 0
}

// Sets whether or not updates can be automatically downloaded in the background.
//
// Note that automatic downloading of updates can be disallowed by the developer
// or by the user's system if silent updates cannot be done (eg: if they require authentication).
// In this case, `sparkle.GetAutomaticallyDownloadsUpdates` will return NO regardless of how this property is set.
//
// Setting this property will persist in the host bundle's user defaults.
func SetAutomaticallyDownloadsUpdates(check bool) {
	C.setAutomaticallyDownloadsUpdates(bool2int(check))
}

// Returns whether or not updates can be automatically downloaded in the background.
//
// Note that automatic downloading of updates can be disallowed by the developer
// or by the user's system if silent updates cannot be done (eg: if they require authentication).
// In this case, -automaticallyDownloadsUpdates will return NO regardless of how this property is set.
func AutomaticallyDownloadsUpdates() bool {
	return C.automaticallyDownloadsUpdates() != 0
}

// Sets the automatic update check interval.
//
// Setting this property will persist in the host bundle's user defaults.
// The update schedule cycle will be reset in a short delay after the property's new value is set.
// This is to allow reverting this property without kicking off a schedule change immediately
func SetUpdateCheckInterval(duration time.Duration) {
	C.setUpdateCheckInterval(C.int(duration.Seconds()))
}

// Returns the current automatic update check interval.
func UpdateCheckInterval() time.Duration {
	return time.Duration(C.updateCheckInterval()) * time.Second
}

// Begins a "probing" check for updates which will not actually offer to
// update to that version.
//
// However, the delegate methods
// SUUpdaterDelegate::updater:didFindValidUpdate: and
// SUUpdaterDelegate::updaterDidNotFindUpdate: will be called,
// so you can use that information in your UI.
//
// Updates that have been skipped by the user will not be found.
func CheckForUpdateInformation() {
	C.checkForUpdateInformation()
}

// Sets the URL of the appcast used to download update information.
//
// Setting this property will persist in the host bundle's user defaults.
// If you don't want persistence, you may want to consider instead implementing
// SUUpdaterDelegate::feedURLStringForUpdater: or SUUpdaterDelegate::feedParametersForUpdater:sendingSystemProfile:
//
// This property must be called on the main thread.
func SetFeedURL(url string) {
	u := C.CString(url)
	defer C.free(unsafe.Pointer(u))
	C.setFeedURL(u)
}

// Returns the URL of the appcast used to download update information.
func FeedURL() string {
	return C.GoString(C.feedURL())
}

// Sets the user agent used when checking for updates.
func SetUserAgentString(ua string) {
	u := C.CString(ua)
	defer C.free(unsafe.Pointer(u))
	C.setUserAgentString(u)
}

// Returns the user agent used when checking for updates.
func UserAgentString() string {
	return C.GoString(C.userAgentString())
}

// Sets whether or not the user's system profile information is sent when checking for updates.
//
// Setting this property will persist in the host bundle's user defaults.
func SetSendsSystemProfile(check bool) {
	C.setSendsSystemProfile(bool2int(check))
}

// Returns whether or not the user's system profile information is sent when checking for updates.
func SendsSystemProfile() bool {
	return C.sendsSystemProfile() != 0
}

// Sets the decryption password used for extracting updates shipped as Apple Disk Images (dmg)
func SetDecryptionPassword(url string) {
	u := C.CString(url)
	defer C.free(unsafe.Pointer(u))
	C.setDecryptionPassword(u)
}

// Returns the decryption password used for extracting updates shipped as Apple Disk Images (dmg)
func DecryptionPassword() string {
	return C.GoString(C.decryptionPassword())
}

// Returns the date of last update check.
func LastUpdateCheckDate() time.Time {
	s, n := math.Modf(float64(C.lastUpdateCheckDate()))
	return time.Unix(int64(s), int64(float64(time.Second)*n))
}

// Appropriately schedules or cancels the update checking timer according to
// the preferences for time interval and automatic checks.
//
// This call does not change the date of the next check,
// but only the internal NSTimer.
func ResetUpdateCycle() {
	C.resetUpdateCycle()
}

// Returns whether or not an update is in progress.
func UpdateInProgress() bool {
	return C.updateInProgress() != 0
}
