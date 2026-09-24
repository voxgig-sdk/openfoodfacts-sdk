"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.OpenfoodfactsError = void 0;
class OpenfoodfactsError extends Error {
    isOpenfoodfactsError = true;
    sdk = 'Openfoodfacts';
    code;
    ctx;
    status = -1;
    // `err.notFound` rather than a magic number at every call site.
    get notFound() { return 404 === this.status; }
    constructor(code, msg, ctx) {
        super(msg);
        this.code = code;
        this.ctx = ctx;
    }
}
exports.OpenfoodfactsError = OpenfoodfactsError;
//# sourceMappingURL=OpenfoodfactsError.js.map