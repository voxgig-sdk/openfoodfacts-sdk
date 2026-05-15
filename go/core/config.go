package core

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
				"prefix": "Bearer",
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
						"name": "code",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "product",
						"req": false,
						"type": "`$OBJECT`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "status",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "status_verbose",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 3,
					},
				},
				"name": "product",
				"op": map[string]any{
					"load": map[string]any{
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
											"active": true,
										},
									},
								},
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
									"res": "`body`",
								},
								"active": true,
								"index$": 0,
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
											"active": true,
										},
									},
								},
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
								"active": true,
								"index$": 1,
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
											"active": true,
										},
									},
								},
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
									"res": "`body`",
								},
								"active": true,
								"index$": 2,
							},
						},
						"input": "data",
						"key$": "load",
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
						"name": "additives_tag",
						"req": false,
						"type": "`$ARRAY`",
						"active": true,
						"index$": 0,
					},
					map[string]any{
						"name": "allergen",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 1,
					},
					map[string]any{
						"name": "brand",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 2,
					},
					map[string]any{
						"name": "category",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 3,
					},
					map[string]any{
						"name": "country",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 4,
					},
					map[string]any{
						"name": "created_t",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 5,
					},
					map[string]any{
						"name": "ecoscore_grade",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 6,
					},
					map[string]any{
						"name": "ecoscore_score",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 7,
					},
					map[string]any{
						"name": "generic_name",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 8,
					},
					map[string]any{
						"name": "image_front_url",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 9,
					},
					map[string]any{
						"name": "image_ingredients_url",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 10,
					},
					map[string]any{
						"name": "image_nutrition_url",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 11,
					},
					map[string]any{
						"name": "image_url",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 12,
					},
					map[string]any{
						"name": "ingredients_analysis_tag",
						"req": false,
						"type": "`$ARRAY`",
						"active": true,
						"index$": 13,
					},
					map[string]any{
						"name": "ingredients_text",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 14,
					},
					map[string]any{
						"name": "label",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 15,
					},
					map[string]any{
						"name": "last_modified_t",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 16,
					},
					map[string]any{
						"name": "manufacturing_place",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 17,
					},
					map[string]any{
						"name": "nova_group",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 18,
					},
					map[string]any{
						"name": "nutriment",
						"req": false,
						"type": "`$OBJECT`",
						"active": true,
						"index$": 19,
					},
					map[string]any{
						"name": "nutriscore_grade",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 20,
					},
					map[string]any{
						"name": "nutriscore_score",
						"req": false,
						"type": "`$INTEGER`",
						"active": true,
						"index$": 21,
					},
					map[string]any{
						"name": "packaging",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 22,
					},
					map[string]any{
						"name": "product_name",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 23,
					},
					map[string]any{
						"name": "quantity",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 24,
					},
					map[string]any{
						"name": "store",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 25,
					},
					map[string]any{
						"name": "trace",
						"req": false,
						"type": "`$STRING`",
						"active": true,
						"index$": 26,
					},
				},
				"name": "search",
				"op": map[string]any{
					"list": map[string]any{
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "brand",
											"orig": "brand",
											"reqd": false,
											"type": "`$STRING`",
											"active": true,
										},
										map[string]any{
											"kind": "query",
											"name": "category",
											"orig": "category",
											"reqd": false,
											"type": "`$STRING`",
											"active": true,
										},
										map[string]any{
											"example": true,
											"kind": "query",
											"name": "json",
											"orig": "json",
											"reqd": false,
											"type": "`$BOOLEAN`",
											"active": true,
										},
										map[string]any{
											"kind": "query",
											"name": "label",
											"orig": "label",
											"reqd": false,
											"type": "`$STRING`",
											"active": true,
										},
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
											"active": true,
										},
										map[string]any{
											"example": 20,
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"reqd": false,
											"type": "`$INTEGER`",
											"active": true,
										},
										map[string]any{
											"kind": "query",
											"name": "search_term",
											"orig": "search_term",
											"reqd": false,
											"type": "`$STRING`",
											"active": true,
										},
									},
								},
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
									"res": "`body`",
								},
								"active": true,
								"index$": 0,
							},
						},
						"input": "data",
						"key$": "list",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
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
