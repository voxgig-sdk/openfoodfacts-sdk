-- Openfoodfacts SDK configuration

-- Build a fresh, fully materialised config table. Every call rebuilds the
-- whole structure, so prefer require("config_shared") unless you need a
-- private copy you intend to mutate.
local function make_config()
  return {
    main = {
      name = "Openfoodfacts",
    },
    feature = {
      ["test"] = {
        ["options"] = {
          ["active"] = false,
        },
      },
    },
    options = {
      base = "https://world.openfoodfacts.org/api/v2",
      auth = {
        prefix = "",
      },
      headers = {
        ["content-type"] = "application/json",
      },
      entity = {
        ["product"] = {},
        ["search"] = {},
      },
    },
    entity = {
      ["product"] = {
        ["fields"] = {
          {
            ["name"] = "additives_tags",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "allergens",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "brands",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "categories",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "countries",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "created_t",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ecoscore_grade",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ecoscore_score",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "generic_name",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "image_front_url",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "image_ingredients_url",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "image_nutrition_url",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "image_url",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ingredients_analysis_tags",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ingredients_text",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "labels",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "last_modified_t",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "manufacturing_places",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "nova_group",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "nutriments",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "nutriscore_grade",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "nutriscore_score",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "packaging",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "product_name",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "quantity",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "stores",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "traces",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "product",
        ["op"] = {
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["example"] = "737628064502",
                      ["kind"] = "param",
                      ["name"] = "barcode",
                      ["orig"] = "barcode",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/product/{barcode}.json",
                ["parts"] = {
                  "product",
                  "{barcode}.json",
                },
                ["select"] = {
                  ["$action"] = "barcode",
                  ["exist"] = {
                    "barcode",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.product`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["example"] = "737628064502",
                      ["kind"] = "param",
                      ["name"] = "barcode",
                      ["orig"] = "barcode",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/product/{barcode}.xml",
                ["parts"] = {
                  "product",
                  "{barcode}.xml",
                },
                ["select"] = {
                  ["$action"] = "barcode",
                  ["exist"] = {
                    "barcode",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["example"] = "737628064502",
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "barcode",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/product/{barcode}",
                ["parts"] = {
                  "product",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["barcode"] = "id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.product`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "product",
            },
          },
        },
      },
      ["search"] = {
        ["fields"] = {
          {
            ["name"] = "additives_tags",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "allergens",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "brands",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "categories",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "countries",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "created_t",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ecoscore_grade",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ecoscore_score",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "generic_name",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "image_front_url",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "image_ingredients_url",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "image_nutrition_url",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "image_url",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ingredients_analysis_tags",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ingredients_text",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "labels",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "last_modified_t",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "manufacturing_places",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "nova_group",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "nutriments",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "nutriscore_grade",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "nutriscore_score",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "packaging",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "product_name",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "quantity",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "stores",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "traces",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "search",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "brand",
                      ["orig"] = "brand",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "category",
                      ["orig"] = "category",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = true,
                      ["kind"] = "query",
                      ["name"] = "json",
                      ["orig"] = "json",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "label",
                      ["orig"] = "label",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = 1,
                      ["kind"] = "query",
                      ["name"] = "page",
                      ["orig"] = "page",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["example"] = 20,
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search_term",
                      ["orig"] = "search_term",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/search",
                ["parts"] = {
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "brand",
                    "category",
                    "json",
                    "label",
                    "page",
                    "page_size",
                    "search_term",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.products`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
    },
  }
end


local function make_feature(name)
  local features = require("features")
  local factory = features[name]
  if factory ~= nil then
    return factory()
  end
  return features.base()
end


-- Attach make_feature to the SDK class
local function setup_sdk(SDK)
  SDK._make_feature = make_feature
end


return make_config
