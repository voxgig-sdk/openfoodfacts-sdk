import { OpenfoodfactsEntityBase } from '../OpenfoodfactsEntityBase';
import type { OpenfoodfactsSDK } from '../OpenfoodfactsSDK';
import type { Control } from '../types';
import type { Search, SearchListMatch } from '../OpenfoodfactsTypes';
declare class SearchEntity extends OpenfoodfactsEntityBase<Search> {
    constructor(client: OpenfoodfactsSDK, entopts: any);
    make(this: SearchEntity): SearchEntity;
    list(this: any, reqmatch?: SearchListMatch, ctrl?: Control): Promise<SearchEntity[]>;
}
export { SearchEntity };
