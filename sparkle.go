// +build darwin

// Package winsparkle provides go bindings for WinSparkle.
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
int getAutomaticallyChecksForUpdates();

void setAutomaticallyDownloadsUpdates(int);
int getAutomaticallyDownloadsUpdates();

void setUpdateCheckInterval(int);
double getUpdateCheckInterval();

void checkForUpdateInformation();

void setFeedURL(const char*);
const char* getFeedURL();

void setUserAgentString(const char*);
const char* getUserAgentString();

void setSendsSystemProfile(int);
int getSendsSystemProfile();

void setDecryptionPassword(const char*);
const char* getDecryptionPassword();

void installUpdatesIfAvailable();

double getLastUpdateCheckDate();

void resetUpdateCycle();

int getUpdateInProgress();

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
func GetAutomaticallyChecksForUpdates() bool {
	return C.getAutomaticallyChecksForUpdates() != 0
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
func GetAutomaticallyDownloadsUpdates() bool {
	return C.getAutomaticallyDownloadsUpdates() != 0
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
func GetUpdateCheckInterval() time.Duration {
	return time.Duration(C.getUpdateCheckInterval()) * time.Second
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
func GetFeedURL() string {
	return C.GoString(C.getFeedURL())
}

// Sets the user agent used when checking for updates.
func SetUserAgentString(ua string) {
	u := C.CString(ua)
	defer C.free(unsafe.Pointer(u))
	C.setUserAgentString(u)
}

// Returns the user agent used when checking for updates.
func GetUserAgentString() string {
	return C.GoString(C.getUserAgentString())
}

// Sets whether or not the user's system profile information is sent when checking for updates.
//
// Setting this property will persist in the host bundle's user defaults.
func SetSendsSystemProfile(check bool) {
	C.setSendsSystemProfile(bool2int(check))
}

// Returns whether or not the user's system profile information is sent when checking for updates.
func GetSendsSystemProfile() bool {
	return C.getSendsSystemProfile() != 0
}

// Sets the decryption password used for extracting updates shipped as Apple Disk Images (dmg)
func SetDecryptionPassword(url string) {
	u := C.CString(url)
	defer C.free(unsafe.Pointer(u))
	C.setDecryptionPassword(u)
}

// Returns the decryption password used for extracting updates shipped as Apple Disk Images (dmg)
func GetDecryptionPassword() string {
	return C.GoString(C.getDecryptionPassword())
}

// This function ignores normal update schedule, ignores user preferences,
// and interrupts users with an unwanted immediate app update.
//
// WARNING: this function should not be used in regular apps. This function
// is a user-unfriendly hack only for very special cases, like unstable
// rapidly-changing beta builds that would not run correctly if they were
// even one day out of date.
//
// Instead of this function you should set `SUAutomaticallyUpdate` to `YES`,
// which will gracefully install updates when the app quits.
//
// For UI-less/daemon apps that aren't usually quit, instead of this function,
// you can use the delegate method
// SUUpdaterDelegate::updater:willInstallUpdateOnQuit:immediateInstallationInvocation:
// or
// SUUpdaterDelegate::updater:willInstallUpdateOnQuit:immediateInstallationBlock:
// to immediately start installation when an update was found.
//
// A progress dialog is shown but the user will never be prompted to read the
// release notes.
//
// This function will cause update to be downloaded twice if automatic updates are
// enabled.
//
// You may want to respond to the userDidCancelDownload delegate method in case
// the user clicks the "Cancel" button while the update is downloading.
func InstallUpdatesIfAvailable() {
	C.installUpdatesIfAvailable()
}

// Returns the date of last update check.
func GetLastUpdateCheckDate() time.Time {
	s, n := math.Modf(float64(C.getLastUpdateCheckDate()))
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
func GetUpdateInProgress() bool {
	return C.getAutomaticallyChecksForUpdates() != 0
}
