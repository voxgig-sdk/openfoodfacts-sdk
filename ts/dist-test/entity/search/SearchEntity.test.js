"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const node_path_1 = __importDefault(require("node:path"));
const Fs = __importStar(require("node:fs"));
const node_test_1 = require("node:test");
const node_assert_1 = __importDefault(require("node:assert"));
const live_runner_1 = require("../../live-runner");
const live_entity_1 = require("../../live-entity");
const __1 = require("../../..");
const utility_1 = require("../../utility");
// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
(0, utility_1.loadEnvLocal)(__dirname + '/../../../.env.local');
(0, node_test_1.describe)('SearchEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when OPENFOODFACTS_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('OPENFOODFACTS_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.OpenfoodfactsSDK.test();
        const ent = testsdk.Search();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.OPENFOODFACTS_TEST_LIVE;
        for (const op of ['list']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'search.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "name": "additives_tags", "req": false, "short": "List of additives", "type": "`$ARRAY`", "index$": 0 }, { "active": true, "name": "allergens", "req": false, "short": "Allergens present in the product", "type": "`$STRING`", "index$": 1 }, { "active": true, "name": "brands", "req": false, "short": "Brands of the product", "type": "`$STRING`", "index$": 2 }, { "active": true, "name": "categories", "req": false, "short": "Categories the product belongs to", "type": "`$STRING`", "index$": 3 }, { "active": true, "name": "countries", "req": false, "short": "Countries where the product is sold", "type": "`$STRING`", "index$": 4 }, { "active": true, "name": "created_t", "req": false, "short": "Creation timestamp", "type": "`$INTEGER`", "index$": 5 }, { "active": true, "name": "ecoscore_grade", "req": false, "short": "Eco-Score grade for environmental impact (a, b, c, d, e)", "type": "`$STRING`", "index$": 6 }, { "active": true, "name": "ecoscore_score", "req": false, "short": "Eco-Score numerical score", "type": "`$INTEGER`", "index$": 7 }, { "active": true, "name": "generic_name", "req": false, "short": "Generic name of the product", "type": "`$STRING`", "index$": 8 }, { "active": true, "name": "image_front_url", "req": false, "short": "URL of the front image", "type": "`$STRING`", "index$": 9 }, { "active": true, "name": "image_ingredients_url", "req": false, "short": "URL of the ingredients image", "type": "`$STRING`", "index$": 10 }, { "active": true, "name": "image_nutrition_url", "req": false, "short": "URL of the nutrition facts image", "type": "`$STRING`", "index$": 11 }, { "active": true, "name": "image_url", "req": false, "short": "URL of the product's front image", "type": "`$STRING`", "index$": 12 }, { "active": true, "name": "ingredients_analysis_tags", "req": false, "short": "Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.)", "type": "`$ARRAY`", "index$": 13 }, { "active": true, "name": "ingredients_text", "req": false, "short": "List of ingredients as text", "type": "`$STRING`", "index$": 14 }, { "active": true, "name": "labels", "req": false, "short": "Labels associated with the product (e.g., Organic, Fair Trade)", "type": "`$STRING`", "index$": 15 }, { "active": true, "name": "last_modified_t", "req": false, "short": "Last modification timestamp", "type": "`$INTEGER`", "index$": 16 }, { "active": true, "name": "manufacturing_places", "req": false, "short": "Manufacturing or processing places", "type": "`$STRING`", "index$": 17 }, { "active": true, "name": "nova_group", "req": false, "short": "NOVA group for food processing level (1-4)", "type": "`$INTEGER`", "index$": 18 }, { "active": true, "name": "nutriments", "req": false, "short": "Nutritional information", "type": "`$OBJECT`", "index$": 19 }, { "active": true, "name": "nutriscore_grade", "req": false, "short": "Nutri-Score grade (a, b, c, d, e)", "type": "`$STRING`", "index$": 20 }, { "active": true, "name": "nutriscore_score", "req": false, "short": "Nutri-Score numerical score", "type": "`$INTEGER`", "index$": 21 }, { "active": true, "name": "packaging", "req": false, "short": "Packaging type", "type": "`$STRING`", "index$": 22 }, { "active": true, "name": "product_name", "req": false, "short": "Name of the product", "type": "`$STRING`", "index$": 23 }, { "active": true, "name": "quantity", "req": false, "short": "Quantity or volume of the product", "type": "`$STRING`", "index$": 24 }, { "active": true, "name": "stores", "req": false, "short": "Stores where the product is available", "type": "`$STRING`", "index$": 25 }, { "active": true, "name": "traces", "req": false, "short": "Traces of allergens", "type": "`$STRING`", "index$": 26 }], "name": "search", "op": { "list": { "input": "data", "name": "list", "points": [{ "active": true, "args": { "query": [{ "active": true, "kind": "query", "name": "brand", "orig": "brand", "reqd": false, "type": "`$STRING`", "index$": 0 }, { "active": true, "kind": "query", "name": "category", "orig": "category", "reqd": false, "type": "`$STRING`", "index$": 1 }, { "active": true, "example": true, "kind": "query", "name": "json", "orig": "json", "reqd": false, "type": "`$BOOLEAN`", "index$": 2 }, { "active": true, "kind": "query", "name": "label", "orig": "label", "reqd": false, "type": "`$STRING`", "index$": 3 }, { "active": true, "example": 1, "kind": "query", "name": "page", "orig": "page", "reqd": false, "type": "`$INTEGER`", "index$": 4 }, { "active": true, "example": 20, "kind": "query", "name": "page_size", "orig": "page_size", "reqd": false, "type": "`$INTEGER`", "index$": 5 }, { "active": true, "kind": "query", "name": "search_term", "orig": "search_term", "reqd": false, "type": "`$STRING`", "index$": 6 }] }, "contract": { "id": "GET /search", "json": "{\"operationId\":\"searchProducts\",\"parameters\":[{\"description\":\"Search terms for product name, brand, or other text fields\",\"in\":\"query\",\"name\":\"search_terms\",\"schema\":{\"type\":\"string\"}},{\"description\":\"Filter by product category\",\"in\":\"query\",\"name\":\"categories\",\"schema\":{\"type\":\"string\"}},{\"description\":\"Filter by brand name\",\"in\":\"query\",\"name\":\"brands\",\"schema\":{\"type\":\"string\"}},{\"description\":\"Filter by labels (e.g., organic, fair-trade)\",\"in\":\"query\",\"name\":\"labels\",\"schema\":{\"type\":\"string\"}},{\"description\":\"Page number for pagination\",\"in\":\"query\",\"name\":\"page\",\"schema\":{\"default\":1,\"minimum\":1,\"type\":\"integer\"}},{\"description\":\"Number of products per page\",\"in\":\"query\",\"name\":\"page_size\",\"schema\":{\"default\":20,\"maximum\":100,\"minimum\":1,\"type\":\"integer\"}},{\"description\":\"Return results in JSON format\",\"in\":\"query\",\"name\":\"json\",\"schema\":{\"default\":true,\"type\":\"boolean\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"count\":{\"description\":\"Total number of products matching the search\",\"type\":\"integer\"},\"page\":{\"description\":\"Current page number\",\"type\":\"integer\"},\"page_count\":{\"description\":\"Total number of pages\",\"type\":\"integer\"},\"page_size\":{\"description\":\"Number of products per page\",\"type\":\"integer\"},\"products\":{\"description\":\"Array of products matching the search criteria\",\"items\":{\"properties\":{\"additives_tags\":{\"description\":\"List of additives\",\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"allergens\":{\"description\":\"Allergens present in the product\",\"type\":\"string\"},\"brands\":{\"description\":\"Brands of the product\",\"type\":\"string\"},\"categories\":{\"description\":\"Categories the product belongs to\",\"type\":\"string\"},\"countries\":{\"description\":\"Countries where the product is sold\",\"type\":\"string\"},\"created_t\":{\"description\":\"Creation timestamp\",\"type\":\"integer\"},\"ecoscore_grade\":{\"description\":\"Eco-Score grade for environmental impact (a, b, c, d, e)\",\"enum\":[\"a\",\"b\",\"c\",\"d\",\"e\"],\"type\":\"string\"},\"ecoscore_score\":{\"description\":\"Eco-Score numerical score\",\"type\":\"integer\"},\"generic_name\":{\"description\":\"Generic name of the product\",\"type\":\"string\"},\"image_front_url\":{\"description\":\"URL of the front image\",\"type\":\"string\"},\"image_ingredients_url\":{\"description\":\"URL of the ingredients image\",\"type\":\"string\"},\"image_nutrition_url\":{\"description\":\"URL of the nutrition facts image\",\"type\":\"string\"},\"image_url\":{\"description\":\"URL of the product's front image\",\"type\":\"string\"},\"ingredients_analysis_tags\":{\"description\":\"Tags for ingredient analysis (vegan, vegetarian, palm oil, etc.)\",\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"ingredients_text\":{\"description\":\"List of ingredients as text\",\"type\":\"string\"},\"labels\":{\"description\":\"Labels associated with the product (e.g., Organic, Fair Trade)\",\"type\":\"string\"},\"last_modified_t\":{\"description\":\"Last modification timestamp\",\"type\":\"integer\"},\"manufacturing_places\":{\"description\":\"Manufacturing or processing places\",\"type\":\"string\"},\"nova_group\":{\"description\":\"NOVA group for food processing level (1-4)\",\"maximum\":4,\"minimum\":1,\"type\":\"integer\"},\"nutriments\":{\"description\":\"Nutritional information\",\"properties\":{\"carbohydrates\":{\"description\":\"Carbohydrates per 100g\",\"type\":\"number\"},\"energy-kcal\":{\"description\":\"Energy in kcal per 100g\",\"type\":\"number\"},\"energy-kj\":{\"description\":\"Energy in kJ per 100g\",\"type\":\"number\"},\"fat\":{\"description\":\"Fat content per 100g\",\"type\":\"number\"},\"fiber\":{\"description\":\"Dietary fiber per 100g\",\"type\":\"number\"},\"proteins\":{\"description\":\"Proteins per 100g\",\"type\":\"number\"},\"salt\":{\"description\":\"Salt content per 100g\",\"type\":\"number\"},\"saturated-fat\":{\"description\":\"Saturated fat content per 100g\",\"type\":\"number\"},\"sodium\":{\"description\":\"Sodium content per 100g\",\"type\":\"number\"},\"sugars\":{\"description\":\"Sugars per 100g\",\"type\":\"number\"}},\"type\":\"object\"},\"nutriscore_grade\":{\"description\":\"Nutri-Score grade (a, b, c, d, e)\",\"enum\":[\"a\",\"b\",\"c\",\"d\",\"e\"],\"type\":\"string\"},\"nutriscore_score\":{\"description\":\"Nutri-Score numerical score\",\"type\":\"integer\"},\"packaging\":{\"description\":\"Packaging type\",\"type\":\"string\"},\"product_name\":{\"description\":\"Name of the product\",\"type\":\"string\"},\"quantity\":{\"description\":\"Quantity or volume of the product\",\"type\":\"string\"},\"stores\":{\"description\":\"Stores where the product is available\",\"type\":\"string\"},\"traces\":{\"description\":\"Traces of allergens\",\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"}},\"type\":\"object\"}}},\"description\":\"Successful search response\"}},\"securitySchemes\":{\"UserAgent\":{\"description\":\"It is recommended to send a User-Agent header with your API calls to help identify your application\",\"in\":\"header\",\"name\":\"User-Agent\",\"type\":\"apiKey\"}},\"securitySource\":\"unspecified\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/search", "segments": [{ "lit": "search" }], "select": { "exist": ["brand", "category", "json", "label", "page", "page_size", "search_term"] }, "transform": { "req": "`reqdata`", "res": "`body.products`" }, "index$": 0 }], "key$": "list" } }, "relations": { "ancestors": [] }, "key$": "search", "name__orig": "search", "Name": "Search", "name_": "search", "name-": "search", "NAME": "SEARCH", "index$": 1 }, { "active": true, "entity": "search", "key$": "BasicSearchFlow", "kind": "basic", "name": "BasicSearchFlow", "param": {}, "step": [{ "active": true, "data": {}, "input": {}, "match": {}, "op": "list", "spec": [], "valid": [{ "apply": "ItemExists", "def": { "ref": "search_ref01" } }], "index$": 0 }] }, 'Search');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let search_ref01_data = Object.values(setup.data.existing.search)[0];
        // LIST
        const search_ref01_ent = client.Search();
        const search_ref01_match = {};
        const search_ref01_list = (await search_ref01_ent.list(search_ref01_match)).map((e) => e.data());
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/search/SearchTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.OpenfoodfactsSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['search01', 'search02', 'search03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'OPENFOODFACTS_TEST_SEARCH_ENTID': idmap,
        'OPENFOODFACTS_TEST_LIVE': 'FALSE',
        'OPENFOODFACTS_TEST_EXPLAIN': 'FALSE',
        'OPENFOODFACTS_APIKEY': '',
    });
    idmap = env['OPENFOODFACTS_TEST_SEARCH_ENTID'];
    const live = 'TRUE' === env.OPENFOODFACTS_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['OPENFOODFACTS_TEST_SEARCH_ENTID'];
        idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {};
        if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
            throw new Error('Live ENTID must be a JSON object');
        }
        client = new __1.OpenfoodfactsSDK(merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            (0, utility_1.liveClientOptions)(),
            {
                apikey: env.OPENFOODFACTS_APIKEY,
            },
            // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
            // last entry is undefined, and basicSetup is normally called with no
            // argument at all - so a bare 'extra' silently discarded the apikey
            // and server values above and handed the SDK undefined. Harmless
            // while there was nothing in that object; not harmless now.
            extra || {},
            { system: { fetch: transport.fetch } }
        ]));
    }
    const setup = {
        idmap,
        env,
        options,
        client,
        struct,
        data: entityData,
        explain: 'TRUE' === env.OPENFOODFACTS_TEST_EXPLAIN,
        live,
        transport,
        now: Date.now(),
    };
    return setup;
}
//# sourceMappingURL=SearchEntity.test.js.map