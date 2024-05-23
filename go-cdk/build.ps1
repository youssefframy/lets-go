# move into lambda directory
cd lambda

# remove any existing function.zip or bootstrap if they exist
Remove-Item -Path function.zip -ErrorAction SilentlyContinue
Remove-Item -Path bootstrap -ErrorAction SilentlyContinue

# Set environment variables
$env:GOOS = "linux"
$env:GOARCH = "amd64"

# Build the binary
go build -o bootstrap

# Create a zip archive
Compress-Archive bootstrap function.zip

# get back to the current directory
cd ..
