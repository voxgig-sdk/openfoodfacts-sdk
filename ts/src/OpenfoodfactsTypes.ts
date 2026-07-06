// Typed models for the Openfoodfacts SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface Product {
  code?: string
  product?: Record<string, any>
  status?: number
  status_verbose?: string
}

export interface ProductLoadMatch {
  barcode: string
  id: string
}

export interface Search {
  additives_tag?: any[]
  allergen?: string
  brand?: string
  category?: string
  country?: string
  created_t?: number
  ecoscore_grade?: string
  ecoscore_score?: number
  generic_name?: string
  image_front_url?: string
  image_ingredients_url?: string
  image_nutrition_url?: string
  image_url?: string
  ingredients_analysis_tag?: any[]
  ingredients_text?: string
  label?: string
  last_modified_t?: number
  manufacturing_place?: string
  nova_group?: number
  nutriment?: Record<string, any>
  nutriscore_grade?: string
  nutriscore_score?: number
  packaging?: string
  product_name?: string
  quantity?: string
  store?: string
  trace?: string
}

export interface SearchListMatch {
  additives_tag?: any[]
  allergen?: string
  brand?: string
  category?: string
  country?: string
  created_t?: number
  ecoscore_grade?: string
  ecoscore_score?: number
  generic_name?: string
  image_front_url?: string
  image_ingredients_url?: string
  image_nutrition_url?: string
  image_url?: string
  ingredients_analysis_tag?: any[]
  ingredients_text?: string
  label?: string
  last_modified_t?: number
  manufacturing_place?: string
  nova_group?: number
  nutriment?: Record<string, any>
  nutriscore_grade?: string
  nutriscore_score?: number
  packaging?: string
  product_name?: string
  quantity?: string
  store?: string
  trace?: string
}

