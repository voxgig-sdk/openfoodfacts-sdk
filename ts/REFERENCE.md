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
| `additives_tags` | `any[]` | No |  |
| `allergens` | `string` | No |  |
| `brands` | `string` | No |  |
| `categories` | `string` | No |  |
| `countries` | `string` | No |  |
| `created_t` | `number` | No |  |
| `ecoscore_grade` | `string` | No |  |
| `ecoscore_score` | `number` | No |  |
| `generic_name` | `string` | No |  |
| `image_front_url` | `string` | No |  |
| `image_ingredients_url` | `string` | No |  |
| `image_nutrition_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `ingredients_analysis_tags` | `any[]` | No |  |
| `ingredients_text` | `string` | No |  |
| `labels` | `string` | No |  |
| `last_modified_t` | `number` | No |  |
| `manufacturing_places` | `string` | No |  |
| `nova_group` | `number` | No |  |
| `nutriments` | `Record<string, any>` | No |  |
| `nutriscore_grade` | `string` | No |  |
| `nutriscore_score` | `number` | No |  |
| `packaging` | `string` | No |  |
| `product_name` | `string` | No |  |
| `quantity` | `string` | No |  |
| `stores` | `string` | No |  |
| `traces` | `string` | No |  |

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
| `additives_tags` | `any[]` | No |  |
| `allergens` | `string` | No |  |
| `brands` | `string` | No |  |
| `categories` | `string` | No |  |
| `countries` | `string` | No |  |
| `created_t` | `number` | No |  |
| `ecoscore_grade` | `string` | No |  |
| `ecoscore_score` | `number` | No |  |
| `generic_name` | `string` | No |  |
| `image_front_url` | `string` | No |  |
| `image_ingredients_url` | `string` | No |  |
| `image_nutrition_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `ingredients_analysis_tags` | `any[]` | No |  |
| `ingredients_text` | `string` | No |  |
| `labels` | `string` | No |  |
| `last_modified_t` | `number` | No |  |
| `manufacturing_places` | `string` | No |  |
| `nova_group` | `number` | No |  |
| `nutriments` | `Record<string, any>` | No |  |
| `nutriscore_grade` | `string` | No |  |
| `nutriscore_score` | `number` | No |  |
| `packaging` | `string` | No |  |
| `product_name` | `string` | No |  |
| `quantity` | `string` | No |  |
| `stores` | `string` | No |  |
| `traces` | `string` | No |  |

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

