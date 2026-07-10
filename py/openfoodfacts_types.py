# Typed models for the Openfoodfacts SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class Product(TypedDict, total=False):
    code: str
    product: dict
    status: int
    status_verbose: str


class ProductLoadMatch(TypedDict):
    id: str


class Search(TypedDict, total=False):
    additives_tag: list
    allergen: str
    brand: str
    category: str
    country: str
    created_t: int
    ecoscore_grade: str
    ecoscore_score: int
    generic_name: str
    image_front_url: str
    image_ingredients_url: str
    image_nutrition_url: str
    image_url: str
    ingredients_analysis_tag: list
    ingredients_text: str
    label: str
    last_modified_t: int
    manufacturing_place: str
    nova_group: int
    nutriment: dict
    nutriscore_grade: str
    nutriscore_score: int
    packaging: str
    product_name: str
    quantity: str
    store: str
    trace: str


class SearchListMatch(TypedDict, total=False):
    additives_tag: list
    allergen: str
    brand: str
    category: str
    country: str
    created_t: int
    ecoscore_grade: str
    ecoscore_score: int
    generic_name: str
    image_front_url: str
    image_ingredients_url: str
    image_nutrition_url: str
    image_url: str
    ingredients_analysis_tag: list
    ingredients_text: str
    label: str
    last_modified_t: int
    manufacturing_place: str
    nova_group: int
    nutriment: dict
    nutriscore_grade: str
    nutriscore_score: int
    packaging: str
    product_name: str
    quantity: str
    store: str
    trace: str
