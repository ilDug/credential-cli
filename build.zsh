#! /bin/zsh

# This script is used to build the project. It will compile the source code and
# generate the executable file, using linix , darwin and windows, (both arm64 and amd64).
# The script will generate the executable file in the `bin` directory.
# The script will also generate the `bin` directory if it does not exist.

# Check for Go installation and display the version
if ! command -v go &>/dev/null; then
    echo "Go is not installed. Please install Go to proceed."
    exit 1
else
    echo "Go version:"
    go version
fi

# Define an array with the target operating systems
os_list=("linux" "mac" "windows")

# Define an array with the target architectures
arch_list=("amd64" "arm64")

# Create the bin directory if it does not exist and subdirectories for each OS and architecture
for os in $os_list; do
    for arch in $arch_list; do
        mkdir -p bin/$os/$arch
    done
done

# Loop through each OS and architecture and build the project

GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/cre .
GOOS=linux GOARCH=arm64 go build -o bin/linux/arm64/cre .
GOOS=darwin GOARCH=amd64 go build -o bin/mac/amd64/cre .
GOOS=darwin GOARCH=arm64 go build -o bin/mac/arm64/cre .
GOOS=windows GOARCH=amd64 go build -o bin/windows/amd64/cre.exe .
GOOS=windows GOARCH=arm64 go build -o bin/windows/arm64/cre.exe .

# Display a message when the build is complete
echo "Build completed successfully."
