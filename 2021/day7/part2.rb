raw = File.read('input')

positions = raw.chomp.split(',').map(&:to_i)

mean = positions.sum / positions.count

def cost_for_pos(positions, position)
  positions.sum do |pos|
    (0..(pos - position).abs).sum
  end
end

costs = []
while mean != positions.size
  costs.append(cost_for_pos(positions, mean))
  break costs[-1] if costs.size > 2 && costs[-2] > costs[-1]
  mean +=1
end
p costs.min
