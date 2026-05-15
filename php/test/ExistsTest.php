<?php
declare(strict_types=1);

// Openfoodfacts SDK exists test

require_once __DIR__ . '/../openfoodfacts_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = OpenfoodfactsSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
