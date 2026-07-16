<?php
declare(strict_types=1);

// Openfoodfacts SDK base feature

class OpenfoodfactsBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(OpenfoodfactsContext $ctx, array $options): void {}
    public function PostConstruct(OpenfoodfactsContext $ctx): void {}
    public function PostConstructEntity(OpenfoodfactsContext $ctx): void {}
    public function SetData(OpenfoodfactsContext $ctx): void {}
    public function GetData(OpenfoodfactsContext $ctx): void {}
    public function GetMatch(OpenfoodfactsContext $ctx): void {}
    public function SetMatch(OpenfoodfactsContext $ctx): void {}
    public function PrePoint(OpenfoodfactsContext $ctx): void {}
    public function PreSpec(OpenfoodfactsContext $ctx): void {}
    public function PreRequest(OpenfoodfactsContext $ctx): void {}
    public function PreResponse(OpenfoodfactsContext $ctx): void {}
    public function PreResult(OpenfoodfactsContext $ctx): void {}
    public function PreDone(OpenfoodfactsContext $ctx): void {}
    public function PreUnexpected(OpenfoodfactsContext $ctx): void {}
}
