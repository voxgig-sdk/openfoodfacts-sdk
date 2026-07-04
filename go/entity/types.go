// Typed models for the Openfoodfacts SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// Product is the typed data model for the product entity.
type Product struct {
	Code *string `json:"code,omitempty"`
	Product *map[string]any `json:"product,omitempty"`
	Status *int `json:"status,omitempty"`
	StatusVerbose *string `json:"status_verbose,omitempty"`
}

// ProductLoadMatch is the typed request payload for Product.LoadTyped.
type ProductLoadMatch struct {
	Barcode string `json:"barcode"`
	Id string `json:"id"`
}

// Search is the typed data model for the search entity.
type Search struct {
	AdditivesTag *[]any `json:"additives_tag,omitempty"`
	Allergen *string `json:"allergen,omitempty"`
	Brand *string `json:"brand,omitempty"`
	Category *string `json:"category,omitempty"`
	Country *string `json:"country,omitempty"`
	CreatedT *int `json:"created_t,omitempty"`
	EcoscoreGrade *string `json:"ecoscore_grade,omitempty"`
	EcoscoreScore *int `json:"ecoscore_score,omitempty"`
	GenericName *string `json:"generic_name,omitempty"`
	ImageFrontUrl *string `json:"image_front_url,omitempty"`
	ImageIngredientsUrl *string `json:"image_ingredients_url,omitempty"`
	ImageNutritionUrl *string `json:"image_nutrition_url,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
	IngredientsAnalysisTag *[]any `json:"ingredients_analysis_tag,omitempty"`
	IngredientsText *string `json:"ingredients_text,omitempty"`
	Label *string `json:"label,omitempty"`
	LastModifiedT *int `json:"last_modified_t,omitempty"`
	ManufacturingPlace *string `json:"manufacturing_place,omitempty"`
	NovaGroup *int `json:"nova_group,omitempty"`
	Nutriment *map[string]any `json:"nutriment,omitempty"`
	NutriscoreGrade *string `json:"nutriscore_grade,omitempty"`
	NutriscoreScore *int `json:"nutriscore_score,omitempty"`
	Packaging *string `json:"packaging,omitempty"`
	ProductName *string `json:"product_name,omitempty"`
	Quantity *string `json:"quantity,omitempty"`
	Store *string `json:"store,omitempty"`
	Trace *string `json:"trace,omitempty"`
}

// SearchListMatch mirrors the search fields as an all-optional match
// filter (Go analog of Partial<Search>).
type SearchListMatch struct {
	AdditivesTag *[]any `json:"additives_tag,omitempty"`
	Allergen *string `json:"allergen,omitempty"`
	Brand *string `json:"brand,omitempty"`
	Category *string `json:"category,omitempty"`
	Country *string `json:"country,omitempty"`
	CreatedT *int `json:"created_t,omitempty"`
	EcoscoreGrade *string `json:"ecoscore_grade,omitempty"`
	EcoscoreScore *int `json:"ecoscore_score,omitempty"`
	GenericName *string `json:"generic_name,omitempty"`
	ImageFrontUrl *string `json:"image_front_url,omitempty"`
	ImageIngredientsUrl *string `json:"image_ingredients_url,omitempty"`
	ImageNutritionUrl *string `json:"image_nutrition_url,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
	IngredientsAnalysisTag *[]any `json:"ingredients_analysis_tag,omitempty"`
	IngredientsText *string `json:"ingredients_text,omitempty"`
	Label *string `json:"label,omitempty"`
	LastModifiedT *int `json:"last_modified_t,omitempty"`
	ManufacturingPlace *string `json:"manufacturing_place,omitempty"`
	NovaGroup *int `json:"nova_group,omitempty"`
	Nutriment *map[string]any `json:"nutriment,omitempty"`
	NutriscoreGrade *string `json:"nutriscore_grade,omitempty"`
	NutriscoreScore *int `json:"nutriscore_score,omitempty"`
	Packaging *string `json:"packaging,omitempty"`
	ProductName *string `json:"product_name,omitempty"`
	Quantity *string `json:"quantity,omitempty"`
	Store *string `json:"store,omitempty"`
	Trace *string `json:"trace,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
