input = File.readlines('input')

init_fish = input[0].split(',').map(&:to_i)

# How many lanternfish would there be after 80 days?

fishies = Array.new(9, 0)
init_fish.each { |f| fishies[f] += 1 }

def simulate_days(fish, days = 0)
  return fish if days.zero?
  spawning = fish[0]
  (0..7).each do |i|
    fish[i] = fish[i+1]
  end
  fish[6] += spawning
  fish[8] = spawning

  return simulate_days(fish, days-1)
end

fish = simulate_days(fishies.clone, 80)
puts "Part1: #{fish.sum}"

fish = simulate_days(fishies.clone, 256)
# fishies.each_with_index do |amount, spawn_count|
#   printf "#{spawn_count}: #{amount} \n"
# end

puts "Part2: #{fish.sum}"
