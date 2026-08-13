# Openfoodfacts SDK exists test

import pytest
from openfoodfacts_sdk import OpenfoodfactsSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = OpenfoodfactsSDK.test(None, None)
        assert testsdk is not None
