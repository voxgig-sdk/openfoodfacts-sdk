package = "voxgig-sdk-openfoodfacts"
version = "0.0-1"
source = {
  url = "git://github.com/voxgig-sdk/openfoodfacts-sdk.git"
}
description = {
  summary = "Openfoodfacts SDK for Lua",
  license = "MIT"
}
dependencies = {
  "lua >= 5.3",
  "dkjson >= 2.5",
  "dkjson >= 2.5",
}
build = {
  type = "builtin",
  modules = {
    ["openfoodfacts_sdk"] = "openfoodfacts_sdk.lua",
    ["config"] = "config.lua",
    ["features"] = "features.lua",
  }
}
