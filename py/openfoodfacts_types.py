# Typed models for the Openfoodfacts SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.

from __future__ import annotations

from dataclasses import dataclass
from typing import Optional, Any


@dataclass
class Product:
    code: Optional[str] = None
    product: Optional[dict] = None
    status: Optional[int] = None
    status_verbose: Optional[str] = None


@dataclass
class ProductLoadMatch:
    barcode: str
    id: str


@dataclass
class Search:
    additives_tag: Optional[list] = None
    allergen: Optional[str] = None
    brand: Optional[str] = None
    category: Optional[str] = None
    country: Optional[str] = None
    created_t: Optional[int] = None
    ecoscore_grade: Optional[str] = None
    ecoscore_score: Optional[int] = None
    generic_name: Optional[str] = None
    image_front_url: Optional[str] = None
    image_ingredients_url: Optional[str] = None
    image_nutrition_url: Optional[str] = None
    image_url: Optional[str] = None
    ingredients_analysis_tag: Optional[list] = None
    ingredients_text: Optional[str] = None
    label: Optional[str] = None
    last_modified_t: Optional[int] = None
    manufacturing_place: Optional[str] = None
    nova_group: Optional[int] = None
    nutriment: Optional[dict] = None
    nutriscore_grade: Optional[str] = None
    nutriscore_score: Optional[int] = None
    packaging: Optional[str] = None
    product_name: Optional[str] = None
    quantity: Optional[str] = None
    store: Optional[str] = None
    trace: Optional[str] = None


@dataclass
class SearchListMatch:
    additives_tag: Optional[list] = None
    allergen: Optional[str] = None
    brand: Optional[str] = None
    category: Optional[str] = None
    country: Optional[str] = None
    created_t: Optional[int] = None
    ecoscore_grade: Optional[str] = None
    ecoscore_score: Optional[int] = None
    generic_name: Optional[str] = None
    image_front_url: Optional[str] = None
    image_ingredients_url: Optional[str] = None
    image_nutrition_url: Optional[str] = None
    image_url: Optional[str] = None
    ingredients_analysis_tag: Optional[list] = None
    ingredients_text: Optional[str] = None
    label: Optional[str] = None
    last_modified_t: Optional[int] = None
    manufacturing_place: Optional[str] = None
    nova_group: Optional[int] = None
    nutriment: Optional[dict] = None
    nutriscore_grade: Optional[str] = None
    nutriscore_score: Optional[int] = None
    packaging: Optional[str] = None
    product_name: Optional[str] = None
    quantity: Optional[str] = None
    store: Optional[str] = None
    trace: Optional[str] = None

