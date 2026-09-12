export interface Product {
    additives_tags?: any[];
    allergens?: string;
    brands?: string;
    categories?: string;
    countries?: string;
    created_t?: number;
    ecoscore_grade?: string;
    ecoscore_score?: number;
    generic_name?: string;
    id?: string;
    image_front_url?: string;
    image_ingredients_url?: string;
    image_nutrition_url?: string;
    image_url?: string;
    ingredients_analysis_tags?: any[];
    ingredients_text?: string;
    labels?: string;
    last_modified_t?: number;
    manufacturing_places?: string;
    nova_group?: number;
    nutriments?: Record<string, any>;
    nutriscore_grade?: string;
    nutriscore_score?: number;
    packaging?: string;
    product_name?: string;
    quantity?: string;
    stores?: string;
    traces?: string;
}
export interface ProductLoadMatch {
    id: string;
    $action?: string;
    [action: string]: any;
}
export interface Search {
    additives_tags?: any[];
    allergens?: string;
    brands?: string;
    categories?: string;
    countries?: string;
    created_t?: number;
    ecoscore_grade?: string;
    ecoscore_score?: number;
    generic_name?: string;
    image_front_url?: string;
    image_ingredients_url?: string;
    image_nutrition_url?: string;
    image_url?: string;
    ingredients_analysis_tags?: any[];
    ingredients_text?: string;
    labels?: string;
    last_modified_t?: number;
    manufacturing_places?: string;
    nova_group?: number;
    nutriments?: Record<string, any>;
    nutriscore_grade?: string;
    nutriscore_score?: number;
    packaging?: string;
    product_name?: string;
    quantity?: string;
    stores?: string;
    traces?: string;
}
export interface SearchListMatch {
    brand?: string;
    category?: string;
    json?: boolean;
    label?: string;
    page?: number;
    page_size?: number;
    search_term?: string;
}
