package main

import (
	"log"
	"net/http"

	"github.com/abemedia/go-sparkle"
	webview "github.com/webview/webview_go"
)

func main() {
	// Start an embedded server for your appcast feed.
	// In a real application this would come from a remote source.
	go func() {
		log.Fatal(http.ListenAndServe(":3001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/xml")
			_, _ = w.Write([]byte(appcastFeed))
		})))
	}()

	// This is not actually needed and is just designed to manually trigger updates.
	// Importing github.com/abemedia/go-sparkle is enough to make your app check for updates on startup.
	sparkle.CheckForUpdates()

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
	</channel>
</rss>`
