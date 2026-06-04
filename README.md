# Openfoodfacts SDK

Crowdsourced food product database with ingredients, nutrition facts, Nutri-Score and Nova group classifications

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

## About OpenFoodFacts API

[Open Food Facts](https://world.openfoodfacts.org) is a collaborative, non-profit food products database built by a global community of contributors. Anyone can scan a barcode, photograph a label, and add or correct product information, which is then made freely available as open data.

What you get from the API:

- Product lookups by barcode, returning ingredients, allergens, categories, brands, labels, packaging, and origin.
- Nutrition facts (energy, fats, sugars, salt, etc.) and computed scores such as **Nutri-Score**, **NOVA group**, and **Green-Score / Eco-Score**.
- Structured search across the catalogue with filters on tags, nutrients, and other fields.
- Bulk exports (CSV, JSONL, MongoDB dumps) are available separately for large-scale analysis.

Operational notes: read endpoints are unauthenticated but require a descriptive `User-Agent` header in the form `AppName/Version (ContactEmail)`. Public rate limits are roughly 100 product reads and 10 searches per minute per IP (subject to change — see the docs). Writes need an Open Food Facts account. CORS is enabled. A staging environment is available at `world.openfoodfacts.net` for testing without polluting production data.

## Try it

**TypeScript**
```bash
npm install openfoodfacts
```

**Python**
```bash
pip install openfoodfacts-sdk
```

**PHP**
```bash
composer require voxgig/openfoodfacts-sdk
```

**Golang**
```bash
go get github.com/voxgig-sdk/openfoodfacts-sdk/go
```

**Ruby**
```bash
gem install openfoodfacts-sdk
```

**Lua**
```bash
luarocks install openfoodfacts-sdk
```

## 30-second quickstart

### TypeScript

```ts
import { OpenfoodfactsSDK } from 'openfoodfacts'

const client = new OpenfoodfactsSDK({})

```

See the [TypeScript README](ts/README.md) for the
full guide, or scroll down for the same example in other languages.

## What's in the box

| Surface | Use it for | Path |
| --- | --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | App integration | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | Scripts, CI, ops, one-off API calls | `go-cli/` |
| **MCP server** | AI agents (Claude, Cursor, Cline) | `go-mcp/` |

## Use it from an AI agent (MCP)

The generated MCP server exposes every operation in this SDK as an
[MCP](https://modelcontextprotocol.io) tool that Claude, Cursor or Cline
can call directly. Build and register it:

```bash
cd go-mcp && go build -o openfoodfacts-mcp .
```

Then add it to your agent's MCP config (Claude Desktop, Cursor, etc.):

```json
{
  "mcpServers": {
    "openfoodfacts": {
      "command": "/abs/path/to/openfoodfacts-mcp"
    }
  }
}
```

## Entities

The API exposes 2 entities:

| Entity | Description | API path |
| --- | --- | --- |
| **Product** | A food product identified by its barcode, with ingredients, nutrition facts, allergens, categories and computed scores. Typical paths: `GET /api/v2/product/{barcode}` and `GET /api/v2/product/{barcode}.json`. | `/product/{barcode}.json` |
| **Search** | Structured, filter-based search over the product catalogue, returning paginated product summaries. Path: `GET /api/v2/search` with tag and field query parameters. | `/search` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
from openfoodfacts_sdk import OpenfoodfactsSDK

client = OpenfoodfactsSDK({})


# Load a specific product
product, err = client.Product(None).load(
    {"id": "example_id"}, None
)
```

### PHP

```php
<?php
require_once 'openfoodfacts_sdk.php';

$client = new OpenfoodfactsSDK([]);


// Load a specific product
[$product, $err] = $client->Product(null)->load(
    ["id" => "example_id"], null
);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/openfoodfacts-sdk/go"

client := sdk.NewOpenfoodfactsSDK(map[string]any{})

```

### Ruby

```ruby
require_relative "Openfoodfacts_sdk"

client = OpenfoodfactsSDK.new({})


# Load a specific product
product, err = client.Product(nil).load(
  { "id" => "example_id" }, nil
)
```

### Lua

```lua
local sdk = require("openfoodfacts_sdk")

local client = sdk.new({})


-- Load a specific product
local product, err = client:Product(nil):load(
  { id = "example_id" }, nil
)
```

## Unit testing in offline mode

Every SDK ships a test mode that swaps the HTTP transport for an
in-memory mock, so unit tests run offline.

### TypeScript

```ts
const client = OpenfoodfactsSDK.test()
const result = await client.Product().load({ id: 'test01' })
// result.ok === true, result.data contains mock data
```

### Python

```python
client = OpenfoodfactsSDK.test(None, None)
result, err = client.Product(None).load(
    {"id": "test01"}, None
)
```

### PHP

```php
$client = OpenfoodfactsSDK::test(null, null);
[$result, $err] = $client->Product(null)->load(
    ["id" => "test01"], null
);
```

### Golang

```go
client := sdk.TestSDK(nil, nil)
result, err := client.Product(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = OpenfoodfactsSDK.test(nil, nil)
result, err = client.Product(nil).load(
  { "id" => "test01" }, nil
)
```

### Lua

```lua
local client = sdk.test(nil, nil)
local result, err = client:Product(nil):load(
  { id = "test01" }, nil
)
```

## How it works

Every SDK call runs the same five-stage pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

A feature hook fires at each stage (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), so features can inspect or modify the pipeline without
forking the SDK.

### Features

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

Pass custom features via the `extend` option at construction time.

### Direct and Prepare

For endpoints the entity model doesn't cover, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`,
`headers`, and `body`. See the [How-to guides](#how-to-guides) below.

## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
console.log(result.data)
```

**Python:**
```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
```

**PHP:**
```php
[$result, $err] = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);
```

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
```

**Ruby:**
```ruby
result, err = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})
```

**Lua:**
```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
```

## Per-language documentation

- [TypeScript](ts/README.md)
- [Python](py/README.md)
- [PHP](php/README.md)
- [Golang](go/README.md)
- [Ruby](rb/README.md)
- [Lua](lua/README.md)

## Using the OpenFoodFacts API

- Upstream: [https://world.openfoodfacts.org](https://world.openfoodfacts.org)
- API docs: [https://openfoodfacts.github.io/openfoodfacts-server/api/](https://openfoodfacts.github.io/openfoodfacts-server/api/)

- Database is published under the **Open Database License (ODbL)** — you may reuse and redistribute the data, including for commercial use, provided you credit Open Food Facts and share derived databases under the same terms.
- Individual database contents are licensed under the **Database Contents License**.
- Product photos are released under **Creative Commons Attribution-ShareAlike**; some embedded packaging artwork may remain under third-party copyright.
- Data is volunteer-contributed and accuracy is not guaranteed; verify critical values before relying on them.

---

Generated from the OpenFoodFacts API OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).
