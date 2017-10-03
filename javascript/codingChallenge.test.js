const assert = require('chai').assert
const StateCapitals = require('./codingChallenge')

describe('StateCapitals()', () => {
  before( () => {
    sc = new StateCapitals
  })

  it('exists', () => {
    assert(sc)
  })
})
