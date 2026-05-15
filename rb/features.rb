# Openfoodfacts SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module OpenfoodfactsFeatures
  def self.make_feature(name)
    case name
    when "base"
      OpenfoodfactsBaseFeature.new
    when "test"
      OpenfoodfactsTestFeature.new
    else
      OpenfoodfactsBaseFeature.new
    end
  end
end
