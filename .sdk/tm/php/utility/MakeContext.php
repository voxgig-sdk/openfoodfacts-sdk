<?php
declare(strict_types=1);

// Openfoodfacts SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class OpenfoodfactsMakeContext
{
    public static function call(array $ctxmap, ?OpenfoodfactsContext $basectx): OpenfoodfactsContext
    {
        return new OpenfoodfactsContext($ctxmap, $basectx);
    }
}
