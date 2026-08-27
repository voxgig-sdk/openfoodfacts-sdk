<?php
declare(strict_types=1);

// Openfoodfacts SDK configuration

class OpenfoodfactsConfig
{
    /** @var array<string,mixed>|null */
    private static ?array $shared_config = null;

    /**
     * Return the process-wide config, built once on first use. The SDK reads
     * the config on every request and never writes to it, so one instance is
     * shared by every client rather than rebuilt per client.
     *
     * PHP arrays are copy-on-write, so callers that do mutate the result get
     * their own copy and cannot disturb the shared one.
     */
    public static function shared_config(): array
    {
        if (self::$shared_config === null) {
            self::$shared_config = self::make_config();
        }
        return self::$shared_config;
    }

    /**
     * Build a fresh, fully materialised config array. Every call rebuilds the
     * whole structure, so prefer shared_config unless you need a private copy.
     */
    public static function make_config(): array
    {
        return [
            "main" => [
                "name" => "Openfoodfacts",
                "slug" => "openfoodfacts",
                "version" => "0.0.1",
                "target" => "php",
            ],
            "feature" => [
                "test" => [
          'options' => [
            'active' => false,
          ],
          'transport' => 'base',
        ],
            ],
            "options" => [
                "base" => "https://world.openfoodfacts.org/api/v2",
                "auth" => [
                    "prefix" => "",
                ],
                "headers" => [
          'content-type' => 'application/json',
        ],
                "entity" => [
                    "product" => [],
                    "search" => [],
                ],
            ],
            "entity" => [
        'product' => [
          'fields' => [
            [
              'name' => 'additives_tags',
              'short' => 'List of additives',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'allergens',
              'short' => 'Allergens present in the product',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'brands',
              'short' => 'Brands of the product',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'categories',
              'short' => 'Categories the product belongs to',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'countries',
              'short' => 'Countries where the product is sold',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'created_t',
              'short' => 'Creation timestamp',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'ecoscore_grade',
              'short' => 'Eco-Score grade for environmental impact (a, b, c, d, e)',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'ecoscore_score',
              'short' => 'Eco-Score numerical score',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'generic_name',
              'short' => 'Generic name of the product',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'image_front_url',
              'short' => 'URL of the front image',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'image_ingredients_url',
              'short' => 'URL of the ingredients image',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'image_nutrition_url',
              'short' => 'URL of the nutrition facts image',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'image_url',
              'short' => 'URL of the product\'s front image',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'ingredients_analysis_tags',
              'short' => 'Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.)',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'ingredients_text',
              'short' => 'List of ingredients as text',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'labels',
              'short' => 'Labels associated with the product (e.g., Organic, Fair Trade)',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'last_modified_t',
              'short' => 'Last modification timestamp',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'manufacturing_places',
              'short' => 'Manufacturing or processing places',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'nova_group',
              'short' => 'NOVA group for food processing level (1-4)',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'nutriments',
              'short' => 'Nutritional information',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'nutriscore_grade',
              'short' => 'Nutri-Score grade (a, b, c, d, e)',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'nutriscore_score',
              'short' => 'Nutri-Score numerical score',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'packaging',
              'short' => 'Packaging type',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'product_name',
              'short' => 'Name of the product',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'quantity',
              'short' => 'Quantity or volume of the product',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'stores',
              'short' => 'Stores where the product is available',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'traces',
              'short' => 'Traces of allergens',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'product',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => '737628064502',
                        'kind' => 'param',
                        'name' => 'barcode',
                        'orig' => 'barcode',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/product/{barcode}.json',
                  'parts' => [
                    'product',
                    '{barcode}.json',
                  ],
                  'select' => [
                    '$action' => 'barcode',
                    'exist' => [
                      'barcode',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.product`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => '737628064502',
                        'kind' => 'param',
                        'name' => 'barcode',
                        'orig' => 'barcode',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/product/{barcode}.xml',
                  'parts' => [
                    'product',
                    '{barcode}.xml',
                  ],
                  'select' => [
                    '$action' => 'barcode',
                    'exist' => [
                      'barcode',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => '737628064502',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'barcode',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/product/{barcode}',
                  'parts' => [
                    'product',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'barcode' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.product`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'product',
              ],
            ],
          ],
        ],
        'search' => [
          'fields' => [
            [
              'name' => 'additives_tags',
              'short' => 'List of additives',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'allergens',
              'short' => 'Allergens present in the product',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'brands',
              'short' => 'Brands of the product',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'categories',
              'short' => 'Categories the product belongs to',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'countries',
              'short' => 'Countries where the product is sold',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'created_t',
              'short' => 'Creation timestamp',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'ecoscore_grade',
              'short' => 'Eco-Score grade for environmental impact (a, b, c, d, e)',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'ecoscore_score',
              'short' => 'Eco-Score numerical score',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'generic_name',
              'short' => 'Generic name of the product',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'image_front_url',
              'short' => 'URL of the front image',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'image_ingredients_url',
              'short' => 'URL of the ingredients image',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'image_nutrition_url',
              'short' => 'URL of the nutrition facts image',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'image_url',
              'short' => 'URL of the product\'s front image',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'ingredients_analysis_tags',
              'short' => 'Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.)',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'ingredients_text',
              'short' => 'List of ingredients as text',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'labels',
              'short' => 'Labels associated with the product (e.g., Organic, Fair Trade)',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'last_modified_t',
              'short' => 'Last modification timestamp',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'manufacturing_places',
              'short' => 'Manufacturing or processing places',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'nova_group',
              'short' => 'NOVA group for food processing level (1-4)',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'nutriments',
              'short' => 'Nutritional information',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'nutriscore_grade',
              'short' => 'Nutri-Score grade (a, b, c, d, e)',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'nutriscore_score',
              'short' => 'Nutri-Score numerical score',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'packaging',
              'short' => 'Packaging type',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'product_name',
              'short' => 'Name of the product',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'quantity',
              'short' => 'Quantity or volume of the product',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'stores',
              'short' => 'Stores where the product is available',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'traces',
              'short' => 'Traces of allergens',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'search',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'brand',
                        'orig' => 'brand',
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'category',
                        'orig' => 'category',
                        'type' => '`$STRING`',
                      ],
                      [
                        'example' => true,
                        'kind' => 'query',
                        'name' => 'json',
                        'orig' => 'json',
                        'type' => '`$BOOLEAN`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'label',
                        'orig' => 'label',
                        'type' => '`$STRING`',
                      ],
                      [
                        'example' => 1,
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'example' => 20,
                        'kind' => 'query',
                        'name' => 'page_size',
                        'orig' => 'page_size',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'search_term',
                        'orig' => 'search_term',
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/search',
                  'parts' => [
                    'search',
                  ],
                  'select' => [
                    'exist' => [
                      'brand',
                      'category',
                      'json',
                      'label',
                      'page',
                      'page_size',
                      'search_term',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.products`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
      ],
        ];
    }


    public static function make_feature(string $name)
    {
        require_once __DIR__ . '/features.php';
        return OpenfoodfactsFeatures::make_feature($name);
    }
}
