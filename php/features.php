<?php
declare(strict_types=1);

// Openfoodfacts SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class OpenfoodfactsFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new OpenfoodfactsBaseFeature();
            case "test":
                return new OpenfoodfactsTestFeature();
            default:
                return new OpenfoodfactsBaseFeature();
        }
    }
}
