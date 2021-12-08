raw = File.read('input')

positions = raw.chomp.split(',').map(&:to_i)

range = positions.max - positions.min

def cost_for_pos(positions, position)
  positions.sum do |pos|
    (0..(pos - position).abs).sum
  end
end

costs = []
while range != positions.size
  costs.append(cost_for_pos(positions, range))
  return costs[-2] if costs.size > 2 && costs[-2] > costs[-1]
  rangee +=1
end
p costs.min
