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

#### `optionsMap(): array`

Return a deep copy of the current SDK options.

#### `getUtility(): ProjectNameUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. Returns `[$result, $err]`.

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

**Returns:** `array [$result, $err]`

#### `prepare(array $fetchargs = []): array`

Prepare a fetch definition without sending the request. Returns `[$fetchdef, $err]`.


---

## ProductEntity

```php
$product = $client->Product();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `code` | ``$STRING`` | No |  |
| `product` | ``$OBJECT`` | No |  |
| `status` | ``$INTEGER`` | No |  |
| `status_verbose` | ``$STRING`` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Product()->load(["id" => "product_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): ProductEntity`

Create a new `ProductEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## SearchEntity

```php
$search = $client->Search();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additives_tag` | ``$ARRAY`` | No |  |
| `allergen` | ``$STRING`` | No |  |
| `brand` | ``$STRING`` | No |  |
| `category` | ``$STRING`` | No |  |
| `country` | ``$STRING`` | No |  |
| `created_t` | ``$INTEGER`` | No |  |
| `ecoscore_grade` | ``$STRING`` | No |  |
| `ecoscore_score` | ``$INTEGER`` | No |  |
| `generic_name` | ``$STRING`` | No |  |
| `image_front_url` | ``$STRING`` | No |  |
| `image_ingredients_url` | ``$STRING`` | No |  |
| `image_nutrition_url` | ``$STRING`` | No |  |
| `image_url` | ``$STRING`` | No |  |
| `ingredients_analysis_tag` | ``$ARRAY`` | No |  |
| `ingredients_text` | ``$STRING`` | No |  |
| `label` | ``$STRING`` | No |  |
| `last_modified_t` | ``$INTEGER`` | No |  |
| `manufacturing_place` | ``$STRING`` | No |  |
| `nova_group` | ``$INTEGER`` | No |  |
| `nutriment` | ``$OBJECT`` | No |  |
| `nutriscore_grade` | ``$STRING`` | No |  |
| `nutriscore_score` | ``$INTEGER`` | No |  |
| `packaging` | ``$STRING`` | No |  |
| `product_name` | ``$STRING`` | No |  |
| `quantity` | ``$STRING`` | No |  |
| `store` | ``$STRING`` | No |  |
| `trace` | ``$STRING`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Search()->list([]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): SearchEntity`

Create a new `SearchEntity` instance with the same client and
options.

#### `getName(): string`

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

