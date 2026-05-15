
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { OpenfoodfactsSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await OpenfoodfactsSDK.test()
    equal(null !== testsdk, true)
  })

})
