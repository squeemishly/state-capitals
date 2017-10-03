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
end
