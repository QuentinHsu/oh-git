#!/bin/zsh
# Used to generate command info
#
# chmod +x ./release.sh && ./release.sh
#
cd "$(dirname "$0")" # Switch to the directory where the script is located.
echo "Generating Release Info\n"

Version="0.0.1 beta2"
Release=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
echo "Version: $Version, Release: $Release\n"

mkdir -p ./pkg/info

echo "package info" > ./pkg/info/main.go
echo "" >> ./pkg/info/main.go
echo "var Version = \"$Version\"" >> ./pkg/info/main.go
echo "var Release = \"$Release\"" >> ./pkg/info/main.go
if [ -f ./pkg/info/main.go ]; then
  echo "File exists, overwriting...\n"
  echo "package info" > ./pkg/info/main.go
  echo "" >> ./pkg/info/main.go
  echo "var Version = \"$Version\"" >> ./pkg/info/main.go
  echo "var Release = \"$Release\"" >> ./pkg/info/main.go
fi
echo "To be reinstalled ohgit\n"
if go install ./; then
  echo "Ok!"
else
  echo "\ngo install failed!"
fi