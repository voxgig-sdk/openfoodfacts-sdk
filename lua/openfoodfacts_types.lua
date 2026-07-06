-- Typed models for the Openfoodfacts SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
-- params (op.<name>.points[].args.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class Product
---@field code? string
---@field product? table
---@field status? number
---@field status_verbose? string

---@class ProductLoadMatch
---@field barcode string
---@field id string

---@class Search
---@field additives_tag? table
---@field allergen? string
---@field brand? string
---@field category? string
---@field country? string
---@field created_t? number
---@field ecoscore_grade? string
---@field ecoscore_score? number
---@field generic_name? string
---@field image_front_url? string
---@field image_ingredients_url? string
---@field image_nutrition_url? string
---@field image_url? string
---@field ingredients_analysis_tag? table
---@field ingredients_text? string
---@field label? string
---@field last_modified_t? number
---@field manufacturing_place? string
---@field nova_group? number
---@field nutriment? table
---@field nutriscore_grade? string
---@field nutriscore_score? number
---@field packaging? string
---@field product_name? string
---@field quantity? string
---@field store? string
---@field trace? string

---@class SearchListMatch
---@field additives_tag? table
---@field allergen? string
---@field brand? string
---@field category? string
---@field country? string
---@field created_t? number
---@field ecoscore_grade? string
---@field ecoscore_score? number
---@field generic_name? string
---@field image_front_url? string
---@field image_ingredients_url? string
---@field image_nutrition_url? string
---@field image_url? string
---@field ingredients_analysis_tag? table
---@field ingredients_text? string
---@field label? string
---@field last_modified_t? number
---@field manufacturing_place? string
---@field nova_group? number
---@field nutriment? table
---@field nutriscore_grade? string
---@field nutriscore_score? number
---@field packaging? string
---@field product_name? string
---@field quantity? string
---@field store? string
---@field trace? string

local M = {}

return M
