# Openfoodfacts SDK context

require_relative '../utility/struct/voxgig_struct'
require_relative 'control'
require_relative 'operation'
require_relative 'spec'
require_relative 'result'
require_relative 'response'
require_relative 'error'
require_relative 'helpers'

class OpenfoodfactsContext
  attr_accessor :id, :out, :client, :utility, :ctrl, :meta, :config,
                :entopts, :options, :entity, :shared, :opmap,
                :data, :reqdata, :match, :reqmatch, :point,
                :spec, :result, :response, :op

  def initialize(ctxmap = {}, basectx = nil)
    ctxmap ||= {}
    @id = "C#{rand(10000000..99999999)}"
    @out = {}

    @client = OpenfoodfactsHelpers.get_ctx_prop(ctxmap, "client") || basectx&.client
    @utility = OpenfoodfactsHelpers.get_ctx_prop(ctxmap, "utility") || basectx&.utility

    @ctrl = OpenfoodfactsControl.new
    ctrl_raw = OpenfoodfactsHelpers.get_ctx_prop(ctxmap, "ctrl")
    if ctrl_raw.is_a?(Hash)
      @ctrl.throw_err = ctrl_raw["throw"] if ctrl_raw.key?("throw")
      @ctrl.explain = ctrl_raw["explain"] if ctrl_raw["explain"].is_a?(Hash)
    elsif basectx&.ctrl
      @ctrl = basectx.ctrl
    end

    m = OpenfoodfactsHelpers.get_ctx_prop(ctxmap, "meta")
    @meta = m.is_a?(Hash) ? m : (basectx&.meta || {})

    cfg = OpenfoodfactsHelpers.get_ctx_prop(ctxmap, "config")
    @config = cfg.is_a?(Hash) ? cfg : basectx&.config

    eo = OpenfoodfactsHelpers.get_ctx_prop(ctxmap, "entopts")
    @entopts = eo.is_a?(Hash) ? eo : basectx&.entopts

    o = OpenfoodfactsHelpers.get_ctx_prop(ctxmap, "options")
    @options = o.is_a?(Hash) ? o : basectx&.options

    e = OpenfoodfactsHelpers.get_ctx_prop(ctxmap, "entity")
    @entity = e || basectx&.entity

    s = OpenfoodfactsHelpers.get_ctx_prop(ctxmap, "shared")
    @shared = s.is_a?(Hash) ? s : basectx&.shared

    om = OpenfoodfactsHelpers.get_ctx_prop(ctxmap, "opmap")
    @opmap = om.is_a?(Hash) ? om : (basectx&.opmap || {})

    @data = OpenfoodfactsHelpers.to_map(OpenfoodfactsHelpers.get_ctx_prop(ctxmap, "data")) || {}
    @reqdata = OpenfoodfactsHelpers.to_map(OpenfoodfactsHelpers.get_ctx_prop(ctxmap, "reqdata")) || {}
    @match = OpenfoodfactsHelpers.to_map(OpenfoodfactsHelpers.get_ctx_prop(ctxmap, "match")) || {}
    @reqmatch = OpenfoodfactsHelpers.to_map(OpenfoodfactsHelpers.get_ctx_prop(ctxmap, "reqmatch")) || {}

    pt = OpenfoodfactsHelpers.get_ctx_prop(ctxmap, "point")
    @point = pt.is_a?(Hash) ? pt : basectx&.point

    sp = OpenfoodfactsHelpers.get_ctx_prop(ctxmap, "spec")
    @spec = sp.is_a?(OpenfoodfactsSpec) ? sp : basectx&.spec

    r = OpenfoodfactsHelpers.get_ctx_prop(ctxmap, "result")
    @result = r.is_a?(OpenfoodfactsResult) ? r : basectx&.result

    rp = OpenfoodfactsHelpers.get_ctx_prop(ctxmap, "response")
    @response = rp.is_a?(OpenfoodfactsResponse) ? rp : basectx&.response

    opname = OpenfoodfactsHelpers.get_ctx_prop(ctxmap, "opname") || ""
    @op = resolve_op(opname)
  end

  def resolve_op(opname)
    return @opmap[opname] if @opmap[opname]
    return OpenfoodfactsOperation.new({}) if opname.empty?

    entname = @entity&.respond_to?(:get_name) ? @entity.get_name : "_"
    opcfg = VoxgigStruct.getpath(@config, "entity.#{entname}.op.#{opname}")

    input = (opname == "update" || opname == "create") ? "data" : "match"

    points = []
    if opcfg.is_a?(Hash)
      t = VoxgigStruct.getprop(opcfg, "points")
      points = t if t.is_a?(Array)
    end

    op = OpenfoodfactsOperation.new({
      "entity" => entname,
      "name" => opname,
      "input" => input,
      "points" => points,
    })
    @opmap[opname] = op
    op
  end

  def make_error(code, msg)
    OpenfoodfactsError.new(code, msg, self)
  end
end
