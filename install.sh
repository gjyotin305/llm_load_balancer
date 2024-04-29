# Check if script is run as root
if [ "$(id -u)" -ne 0 ]; then
    echo "This script must be run as root. Restarting as root..."
    sudo "$0" "$@"
    exit $?
fi

# Define directories
binDir="bin"
installDir="/usr/bin"

# Create installation directories if they don't exist
mkdir -p "$configInstallDir"
mkdir -p "$readmeDir"
mkdir -p "$licenseDir"

# Copy binaries
cp "$binDir/jpt" "$installDir"

echo "Installation of jpt completed successfully!"

