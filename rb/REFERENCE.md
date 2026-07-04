# Openfoodfacts Ruby SDK Reference

Complete API reference for the Openfoodfacts Ruby SDK.


## OpenfoodfactsSDK

### Constructor

```ruby
require_relative 'openfoodfacts_sdk'

client = OpenfoodfactsSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["apikey"]` | `String` | API key for authentication. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `OpenfoodfactsSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = OpenfoodfactsSDK.test
```


### Instance Methods

#### `Product(data = nil)`

Create a new `Product` entity instance. Pass `nil` for no initial data.

#### `Search(data = nil)`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## ProductEntity

```ruby
product = client.Product
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `code` | ``$STRING`` | No |  |
| `product` | ``$OBJECT`` | No |  |
| `status` | ``$INTEGER`` | No |  |
| `status_verbose` | ``$STRING`` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Product.load({ "id" => "product_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ProductEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SearchEntity

```ruby
search = client.Search
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

#### `list(reqmatch, ctrl = nil) -> Array`

List entities matching the given criteria. Returns an array. Raises on error.

```ruby
results = client.Search.list(nil)
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ruby
client = OpenfoodfactsSDK.new({
  "feature" => {
    "test" => { "active" => true },
  },
})
```

