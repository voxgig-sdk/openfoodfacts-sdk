# Openfoodfacts SDK feature factory

from feature.base_feature import OpenfoodfactsBaseFeature
from feature.test_feature import OpenfoodfactsTestFeature


def _make_feature(name):
    features = {
        "base": lambda: OpenfoodfactsBaseFeature(),
        "test": lambda: OpenfoodfactsTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
