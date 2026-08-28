# Openfoodfacts TypeScript SDK Reference

Complete API reference for the Openfoodfacts TypeScript SDK.


## OpenfoodfactsSDK

### Constructor

```ts
new OpenfoodfactsSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `OpenfoodfactsSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = OpenfoodfactsSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `OpenfoodfactsSDK` instance in test mode.


### Instance Methods

#### `Product(data?: object)`

Create a new `Product` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ProductEntity` instance.

#### `Search(data?: object)`

Create a new `Search` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SearchEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `OpenfoodfactsSDK.test()`.

**Returns:** `OpenfoodfactsSDK` instance in test mode.


---

## ProductEntity

```ts
const product = client.Product()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additives_tags` | `any[]` | No | List of additives |
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
| `ingredients_analysis_tags` | `any[]` | No | Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.) |
| `ingredients_text` | `string` | No | List of ingredients as text |
| `labels` | `string` | No | Labels associated with the product (e.g., Organic, Fair Trade) |
| `last_modified_t` | `number` | No | Last modification timestamp |
| `manufacturing_places` | `string` | No | Manufacturing or processing places |
| `nova_group` | `number` | No | NOVA group for food processing level (1-4) |
| `nutriments` | `Record<string, any>` | No | Nutritional information |
| `nutriscore_grade` | `string` | No | Nutri-Score grade (a, b, c, d, e) |
| `nutriscore_score` | `number` | No | Nutri-Score numerical score |
| `packaging` | `string` | No | Packaging type |
| `product_name` | `string` | No | Name of the product |
| `quantity` | `string` | No | Quantity or volume of the product |
| `stores` | `string` | No | Stores where the product is available |
| `traces` | `string` | No | Traces of allergens |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `barcode` | `/product/{barcode}.json` | `client.Product().load({ $action: 'barcode', ... })` |
| `barcode` | `/product/{barcode}.xml` | `client.Product().load({ $action: 'barcode', ... })` |

An action returns that action's OWN response, which is not necessarily a
Product record — check the API definition for its shape.

```ts
const result = await client.Product().load({
  $action: 'barcode',
  /* ...the action's own arguments */
})
```

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Product().load({ id: 'product_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ProductEntity` instance with the same client and
options.

#### `client()`

Return the parent `OpenfoodfactsSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SearchEntity

```ts
const search = client.Search()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additives_tags` | `any[]` | No | List of additives |
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
| `ingredients_analysis_tags` | `any[]` | No | Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.) |
| `ingredients_text` | `string` | No | List of ingredients as text |
| `labels` | `string` | No | Labels associated with the product (e.g., Organic, Fair Trade) |
| `last_modified_t` | `number` | No | Last modification timestamp |
| `manufacturing_places` | `string` | No | Manufacturing or processing places |
| `nova_group` | `number` | No | NOVA group for food processing level (1-4) |
| `nutriments` | `Record<string, any>` | No | Nutritional information |
| `nutriscore_grade` | `string` | No | Nutri-Score grade (a, b, c, d, e) |
| `nutriscore_score` | `number` | No | Nutri-Score numerical score |
| `packaging` | `string` | No | Packaging type |
| `product_name` | `string` | No | Name of the product |
| `quantity` | `string` | No | Quantity or volume of the product |
| `stores` | `string` | No | Stores where the product is available |
| `traces` | `string` | No | Traces of allergens |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Search().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SearchEntity` instance with the same client and
options.

#### `client()`

Return the parent `OpenfoodfactsSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new OpenfoodfactsSDK({
  feature: {
    test: { active: true },
  }
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

