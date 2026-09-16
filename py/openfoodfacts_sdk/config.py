# Openfoodfacts SDK configuration


# The sekreto plugin DEFINITIONS the model selected per feature, imported
# above by name from the modules the catalogue's active `plugin.def`
# entries declare. Handed to each feature (secrets builds its Sekreto
# with them): a provider kind not listed here is unknown to that SDK.
FEATURE_PLUGINS = {
}


_shared_config = None


def shared_config():
    """Return the process-wide config, built once on first use.

    The SDK reads the config on every request and never writes to it, so one
    instance is shared by every client rather than rebuilt per client.

    The returned dict is shared: treat it as read-only. Callers that need to
    mutate should use make_config, which always returns a fresh copy.
    """
    global _shared_config
    if _shared_config is None:
        _shared_config = make_config()
    return _shared_config


def make_config():
    """Build a fresh, fully materialised config dict.

    Every call rebuilds the whole structure, so prefer shared_config unless
    you need a private copy you intend to mutate.
    """
    return {
        "main": {
            "name": "Openfoodfacts",
            "slug": "openfoodfacts",
            "version": "0.0.1",
            "target": "py",
        },
        "feature": {
            "ratelimit": {
        "options": {
          "active": False,
          "burst": 5,
          "rate": 5,
        },
        "optspec": {
          "now": "`$FUNCTION`",
          "sleep": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
            "retry": {
        "options": {
          "active": False,
          "factor": 2,
          "maxDelay": 2000,
          "minDelay": 50,
          "retries": 2,
          "statuses": [
            408,
            425,
            429,
            500,
            502,
            503,
            504,
          ],
        },
        "optspec": {
          "jitter": "`$BOOLEAN`",
          "sleep": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
            "test": {
        "options": {
          "active": False,
        },
        "optspec": {
          "entity": "`$MAP`",
          "net": "`$MAP`",
        },
        "strict": False,
        "transport": "base",
      },
            "timeout": {
        "options": {
          "active": False,
          "ms": 30000,
        },
        "optspec": {
          "clearTimer": "`$FUNCTION`",
          "setTimer": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
        },
        "options": {
            "base": "https://world.openfoodfacts.org/api/v2",
            "auth": {
                "prefix": "",
            },
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "product": {},
                "search": {},
            },
        },
        "entity": {
      "product": {
        "fields": [
          {
            "name": "additives_tags",
            "short": "List of additives",
            "type": "`$ARRAY`",
          },
          {
            "name": "allergens",
            "short": "Allergens present in the product",
            "type": "`$STRING`",
          },
          {
            "name": "brands",
            "short": "Brands of the product",
            "type": "`$STRING`",
          },
          {
            "name": "categories",
            "short": "Categories the product belongs to",
            "type": "`$STRING`",
          },
          {
            "name": "countries",
            "short": "Countries where the product is sold",
            "type": "`$STRING`",
          },
          {
            "name": "created_t",
            "short": "Creation timestamp",
            "type": "`$INTEGER`",
          },
          {
            "name": "ecoscore_grade",
            "short": "Eco-Score grade for environmental impact (a, b, c, d, e)",
            "type": "`$STRING`",
          },
          {
            "name": "ecoscore_score",
            "short": "Eco-Score numerical score",
            "type": "`$INTEGER`",
          },
          {
            "name": "generic_name",
            "short": "Generic name of the product",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "image_front_url",
            "short": "URL of the front image",
            "type": "`$STRING`",
          },
          {
            "name": "image_ingredients_url",
            "short": "URL of the ingredients image",
            "type": "`$STRING`",
          },
          {
            "name": "image_nutrition_url",
            "short": "URL of the nutrition facts image",
            "type": "`$STRING`",
          },
          {
            "name": "image_url",
            "short": "URL of the product's front image",
            "type": "`$STRING`",
          },
          {
            "name": "ingredients_analysis_tags",
            "short": "Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.)",
            "type": "`$ARRAY`",
          },
          {
            "name": "ingredients_text",
            "short": "List of ingredients as text",
            "type": "`$STRING`",
          },
          {
            "name": "labels",
            "short": "Labels associated with the product (e.g., Organic, Fair Trade)",
            "type": "`$STRING`",
          },
          {
            "name": "last_modified_t",
            "short": "Last modification timestamp",
            "type": "`$INTEGER`",
          },
          {
            "name": "manufacturing_places",
            "short": "Manufacturing or processing places",
            "type": "`$STRING`",
          },
          {
            "name": "nova_group",
            "short": "NOVA group for food processing level (1-4)",
            "type": "`$INTEGER`",
          },
          {
            "name": "nutriments",
            "short": "Nutritional information",
            "type": "`$OBJECT`",
          },
          {
            "name": "nutriscore_grade",
            "short": "Nutri-Score grade (a, b, c, d, e)",
            "type": "`$STRING`",
          },
          {
            "name": "nutriscore_score",
            "short": "Nutri-Score numerical score",
            "type": "`$INTEGER`",
          },
          {
            "name": "packaging",
            "short": "Packaging type",
            "type": "`$STRING`",
          },
          {
            "name": "product_name",
            "short": "Name of the product",
            "type": "`$STRING`",
          },
          {
            "name": "quantity",
            "short": "Quantity or volume of the product",
            "type": "`$STRING`",
          },
          {
            "name": "stores",
            "short": "Stores where the product is available",
            "type": "`$STRING`",
          },
          {
            "name": "traces",
            "short": "Traces of allergens",
            "type": "`$STRING`",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "product",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "737628064502",
                      "kind": "param",
                      "name": "barcode",
                      "orig": "barcode",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/product/{barcode}.json",
                "segments": [
                  {
                    "lit": "product",
                  },
                  {
                    "lit": "{barcode}.json",
                  },
                ],
                "select": {
                  "$action": "barcode",
                  "exist": [
                    "barcode",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.product`",
                },
                "parts": [
                  "product",
                  "{barcode}.json",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "example": "737628064502",
                      "kind": "param",
                      "name": "barcode",
                      "orig": "barcode",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/product/{barcode}.xml",
                "segments": [
                  {
                    "lit": "product",
                  },
                  {
                    "lit": "{barcode}.xml",
                  },
                ],
                "select": {
                  "$action": "barcode",
                  "exist": [
                    "barcode",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "product",
                  "{barcode}.xml",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "example": "737628064502",
                      "kind": "param",
                      "name": "id",
                      "orig": "barcode",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/product/{barcode}",
                "rename": {
                  "param": {
                    "barcode": "id",
                  },
                },
                "segments": [
                  {
                    "lit": "product",
                  },
                  {
                    "var": "id",
                  },
                ],
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.product`",
                },
                "parts": [
                  "product",
                  "{id}",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "search": {
        "fields": [
          {
            "name": "additives_tags",
            "short": "List of additives",
            "type": "`$ARRAY`",
          },
          {
            "name": "allergens",
            "short": "Allergens present in the product",
            "type": "`$STRING`",
          },
          {
            "name": "brands",
            "short": "Brands of the product",
            "type": "`$STRING`",
          },
          {
            "name": "categories",
            "short": "Categories the product belongs to",
            "type": "`$STRING`",
          },
          {
            "name": "countries",
            "short": "Countries where the product is sold",
            "type": "`$STRING`",
          },
          {
            "name": "created_t",
            "short": "Creation timestamp",
            "type": "`$INTEGER`",
          },
          {
            "name": "ecoscore_grade",
            "short": "Eco-Score grade for environmental impact (a, b, c, d, e)",
            "type": "`$STRING`",
          },
          {
            "name": "ecoscore_score",
            "short": "Eco-Score numerical score",
            "type": "`$INTEGER`",
          },
          {
            "name": "generic_name",
            "short": "Generic name of the product",
            "type": "`$STRING`",
          },
          {
            "name": "image_front_url",
            "short": "URL of the front image",
            "type": "`$STRING`",
          },
          {
            "name": "image_ingredients_url",
            "short": "URL of the ingredients image",
            "type": "`$STRING`",
          },
          {
            "name": "image_nutrition_url",
            "short": "URL of the nutrition facts image",
            "type": "`$STRING`",
          },
          {
            "name": "image_url",
            "short": "URL of the product's front image",
            "type": "`$STRING`",
          },
          {
            "name": "ingredients_analysis_tags",
            "short": "Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.)",
            "type": "`$ARRAY`",
          },
          {
            "name": "ingredients_text",
            "short": "List of ingredients as text",
            "type": "`$STRING`",
          },
          {
            "name": "labels",
            "short": "Labels associated with the product (e.g., Organic, Fair Trade)",
            "type": "`$STRING`",
          },
          {
            "name": "last_modified_t",
            "short": "Last modification timestamp",
            "type": "`$INTEGER`",
          },
          {
            "name": "manufacturing_places",
            "short": "Manufacturing or processing places",
            "type": "`$STRING`",
          },
          {
            "name": "nova_group",
            "short": "NOVA group for food processing level (1-4)",
            "type": "`$INTEGER`",
          },
          {
            "name": "nutriments",
            "short": "Nutritional information",
            "type": "`$OBJECT`",
          },
          {
            "name": "nutriscore_grade",
            "short": "Nutri-Score grade (a, b, c, d, e)",
            "type": "`$STRING`",
          },
          {
            "name": "nutriscore_score",
            "short": "Nutri-Score numerical score",
            "type": "`$INTEGER`",
          },
          {
            "name": "packaging",
            "short": "Packaging type",
            "type": "`$STRING`",
          },
          {
            "name": "product_name",
            "short": "Name of the product",
            "type": "`$STRING`",
          },
          {
            "name": "quantity",
            "short": "Quantity or volume of the product",
            "type": "`$STRING`",
          },
          {
            "name": "stores",
            "short": "Stores where the product is available",
            "type": "`$STRING`",
          },
          {
            "name": "traces",
            "short": "Traces of allergens",
            "type": "`$STRING`",
          },
        ],
        "name": "search",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "query": [
                    {
                      "kind": "query",
                      "name": "brand",
                      "orig": "brand",
                      "type": "`$STRING`",
                    },
                    {
                      "kind": "query",
                      "name": "category",
                      "orig": "category",
                      "type": "`$STRING`",
                    },
                    {
                      "example": True,
                      "kind": "query",
                      "name": "json",
                      "orig": "json",
                      "type": "`$BOOLEAN`",
                    },
                    {
                      "kind": "query",
                      "name": "label",
                      "orig": "label",
                      "type": "`$STRING`",
                    },
                    {
                      "example": 1,
                      "kind": "query",
                      "name": "page",
                      "orig": "page",
                      "type": "`$INTEGER`",
                    },
                    {
                      "example": 20,
                      "kind": "query",
                      "name": "page_size",
                      "orig": "page_size",
                      "type": "`$INTEGER`",
                    },
                    {
                      "kind": "query",
                      "name": "search_term",
                      "orig": "search_term",
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/search",
                "segments": [
                  {
                    "lit": "search",
                  },
                ],
                "select": {
                  "exist": [
                    "brand",
                    "category",
                    "json",
                    "label",
                    "page",
                    "page_size",
                    "search_term",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.products`",
                },
                "parts": [
                  "search",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }
