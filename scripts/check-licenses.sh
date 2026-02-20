#!/bin/bash
# License Compliance Check Script
# This script checks Go dependencies for license compliance
# Alert ID: SB-LICENSE-POLICY-VIOLATION::bsd-2-clause

set -e

echo "=== License Compliance Check ==="
echo "Checking dependencies for license policy violations..."
echo ""

# Check if go-licenses tool is installed
if ! command -v go-licenses &> /dev/null; then
    echo "Installing go-licenses tool..."
    go install github.com/google/go-licenses@latest
fi

# Generate license report
echo "Generating license report..."
go-licenses report ./... 2>/dev/null || true

echo ""
echo "=== Checking for prohibited licenses ==="

# Check for BSD-2-Clause licenses
if go-licenses report ./... 2>/dev/null | grep -i "BSD-2-Clause"; then
    echo "❌ VIOLATION: BSD-2-Clause license detected"
    echo "   This license is prohibited by organizational policy"
    echo "   See .licenserc.yaml for approved alternatives"
    exit 1
fi

# Check for GPL licenses
if go-licenses report ./... 2>/dev/null | grep -iE "GPL-[0-9]"; then
    echo "❌ VIOLATION: GPL license detected"
    echo "   GPL licenses are prohibited by organizational policy"
    exit 1
fi

echo "✅ No prohibited licenses detected"
echo ""
echo "For full compliance review, see .licenserc.yaml"
