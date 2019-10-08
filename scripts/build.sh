#!/bin/bash

echo "Preparing..."

cd "$(dirname "$0")"
cd "$(realpath $PWD/..)"
MYDIR=$PWD

# Make sure we have the go binary
go version > /dev/null

# GOPATH may contain multiple paths separated by ":"
GOPATH1=$(go env GOPATH | cut -f1 -d:)

if [[ $PWD != *"/src/github.com/shrotavre/filetinder" ]] ; then
	echo "Warning: Building outside of GOPATH will most likely fail."
	echo "         Please rename $PWD to $GOPATH1/src/github.com/shrotavre/filetinder ."
	sleep 5
	echo
fi

# gocryptfs version according to git or a VERSION file
if [[ -d .git ]] ; then
	GITVERSION=$(git describe --tags --dirty)
elif [[ -f VERSION ]] ; then
	GITVERSION=$(cat VERSION)
else
	echo "Warning: could not determine gocryptfs version"
	GITVERSION="[unknown]"
fi

# Build date, something like "2017-09-06". Don't override BUILDDATE
# if it is already set. This may be done for reproducible builds.
if [[ -z ${BUILDDATE:-} ]] ; then
	BUILDDATE=$(date +%Y-%m-%d)
fi

# If SOURCE_DATE_EPOCH is set, it overrides BUILDDATE. This is the
# standard environment variable for faking the date in reproducible builds.
if [[ -n ${SOURCE_DATE_EPOCH:-} ]] ; then
	BUILDDATE=$(date --utc --date="@${SOURCE_DATE_EPOCH}" +%Y-%m-%d)
fi

# For reproducible builds, we get rid of $HOME references in the binary
# using "-trimpath".
# Note: we have to set both -gcflags and -asmflags because otherwise
# "$HOME/go/src/golang.org/x/sys/unix/asm_linux_amd64.s" stays in the binary.
GV=$(go version)
if [[ $GV == *"1.7"* ]] || [[ $GV == *"1.8"* ]] || [[ $GV == *"1.9"* ]] ; then
	TRIM="-trimpath=${GOPATH1}/src"
else
	# Go 1.10 changed the syntax. You now have to prefix "all=" to affect
	# all compiled packages.
	TRIM="all=-trimpath=${GOPATH1}/src"
fi

GO_LDFLAGS="-X main.GitVersion=$GITVERSION -X main.BuildDate=$BUILDDATE"

# If LDFLAGS is set, add it as "-extldflags".
if [[ -n ${LDFLAGS:-} ]] ; then
	GO_LDFLAGS="$GO_LDFLAGS \"-extldflags=$LDFLAGS\""
fi

rm -rf dist/*

echo "(1/4) Building binaries..."
go build "-ldflags=$GO_LDFLAGS" "-gcflags=$TRIM" "-asmflags=$TRIM" "$@" -o dist/tmp/filetinder cmd/filetinder/*

echo "(2/4) Building UI.."
pushd ui > /dev/null && yarn run --silent build:silent && popd > /dev/null && cp -ar ui/public dist/tmp/ui

echo "(3/4) Packing up..."
pushd dist/tmp > /dev/null && tar -czf filetinder.tar.gz ui filetinder && popd > /dev/null

echo "(4/4) Cleaning..."
cp dist/tmp/filetinder.tar.gz dist/filetinder.tar.gz
rm -rf dist/tmp/

echo "Finished!"