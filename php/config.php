<?php
declare(strict_types=1);

// Openfoodfacts SDK configuration

class OpenfoodfactsConfig
{
    public static function make_config(): array
    {
        return [
            "main" => [
                "name" => "Openfoodfacts",
            ],
            "feature" => [
                "test" => [
          'options' => [
            'active' => false,
          ],
        ],
            ],
            "options" => [
                "base" => "https://world.openfoodfacts.org/api/v2",
                "auth" => [
                    "prefix" => "Bearer",
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
              'name' => 'code',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 0,
            ],
            [
              'name' => 'product',
              'req' => false,
              'type' => '`$OBJECT`',
              'active' => true,
              'index$' => 1,
            ],
            [
              'name' => 'status',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 2,
            ],
            [
              'name' => 'status_verbose',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 3,
            ],
          ],
          'name' => 'product',
          'op' => [
            'load' => [
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
                        'active' => true,
                      ],
                    ],
                  ],
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
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
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
                        'active' => true,
                      ],
                    ],
                  ],
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
                  'active' => true,
                  'index$' => 1,
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
                        'active' => true,
                      ],
                    ],
                  ],
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
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 2,
                ],
              ],
              'input' => 'data',
              'key$' => 'load',
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
              'name' => 'additives_tag',
              'req' => false,
              'type' => '`$ARRAY`',
              'active' => true,
              'index$' => 0,
            ],
            [
              'name' => 'allergen',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 1,
            ],
            [
              'name' => 'brand',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 2,
            ],
            [
              'name' => 'category',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 3,
            ],
            [
              'name' => 'country',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 4,
            ],
            [
              'name' => 'created_t',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 5,
            ],
            [
              'name' => 'ecoscore_grade',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 6,
            ],
            [
              'name' => 'ecoscore_score',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 7,
            ],
            [
              'name' => 'generic_name',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 8,
            ],
            [
              'name' => 'image_front_url',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 9,
            ],
            [
              'name' => 'image_ingredients_url',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 10,
            ],
            [
              'name' => 'image_nutrition_url',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 11,
            ],
            [
              'name' => 'image_url',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 12,
            ],
            [
              'name' => 'ingredients_analysis_tag',
              'req' => false,
              'type' => '`$ARRAY`',
              'active' => true,
              'index$' => 13,
            ],
            [
              'name' => 'ingredients_text',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 14,
            ],
            [
              'name' => 'label',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 15,
            ],
            [
              'name' => 'last_modified_t',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 16,
            ],
            [
              'name' => 'manufacturing_place',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 17,
            ],
            [
              'name' => 'nova_group',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 18,
            ],
            [
              'name' => 'nutriment',
              'req' => false,
              'type' => '`$OBJECT`',
              'active' => true,
              'index$' => 19,
            ],
            [
              'name' => 'nutriscore_grade',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 20,
            ],
            [
              'name' => 'nutriscore_score',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 21,
            ],
            [
              'name' => 'packaging',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 22,
            ],
            [
              'name' => 'product_name',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 23,
            ],
            [
              'name' => 'quantity',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 24,
            ],
            [
              'name' => 'store',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 25,
            ],
            [
              'name' => 'trace',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 26,
            ],
          ],
          'name' => 'search',
          'op' => [
            'list' => [
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'brand',
                        'orig' => 'brand',
                        'reqd' => false,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'category',
                        'orig' => 'category',
                        'reqd' => false,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => true,
                        'kind' => 'query',
                        'name' => 'json',
                        'orig' => 'json',
                        'reqd' => false,
                        'type' => '`$BOOLEAN`',
                        'active' => true,
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'label',
                        'orig' => 'label',
                        'reqd' => false,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 1,
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'reqd' => false,
                        'type' => '`$INTEGER`',
                        'active' => true,
                      ],
                      [
                        'example' => 20,
                        'kind' => 'query',
                        'name' => 'page_size',
                        'orig' => 'page_size',
                        'reqd' => false,
                        'type' => '`$INTEGER`',
                        'active' => true,
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'search_term',
                        'orig' => 'search_term',
                        'reqd' => false,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
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
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
              ],
              'input' => 'data',
              'key$' => 'list',
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
