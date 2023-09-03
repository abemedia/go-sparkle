#!/bin/sh

VERSION="2.4.2"

wget -O sparkle.tar.xz "https://github.com/sparkle-project/Sparkle/releases/download/$VERSION/Sparkle-$VERSION.tar.xz"
mkdir sparkle
tar -xf sparkle.tar.xz --directory ./sparkle
rm -f sparkle.tar.xz

find Sparkle.framework -type f ! -name '*.go' -delete
cp -Lr sparkle/Sparkle.framework/Sparkle Sparkle.framework/
cp -Lr sparkle/Sparkle.framework/Headers Sparkle.framework/
cp -Lr sparkle/Sparkle.framework/PrivateHeaders Sparkle.framework/
rm -rf sparkle

sed -i "s#Sparkle/releases/.*>#Sparkle/releases/$VERSION>#g" README.md
