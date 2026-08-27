# Openfoodfacts PHP SDK Reference

Complete API reference for the Openfoodfacts PHP SDK.


## OpenfoodfactsSDK

### Constructor

```php
require_once __DIR__ . '/openfoodfacts_sdk.php';

$client = new OpenfoodfactsSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["apikey"]` | `string` | API key for authentication. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `OpenfoodfactsSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = OpenfoodfactsSDK::test();
```


### Instance Methods

#### `Product($data = null)`

Create a new `ProductEntity` instance. Pass `null` for no initial data.

#### `Search($data = null)`

Create a new `SearchEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): OpenfoodfactsUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## ProductEntity

```php
$product = $client->Product();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additives_tags` | `array` | No | List of additives |
| `allergens` | `string` | No | Allergens present in the product |
| `brands` | `string` | No | Brands of the product |
| `categories` | `string` | No | Categories the product belongs to |
| `countries` | `string` | No | Countries where the product is sold |
| `created_t` | `int` | No | Creation timestamp |
| `ecoscore_grade` | `string` | No | Eco-Score grade for environmental impact (a, b, c, d, e) |
| `ecoscore_score` | `int` | No | Eco-Score numerical score |
| `generic_name` | `string` | No | Generic name of the product |
| `id` | `string` | No |  |
| `image_front_url` | `string` | No | URL of the front image |
| `image_ingredients_url` | `string` | No | URL of the ingredients image |
| `image_nutrition_url` | `string` | No | URL of the nutrition facts image |
| `image_url` | `string` | No | URL of the product's front image |
| `ingredients_analysis_tags` | `array` | No | Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.) |
| `ingredients_text` | `string` | No | List of ingredients as text |
| `labels` | `string` | No | Labels associated with the product (e.g., Organic, Fair Trade) |
| `last_modified_t` | `int` | No | Last modification timestamp |
| `manufacturing_places` | `string` | No | Manufacturing or processing places |
| `nova_group` | `int` | No | NOVA group for food processing level (1-4) |
| `nutriments` | `array` | No | Nutritional information |
| `nutriscore_grade` | `string` | No | Nutri-Score grade (a, b, c, d, e) |
| `nutriscore_score` | `int` | No | Nutri-Score numerical score |
| `packaging` | `string` | No | Packaging type |
| `product_name` | `string` | No | Name of the product |
| `quantity` | `string` | No | Quantity or volume of the product |
| `stores` | `string` | No | Stores where the product is available |
| `traces` | `string` | No | Traces of allergens |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Product()->load(["id" => "product_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ProductEntity`

Create a new `ProductEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SearchEntity

```php
$search = $client->Search();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additives_tags` | `array` | No | List of additives |
| `allergens` | `string` | No | Allergens present in the product |
| `brands` | `string` | No | Brands of the product |
| `categories` | `string` | No | Categories the product belongs to |
| `countries` | `string` | No | Countries where the product is sold |
| `created_t` | `int` | No | Creation timestamp |
| `ecoscore_grade` | `string` | No | Eco-Score grade for environmental impact (a, b, c, d, e) |
| `ecoscore_score` | `int` | No | Eco-Score numerical score |
| `generic_name` | `string` | No | Generic name of the product |
| `image_front_url` | `string` | No | URL of the front image |
| `image_ingredients_url` | `string` | No | URL of the ingredients image |
| `image_nutrition_url` | `string` | No | URL of the nutrition facts image |
| `image_url` | `string` | No | URL of the product's front image |
| `ingredients_analysis_tags` | `array` | No | Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.) |
| `ingredients_text` | `string` | No | List of ingredients as text |
| `labels` | `string` | No | Labels associated with the product (e.g., Organic, Fair Trade) |
| `last_modified_t` | `int` | No | Last modification timestamp |
| `manufacturing_places` | `string` | No | Manufacturing or processing places |
| `nova_group` | `int` | No | NOVA group for food processing level (1-4) |
| `nutriments` | `array` | No | Nutritional information |
| `nutriscore_grade` | `string` | No | Nutri-Score grade (a, b, c, d, e) |
| `nutriscore_score` | `int` | No | Nutri-Score numerical score |
| `packaging` | `string` | No | Packaging type |
| `product_name` | `string` | No | Name of the product |
| `quantity` | `string` | No | Quantity or volume of the product |
| `stores` | `string` | No | Stores where the product is available |
| `traces` | `string` | No | Traces of allergens |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Search()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SearchEntity`

Create a new `SearchEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```php
$client = new OpenfoodfactsSDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
```

