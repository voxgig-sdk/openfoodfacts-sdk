<?php
declare(strict_types=1);

// Openfoodfacts SDK utility: result_body

class OpenfoodfactsResultBody
{
    public static function call(OpenfoodfactsContext $ctx): ?OpenfoodfactsResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
