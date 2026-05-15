# Openfoodfacts SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

OpenfoodfactsUtility.registrar = ->(u) {
  u.clean = OpenfoodfactsUtilities::Clean
  u.done = OpenfoodfactsUtilities::Done
  u.make_error = OpenfoodfactsUtilities::MakeError
  u.feature_add = OpenfoodfactsUtilities::FeatureAdd
  u.feature_hook = OpenfoodfactsUtilities::FeatureHook
  u.feature_init = OpenfoodfactsUtilities::FeatureInit
  u.fetcher = OpenfoodfactsUtilities::Fetcher
  u.make_fetch_def = OpenfoodfactsUtilities::MakeFetchDef
  u.make_context = OpenfoodfactsUtilities::MakeContext
  u.make_options = OpenfoodfactsUtilities::MakeOptions
  u.make_request = OpenfoodfactsUtilities::MakeRequest
  u.make_response = OpenfoodfactsUtilities::MakeResponse
  u.make_result = OpenfoodfactsUtilities::MakeResult
  u.make_point = OpenfoodfactsUtilities::MakePoint
  u.make_spec = OpenfoodfactsUtilities::MakeSpec
  u.make_url = OpenfoodfactsUtilities::MakeUrl
  u.param = OpenfoodfactsUtilities::Param
  u.prepare_auth = OpenfoodfactsUtilities::PrepareAuth
  u.prepare_body = OpenfoodfactsUtilities::PrepareBody
  u.prepare_headers = OpenfoodfactsUtilities::PrepareHeaders
  u.prepare_method = OpenfoodfactsUtilities::PrepareMethod
  u.prepare_params = OpenfoodfactsUtilities::PrepareParams
  u.prepare_path = OpenfoodfactsUtilities::PreparePath
  u.prepare_query = OpenfoodfactsUtilities::PrepareQuery
  u.result_basic = OpenfoodfactsUtilities::ResultBasic
  u.result_body = OpenfoodfactsUtilities::ResultBody
  u.result_headers = OpenfoodfactsUtilities::ResultHeaders
  u.transform_request = OpenfoodfactsUtilities::TransformRequest
  u.transform_response = OpenfoodfactsUtilities::TransformResponse
}
