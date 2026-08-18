# Openfoodfacts SDK configuration


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
        },
        "feature": {
            "test": {
        "options": {
          "active": False,
        },
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
            "type": "`$ARRAY`",
          },
          {
            "name": "allergens",
            "type": "`$STRING`",
          },
          {
            "name": "brands",
            "type": "`$STRING`",
          },
          {
            "name": "categories",
            "type": "`$STRING`",
          },
          {
            "name": "countries",
            "type": "`$STRING`",
          },
          {
            "name": "created_t",
            "type": "`$INTEGER`",
          },
          {
            "name": "ecoscore_grade",
            "type": "`$STRING`",
          },
          {
            "name": "ecoscore_score",
            "type": "`$INTEGER`",
          },
          {
            "name": "generic_name",
            "type": "`$STRING`",
          },
          {
            "name": "image_front_url",
            "type": "`$STRING`",
          },
          {
            "name": "image_ingredients_url",
            "type": "`$STRING`",
          },
          {
            "name": "image_nutrition_url",
            "type": "`$STRING`",
          },
          {
            "name": "image_url",
            "type": "`$STRING`",
          },
          {
            "name": "ingredients_analysis_tags",
            "type": "`$ARRAY`",
          },
          {
            "name": "ingredients_text",
            "type": "`$STRING`",
          },
          {
            "name": "labels",
            "type": "`$STRING`",
          },
          {
            "name": "last_modified_t",
            "type": "`$INTEGER`",
          },
          {
            "name": "manufacturing_places",
            "type": "`$STRING`",
          },
          {
            "name": "nova_group",
            "type": "`$INTEGER`",
          },
          {
            "name": "nutriments",
            "type": "`$OBJECT`",
          },
          {
            "name": "nutriscore_grade",
            "type": "`$STRING`",
          },
          {
            "name": "nutriscore_score",
            "type": "`$INTEGER`",
          },
          {
            "name": "packaging",
            "type": "`$STRING`",
          },
          {
            "name": "product_name",
            "type": "`$STRING`",
          },
          {
            "name": "quantity",
            "type": "`$STRING`",
          },
          {
            "name": "stores",
            "type": "`$STRING`",
          },
          {
            "name": "traces",
            "type": "`$STRING`",
          },
        ],
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
                "parts": [
                  "product",
                  "{barcode}.json",
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
                "parts": [
                  "product",
                  "{barcode}.xml",
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
                "parts": [
                  "product",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "barcode": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.product`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [
            [
              "product",
            ],
          ],
        },
      },
      "search": {
        "fields": [
          {
            "name": "additives_tags",
            "type": "`$ARRAY`",
          },
          {
            "name": "allergens",
            "type": "`$STRING`",
          },
          {
            "name": "brands",
            "type": "`$STRING`",
          },
          {
            "name": "categories",
            "type": "`$STRING`",
          },
          {
            "name": "countries",
            "type": "`$STRING`",
          },
          {
            "name": "created_t",
            "type": "`$INTEGER`",
          },
          {
            "name": "ecoscore_grade",
            "type": "`$STRING`",
          },
          {
            "name": "ecoscore_score",
            "type": "`$INTEGER`",
          },
          {
            "name": "generic_name",
            "type": "`$STRING`",
          },
          {
            "name": "image_front_url",
            "type": "`$STRING`",
          },
          {
            "name": "image_ingredients_url",
            "type": "`$STRING`",
          },
          {
            "name": "image_nutrition_url",
            "type": "`$STRING`",
          },
          {
            "name": "image_url",
            "type": "`$STRING`",
          },
          {
            "name": "ingredients_analysis_tags",
            "type": "`$ARRAY`",
          },
          {
            "name": "ingredients_text",
            "type": "`$STRING`",
          },
          {
            "name": "labels",
            "type": "`$STRING`",
          },
          {
            "name": "last_modified_t",
            "type": "`$INTEGER`",
          },
          {
            "name": "manufacturing_places",
            "type": "`$STRING`",
          },
          {
            "name": "nova_group",
            "type": "`$INTEGER`",
          },
          {
            "name": "nutriments",
            "type": "`$OBJECT`",
          },
          {
            "name": "nutriscore_grade",
            "type": "`$STRING`",
          },
          {
            "name": "nutriscore_score",
            "type": "`$INTEGER`",
          },
          {
            "name": "packaging",
            "type": "`$STRING`",
          },
          {
            "name": "product_name",
            "type": "`$STRING`",
          },
          {
            "name": "quantity",
            "type": "`$STRING`",
          },
          {
            "name": "stores",
            "type": "`$STRING`",
          },
          {
            "name": "traces",
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
                "parts": [
                  "search",
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
