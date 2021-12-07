raw = File.read('input')

positions = raw.chomp.split(',').map(&:to_i)

mean = positions.sum / positions.count

def cost_for_pos(positions, position)
  positions.sum do |pos|
    diff = (pos - position).abs
    sum = 0
    diff.times do |i|
      sum += i+1
    end
    sum
  end
end

costs = []
while mean != positions.size
  costs.append(cost_for_pos(positions, mean))
  mean +=1
end
p costs.min
