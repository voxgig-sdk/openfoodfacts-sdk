# Openfoodfacts SDK utility: feature_add
module OpenfoodfactsUtilities
  FeatureAdd = ->(ctx, f) {
    ctx.client.features << f
  }
end
