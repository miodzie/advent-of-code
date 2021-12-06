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

fish = simulate_days(fishies, 256)
p fish[8]
fishies.each_with_index do |amount, spawn_count|
  printf "#{spawn_count}: #{amount} \n"
end
# That's not the right answer; your answer is too low.
# 26984457539
# 26984457539
p fish.sum
