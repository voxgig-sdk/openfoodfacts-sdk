# Openfoodfacts SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/ratelimit_feature'
require_relative 'feature/retry_feature'
require_relative 'feature/test_feature'
require_relative 'feature/timeout_feature'


module OpenfoodfactsFeatures
  def self.make_feature(name)
    case name
    when "base"
      OpenfoodfactsBaseFeature.new
    when "ratelimit"
      OpenfoodfactsRatelimitFeature.new
    when "retry"
      OpenfoodfactsRetryFeature.new
    when "test"
      OpenfoodfactsTestFeature.new
    when "timeout"
      OpenfoodfactsTimeoutFeature.new
    else
      OpenfoodfactsBaseFeature.new
    end
  end
end
