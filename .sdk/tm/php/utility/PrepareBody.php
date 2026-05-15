<?php
declare(strict_types=1);

// Openfoodfacts SDK utility: prepare_body

class OpenfoodfactsPrepareBody
{
    public static function call(OpenfoodfactsContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
