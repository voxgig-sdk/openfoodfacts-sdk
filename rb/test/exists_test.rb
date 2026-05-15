# Openfoodfacts SDK exists test

require "minitest/autorun"
require_relative "../Openfoodfacts_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = OpenfoodfactsSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
