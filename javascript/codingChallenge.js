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
    return this.capitals[this.state_to_abbreviation(state)]
  }
}

module.exports = StateCapitals;
