// Typed models for the Openfoodfacts SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/openfoodfacts-sdk/go/core"
)

// Product is the typed data model for the product entity.
type Product struct {
	AdditivesTags *[]any `json:"additives_tags,omitempty"`
	Allergens *string `json:"allergens,omitempty"`
	Brands *string `json:"brands,omitempty"`
	Categories *string `json:"categories,omitempty"`
	Countries *string `json:"countries,omitempty"`
	CreatedT *int `json:"created_t,omitempty"`
	EcoscoreGrade *string `json:"ecoscore_grade,omitempty"`
	EcoscoreScore *int `json:"ecoscore_score,omitempty"`
	GenericName *string `json:"generic_name,omitempty"`
	ImageFrontUrl *string `json:"image_front_url,omitempty"`
	ImageIngredientsUrl *string `json:"image_ingredients_url,omitempty"`
	ImageNutritionUrl *string `json:"image_nutrition_url,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
	IngredientsAnalysisTags *[]any `json:"ingredients_analysis_tags,omitempty"`
	IngredientsText *string `json:"ingredients_text,omitempty"`
	Labels *string `json:"labels,omitempty"`
	LastModifiedT *int `json:"last_modified_t,omitempty"`
	ManufacturingPlaces *string `json:"manufacturing_places,omitempty"`
	NovaGroup *int `json:"nova_group,omitempty"`
	Nutriments *map[string]any `json:"nutriments,omitempty"`
	NutriscoreGrade *string `json:"nutriscore_grade,omitempty"`
	NutriscoreScore *int `json:"nutriscore_score,omitempty"`
	Packaging *string `json:"packaging,omitempty"`
	ProductName *string `json:"product_name,omitempty"`
	Quantity *string `json:"quantity,omitempty"`
	Stores *string `json:"stores,omitempty"`
	Traces *string `json:"traces,omitempty"`
}

// ProductLoadMatch is the typed request payload for Product.LoadTyped.
type ProductLoadMatch struct {
	Id string `json:"id"`
}

// Search is the typed data model for the search entity.
type Search struct {
	AdditivesTags *[]any `json:"additives_tags,omitempty"`
	Allergens *string `json:"allergens,omitempty"`
	Brands *string `json:"brands,omitempty"`
	Categories *string `json:"categories,omitempty"`
	Countries *string `json:"countries,omitempty"`
	CreatedT *int `json:"created_t,omitempty"`
	EcoscoreGrade *string `json:"ecoscore_grade,omitempty"`
	EcoscoreScore *int `json:"ecoscore_score,omitempty"`
	GenericName *string `json:"generic_name,omitempty"`
	ImageFrontUrl *string `json:"image_front_url,omitempty"`
	ImageIngredientsUrl *string `json:"image_ingredients_url,omitempty"`
	ImageNutritionUrl *string `json:"image_nutrition_url,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
	IngredientsAnalysisTags *[]any `json:"ingredients_analysis_tags,omitempty"`
	IngredientsText *string `json:"ingredients_text,omitempty"`
	Labels *string `json:"labels,omitempty"`
	LastModifiedT *int `json:"last_modified_t,omitempty"`
	ManufacturingPlaces *string `json:"manufacturing_places,omitempty"`
	NovaGroup *int `json:"nova_group,omitempty"`
	Nutriments *map[string]any `json:"nutriments,omitempty"`
	NutriscoreGrade *string `json:"nutriscore_grade,omitempty"`
	NutriscoreScore *int `json:"nutriscore_score,omitempty"`
	Packaging *string `json:"packaging,omitempty"`
	ProductName *string `json:"product_name,omitempty"`
	Quantity *string `json:"quantity,omitempty"`
	Stores *string `json:"stores,omitempty"`
	Traces *string `json:"traces,omitempty"`
}

// SearchListMatch is the typed request payload for Search.ListTyped.
type SearchListMatch struct {
	AdditivesTags *[]any `json:"additives_tags,omitempty"`
	Allergens *string `json:"allergens,omitempty"`
	Brands *string `json:"brands,omitempty"`
	Categories *string `json:"categories,omitempty"`
	Countries *string `json:"countries,omitempty"`
	CreatedT *int `json:"created_t,omitempty"`
	EcoscoreGrade *string `json:"ecoscore_grade,omitempty"`
	EcoscoreScore *int `json:"ecoscore_score,omitempty"`
	GenericName *string `json:"generic_name,omitempty"`
	ImageFrontUrl *string `json:"image_front_url,omitempty"`
	ImageIngredientsUrl *string `json:"image_ingredients_url,omitempty"`
	ImageNutritionUrl *string `json:"image_nutrition_url,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
	IngredientsAnalysisTags *[]any `json:"ingredients_analysis_tags,omitempty"`
	IngredientsText *string `json:"ingredients_text,omitempty"`
	Labels *string `json:"labels,omitempty"`
	LastModifiedT *int `json:"last_modified_t,omitempty"`
	ManufacturingPlaces *string `json:"manufacturing_places,omitempty"`
	NovaGroup *int `json:"nova_group,omitempty"`
	Nutriments *map[string]any `json:"nutriments,omitempty"`
	NutriscoreGrade *string `json:"nutriscore_grade,omitempty"`
	NutriscoreScore *int `json:"nutriscore_score,omitempty"`
	Packaging *string `json:"packaging,omitempty"`
	ProductName *string `json:"product_name,omitempty"`
	Quantity *string `json:"quantity,omitempty"`
	Stores *string `json:"stores,omitempty"`
	Traces *string `json:"traces,omitempty"`
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

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
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

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
