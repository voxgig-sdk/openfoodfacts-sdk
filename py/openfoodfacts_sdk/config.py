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
                "name": "User-Agent",
            },
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "product": {},
            },
        },
        "entity": {
      "product": {
        "fields": [
          {
            "name": "additives_tags",
            "title": "Additives Tags",
            "type": "`$ARRAY`",
            "short": "List of additives",
          },
          {
            "name": "allergens",
            "title": "Allergens",
            "type": "`$STRING`",
            "short": "Allergens present in the product",
          },
          {
            "name": "brands",
            "title": "Brands",
            "type": "`$STRING`",
            "short": "Brands of the product",
          },
          {
            "name": "categories",
            "title": "Categories",
            "type": "`$STRING`",
            "short": "Categories the product belongs to",
          },
          {
            "name": "countries",
            "title": "Countries",
            "type": "`$STRING`",
            "short": "Countries where the product is sold",
          },
          {
            "name": "created_t",
            "title": "Created T",
            "type": "`$INTEGER`",
            "short": "Creation timestamp",
          },
          {
            "name": "ecoscore_grade",
            "title": "Ecoscore Grade",
            "type": "`$STRING`",
            "short": "Eco-Score grade for environmental impact (a, b, c, d, e)",
          },
          {
            "name": "ecoscore_score",
            "title": "Ecoscore Score",
            "type": "`$INTEGER`",
            "short": "Eco-Score numerical score",
          },
          {
            "name": "generic_name",
            "title": "Generic Name",
            "type": "`$STRING`",
            "short": "Generic name of the product",
          },
          {
            "name": "id",
            "title": "Id",
            "type": "`$STRING`",
          },
          {
            "name": "image_front_url",
            "title": "Image Front Url",
            "type": "`$STRING`",
            "short": "URL of the front image",
          },
          {
            "name": "image_ingredients_url",
            "title": "Image Ingredients Url",
            "type": "`$STRING`",
            "short": "URL of the ingredients image",
          },
          {
            "name": "image_nutrition_url",
            "title": "Image Nutrition Url",
            "type": "`$STRING`",
            "short": "URL of the nutrition facts image",
          },
          {
            "name": "image_url",
            "title": "Image Url",
            "type": "`$STRING`",
            "short": "URL of the product's front image",
          },
          {
            "name": "ingredients_analysis_tags",
            "title": "Ingredients Analysis Tags",
            "type": "`$ARRAY`",
            "short": "Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.)",
          },
          {
            "name": "ingredients_text",
            "title": "Ingredients Text",
            "type": "`$STRING`",
            "short": "List of ingredients as text",
          },
          {
            "name": "labels",
            "title": "Labels",
            "type": "`$STRING`",
            "short": "Labels associated with the product (e.g., Organic, Fair Trade)",
          },
          {
            "name": "last_modified_t",
            "title": "Last Modified T",
            "type": "`$INTEGER`",
            "short": "Last modification timestamp",
          },
          {
            "name": "manufacturing_places",
            "title": "Manufacturing Places",
            "type": "`$STRING`",
            "short": "Manufacturing or processing places",
          },
          {
            "name": "nova_group",
            "title": "Nova Group",
            "type": "`$INTEGER`",
            "short": "NOVA group for food processing level (1-4)",
          },
          {
            "name": "nutriments",
            "title": "Nutriments",
            "type": "`$OBJECT`",
            "short": "Nutritional information",
          },
          {
            "name": "nutriscore_grade",
            "title": "Nutriscore Grade",
            "type": "`$STRING`",
            "short": "Nutri-Score grade (a, b, c, d, e)",
          },
          {
            "name": "nutriscore_score",
            "title": "Nutriscore Score",
            "type": "`$INTEGER`",
            "short": "Nutri-Score numerical score",
          },
          {
            "name": "packaging",
            "title": "Packaging",
            "type": "`$STRING`",
            "short": "Packaging type",
          },
          {
            "name": "product_name",
            "title": "Product Name",
            "type": "`$STRING`",
            "short": "Name of the product",
          },
          {
            "name": "quantity",
            "title": "Quantity",
            "type": "`$STRING`",
            "short": "Quantity or volume of the product",
          },
          {
            "name": "stores",
            "title": "Stores",
            "type": "`$STRING`",
            "short": "Stores where the product is available",
          },
          {
            "name": "traces",
            "title": "Traces",
            "type": "`$STRING`",
            "short": "Traces of allergens",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "product",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/search",
                "segments": [
                  {
                    "lit": "search",
                  },
                ],
                "parts": [
                  "search",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.products`",
                },
                "args": {
                  "query": [
                    {
                      "name": "brand",
                      "orig": "brand",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "category",
                      "orig": "category",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "json",
                      "orig": "json",
                      "type": "`$BOOLEAN`",
                      "kind": "query",
                      "example": True,
                    },
                    {
                      "name": "label",
                      "orig": "label",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "page",
                      "orig": "page",
                      "type": "`$INTEGER`",
                      "kind": "query",
                      "example": 1,
                    },
                    {
                      "name": "page_size",
                      "orig": "page_size",
                      "type": "`$INTEGER`",
                      "kind": "query",
                      "example": 20,
                    },
                    {
                      "name": "search_term",
                      "orig": "search_term",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                  ],
                },
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
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
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
                "parts": [
                  "product",
                  "{barcode}.json",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.product`",
                },
                "args": {
                  "params": [
                    {
                      "name": "barcode",
                      "orig": "barcode",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                      "example": "737628064502",
                    },
                  ],
                },
                "select": {
                  "$action": "barcode",
                  "exist": [
                    "barcode",
                  ],
                },
              },
              {
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
                "parts": [
                  "product",
                  "{barcode}.xml",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "barcode",
                      "orig": "barcode",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                      "example": "737628064502",
                    },
                  ],
                },
                "select": {
                  "$action": "barcode",
                  "exist": [
                    "barcode",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/product/{barcode}",
                "segments": [
                  {
                    "lit": "product",
                  },
                  {
                    "var": "id",
                  },
                ],
                "parts": [
                  "product",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "barcode": "id",
                  },
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.product`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "barcode",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                      "example": "737628064502",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
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
