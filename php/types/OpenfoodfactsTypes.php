<?php
declare(strict_types=1);

// Typed models for the Openfoodfacts SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** Product entity data model. */
class Product
{
    public ?string $code = null;
    public ?array $product = null;
    public ?int $status = null;
    public ?string $status_verbose = null;
}

/** Request payload for Product#load. */
class ProductLoadMatch
{
    public string $barcode;
    public string $id;
}

/** Search entity data model. */
class Search
{
    public ?array $additives_tag = null;
    public ?string $allergen = null;
    public ?string $brand = null;
    public ?string $category = null;
    public ?string $country = null;
    public ?int $created_t = null;
    public ?string $ecoscore_grade = null;
    public ?int $ecoscore_score = null;
    public ?string $generic_name = null;
    public ?string $image_front_url = null;
    public ?string $image_ingredients_url = null;
    public ?string $image_nutrition_url = null;
    public ?string $image_url = null;
    public ?array $ingredients_analysis_tag = null;
    public ?string $ingredients_text = null;
    public ?string $label = null;
    public ?int $last_modified_t = null;
    public ?string $manufacturing_place = null;
    public ?int $nova_group = null;
    public ?array $nutriment = null;
    public ?string $nutriscore_grade = null;
    public ?int $nutriscore_score = null;
    public ?string $packaging = null;
    public ?string $product_name = null;
    public ?string $quantity = null;
    public ?string $store = null;
    public ?string $trace = null;
}

/** Match filter for Search#list (any subset of Search fields). */
class SearchListMatch
{
    public ?array $additives_tag = null;
    public ?string $allergen = null;
    public ?string $brand = null;
    public ?string $category = null;
    public ?string $country = null;
    public ?int $created_t = null;
    public ?string $ecoscore_grade = null;
    public ?int $ecoscore_score = null;
    public ?string $generic_name = null;
    public ?string $image_front_url = null;
    public ?string $image_ingredients_url = null;
    public ?string $image_nutrition_url = null;
    public ?string $image_url = null;
    public ?array $ingredients_analysis_tag = null;
    public ?string $ingredients_text = null;
    public ?string $label = null;
    public ?int $last_modified_t = null;
    public ?string $manufacturing_place = null;
    public ?int $nova_group = null;
    public ?array $nutriment = null;
    public ?string $nutriscore_grade = null;
    public ?int $nutriscore_score = null;
    public ?string $packaging = null;
    public ?string $product_name = null;
    public ?string $quantity = null;
    public ?string $store = null;
    public ?string $trace = null;
}

