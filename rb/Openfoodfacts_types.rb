# frozen_string_literal: true

# Typed models for the Openfoodfacts SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# Product entity data model.
#
# @!attribute [rw] additives_tags
#   @return [Array, nil]
#
# @!attribute [rw] allergens
#   @return [String, nil]
#
# @!attribute [rw] brands
#   @return [String, nil]
#
# @!attribute [rw] categories
#   @return [String, nil]
#
# @!attribute [rw] countries
#   @return [String, nil]
#
# @!attribute [rw] created_t
#   @return [Integer, nil]
#
# @!attribute [rw] ecoscore_grade
#   @return [String, nil]
#
# @!attribute [rw] ecoscore_score
#   @return [Integer, nil]
#
# @!attribute [rw] generic_name
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] image_front_url
#   @return [String, nil]
#
# @!attribute [rw] image_ingredients_url
#   @return [String, nil]
#
# @!attribute [rw] image_nutrition_url
#   @return [String, nil]
#
# @!attribute [rw] image_url
#   @return [String, nil]
#
# @!attribute [rw] ingredients_analysis_tags
#   @return [Array, nil]
#
# @!attribute [rw] ingredients_text
#   @return [String, nil]
#
# @!attribute [rw] labels
#   @return [String, nil]
#
# @!attribute [rw] last_modified_t
#   @return [Integer, nil]
#
# @!attribute [rw] manufacturing_places
#   @return [String, nil]
#
# @!attribute [rw] nova_group
#   @return [Integer, nil]
#
# @!attribute [rw] nutriments
#   @return [Hash, nil]
#
# @!attribute [rw] nutriscore_grade
#   @return [String, nil]
#
# @!attribute [rw] nutriscore_score
#   @return [Integer, nil]
#
# @!attribute [rw] packaging
#   @return [String, nil]
#
# @!attribute [rw] product_name
#   @return [String, nil]
#
# @!attribute [rw] quantity
#   @return [String, nil]
#
# @!attribute [rw] stores
#   @return [String, nil]
#
# @!attribute [rw] traces
#   @return [String, nil]
Product = Struct.new(
  :additives_tags,
  :allergens,
  :brands,
  :categories,
  :countries,
  :created_t,
  :ecoscore_grade,
  :ecoscore_score,
  :generic_name,
  :id,
  :image_front_url,
  :image_ingredients_url,
  :image_nutrition_url,
  :image_url,
  :ingredients_analysis_tags,
  :ingredients_text,
  :labels,
  :last_modified_t,
  :manufacturing_places,
  :nova_group,
  :nutriments,
  :nutriscore_grade,
  :nutriscore_score,
  :packaging,
  :product_name,
  :quantity,
  :stores,
  :traces,
  keyword_init: true
)

# Request payload for Product#load.
#
# @!attribute [rw] id
#   @return [String]
ProductLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Search entity data model.
#
# @!attribute [rw] additives_tags
#   @return [Array, nil]
#
# @!attribute [rw] allergens
#   @return [String, nil]
#
# @!attribute [rw] brands
#   @return [String, nil]
#
# @!attribute [rw] categories
#   @return [String, nil]
#
# @!attribute [rw] countries
#   @return [String, nil]
#
# @!attribute [rw] created_t
#   @return [Integer, nil]
#
# @!attribute [rw] ecoscore_grade
#   @return [String, nil]
#
# @!attribute [rw] ecoscore_score
#   @return [Integer, nil]
#
# @!attribute [rw] generic_name
#   @return [String, nil]
#
# @!attribute [rw] image_front_url
#   @return [String, nil]
#
# @!attribute [rw] image_ingredients_url
#   @return [String, nil]
#
# @!attribute [rw] image_nutrition_url
#   @return [String, nil]
#
# @!attribute [rw] image_url
#   @return [String, nil]
#
# @!attribute [rw] ingredients_analysis_tags
#   @return [Array, nil]
#
# @!attribute [rw] ingredients_text
#   @return [String, nil]
#
# @!attribute [rw] labels
#   @return [String, nil]
#
# @!attribute [rw] last_modified_t
#   @return [Integer, nil]
#
# @!attribute [rw] manufacturing_places
#   @return [String, nil]
#
# @!attribute [rw] nova_group
#   @return [Integer, nil]
#
# @!attribute [rw] nutriments
#   @return [Hash, nil]
#
# @!attribute [rw] nutriscore_grade
#   @return [String, nil]
#
# @!attribute [rw] nutriscore_score
#   @return [Integer, nil]
#
# @!attribute [rw] packaging
#   @return [String, nil]
#
# @!attribute [rw] product_name
#   @return [String, nil]
#
# @!attribute [rw] quantity
#   @return [String, nil]
#
# @!attribute [rw] stores
#   @return [String, nil]
#
# @!attribute [rw] traces
#   @return [String, nil]
Search = Struct.new(
  :additives_tags,
  :allergens,
  :brands,
  :categories,
  :countries,
  :created_t,
  :ecoscore_grade,
  :ecoscore_score,
  :generic_name,
  :image_front_url,
  :image_ingredients_url,
  :image_nutrition_url,
  :image_url,
  :ingredients_analysis_tags,
  :ingredients_text,
  :labels,
  :last_modified_t,
  :manufacturing_places,
  :nova_group,
  :nutriments,
  :nutriscore_grade,
  :nutriscore_score,
  :packaging,
  :product_name,
  :quantity,
  :stores,
  :traces,
  keyword_init: true
)

# Request payload for Search#list.
#
# @!attribute [rw] additives_tags
#   @return [Array, nil]
#
# @!attribute [rw] allergens
#   @return [String, nil]
#
# @!attribute [rw] brands
#   @return [String, nil]
#
# @!attribute [rw] categories
#   @return [String, nil]
#
# @!attribute [rw] countries
#   @return [String, nil]
#
# @!attribute [rw] created_t
#   @return [Integer, nil]
#
# @!attribute [rw] ecoscore_grade
#   @return [String, nil]
#
# @!attribute [rw] ecoscore_score
#   @return [Integer, nil]
#
# @!attribute [rw] generic_name
#   @return [String, nil]
#
# @!attribute [rw] image_front_url
#   @return [String, nil]
#
# @!attribute [rw] image_ingredients_url
#   @return [String, nil]
#
# @!attribute [rw] image_nutrition_url
#   @return [String, nil]
#
# @!attribute [rw] image_url
#   @return [String, nil]
#
# @!attribute [rw] ingredients_analysis_tags
#   @return [Array, nil]
#
# @!attribute [rw] ingredients_text
#   @return [String, nil]
#
# @!attribute [rw] labels
#   @return [String, nil]
#
# @!attribute [rw] last_modified_t
#   @return [Integer, nil]
#
# @!attribute [rw] manufacturing_places
#   @return [String, nil]
#
# @!attribute [rw] nova_group
#   @return [Integer, nil]
#
# @!attribute [rw] nutriments
#   @return [Hash, nil]
#
# @!attribute [rw] nutriscore_grade
#   @return [String, nil]
#
# @!attribute [rw] nutriscore_score
#   @return [Integer, nil]
#
# @!attribute [rw] packaging
#   @return [String, nil]
#
# @!attribute [rw] product_name
#   @return [String, nil]
#
# @!attribute [rw] quantity
#   @return [String, nil]
#
# @!attribute [rw] stores
#   @return [String, nil]
#
# @!attribute [rw] traces
#   @return [String, nil]
SearchListMatch = Struct.new(
  :additives_tags,
  :allergens,
  :brands,
  :categories,
  :countries,
  :created_t,
  :ecoscore_grade,
  :ecoscore_score,
  :generic_name,
  :image_front_url,
  :image_ingredients_url,
  :image_nutrition_url,
  :image_url,
  :ingredients_analysis_tags,
  :ingredients_text,
  :labels,
  :last_modified_t,
  :manufacturing_places,
  :nova_group,
  :nutriments,
  :nutriscore_grade,
  :nutriscore_score,
  :packaging,
  :product_name,
  :quantity,
  :stores,
  :traces,
  keyword_init: true
)

