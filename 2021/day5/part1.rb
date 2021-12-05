input = File.readlines('input')
require './lib/plane'

pairs = parse_point_pairs(input)
hv_pairs = pairs.select do |pair|
  pair[0].x == pair[1].x || pair[0].y == pair[1].y
end

plane = Plane.new
hv_pairs.each {|pair| plane.graph_line(pair)}

# plane.display
p plane.get_overlap
