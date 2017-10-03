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
end
