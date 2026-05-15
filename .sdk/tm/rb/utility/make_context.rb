# Openfoodfacts SDK utility: make_context
require_relative '../core/context'
module OpenfoodfactsUtilities
  MakeContext = ->(ctxmap, basectx) {
    OpenfoodfactsContext.new(ctxmap, basectx)
  }
end
