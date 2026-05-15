# Openfoodfacts Lua SDK Reference

Complete API reference for the Openfoodfacts Lua SDK.


## OpenfoodfactsSDK

### Constructor

```lua
local sdk = require("openfoodfacts_sdk")
local client = sdk.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `table` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts, sdkopts)`

Create a test client with mock features active. Both arguments may be `nil`.

```lua
local client = sdk.test(nil, nil)
```


### Instance Methods

#### `Product(data)`

Create a new `Product` entity instance. Pass `nil` for no initial data.

#### `Search(data)`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `options_map() -> table`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs) -> table, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs.params` | `table` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `table` | Query string parameters. |
| `fetchargs.headers` | `table` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (tables are JSON-serialized). |
| `fetchargs.ctrl` | `table` | Control options (e.g. `{ explain = true }`). |

**Returns:** `table, err`

#### `prepare(fetchargs) -> table, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `table, err`


---

## ProductEntity

```lua
local product = client:Product(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `code` | ``$STRING`` | No |  |
| `product` | ``$OBJECT`` | No |  |
| `status` | ``$INTEGER`` | No |  |
| `status_verbose` | ``$STRING`` | No |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Product(nil):load({ id = "product_id" }, nil)
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ProductEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SearchEntity

```lua
local search = client:Search(nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Search(nil):list(nil, nil)
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```lua
local client = sdk.new({
  feature = {
    test = { active = true },
  },
})
```

