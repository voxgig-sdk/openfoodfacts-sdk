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
			"ratelimit": map[string]any{
				"options": map[string]any{
					"active": false,
					"burst": 5,
					"rate": 5,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"retry": map[string]any{
				"options": map[string]any{
					"active": false,
					"factor": 2,
					"maxDelay": 2000,
					"minDelay": 50,
					"retries": 2,
					"statuses": []any{
						408,
						425,
						429,
						500,
						502,
						503,
						504,
					},
				},
				"optspec": map[string]any{
					"jitter": "`$BOOLEAN`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
			"timeout": map[string]any{
				"options": map[string]any{
					"active": false,
					"ms": 30000,
				},
				"optspec": map[string]any{
					"clearTimer": "`$FUNCTION`",
					"setTimer": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
		},
		"options": map[string]any{
			"base": "https://world.openfoodfacts.org/api/v2",
			"auth": map[string]any{
				"prefix": "",
				"name": "User-Agent",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"product": map[string]any{},
			},
		},
		"entity": map[string]any{
			"product": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "additives_tags",
						"title": "Additives Tags",
						"type": "`$ARRAY`",
						"short": "List of additives",
					},
					map[string]any{
						"name": "allergens",
						"title": "Allergens",
						"type": "`$STRING`",
						"short": "Allergens present in the product",
					},
					map[string]any{
						"name": "brands",
						"title": "Brands",
						"type": "`$STRING`",
						"short": "Brands of the product",
					},
					map[string]any{
						"name": "categories",
						"title": "Categories",
						"type": "`$STRING`",
						"short": "Categories the product belongs to",
					},
					map[string]any{
						"name": "countries",
						"title": "Countries",
						"type": "`$STRING`",
						"short": "Countries where the product is sold",
					},
					map[string]any{
						"name": "created_t",
						"title": "Created T",
						"type": "`$INTEGER`",
						"short": "Creation timestamp",
					},
					map[string]any{
						"name": "ecoscore_grade",
						"title": "Ecoscore Grade",
						"type": "`$STRING`",
						"short": "Eco-Score grade for environmental impact (a, b, c, d, e)",
					},
					map[string]any{
						"name": "ecoscore_score",
						"title": "Ecoscore Score",
						"type": "`$INTEGER`",
						"short": "Eco-Score numerical score",
					},
					map[string]any{
						"name": "generic_name",
						"title": "Generic Name",
						"type": "`$STRING`",
						"short": "Generic name of the product",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_front_url",
						"title": "Image Front Url",
						"type": "`$STRING`",
						"short": "URL of the front image",
					},
					map[string]any{
						"name": "image_ingredients_url",
						"title": "Image Ingredients Url",
						"type": "`$STRING`",
						"short": "URL of the ingredients image",
					},
					map[string]any{
						"name": "image_nutrition_url",
						"title": "Image Nutrition Url",
						"type": "`$STRING`",
						"short": "URL of the nutrition facts image",
					},
					map[string]any{
						"name": "image_url",
						"title": "Image Url",
						"type": "`$STRING`",
						"short": "URL of the product's front image",
					},
					map[string]any{
						"name": "ingredients_analysis_tags",
						"title": "Ingredients Analysis Tags",
						"type": "`$ARRAY`",
						"short": "Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.)",
					},
					map[string]any{
						"name": "ingredients_text",
						"title": "Ingredients Text",
						"type": "`$STRING`",
						"short": "List of ingredients as text",
					},
					map[string]any{
						"name": "labels",
						"title": "Labels",
						"type": "`$STRING`",
						"short": "Labels associated with the product (e.g., Organic, Fair Trade)",
					},
					map[string]any{
						"name": "last_modified_t",
						"title": "Last Modified T",
						"type": "`$INTEGER`",
						"short": "Last modification timestamp",
					},
					map[string]any{
						"name": "manufacturing_places",
						"title": "Manufacturing Places",
						"type": "`$STRING`",
						"short": "Manufacturing or processing places",
					},
					map[string]any{
						"name": "nova_group",
						"title": "Nova Group",
						"type": "`$INTEGER`",
						"short": "NOVA group for food processing level (1-4)",
					},
					map[string]any{
						"name": "nutriments",
						"title": "Nutriments",
						"type": "`$OBJECT`",
						"short": "Nutritional information",
					},
					map[string]any{
						"name": "nutriscore_grade",
						"title": "Nutriscore Grade",
						"type": "`$STRING`",
						"short": "Nutri-Score grade (a, b, c, d, e)",
					},
					map[string]any{
						"name": "nutriscore_score",
						"title": "Nutriscore Score",
						"type": "`$INTEGER`",
						"short": "Nutri-Score numerical score",
					},
					map[string]any{
						"name": "packaging",
						"title": "Packaging",
						"type": "`$STRING`",
						"short": "Packaging type",
					},
					map[string]any{
						"name": "product_name",
						"title": "Product Name",
						"type": "`$STRING`",
						"short": "Name of the product",
					},
					map[string]any{
						"name": "quantity",
						"title": "Quantity",
						"type": "`$STRING`",
						"short": "Quantity or volume of the product",
					},
					map[string]any{
						"name": "stores",
						"title": "Stores",
						"type": "`$STRING`",
						"short": "Stores where the product is available",
					},
					map[string]any{
						"name": "traces",
						"title": "Traces",
						"type": "`$STRING`",
						"short": "Traces of allergens",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "product",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/search",
								"segments": []any{
									map[string]any{
										"lit": "search",
									},
								},
								"parts": []any{
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.products`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "brand",
											"orig": "brand",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "category",
											"orig": "category",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "json",
											"orig": "json",
											"type": "`$BOOLEAN`",
											"kind": "query",
											"example": true,
										},
										map[string]any{
											"name": "label",
											"orig": "label",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 20,
										},
										map[string]any{
											"name": "search_term",
											"orig": "search_term",
											"type": "`$STRING`",
											"kind": "query",
										},
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
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
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
								"parts": []any{
									"product",
									"{barcode}.json",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.product`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "barcode",
											"orig": "barcode",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "737628064502",
										},
									},
								},
								"select": map[string]any{
									"$action": "barcode",
									"exist": []any{
										"barcode",
									},
								},
							},
							map[string]any{
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
								"parts": []any{
									"product",
									"{barcode}.xml",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "barcode",
											"orig": "barcode",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "737628064502",
										},
									},
								},
								"select": map[string]any{
									"$action": "barcode",
									"exist": []any{
										"barcode",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/product/{barcode}",
								"segments": []any{
									map[string]any{
										"lit": "product",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"product",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"barcode": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.product`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "barcode",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "737628064502",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
	case "ratelimit":
		if NewRatelimitFeatureFunc != nil {
			return NewRatelimitFeatureFunc()
		}
	case "retry":
		if NewRetryFeatureFunc != nil {
			return NewRetryFeatureFunc()
		}
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	case "timeout":
		if NewTimeoutFeatureFunc != nil {
			return NewTimeoutFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
