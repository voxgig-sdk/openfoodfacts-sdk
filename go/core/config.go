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
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
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
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "allergens",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "brands",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "categories",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "countries",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "created_t",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ecoscore_grade",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ecoscore_score",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "generic_name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_front_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_ingredients_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_nutrition_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ingredients_analysis_tags",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "ingredients_text",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "labels",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "last_modified_t",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "manufacturing_places",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nova_group",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "nutriments",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "nutriscore_grade",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nutriscore_score",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "packaging",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "product_name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "quantity",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "stores",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "traces",
						"type": "`$STRING`",
					},
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
								"parts": []any{
									"product",
									"{barcode}.json",
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
								"parts": []any{
									"product",
									"{barcode}.xml",
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
								"parts": []any{
									"product",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"barcode": "id",
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
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"product",
						},
					},
				},
			},
			"search": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "additives_tags",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "allergens",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "brands",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "categories",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "countries",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "created_t",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ecoscore_grade",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ecoscore_score",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "generic_name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_front_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_ingredients_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_nutrition_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ingredients_analysis_tags",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "ingredients_text",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "labels",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "last_modified_t",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "manufacturing_places",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nova_group",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "nutriments",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "nutriscore_grade",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nutriscore_score",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "packaging",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "product_name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "quantity",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "stores",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "traces",
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
								"parts": []any{
									"search",
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
