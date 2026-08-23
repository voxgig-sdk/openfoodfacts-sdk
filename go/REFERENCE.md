# Openfoodfacts Golang SDK Reference

Complete API reference for the Openfoodfacts Golang SDK.


## OpenfoodfactsSDK

### Constructor

```go
func NewOpenfoodfactsSDK(options map[string]any) *OpenfoodfactsSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["apikey"]` | `string` | API key for authentication. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *OpenfoodfactsSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *OpenfoodfactsSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `Product(data map[string]any) OpenfoodfactsEntity`

Create a new `Product` entity instance. Pass `nil` for no initial data.

#### `Search(data map[string]any) OpenfoodfactsEntity`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## ProductEntity

```go
product := client.Product(nil)
fmt.Println(product.GetName()) // "product"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additives_tags` | `[]any` | No | List of additives |
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
| `ingredients_analysis_tags` | `[]any` | No | Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.) |
| `ingredients_text` | `string` | No | List of ingredients as text |
| `labels` | `string` | No | Labels associated with the product (e.g., Organic, Fair Trade) |
| `last_modified_t` | `int` | No | Last modification timestamp |
| `manufacturing_places` | `string` | No | Manufacturing or processing places |
| `nova_group` | `int` | No | NOVA group for food processing level (1-4) |
| `nutriments` | `map[string]any` | No | Nutritional information |
| `nutriscore_grade` | `string` | No | Nutri-Score grade (a, b, c, d, e) |
| `nutriscore_score` | `int` | No | Nutri-Score numerical score |
| `packaging` | `string` | No | Packaging type |
| `product_name` | `string` | No | Name of the product |
| `quantity` | `string` | No | Quantity or volume of the product |
| `stores` | `string` | No | Stores where the product is available |
| `traces` | `string` | No | Traces of allergens |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Product(nil).Load(map[string]any{"id": "product_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ProductEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SearchEntity

```go
search := client.Search(nil)
fmt.Println(search.GetName()) // "search"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additives_tags` | `[]any` | No | List of additives |
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
| `ingredients_analysis_tags` | `[]any` | No | Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.) |
| `ingredients_text` | `string` | No | List of ingredients as text |
| `labels` | `string` | No | Labels associated with the product (e.g., Organic, Fair Trade) |
| `last_modified_t` | `int` | No | Last modification timestamp |
| `manufacturing_places` | `string` | No | Manufacturing or processing places |
| `nova_group` | `int` | No | NOVA group for food processing level (1-4) |
| `nutriments` | `map[string]any` | No | Nutritional information |
| `nutriscore_grade` | `string` | No | Nutri-Score grade (a, b, c, d, e) |
| `nutriscore_score` | `int` | No | Nutri-Score numerical score |
| `packaging` | `string` | No | Packaging type |
| `product_name` | `string` | No | Name of the product |
| `quantity` | `string` | No | Quantity or volume of the product |
| `stores` | `string` | No | Stores where the product is available |
| `traces` | `string` | No | Traces of allergens |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Search(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```go
client := sdk.NewOpenfoodfactsSDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
    },
})
```

