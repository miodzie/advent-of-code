raw = File.read('input')

positions = raw.chomp.split(',').map(&:to_i)
p positions

mean = positions.sum / positions.count

def cost_for_pos(positions, position)
  positions.sum do |pos|
    (pos - position).abs
  end
end

p cost_for_pos(positions, mean)

costs = []
while mean != 0
  costs.append(cost_for_pos(positions, mean))
  mean -=1
end
p costs.min
