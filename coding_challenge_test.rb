gem 'minitest'
require 'pry'
require 'minitest/autorun'
require 'minitest/pride'
require_relative './coding_challenge'

class ObjectTest < Minitest::Test
  def test_it_exists
    
    template = TemplateClass.new
    assert_instance_of TemplateClass, template

  end
end
