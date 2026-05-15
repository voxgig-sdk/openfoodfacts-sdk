<?php
declare(strict_types=1);

// Openfoodfacts SDK utility registration

require_once __DIR__ . '/../core/UtilityType.php';
require_once __DIR__ . '/Clean.php';
require_once __DIR__ . '/Done.php';
require_once __DIR__ . '/MakeError.php';
require_once __DIR__ . '/FeatureAdd.php';
require_once __DIR__ . '/FeatureHook.php';
require_once __DIR__ . '/FeatureInit.php';
require_once __DIR__ . '/Fetcher.php';
require_once __DIR__ . '/MakeFetchDef.php';
require_once __DIR__ . '/MakeContext.php';
require_once __DIR__ . '/MakeOptions.php';
require_once __DIR__ . '/MakeRequest.php';
require_once __DIR__ . '/MakeResponse.php';
require_once __DIR__ . '/MakeResult.php';
require_once __DIR__ . '/MakePoint.php';
require_once __DIR__ . '/MakeSpec.php';
require_once __DIR__ . '/MakeUrl.php';
require_once __DIR__ . '/Param.php';
require_once __DIR__ . '/PrepareAuth.php';
require_once __DIR__ . '/PrepareBody.php';
require_once __DIR__ . '/PrepareHeaders.php';
require_once __DIR__ . '/PrepareMethod.php';
require_once __DIR__ . '/PrepareParams.php';
require_once __DIR__ . '/PreparePath.php';
require_once __DIR__ . '/PrepareQuery.php';
require_once __DIR__ . '/ResultBasic.php';
require_once __DIR__ . '/ResultBody.php';
require_once __DIR__ . '/ResultHeaders.php';
require_once __DIR__ . '/TransformRequest.php';
require_once __DIR__ . '/TransformResponse.php';

OpenfoodfactsUtility::setRegistrar(function (OpenfoodfactsUtility $u): void {
    $u->clean = [OpenfoodfactsClean::class, 'call'];
    $u->done = [OpenfoodfactsDone::class, 'call'];
    $u->make_error = [OpenfoodfactsMakeError::class, 'call'];
    $u->feature_add = [OpenfoodfactsFeatureAdd::class, 'call'];
    $u->feature_hook = [OpenfoodfactsFeatureHook::class, 'call'];
    $u->feature_init = [OpenfoodfactsFeatureInit::class, 'call'];
    $u->fetcher = [OpenfoodfactsFetcher::class, 'call'];
    $u->make_fetch_def = [OpenfoodfactsMakeFetchDef::class, 'call'];
    $u->make_context = [OpenfoodfactsMakeContext::class, 'call'];
    $u->make_options = [OpenfoodfactsMakeOptions::class, 'call'];
    $u->make_request = [OpenfoodfactsMakeRequest::class, 'call'];
    $u->make_response = [OpenfoodfactsMakeResponse::class, 'call'];
    $u->make_result = [OpenfoodfactsMakeResult::class, 'call'];
    $u->make_point = [OpenfoodfactsMakePoint::class, 'call'];
    $u->make_spec = [OpenfoodfactsMakeSpec::class, 'call'];
    $u->make_url = [OpenfoodfactsMakeUrl::class, 'call'];
    $u->param = [OpenfoodfactsParam::class, 'call'];
    $u->prepare_auth = [OpenfoodfactsPrepareAuth::class, 'call'];
    $u->prepare_body = [OpenfoodfactsPrepareBody::class, 'call'];
    $u->prepare_headers = [OpenfoodfactsPrepareHeaders::class, 'call'];
    $u->prepare_method = [OpenfoodfactsPrepareMethod::class, 'call'];
    $u->prepare_params = [OpenfoodfactsPrepareParams::class, 'call'];
    $u->prepare_path = [OpenfoodfactsPreparePath::class, 'call'];
    $u->prepare_query = [OpenfoodfactsPrepareQuery::class, 'call'];
    $u->result_basic = [OpenfoodfactsResultBasic::class, 'call'];
    $u->result_body = [OpenfoodfactsResultBody::class, 'call'];
    $u->result_headers = [OpenfoodfactsResultHeaders::class, 'call'];
    $u->transform_request = [OpenfoodfactsTransformRequest::class, 'call'];
    $u->transform_response = [OpenfoodfactsTransformResponse::class, 'call'];
});
