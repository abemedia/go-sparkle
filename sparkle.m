#define BUILDING_SPARKLE_SOURCES_EXTERNALLY

#import "_cgo_export.h"
#import <Foundation/Foundation.h>
#import <Headers/SPUStandardUpdaterController.h>
#import <Headers/SPUUpdater.h>
#import <Headers/SPUUpdaterDelegate.h>
#import <objc/runtime.h> // Required for class_addMethod()

@interface SparkleUpdaterDelegate : NSObject <SPUUpdaterDelegate>

@property(nonatomic, strong) NSString *decryptionPassword;

@property(nonatomic, strong) NSSet<NSString *> *allowedChannelsForUpdater;

@end

static SparkleUpdaterDelegate *delegate = nil;
static SPUStandardUpdaterController *updateController = nil;

@implementation SparkleUpdaterDelegate

- (NSString *)feedURLStringForUpdater:(SPUUpdater *)updater {
  char *url = CGOFeedURL();
  if (url != NULL) {
    NSString *u = @(url);
    free(url);
    return u;
  }
  return nil;
}

- (NSString *)decryptionPasswordForUpdater:(SPUUpdater *)updater {
  return self.decryptionPassword;
}

- (NSSet<NSString *> *)allowedChannelsForUpdater:(SPUUpdater *)updater {
  return self.allowedChannelsForUpdater;
}

@end

void sparkle_initialize() {
  if (!updateController) {
    delegate = [[SparkleUpdaterDelegate alloc] init];
    updateController =
        [[SPUStandardUpdaterController alloc] initWithStartingUpdater:true
                                                      updaterDelegate:delegate
                                                   userDriverDelegate:nil];
    /*
    Clears any feed URL from the host bundle’s user defaults that was set via
    -setFeedURL:

    You should call this method if you have used -setFeedURL: in the past and
    want to stop using that API. Otherwise for compatibility Sparkle will prefer
    to use the feed URL that was set in the user defaults over the one that was
    specified in the host bundle’s Info.plist, which is often undesirable
    (except for testing purposes).

    If a feed URL is found stored in the host bundle’s user defaults (from
    calling -setFeedURL:) before it gets cleared, then that previously set URL
    is returned from this method.

    For dynamic URL setting we use now delegate method feedURLStringForUpdater:
    */
    [updateController.updater clearFeedURLFromUserDefaults];
  }
}

void sparkle_checkForUpdates() { [updateController checkForUpdates:nil]; }

void sparkle_checkForUpdatesInBackground() {
  [updateController.updater checkForUpdatesInBackground];
}

void sparkle_setAutomaticallyChecksForUpdates(int check) {
  [updateController.updater setAutomaticallyChecksForUpdates:check];
}

int sparkle_automaticallyChecksForUpdates() {
  return [updateController.updater automaticallyChecksForUpdates];
}

void sparkle_setAutomaticallyDownloadsUpdates(int check) {
  [updateController.updater setAutomaticallyDownloadsUpdates:check];
}

int sparkle_automaticallyDownloadsUpdates() {
  return [updateController.updater automaticallyDownloadsUpdates];
}

void sparkle_setUpdateCheckInterval(int interval) {
  [updateController.updater setUpdateCheckInterval:interval];
}

int sparkle_updateCheckInterval() {
  return [updateController.updater updateCheckInterval];
}

void sparkle_checkForUpdateInformation() {
  [updateController.updater checkForUpdateInformation];
}

void sparkle_setUserAgentString(const char *ua) {
  [updateController.updater setUserAgentString:@(ua)];
}

const char *sparkle_userAgentString() {
  return [[updateController.updater userAgentString] UTF8String];
}

void sparkle_setSendsSystemProfile(int check) {
  [updateController.updater setSendsSystemProfile:check];
}

int sparkle_sendsSystemProfile() {
  return [updateController.updater sendsSystemProfile];
}

void sparkle_setDecryptionPassword(const char *pw) {
  delegate.decryptionPassword = @(pw);
}

const char *sparkle_decryptionPassword() {
  return [delegate.decryptionPassword UTF8String];
}

void sparkle_setAllowedChannelsForUpdater(const char **channels, int length) {
  NSMutableArray *array = [NSMutableArray arrayWithCapacity:length];
  for (int i = 0; i < length; i++) {
    [array addObject:[NSString stringWithUTF8String:channels[i]]];
  }
  NSSet *set = [NSSet setWithArray:array];
  delegate.allowedChannelsForUpdater = set;
}

double sparkle_lastUpdateCheckDate() {
  return [[updateController.updater lastUpdateCheckDate] timeIntervalSince1970];
}

void sparkle_resetUpdateCycle() { [updateController.updater resetUpdateCycle]; }

int sparkle_updateInProgress() {
  return [updateController.updater sessionInProgress];
}
