<?php
declare(strict_types=1);

// Openfoodfacts SDK utility: result_headers

class OpenfoodfactsResultHeaders
{
    public static function call(OpenfoodfactsContext $ctx): ?OpenfoodfactsResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
