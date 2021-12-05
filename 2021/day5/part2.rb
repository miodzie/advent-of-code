input = File.readlines('input')
require './lib/plane'

pairs = parse_point_pairs(input)
plane = Plane.new
pairs.each { |pair| plane.graph_line(pair) }

# plane.display
p plane.get_overlap
