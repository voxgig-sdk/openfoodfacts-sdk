import { OpenfoodfactsEntityBase } from '../OpenfoodfactsEntityBase';
import type { OpenfoodfactsSDK } from '../OpenfoodfactsSDK';
import type { Control } from '../types';
import type { Product, ProductLoadMatch, ProductListMatch } from '../OpenfoodfactsTypes';
declare class ProductEntity extends OpenfoodfactsEntityBase<Product> {
    constructor(client: OpenfoodfactsSDK, entopts: any);
    make(this: ProductEntity): ProductEntity;
    load(this: any, reqmatch?: ProductLoadMatch, ctrl?: Control): Promise<ProductEntity>;
    list(this: any, reqmatch?: ProductListMatch, ctrl?: Control): Promise<ProductEntity[]>;
}
export { ProductEntity };
