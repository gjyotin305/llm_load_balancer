# Check if script is run as root
if [ "$(id -u)" -ne 0 ]; then
    echo "This script must be run as root. Restarting as root..."
    sudo "$0" "$@"
    exit $?
fi

# Define directories
binDir="bin"
installDir="/usr/bin"
readmeDir="/usr/share/doc/jpt"
licenseDir="/usr/share/licenses/jpt"

# Create installation directories if they don't exist
mkdir -p "$configInstallDir"
mkdir -p "$readmeDir"
mkdir -p "$licenseDir"

# Move Binaries and README and license
mv "$binDir/jpt" "$installDir"
mv "README.md" "$readmeDir"
mv "LICENSE" "$licenseDir"

echo "Installation of jpt completed successfully!"

