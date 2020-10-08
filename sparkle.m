#import <Headers/SUUpdater.h>
#import <Foundation/Foundation.h>
#import <objc/runtime.h> // Required for class_addMethod()

static SUUpdater* updater = nil;

void initialize() {
    if(!updater) updater = [[SUUpdater sharedUpdater] retain];
}

void checkForUpdates() {
	[updater checkForUpdates:updater];
}

void checkForUpdatesInBackground() {
	[updater checkForUpdatesInBackground];
}

void setAutomaticallyChecksForUpdates(int check) {
	[updater setAutomaticallyChecksForUpdates:check];
}

int getAutomaticallyChecksForUpdates() {
	return [updater automaticallyChecksForUpdates];
}

void setAutomaticallyDownloadsUpdates(int check) {
	[updater setAutomaticallyDownloadsUpdates:check];
}

int getAutomaticallyDownloadsUpdates() {
	return [updater automaticallyDownloadsUpdates];
}

void setUpdateCheckInterval(int interval) {
	[updater setUpdateCheckInterval:interval];
}

int getUpdateCheckInterval() {
	return [updater updateCheckInterval];
}

void checkForUpdateInformation() {
	[updater checkForUpdateInformation];
}

void setFeedURL(const char *feedURL) {
	[updater setFeedURL:[NSURL URLWithString:@(feedURL)]];
}

const char* getFeedURL() {
	return [[[updater feedURL] absoluteString] UTF8String];
}

void setUserAgentString(const char *ua) {
	[updater setUserAgentString:@(ua)];
}

const char* getUserAgentString() {
	return [[updater userAgentString] UTF8String];
}

void setSendsSystemProfile(int check) {
	[updater setSendsSystemProfile:check];
}

int getSendsSystemProfile() {
	return [updater sendsSystemProfile];
}

void setDecryptionPassword(const char *pw) {
	[updater setDecryptionPassword:@(pw)];
}

const char* getDecryptionPassword() {
	return [[updater decryptionPassword] UTF8String];
}

void installUpdatesIfAvailable() {
	[updater installUpdatesIfAvailable];
}

double getLastUpdateCheckDate() {
	return [[updater lastUpdateCheckDate] timeIntervalSince1970];
}

void resetUpdateCycle() {
	[updater resetUpdateCycle];
}

int getUpdateInProgress() {
	return [updater updateInProgress];
}

// typedef int (*updaterMayCheckForUpdatesCallback_t)();
// typedef BOOL (^updaterMayCheckForUpdates_t)(id, SUUpdater*);

// // void setUpdaterMayCheckForUpdates(updaterMayCheckForUpdatesCallback_t callback) 
// // {
// // 	updaterMayCheckForUpdates_t updaterMayCheckForUpdates = ^(id self, SEL _cmd, SUUpdater* updater) { return callback(); };
// // 	class_addMethod([SUUpdater class], @selector(updaterMayCheckForUpdates:), (IMP)updaterMayCheckForUpdates, "v@:@");

// // 	// updater.updaterMayCheckForUpdates:updaterMayCheckForUpdates];
// // }

// void setUpdaterMayCheckForUpdates(updaterMayCheckForUpdatesCallback_t *callback) 
// {
// 	updaterMayCheckForUpdates_t updaterMayCheckForUpdates = ^(id self, SUUpdater *updater) { 
		
// 		FILE * fp;
// 		int i;
// 		/* open the file for writing*/
// 		fp = fopen ("/Users/adam/Work/go-sparkle/jah.log","w");

// 		/* write 10 lines of text into the file stream*/
// 		for(i = 0; i < 10;i++){
// 			fprintf (fp, "This is line %d\n",i + 1);
// 		}

// 		/* close the file*/  
// 		fclose (fp);

// 		return YES;

// 	 };

// 	// [[updater.delegate] updaterMayCheckForUpdates:updaterMayCheckForUpdates];

// 	// [updater.delegate 
// 	// 	updaterMayCheckForUpdates:^(id receiver, SEL _cmd, SUUpdater* updater) {
// 	// 		FILE * fp;
// 	// 		int i;
// 	// 		/* open the file for writing*/
// 	// 		fp = fopen ("/Users/adam/Work/go-sparkle/jah.log","w");

// 	// 		/* write 10 lines of text into the file stream*/
// 	// 		for(i = 0; i < 10;i++){
// 	// 			fprintf (fp, "This is line %d\n",i + 1);
// 	// 		}

// 	// 		/* close the file*/  
// 	// 		fclose (fp);

// 	// 		return 1;
// 	// 	}
// 	// ];
// 	// class_addMethod([SUUpdater class], @selector(updaterMayCheckForUpdates:), (IMP)updaterMayCheckForUpdates, "v@:@");

// 	// updater.updaterMayCheckForUpdates:updaterMayCheckForUpdates];
// }