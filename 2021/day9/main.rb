raw = File.readlines('input')

heightmap = Array.new([])
raw.each_with_index do |x, i|
  heightmap[i] = x.chomp.split('').map(&:to_i) if heightmap[i].nil?
end

def find_low_points(heightmap)
  lows = []
  heightmap.each_with_index do |row, y|
    row.each_with_index do |point, x|
      # check top, bottom, left right
      next if y - 1 >= 0 && heightmap[y - 1][x] <= point
      next if y + 1 < heightmap.size && heightmap[y + 1][x] <= point
      next if x - 1 >= 0 && heightmap[y][x - 1] <= point
      next if x + 1 < row.size && heightmap[y][x + 1] <= point

      lows.push([y, x, point])
    end
  end

  lows
end

lows = find_low_points(heightmap)
puts "Part1: #{lows.sum { |p| p[2] + 1 }}"

def find_basin(map, basin, y, x)
  return true if x < 0 || x >= map.first.length || y < 0 || y >= map.size

  return true if map[y][x] == 9

  basin[y] = [] if basin[y].nil?

  return true unless basin[y][x].nil?

  basin[y][x] = 1

  # puts "Checking (#{y}, #{x})"

  # bold = map.clone
  # bold[y][x] = '*'
  # print_map(bold)
  # gets

  find_basin(map, basin, y + 1, x) &&
    find_basin(map, basin, y - 1, x) &&
    find_basin(map, basin, y, x - 1) &&
    find_basin(map, basin, y, x + 1)
end

def print_map(map)
  map.each do |row|
    puts row.join
  end
  puts
end

basins = []
lows.each do |point|
  basin = []
  # puts "Finding basin: (#{point[0]}, #{point[1]}: #{point[2]})"
  find_basin(heightmap, basin, point[0], point[1])
  # pp basin
  basins.push(basin.flatten.compact.sum)
end

puts "Part 2: #{basins.sort.slice(-3, 3).inject(1) { |product, n| product * n }}"
