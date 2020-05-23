#!/usr/bin/env bash

set -euf -o pipefail

binary="$1"
zip="$2"
config="$(dirname "$binary")/gon.hcl"

cat <<-EOF > "$config"
source = ["$binary"]
bundle_id = "com.bendrucker.terraform-credentials-keychain"

apple_id {
  username = "bvdrucker@gmail.com"
}

sign {
  application_identity = "Developer ID Application: Benjamin V Drucker"
}

zip {
  output_path = "$zip"
}
EOF

gon "$config"