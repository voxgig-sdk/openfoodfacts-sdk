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
# @!attribute [rw] code
#   @return [String, nil]
#
# @!attribute [rw] product
#   @return [Hash, nil]
#
# @!attribute [rw] status
#   @return [Integer, nil]
#
# @!attribute [rw] status_verbose
#   @return [String, nil]
Product = Struct.new(
  :code,
  :product,
  :status,
  :status_verbose,
  keyword_init: true
)

# Request payload for Product#load.
#
# @!attribute [rw] barcode
#   @return [String]
#
# @!attribute [rw] id
#   @return [String]
ProductLoadMatch = Struct.new(
  :barcode,
  :id,
  keyword_init: true
)

# Search entity data model.
#
# @!attribute [rw] additives_tag
#   @return [Array, nil]
#
# @!attribute [rw] allergen
#   @return [String, nil]
#
# @!attribute [rw] brand
#   @return [String, nil]
#
# @!attribute [rw] category
#   @return [String, nil]
#
# @!attribute [rw] country
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
# @!attribute [rw] ingredients_analysis_tag
#   @return [Array, nil]
#
# @!attribute [rw] ingredients_text
#   @return [String, nil]
#
# @!attribute [rw] label
#   @return [String, nil]
#
# @!attribute [rw] last_modified_t
#   @return [Integer, nil]
#
# @!attribute [rw] manufacturing_place
#   @return [String, nil]
#
# @!attribute [rw] nova_group
#   @return [Integer, nil]
#
# @!attribute [rw] nutriment
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
# @!attribute [rw] store
#   @return [String, nil]
#
# @!attribute [rw] trace
#   @return [String, nil]
Search = Struct.new(
  :additives_tag,
  :allergen,
  :brand,
  :category,
  :country,
  :created_t,
  :ecoscore_grade,
  :ecoscore_score,
  :generic_name,
  :image_front_url,
  :image_ingredients_url,
  :image_nutrition_url,
  :image_url,
  :ingredients_analysis_tag,
  :ingredients_text,
  :label,
  :last_modified_t,
  :manufacturing_place,
  :nova_group,
  :nutriment,
  :nutriscore_grade,
  :nutriscore_score,
  :packaging,
  :product_name,
  :quantity,
  :store,
  :trace,
  keyword_init: true
)

# Match filter for Search#list (any subset of Search fields).
#
# @!attribute [rw] additives_tag
#   @return [Array, nil]
#
# @!attribute [rw] allergen
#   @return [String, nil]
#
# @!attribute [rw] brand
#   @return [String, nil]
#
# @!attribute [rw] category
#   @return [String, nil]
#
# @!attribute [rw] country
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
# @!attribute [rw] ingredients_analysis_tag
#   @return [Array, nil]
#
# @!attribute [rw] ingredients_text
#   @return [String, nil]
#
# @!attribute [rw] label
#   @return [String, nil]
#
# @!attribute [rw] last_modified_t
#   @return [Integer, nil]
#
# @!attribute [rw] manufacturing_place
#   @return [String, nil]
#
# @!attribute [rw] nova_group
#   @return [Integer, nil]
#
# @!attribute [rw] nutriment
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
# @!attribute [rw] store
#   @return [String, nil]
#
# @!attribute [rw] trace
#   @return [String, nil]
SearchListMatch = Struct.new(
  :additives_tag,
  :allergen,
  :brand,
  :category,
  :country,
  :created_t,
  :ecoscore_grade,
  :ecoscore_score,
  :generic_name,
  :image_front_url,
  :image_ingredients_url,
  :image_nutrition_url,
  :image_url,
  :ingredients_analysis_tag,
  :ingredients_text,
  :label,
  :last_modified_t,
  :manufacturing_place,
  :nova_group,
  :nutriment,
  :nutriscore_grade,
  :nutriscore_score,
  :packaging,
  :product_name,
  :quantity,
  :store,
  :trace,
  keyword_init: true
)

