input = File.readlines('input')
require 'ostruct'
require './helpers'

Point = Struct.new :x, :y
Pair = Struct.new :first, :second

pairs = input.map do |line|
  points = line.chomp.split('->').map do |x|
    p = x.split(',')
    point = Point.new p[0].to_i, p[1].to_i
  end
  Pair.new points[0], points[1]
end

graph = []
pairs.each {|pair| graph_line(pair, graph)}

# print_graph(graph)

p graph.flatten.compact.select {|x| x > 1}.size
