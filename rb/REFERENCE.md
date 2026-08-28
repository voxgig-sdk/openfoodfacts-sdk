# Openfoodfacts Ruby SDK Reference

Complete API reference for the Openfoodfacts Ruby SDK.


## OpenfoodfactsSDK

### Constructor

```ruby
require_relative 'Openfoodfacts_sdk'

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
| `additives_tags` | `Array` | No | List of additives |
| `allergens` | `String` | No | Allergens present in the product |
| `brands` | `String` | No | Brands of the product |
| `categories` | `String` | No | Categories the product belongs to |
| `countries` | `String` | No | Countries where the product is sold |
| `created_t` | `Integer` | No | Creation timestamp |
| `ecoscore_grade` | `String` | No | Eco-Score grade for environmental impact (a, b, c, d, e) |
| `ecoscore_score` | `Integer` | No | Eco-Score numerical score |
| `generic_name` | `String` | No | Generic name of the product |
| `id` | `String` | No |  |
| `image_front_url` | `String` | No | URL of the front image |
| `image_ingredients_url` | `String` | No | URL of the ingredients image |
| `image_nutrition_url` | `String` | No | URL of the nutrition facts image |
| `image_url` | `String` | No | URL of the product's front image |
| `ingredients_analysis_tags` | `Array` | No | Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.) |
| `ingredients_text` | `String` | No | List of ingredients as text |
| `labels` | `String` | No | Labels associated with the product (e.g., Organic, Fair Trade) |
| `last_modified_t` | `Integer` | No | Last modification timestamp |
| `manufacturing_places` | `String` | No | Manufacturing or processing places |
| `nova_group` | `Integer` | No | NOVA group for food processing level (1-4) |
| `nutriments` | `Hash` | No | Nutritional information |
| `nutriscore_grade` | `String` | No | Nutri-Score grade (a, b, c, d, e) |
| `nutriscore_score` | `Integer` | No | Nutri-Score numerical score |
| `packaging` | `String` | No | Packaging type |
| `product_name` | `String` | No | Name of the product |
| `quantity` | `String` | No | Quantity or volume of the product |
| `stores` | `String` | No | Stores where the product is available |
| `traces` | `String` | No | Traces of allergens |

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
| `additives_tags` | `Array` | No | List of additives |
| `allergens` | `String` | No | Allergens present in the product |
| `brands` | `String` | No | Brands of the product |
| `categories` | `String` | No | Categories the product belongs to |
| `countries` | `String` | No | Countries where the product is sold |
| `created_t` | `Integer` | No | Creation timestamp |
| `ecoscore_grade` | `String` | No | Eco-Score grade for environmental impact (a, b, c, d, e) |
| `ecoscore_score` | `Integer` | No | Eco-Score numerical score |
| `generic_name` | `String` | No | Generic name of the product |
| `image_front_url` | `String` | No | URL of the front image |
| `image_ingredients_url` | `String` | No | URL of the ingredients image |
| `image_nutrition_url` | `String` | No | URL of the nutrition facts image |
| `image_url` | `String` | No | URL of the product's front image |
| `ingredients_analysis_tags` | `Array` | No | Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.) |
| `ingredients_text` | `String` | No | List of ingredients as text |
| `labels` | `String` | No | Labels associated with the product (e.g., Organic, Fair Trade) |
| `last_modified_t` | `Integer` | No | Last modification timestamp |
| `manufacturing_places` | `String` | No | Manufacturing or processing places |
| `nova_group` | `Integer` | No | NOVA group for food processing level (1-4) |
| `nutriments` | `Hash` | No | Nutritional information |
| `nutriscore_grade` | `String` | No | Nutri-Score grade (a, b, c, d, e) |
| `nutriscore_score` | `Integer` | No | Nutri-Score numerical score |
| `packaging` | `String` | No | Packaging type |
| `product_name` | `String` | No | Name of the product |
| `quantity` | `String` | No | Quantity or volume of the product |
| `stores` | `String` | No | Stores where the product is available |
| `traces` | `String` | No | Traces of allergens |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Search.list
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


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### `test`

In-memory mock transport for testing without a live server.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

Options above are those the model carries a default for. A feature may
also accept callback options — a `sink` to receive each record, for
instance — which have no default and are covered in the full feature
reference.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

