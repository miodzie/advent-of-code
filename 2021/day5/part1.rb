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

hv_pairs = pairs.select do |pair|
  pair[0].x == pair[1].x || pair[0].y == pair[1].y
end

graph = []
hv_pairs.each {|pair| graph_line(pair, graph)}
# print_graph(graph)
p graph.flatten.compact.select {|x| x > 1}.size