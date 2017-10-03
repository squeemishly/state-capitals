gem 'minitest'
require 'pry'
require 'minitest/autorun'
require_relative './coding_challenge'

class StateCapitalsTest < Minitest::Test
  attr_reader :sc,
              :states,
              :capitals

  def setup
    @sc = StateCapitals.new
    @states = {"Oregon" => "OR",
              "Alabama" => "AL",
              "New Jersey" => "NJ",
              "Colorado" => "CO"}
    @capitals = {"OR" => "Salem",
                "AL" => "Montgomery",
                "NJ" => "Trenton",
                "CO" => "Denver"}
  end

  def test_it_exists
    assert_instance_of StateCapitals, sc
  end

  def test_it_can_find_the_abbreviation_for_a_state
    assert_equal "OR", sc.find_abbreviation("Oregon")
  end

  def test_it_finds_the_capital_from_a_state_abbreviation
    assert_equal "Salem", sc.find_capital("OR")
  end

end
