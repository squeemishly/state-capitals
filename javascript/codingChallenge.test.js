const assert = require('chai').assert
const StateCapitals = require('./codingChallenge')

describe('StateCapitals()', () => {
  before( () => {
    sc = new StateCapitals
  })

  it('exists', () => {
    assert(sc)
  })

  it('can find the state abbreviation from a state name', () => {
    assert.equal("OR", sc.state_to_abbreviation("Oregon"))
  })

  it('can return the capital from a state name', () => {
    assert.equal('Salem', sc.state_to_capital('Oregon'))
  })

  it('returns Unknown if the capital cannot be found', () => {
    assert.equal('Unknown', sc.state_to_capital("Boudi"))
  })

  it('returns an abbreviated state when given the capital', () => {
    assert.equal('OR', sc.capital_to_abbreviation('Salem'))
  })
})
