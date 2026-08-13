# Openfoodfacts SDK utility: make_context

from openfoodfacts_sdk.core.context import OpenfoodfactsContext


def make_context_util(ctxmap, basectx):
    return OpenfoodfactsContext(ctxmap, basectx)
