import { OpenfoodfactsEntityBase } from '../OpenfoodfactsEntityBase';
import type { OpenfoodfactsSDK } from '../OpenfoodfactsSDK';
import type { Control } from '../types';
import type { Product, ProductLoadMatch } from '../OpenfoodfactsTypes';
declare class ProductEntity extends OpenfoodfactsEntityBase<Product> {
    constructor(client: OpenfoodfactsSDK, entopts: any);
    make(this: ProductEntity): ProductEntity;
    load(this: any, reqmatch?: ProductLoadMatch, ctrl?: Control): Promise<ProductEntity>;
}
export { ProductEntity };
