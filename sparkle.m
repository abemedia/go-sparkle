#define BUILDING_SPARKLE_SOURCES_EXTERNALLY

#import <Foundation/Foundation.h>
#import <Headers/SUUpdater.h>
#import <objc/runtime.h> // Required for class_addMethod()

static SUUpdater *updater = nil;

void initialize() {
  if (!updater)
    updater = [[SUUpdater sharedUpdater] retain];
}

void checkForUpdates() { [updater checkForUpdates:updater]; }

void checkForUpdatesInBackground() { [updater checkForUpdatesInBackground]; }

void setAutomaticallyChecksForUpdates(int check) {
  [updater setAutomaticallyChecksForUpdates:check];
}

int automaticallyChecksForUpdates() {
  return [updater automaticallyChecksForUpdates];
}

void setAutomaticallyDownloadsUpdates(int check) {
  [updater setAutomaticallyDownloadsUpdates:check];
}

int automaticallyDownloadsUpdates() {
  return [updater automaticallyDownloadsUpdates];
}

void setUpdateCheckInterval(int interval) {
  [updater setUpdateCheckInterval:interval];
}

int updateCheckInterval() { return [updater updateCheckInterval]; }

void checkForUpdateInformation() { [updater checkForUpdateInformation]; }

void setFeedURL(const char *feedURL) {
  [updater setFeedURL:[NSURL URLWithString:@(feedURL)]];
}

const char *feedURL() {
  return [[[updater feedURL] absoluteString] UTF8String];
}

void setUserAgentString(const char *ua) { [updater setUserAgentString:@(ua)]; }

const char *userAgentString() { return [[updater userAgentString] UTF8String]; }

void setSendsSystemProfile(int check) { [updater setSendsSystemProfile:check]; }

int sendsSystemProfile() { return [updater sendsSystemProfile]; }

void setDecryptionPassword(const char *pw) {
  [updater setDecryptionPassword:@(pw)];
}

const char *decryptionPassword() {
  return [[updater decryptionPassword] UTF8String];
}

double lastUpdateCheckDate() {
  return [[updater lastUpdateCheckDate] timeIntervalSince1970];
}

void resetUpdateCycle() { [updater resetUpdateCycle]; }

int updateInProgress() { return [updater updateInProgress]; }
