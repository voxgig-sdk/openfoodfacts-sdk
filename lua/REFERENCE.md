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

#### `sdk.test(testopts?, sdkopts?)`

Create a test client with mock features active. Both arguments are optional.

```lua
local client = sdk.test()
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
| `additives_tags` | `table` | No | List of additives |
| `allergens` | `string` | No | Allergens present in the product |
| `brands` | `string` | No | Brands of the product |
| `categories` | `string` | No | Categories the product belongs to |
| `countries` | `string` | No | Countries where the product is sold |
| `created_t` | `number` | No | Creation timestamp |
| `ecoscore_grade` | `string` | No | Eco-Score grade for environmental impact (a, b, c, d, e) |
| `ecoscore_score` | `number` | No | Eco-Score numerical score |
| `generic_name` | `string` | No | Generic name of the product |
| `id` | `string` | No |  |
| `image_front_url` | `string` | No | URL of the front image |
| `image_ingredients_url` | `string` | No | URL of the ingredients image |
| `image_nutrition_url` | `string` | No | URL of the nutrition facts image |
| `image_url` | `string` | No | URL of the product's front image |
| `ingredients_analysis_tags` | `table` | No | Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.) |
| `ingredients_text` | `string` | No | List of ingredients as text |
| `labels` | `string` | No | Labels associated with the product (e.g., Organic, Fair Trade) |
| `last_modified_t` | `number` | No | Last modification timestamp |
| `manufacturing_places` | `string` | No | Manufacturing or processing places |
| `nova_group` | `number` | No | NOVA group for food processing level (1-4) |
| `nutriments` | `table` | No | Nutritional information |
| `nutriscore_grade` | `string` | No | Nutri-Score grade (a, b, c, d, e) |
| `nutriscore_score` | `number` | No | Nutri-Score numerical score |
| `packaging` | `string` | No | Packaging type |
| `product_name` | `string` | No | Name of the product |
| `quantity` | `string` | No | Quantity or volume of the product |
| `stores` | `string` | No | Stores where the product is available |
| `traces` | `string` | No | Traces of allergens |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Product():load({ id = "product_id" })
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
| `additives_tags` | `table` | No | List of additives |
| `allergens` | `string` | No | Allergens present in the product |
| `brands` | `string` | No | Brands of the product |
| `categories` | `string` | No | Categories the product belongs to |
| `countries` | `string` | No | Countries where the product is sold |
| `created_t` | `number` | No | Creation timestamp |
| `ecoscore_grade` | `string` | No | Eco-Score grade for environmental impact (a, b, c, d, e) |
| `ecoscore_score` | `number` | No | Eco-Score numerical score |
| `generic_name` | `string` | No | Generic name of the product |
| `image_front_url` | `string` | No | URL of the front image |
| `image_ingredients_url` | `string` | No | URL of the ingredients image |
| `image_nutrition_url` | `string` | No | URL of the nutrition facts image |
| `image_url` | `string` | No | URL of the product's front image |
| `ingredients_analysis_tags` | `table` | No | Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.) |
| `ingredients_text` | `string` | No | List of ingredients as text |
| `labels` | `string` | No | Labels associated with the product (e.g., Organic, Fair Trade) |
| `last_modified_t` | `number` | No | Last modification timestamp |
| `manufacturing_places` | `string` | No | Manufacturing or processing places |
| `nova_group` | `number` | No | NOVA group for food processing level (1-4) |
| `nutriments` | `table` | No | Nutritional information |
| `nutriscore_grade` | `string` | No | Nutri-Score grade (a, b, c, d, e) |
| `nutriscore_score` | `number` | No | Nutri-Score numerical score |
| `packaging` | `string` | No | Packaging type |
| `product_name` | `string` | No | Name of the product |
| `quantity` | `string` | No | Quantity or volume of the product |
| `stores` | `string` | No | Stores where the product is available |
| `traces` | `string` | No | Traces of allergens |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Search():list()
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

