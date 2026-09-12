import { Context } from './Context';
declare class OpenfoodfactsError extends Error {
    isOpenfoodfactsError: boolean;
    sdk: string;
    code: string;
    ctx: Context;
    status: number;
    get notFound(): boolean;
    constructor(code: string, msg: string, ctx: Context);
}
export { OpenfoodfactsError };
