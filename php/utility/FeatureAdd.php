<?php
declare(strict_types=1);

// Openfoodfacts SDK utility: feature_add

class OpenfoodfactsFeatureAdd
{
    public static function call(OpenfoodfactsContext $ctx, mixed $f): void
    {
        $ctx->client->features[] = $f;
    }
}
