# Openfoodfacts SDK feature factory

from openfoodfacts_sdk.feature.base_feature import OpenfoodfactsBaseFeature
from openfoodfacts_sdk.feature.test_feature import OpenfoodfactsTestFeature


def _make_feature(name):
    features = {
        "base": lambda: OpenfoodfactsBaseFeature(),
        "test": lambda: OpenfoodfactsTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
