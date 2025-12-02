#!/bin/bash
# DMMVC Uninstallation Script

set -e

echo "========================================"
echo "DMMVC Uninstallation"
echo "========================================"
echo ""

GOPATH=${GOPATH:-$(go env GOPATH)}
GOBIN=${GOBIN:-$GOPATH/bin}
CLI_PATH="$GOBIN/dmmvc"

if [ -f "$CLI_PATH" ]; then
    echo "Removing CLI from: $CLI_PATH"
    rm -f "$CLI_PATH"
    echo "✓ CLI removed successfully"
else
    echo "CLI not found at: $CLI_PATH"
    echo "Nothing to uninstall"
fi

echo ""
echo "Uninstallation complete!"
echo ""
echo "To remove the project directory, run:"
echo "  rm -rf $(pwd)"
echo ""
