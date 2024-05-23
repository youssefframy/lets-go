# Set environment variables
$env:GOOS = "linux"
$env:GOARCH = "amd64"

# Build the binary
go build -o bootstrap

# Create a zip archive
Compress-Archive bootstrap function.zip
