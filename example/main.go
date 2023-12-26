package main

import (
	"log"
	"net/http"

	webview "github.com/webview/webview_go"

	"github.com/abemedia/go-sparkle"
)

func main() {
	// Start an embedded server for your appcast feed.
	// In a real application this would come from a remote source.
	go func() {
		log.Fatal(http.ListenAndServe(":3001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("Serving appcast feed", r.URL)
			w.Header().Set("Content-Type", "application/xml")
			_, _ = w.Write([]byte(appcastFeed))
		})))
	}()

	// It's encouraged to set the url in your Info.plist file and consider channels,
	// but you can also set them programmatically.
	sparkle.SetFeedURL("http://localhost:3001/appcast.xml")

	// You can set also alternative channels as an addition to the default one.
	// More info: https://sparkle-project.org/documentation/publishing/#channels
	sparkle.SetAllowedChannelsForUpdater("beta", "rc")
	//sparkle.SetAllowedChannelsForUpdater() // call with no args to reset to default behavior

	// This is not actually needed and is just designed to manually trigger updates.
	// Importing github.com/abemedia/go-sparkle is enough to make your app check for updates on startup.
	sparkle.CheckForUpdates()

	log.Println("Updates check", sparkle.FeedURL(), "by", sparkle.UserAgentString())

	// You can also set the decryption password for DMG programmatically
	sparkle.SetDecryptionPassword("password")
	log.Printf("DMG decryption pwd set to <%s>", sparkle.DecryptionPassword())

	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Sparkle Example")
	w.SetSize(800, 600, webview.HintNone)
	w.SetHtml("<center><h1>Hello World</h1></center>")
	w.Run()
}

const appcastFeed = `<?xml version="1.0" encoding="utf-8"?>
<rss version="2.0" xmlns:sparkle="http://www.andymatuschak.org/xml-namespaces/sparkle"	xmlns:dc="http://purl.org/dc/elements/1.1/">
	<channel>
		<title>Your Great App's Changelog</title>
		<link>http://example.com/appcast.xml</link>
		<description>Most recent changes with links to updates.</description>
		<language>en</language>
		<item>
			<title>Version 2.0</title>
			<sparkle:version>2.0.0</sparkle:version>
			<description>This is an update.</description>
			<pubDate>Mon, 28 Jan 2013 14:30:00 +0500</pubDate>
			<enclosure url="http://example.com/my_app_v2.zip" sparkle:edSignature="7cLALFUHSwvEJWSkV8aMreoBe4fhRa4FncC5NoThKxwThL6FDR7hTiPJh1fo2uagnPogisnQsgFgq6mGkt2RBw==" length="1623481" type="application/octet-stream" />
		</item>
		<item>
			<title>Version 2.0.1 (Beta 1)</title>
			<sparkle:version>2.0.1-beta</sparkle:version>
 			<sparkle:channel>beta</sparkle:channel>
			<description>This is a beta update which could be explicitly allowed.</description>
			<pubDate>Mon, 28 Jan 2013 14:30:00 +0500</pubDate>
			<enclosure url="http://example.com/my_app_v2.0.1beta.zip" sparkle:edSignature="7cLALFUHSwvEJWSkV8aMreoBe4fhRa4FncC5NoThKxwThL6FDR7hTiPJh1fo2uagnPogisnQsgFgq6mGkt2RBw==" length="1623481" type="application/octet-stream" />
		</item>
	</channel>
</rss>`
