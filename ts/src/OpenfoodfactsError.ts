
import { Context } from './Context'


class OpenfoodfactsError extends Error {

  isOpenfoodfactsError = true

  sdk = 'Openfoodfacts'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  OpenfoodfactsError
}

