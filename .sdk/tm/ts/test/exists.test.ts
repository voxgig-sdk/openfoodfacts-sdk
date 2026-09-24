
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { OpenfoodfactsSDK } from '..'


describe('exists', async () => {

  test('test-mode', () => {
    const testsdk = OpenfoodfactsSDK.test()
    equal(testsdk instanceof OpenfoodfactsSDK, true,
      'OpenfoodfactsSDK.test() must return a client synchronously')
  })

})
