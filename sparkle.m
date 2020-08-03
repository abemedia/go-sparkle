#import <Headers/SUUpdater.h>
#import <Foundation/Foundation.h>
#import <objc/runtime.h> // Required for class_addMethod()

static SUUpdater* updater = nil;

void initialize() {
    if(!updater) updater = [[SUUpdater sharedUpdater] retain];
}

void checkForUpdates()
{
	[updater checkForUpdates:updater];
}

void checkForUpdatesInBackground()
{
	[updater checkForUpdatesInBackground];
}

void setAutomaticallyChecksForUpdates(int check)
{
	[updater setAutomaticallyChecksForUpdates:check];
}

int getAutomaticallyChecksForUpdates() 
{
	return [updater automaticallyChecksForUpdates];
}

void setAutomaticallyDownloadsUpdates(int check)
{
	[updater setAutomaticallyDownloadsUpdates:check];
}

int getAutomaticallyDownloadsUpdates() 
{
	return [updater automaticallyDownloadsUpdates];
}

void setUpdateCheckInterval(int interval)
{
	[updater setUpdateCheckInterval:interval];
}

int getUpdateCheckInterval() 
{
	return [updater updateCheckInterval];
}

void checkForUpdateInformation()
{
	[updater checkForUpdateInformation];
}

void setFeedURL(const char *feedURL)
{
	[updater setFeedURL:[NSURL URLWithString:@(feedURL)]];
}

const char* getFeedURL()
{
	return [[[updater feedURL] absoluteString] UTF8String];
}

void setUserAgentString(const char *ua)
{
	[updater setUserAgentString:@(ua)];
}

const char* getUserAgentString()
{
	return [[updater userAgentString] UTF8String];
}

void setSendsSystemProfile(int check)
{
	[updater setSendsSystemProfile:check];
}

int getSendsSystemProfile() 
{
	return [updater sendsSystemProfile];
}

void setDecryptionPassword(const char *pw)
{
	[updater setDecryptionPassword:@(pw)];
}

const char* getDecryptionPassword()
{
	return [[updater decryptionPassword] UTF8String];
}

void installUpdatesIfAvailable()
{
	[updater installUpdatesIfAvailable];
}

double getLastUpdateCheckDate()
{
	return [[updater lastUpdateCheckDate] timeIntervalSince1970];
}

void resetUpdateCycle()
{
	[updater resetUpdateCycle];
}

int getUpdateInProgress() 
{
	return [updater updateInProgress];
}

