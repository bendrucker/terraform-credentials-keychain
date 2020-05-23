source = ["./dist/macos_darwin_amd64/terraform-credentials-keychain"]
bundle_id = "com.bendrucker.terraform-credentials-keychain"

apple_id {
  username = "bvdrucker@gmail.com"
}

sign {
  application_identity = "Developer ID Application: Benjamin V Drucker"
}

zip {
  output_path = "./dist/macos_darwin_amd64/terraform-credentials-keychain.zip"
}
