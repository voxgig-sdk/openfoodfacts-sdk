# Openfoodfacts Python SDK Reference

Complete API reference for the Openfoodfacts Python SDK.


## OpenfoodfactsSDK

### Constructor

```python
from openfoodfacts_sdk import OpenfoodfactsSDK

client = OpenfoodfactsSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["apikey"]` | `str` | API key for authentication. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `OpenfoodfactsSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = OpenfoodfactsSDK.test()
```


### Instance Methods

#### `Product(data=None)`

Create a new `ProductEntity` instance. Pass `None` for no initial data.

#### `Search(data=None)`

Create a new `SearchEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## ProductEntity

```python
product = client.Product()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additives_tags` | `list` | No | List of additives |
| `allergens` | `str` | No | Allergens present in the product |
| `brands` | `str` | No | Brands of the product |
| `categories` | `str` | No | Categories the product belongs to |
| `countries` | `str` | No | Countries where the product is sold |
| `created_t` | `int` | No | Creation timestamp |
| `ecoscore_grade` | `str` | No | Eco-Score grade for environmental impact (a, b, c, d, e) |
| `ecoscore_score` | `int` | No | Eco-Score numerical score |
| `generic_name` | `str` | No | Generic name of the product |
| `id` | `str` | No |  |
| `image_front_url` | `str` | No | URL of the front image |
| `image_ingredients_url` | `str` | No | URL of the ingredients image |
| `image_nutrition_url` | `str` | No | URL of the nutrition facts image |
| `image_url` | `str` | No | URL of the product's front image |
| `ingredients_analysis_tags` | `list` | No | Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.) |
| `ingredients_text` | `str` | No | List of ingredients as text |
| `labels` | `str` | No | Labels associated with the product (e.g., Organic, Fair Trade) |
| `last_modified_t` | `int` | No | Last modification timestamp |
| `manufacturing_places` | `str` | No | Manufacturing or processing places |
| `nova_group` | `int` | No | NOVA group for food processing level (1-4) |
| `nutriments` | `dict` | No | Nutritional information |
| `nutriscore_grade` | `str` | No | Nutri-Score grade (a, b, c, d, e) |
| `nutriscore_score` | `int` | No | Nutri-Score numerical score |
| `packaging` | `str` | No | Packaging type |
| `product_name` | `str` | No | Name of the product |
| `quantity` | `str` | No | Quantity or volume of the product |
| `stores` | `str` | No | Stores where the product is available |
| `traces` | `str` | No | Traces of allergens |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Product().load({"id": "product_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ProductEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SearchEntity

```python
search = client.Search()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additives_tags` | `list` | No | List of additives |
| `allergens` | `str` | No | Allergens present in the product |
| `brands` | `str` | No | Brands of the product |
| `categories` | `str` | No | Categories the product belongs to |
| `countries` | `str` | No | Countries where the product is sold |
| `created_t` | `int` | No | Creation timestamp |
| `ecoscore_grade` | `str` | No | Eco-Score grade for environmental impact (a, b, c, d, e) |
| `ecoscore_score` | `int` | No | Eco-Score numerical score |
| `generic_name` | `str` | No | Generic name of the product |
| `image_front_url` | `str` | No | URL of the front image |
| `image_ingredients_url` | `str` | No | URL of the ingredients image |
| `image_nutrition_url` | `str` | No | URL of the nutrition facts image |
| `image_url` | `str` | No | URL of the product's front image |
| `ingredients_analysis_tags` | `list` | No | Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.) |
| `ingredients_text` | `str` | No | List of ingredients as text |
| `labels` | `str` | No | Labels associated with the product (e.g., Organic, Fair Trade) |
| `last_modified_t` | `int` | No | Last modification timestamp |
| `manufacturing_places` | `str` | No | Manufacturing or processing places |
| `nova_group` | `int` | No | NOVA group for food processing level (1-4) |
| `nutriments` | `dict` | No | Nutritional information |
| `nutriscore_grade` | `str` | No | Nutri-Score grade (a, b, c, d, e) |
| `nutriscore_score` | `int` | No | Nutri-Score numerical score |
| `packaging` | `str` | No | Packaging type |
| `product_name` | `str` | No | Name of the product |
| `quantity` | `str` | No | Quantity or volume of the product |
| `stores` | `str` | No | Stores where the product is available |
| `traces` | `str` | No | Traces of allergens |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Search().list()
for search in results:
    print(search)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SearchEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```python
client = OpenfoodfactsSDK({
    "feature": {
        "test": {"active": True},
    },
})
```

