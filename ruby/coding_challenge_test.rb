gem 'minitest'
require 'pry'
require 'minitest/autorun'
require_relative './coding_challenge'

class StateCapitalsTest < Minitest::Test
  attr_reader :sc

  def setup
    @sc = StateCapitals.new
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

  def test_it_can_find_a_capital_for_a_state
    assert_equal "Salem", sc.find_capital_from_state("Oregon")
  end

  def test_it_returns_unknown_for_a_state_not_in_the_list
    assert_equal "Unknown", sc.find_capital_from_state("Boudi")
  end

  def test_it_can_find_a_state_given_the_capital
    assert_equal "CO", sc.capital_to_abbreviation("Denver")
  end

  def test_it_can_return_the_full_state_name_from_the_capital
    assert_equal "Colorado", sc.capital_to_state("Denver")
  end
end
