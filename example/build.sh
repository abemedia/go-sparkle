#/bin/bash -e

# Customise with your own developer ID.
CODE_SIGN_IDENTITY="Developer ID Application: Adam Bouqdib"
BINARY="sparkle-example"
BUNDLE="Example.app"
SPARKLE_VERSION="2.5.1"

# Build Go app.
CGO_LDFLAGS='-Wl,-rpath,@loader_path/../Frameworks' go build -o $BUNDLE/Contents/MacOS/$BINARY .

SPARKLE_PATH="$BUNDLE/Contents/Frameworks/Sparkle.framework"
DIR=$(mktemp -d)

# Download Sparkle Framework
URL=https://github.com/sparkle-project/Sparkle/releases/download/$SPARKLE_VERSION/Sparkle-$SPARKLE_VERSION.tar.xz
curl -fsSL $URL | tar -xJ -C "$DIR" ./Sparkle.framework

# Delete source files
for name in Headers PrivateHeaders Modules; do
  rm -rf "$(readlink -f "$DIR/Sparkle.framework/$name")"
  rm -rf "$DIR/Sparkle.framework/$name"
done

# Add to bundle
rm -rf $SPARKLE_PATH
mkdir -p $BUNDLE/Contents/Frameworks
mv $DIR/Sparkle.framework/ $SPARKLE_PATH

# Sign code
echo "\nSigning..."
# Uncomment for sandboxed apps. See https://sparkle-project.org/documentation/sandboxing/
# codesign --force --sign "$CODE_SIGN_IDENTITY" --options runtime $SPARKLE_PATH/Versions/B/XPCServices/Installer.xpc
# codesign --force --sign "$CODE_SIGN_IDENTITY" --options runtime --entitlements $DIR/Entitlements/Downloader.entitlements $SPARKLE_PATH/Versions/B/XPCServices/Downloader.xpc
# codesign --force --sign "$CODE_SIGN_IDENTITY" --options runtime $SPARKLE_PATH/Versions/B/Autoupdate
# codesign --force --sign "$CODE_SIGN_IDENTITY" --options runtime $SPARKLE_PATH/Versions/B/Updater.app
codesign --force --sign "$CODE_SIGN_IDENTITY" --options runtime $SPARKLE_PATH
codesign --force --sign "$CODE_SIGN_IDENTITY" --options runtime $BUNDLE

echo "\nChecking signature..."
codesign --deep --verify -vvvv $BUNDLE

# Cleanup
rm -rf $DIR
