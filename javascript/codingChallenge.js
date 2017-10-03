const _ = require('lodash')

class StateCapitals {
  constructor() {
    this.states = {"Oregon": "OR",
              "Alabama": "AL",
              "New Jersey": "NJ",
              "Colorado": "CO"};

    this.capitals = {"OR": "Salem",
                "AL": "Montgomery",
                "NJ": "Trenton",
                "CO": "Denver"};
  }

  state_to_abbreviation(state) {
    return this.states[state]
  }

  state_to_capital(state) {
    if (this.capitals[this.state_to_abbreviation(state)]) {
      return this.capitals[this.state_to_abbreviation(state)]
    } else {
      return "Unknown"
    }
  }

  capital_to_abbreviation(city) {
    return _.findKey(this.capitals, _.partial(_.isEqual,city))
  }
}

module.exports = StateCapitals;
