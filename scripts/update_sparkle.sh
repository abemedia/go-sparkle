#!/bin/sh -e

VERSION="2.5.1"
DIR=$(mktemp -d)

wget -q -O $DIR/Sparkle-$VERSION.tar.xz "https://github.com/sparkle-project/Sparkle/releases/download/$VERSION/Sparkle-$VERSION.tar.xz"
tar -xf $DIR/Sparkle-$VERSION.tar.xz --directory $DIR

find Sparkle.framework -type f ! -name '*.go' -delete
cp -Lr $DIR/Sparkle.framework/Sparkle Sparkle.framework/
cp -Lr $DIR/Sparkle.framework/Headers Sparkle.framework/
cp -Lr $DIR/Sparkle.framework/PrivateHeaders Sparkle.framework/
rm -rf $DIR

sed -i '' "s#Sparkle/releases/.*>#Sparkle/releases/$VERSION>#g" README.md
