require 'pry'

class StateCapitals
  attr_reader :states,
              :capitals

  def initialize
    @states = {"Oregon" => "OR",
              "Alabama" => "AL",
              "New Jersey" => "NJ",
              "Colorado" => "CO"}
    @capitals = {"OR" => "Salem",
                "AL" => "Montgomery",
                "NJ" => "Trenton",
                "CO" => "Denver"}
  end

  def find_abbreviation(state)
    states[state]
  end

  def find_capital(abbrev)
    capitals[abbrev]
  end

  def find_capital_from_state(state)
    if capitals[find_abbreviation(state)]
      capitals[find_abbreviation(state)]
    else
      "Unknown"
    end
  end

  def capital_to_abbreviation(city)
    capitals.key(city)
  end

  def capital_to_state(city)
    states.key(capital_to_abbreviation(city))
  end
end
