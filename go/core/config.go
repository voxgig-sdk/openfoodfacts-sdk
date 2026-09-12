package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "Openfoodfacts",
			"slug": "openfoodfacts",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"transport": "base",
			},
		},
		"options": map[string]any{
			"base": "https://world.openfoodfacts.org/api/v2",
			"auth": map[string]any{
				"prefix": "",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"product": map[string]any{},
				"search": map[string]any{},
			},
		},
		"entity": map[string]any{
			"product": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "additives_tags",
						"short": "List of additives",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "allergens",
						"short": "Allergens present in the product",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "brands",
						"short": "Brands of the product",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "categories",
						"short": "Categories the product belongs to",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "countries",
						"short": "Countries where the product is sold",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "created_t",
						"short": "Creation timestamp",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ecoscore_grade",
						"short": "Eco-Score grade for environmental impact (a, b, c, d, e)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ecoscore_score",
						"short": "Eco-Score numerical score",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "generic_name",
						"short": "Generic name of the product",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_front_url",
						"short": "URL of the front image",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_ingredients_url",
						"short": "URL of the ingredients image",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_nutrition_url",
						"short": "URL of the nutrition facts image",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_url",
						"short": "URL of the product's front image",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ingredients_analysis_tags",
						"short": "Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.)",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "ingredients_text",
						"short": "List of ingredients as text",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "labels",
						"short": "Labels associated with the product (e.g., Organic, Fair Trade)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "last_modified_t",
						"short": "Last modification timestamp",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "manufacturing_places",
						"short": "Manufacturing or processing places",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nova_group",
						"short": "NOVA group for food processing level (1-4)",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "nutriments",
						"short": "Nutritional information",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "nutriscore_grade",
						"short": "Nutri-Score grade (a, b, c, d, e)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nutriscore_score",
						"short": "Nutri-Score numerical score",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "packaging",
						"short": "Packaging type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "product_name",
						"short": "Name of the product",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "quantity",
						"short": "Quantity or volume of the product",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "stores",
						"short": "Stores where the product is available",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "traces",
						"short": "Traces of allergens",
						"type": "`$STRING`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "product",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "737628064502",
											"kind": "param",
											"name": "barcode",
											"orig": "barcode",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/product/{barcode}.json",
								"segments": []any{
									map[string]any{
										"lit": "product",
									},
									map[string]any{
										"lit": "{barcode}.json",
									},
								},
								"select": map[string]any{
									"$action": "barcode",
									"exist": []any{
										"barcode",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.product`",
								},
								"parts": []any{
									"product",
									"{barcode}.json",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "737628064502",
											"kind": "param",
											"name": "barcode",
											"orig": "barcode",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/product/{barcode}.xml",
								"segments": []any{
									map[string]any{
										"lit": "product",
									},
									map[string]any{
										"lit": "{barcode}.xml",
									},
								},
								"select": map[string]any{
									"$action": "barcode",
									"exist": []any{
										"barcode",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"product",
									"{barcode}.xml",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "737628064502",
											"kind": "param",
											"name": "id",
											"orig": "barcode",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/product/{barcode}",
								"rename": map[string]any{
									"param": map[string]any{
										"barcode": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "product",
									},
									map[string]any{
										"var": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.product`",
								},
								"parts": []any{
									"product",
									"{id}",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"search": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "additives_tags",
						"short": "List of additives",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "allergens",
						"short": "Allergens present in the product",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "brands",
						"short": "Brands of the product",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "categories",
						"short": "Categories the product belongs to",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "countries",
						"short": "Countries where the product is sold",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "created_t",
						"short": "Creation timestamp",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ecoscore_grade",
						"short": "Eco-Score grade for environmental impact (a, b, c, d, e)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ecoscore_score",
						"short": "Eco-Score numerical score",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "generic_name",
						"short": "Generic name of the product",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_front_url",
						"short": "URL of the front image",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_ingredients_url",
						"short": "URL of the ingredients image",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_nutrition_url",
						"short": "URL of the nutrition facts image",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_url",
						"short": "URL of the product's front image",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ingredients_analysis_tags",
						"short": "Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.)",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "ingredients_text",
						"short": "List of ingredients as text",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "labels",
						"short": "Labels associated with the product (e.g., Organic, Fair Trade)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "last_modified_t",
						"short": "Last modification timestamp",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "manufacturing_places",
						"short": "Manufacturing or processing places",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nova_group",
						"short": "NOVA group for food processing level (1-4)",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "nutriments",
						"short": "Nutritional information",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "nutriscore_grade",
						"short": "Nutri-Score grade (a, b, c, d, e)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nutriscore_score",
						"short": "Nutri-Score numerical score",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "packaging",
						"short": "Packaging type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "product_name",
						"short": "Name of the product",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "quantity",
						"short": "Quantity or volume of the product",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "stores",
						"short": "Stores where the product is available",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "traces",
						"short": "Traces of allergens",
						"type": "`$STRING`",
					},
				},
				"name": "search",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "brand",
											"orig": "brand",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "category",
											"orig": "category",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": true,
											"kind": "query",
											"name": "json",
											"orig": "json",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "label",
											"orig": "label",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 20,
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search_term",
											"orig": "search_term",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/search",
								"segments": []any{
									map[string]any{
										"lit": "search",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"brand",
										"category",
										"json",
										"label",
										"page",
										"page_size",
										"search_term",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.products`",
								},
								"parts": []any{
									"search",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
