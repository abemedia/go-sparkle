# Go Sparkle

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/abemedia/go-sparkle?tab=doc)

This package provides (incomplete) go bindings for
[Sparkle](https://github.com/sparkle-project/Sparkle) created by Andy Matuschak.

Sparkle is a secure and reliable software update framework for Cocoa developers.

See <https://sparkle-project.org/> for more information about Sparkle.

## Documentation

See the [GoDoc](https://pkg.go.dev/github.com/abemedia/go-sparkle?tab=doc) and
[Sparkle's documentation](https://sparkle-project.org/documentation/#3-segue-for-security-concerns)
(you can ignore point 1 & 2 in the docs).

## Getting started

Download the Sparkle framework from <https://github.com/sparkle-project/Sparkle/releases/2.7.0> and
move the `Sparkle.framework` directory from the archive into your bundle's `Frameworks` directory
(e.g. `YourBundle.app/Contents/Frameworks/`).

Configure Sparkle's defaults using your bundle's `Info.plist` (see
<https://sparkle-project.org/documentation/customization/>). Only use the functions this package
exposes to allow your users to change the defaults.

Ensure your binary is built with `CGO_LDFLAGS` set to `-Wl,-rpath,@loader_path/../Frameworks` e.g.

```sh
CGO_LDFLAGS='-Wl,-rpath,@loader_path/../Frameworks' go build .
```

Publish your updates as an appcast. See <https://sparkle-project.org/documentation/publishing/>.

See the [example](./example/) for more details.

## Caveats

Sparkle requires a Cocoa run loop to work and as such can only be used by Go apps with a Cocoa UI
such as [webview/webview](https://github.com/webview/webview) or
[therecipe/qt](https://github.com/therecipe/qt).

## Migrating to v0.1.0

The `Get` prefix was removed from all function names to make them match C functions in the Sparkle
framework e.g. `GetFeedURL` is now `FeedURL` etc.
