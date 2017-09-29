gem 'minitest'
require 'pry'
require 'minitest/autorun'
require_relative './coding_challenge'

class ObjectTest < Minitest::Test
  attr_reader :template

  def setup
    @template = TemplateClass.new
  end

  def test_it_exists
    assert_instance_of TemplateClass, template
  end
end
