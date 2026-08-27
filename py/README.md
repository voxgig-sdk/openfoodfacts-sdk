# Openfoodfacts Python SDK



The Python SDK for the Openfoodfacts API — an entity-oriented client following Pythonic conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.Product()` — each
carrying a small, uniform set of operations (`list`, `load`) instead of raw URL
paths and query strings. You work with named resources and verbs, which
keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/openfoodfacts-sdk/releases)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
import os
from openfoodfacts_sdk import OpenfoodfactsSDK

client = OpenfoodfactsSDK({
    "apikey": os.environ.get("OPENFOODFACTS_APIKEY"),
})
```

### 3. Load a product

`load()` returns the ENTITY — call data_get() for the record — and raises on error.

```python
try:
    product = client.Product().load({"id": "example_id"})
    print(product)
except Exception as err:
    print(f"load failed: {err}")
```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    product = client.Product().load({"id": "example_id"})
    print(product)
except Exception as err:
    print(f"load failed: {err}")
```

`direct()` does **not** raise — it returns the result envelope. Branch
on `ok`; on failure `status` holds the HTTP status (for error responses)
and `err` holds a transport error, so read both defensively:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example_id"},
})

if not result["ok"]:
    print("request failed:", result.get("status"), result.get("err"))
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    # A non-2xx response carries status + data (the error body); a
    # transport-level failure carries err instead. Only one is present, so
    # read both with .get() rather than indexing a key that may be absent.
    print(result.get("status"), result.get("err"))
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = OpenfoodfactsSDK.test()

# Entity ops return the ENTITY and raises on error;
# call data_get() for the record.
product = client.Product().load({"id": "test01"})
# product contains the mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = OpenfoodfactsSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
    },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
OPENFOODFACTS_TEST_LIVE=TRUE
OPENFOODFACTS_APIKEY=<your-key>
```

Then run:

```bash
cd py && pytest test/
```


## Reference

### OpenfoodfactsSDK

```python
from openfoodfacts_sdk import OpenfoodfactsSDK

client = OpenfoodfactsSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `str` | API key for authentication. |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = OpenfoodfactsSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### OpenfoodfactsSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
| `Product` | `(data) -> ProductEntity` | Create a Product entity instance. |
| `Search` | `(data) -> SearchEntity` | Create a Search entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `list` | `(reqmatch, ctrl) -> list` | List entities matching the criteria. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (a `dict` for single-entity
ops, a `list` for `list`) and raise on error. Wrap calls in
`try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

### Entities

#### Product

| Field | Description |
| --- | --- |
| `additives_tags` | List of additives |
| `allergens` | Allergens present in the product |
| `brands` | Brands of the product |
| `categories` | Categories the product belongs to |
| `countries` | Countries where the product is sold |
| `created_t` | Creation timestamp |
| `ecoscore_grade` | Eco-Score grade for environmental impact (a, b, c, d, e) |
| `ecoscore_score` | Eco-Score numerical score |
| `generic_name` | Generic name of the product |
| `id` |  |
| `image_front_url` | URL of the front image |
| `image_ingredients_url` | URL of the ingredients image |
| `image_nutrition_url` | URL of the nutrition facts image |
| `image_url` | URL of the product's front image |
| `ingredients_analysis_tags` | Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.) |
| `ingredients_text` | List of ingredients as text |
| `labels` | Labels associated with the product (e.g., Organic, Fair Trade) |
| `last_modified_t` | Last modification timestamp |
| `manufacturing_places` | Manufacturing or processing places |
| `nova_group` | NOVA group for food processing level (1-4) |
| `nutriments` | Nutritional information |
| `nutriscore_grade` | Nutri-Score grade (a, b, c, d, e) |
| `nutriscore_score` | Nutri-Score numerical score |
| `packaging` | Packaging type |
| `product_name` | Name of the product |
| `quantity` | Quantity or volume of the product |
| `stores` | Stores where the product is available |
| `traces` | Traces of allergens |

Operations: Load.

API path: `/product/{barcode}.json`

#### Search

| Field | Description |
| --- | --- |
| `additives_tags` | List of additives |
| `allergens` | Allergens present in the product |
| `brands` | Brands of the product |
| `categories` | Categories the product belongs to |
| `countries` | Countries where the product is sold |
| `created_t` | Creation timestamp |
| `ecoscore_grade` | Eco-Score grade for environmental impact (a, b, c, d, e) |
| `ecoscore_score` | Eco-Score numerical score |
| `generic_name` | Generic name of the product |
| `image_front_url` | URL of the front image |
| `image_ingredients_url` | URL of the ingredients image |
| `image_nutrition_url` | URL of the nutrition facts image |
| `image_url` | URL of the product's front image |
| `ingredients_analysis_tags` | Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.) |
| `ingredients_text` | List of ingredients as text |
| `labels` | Labels associated with the product (e.g., Organic, Fair Trade) |
| `last_modified_t` | Last modification timestamp |
| `manufacturing_places` | Manufacturing or processing places |
| `nova_group` | NOVA group for food processing level (1-4) |
| `nutriments` | Nutritional information |
| `nutriscore_grade` | Nutri-Score grade (a, b, c, d, e) |
| `nutriscore_score` | Nutri-Score numerical score |
| `packaging` | Packaging type |
| `product_name` | Name of the product |
| `quantity` | Quantity or volume of the product |
| `stores` | Stores where the product is available |
| `traces` | Traces of allergens |

Operations: List.

API path: `/search`



## Entities


### Product

Create an instance: `product = client.Product()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `additives_tags` | `list` | List of additives |
| `allergens` | `str` | Allergens present in the product |
| `brands` | `str` | Brands of the product |
| `categories` | `str` | Categories the product belongs to |
| `countries` | `str` | Countries where the product is sold |
| `created_t` | `int` | Creation timestamp |
| `ecoscore_grade` | `str` | Eco-Score grade for environmental impact (a, b, c, d, e) |
| `ecoscore_score` | `int` | Eco-Score numerical score |
| `generic_name` | `str` | Generic name of the product |
| `id` | `str` |  |
| `image_front_url` | `str` | URL of the front image |
| `image_ingredients_url` | `str` | URL of the ingredients image |
| `image_nutrition_url` | `str` | URL of the nutrition facts image |
| `image_url` | `str` | URL of the product's front image |
| `ingredients_analysis_tags` | `list` | Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.) |
| `ingredients_text` | `str` | List of ingredients as text |
| `labels` | `str` | Labels associated with the product (e.g., Organic, Fair Trade) |
| `last_modified_t` | `int` | Last modification timestamp |
| `manufacturing_places` | `str` | Manufacturing or processing places |
| `nova_group` | `int` | NOVA group for food processing level (1-4) |
| `nutriments` | `dict` | Nutritional information |
| `nutriscore_grade` | `str` | Nutri-Score grade (a, b, c, d, e) |
| `nutriscore_score` | `int` | Nutri-Score numerical score |
| `packaging` | `str` | Packaging type |
| `product_name` | `str` | Name of the product |
| `quantity` | `str` | Quantity or volume of the product |
| `stores` | `str` | Stores where the product is available |
| `traces` | `str` | Traces of allergens |

#### Example: Load

```python
product = client.Product().load({"id": "product_id"})
```


### Search

Create an instance: `search = client.Search()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `additives_tags` | `list` | List of additives |
| `allergens` | `str` | Allergens present in the product |
| `brands` | `str` | Brands of the product |
| `categories` | `str` | Categories the product belongs to |
| `countries` | `str` | Countries where the product is sold |
| `created_t` | `int` | Creation timestamp |
| `ecoscore_grade` | `str` | Eco-Score grade for environmental impact (a, b, c, d, e) |
| `ecoscore_score` | `int` | Eco-Score numerical score |
| `generic_name` | `str` | Generic name of the product |
| `image_front_url` | `str` | URL of the front image |
| `image_ingredients_url` | `str` | URL of the ingredients image |
| `image_nutrition_url` | `str` | URL of the nutrition facts image |
| `image_url` | `str` | URL of the product's front image |
| `ingredients_analysis_tags` | `list` | Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.) |
| `ingredients_text` | `str` | List of ingredients as text |
| `labels` | `str` | Labels associated with the product (e.g., Organic, Fair Trade) |
| `last_modified_t` | `int` | Last modification timestamp |
| `manufacturing_places` | `str` | Manufacturing or processing places |
| `nova_group` | `int` | NOVA group for food processing level (1-4) |
| `nutriments` | `dict` | Nutritional information |
| `nutriscore_grade` | `str` | Nutri-Score grade (a, b, c, d, e) |
| `nutriscore_score` | `int` | Nutri-Score numerical score |
| `packaging` | `str` | Packaging type |
| `product_name` | `str` | Name of the product |
| `quantity` | `str` | Quantity or volume of the product |
| `stores` | `str` | Stores where the product is available |
| `traces` | `str` | Traces of allergens |

#### Example: List

```python
searchs = client.Search().list()
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── openfoodfacts_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`openfoodfacts_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```python
product = client.Product()
product.load({"id": "example_id"})

# product.data_get() now returns the product data from the last load
# product.match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
