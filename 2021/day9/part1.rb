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

      puts "Found low point! #{point} (#{y}, #{x})"
      lows.push(point)
    end
  end

  lows
end

lows = find_low_points(heightmap)

p lows.sum { |p| p + 1 }

# That's not the right answer; your answer is too high. If you're stuck, make sure you're using the full input data; there are also some general tips on the about page, or you can ask for hints on the subreddit. Please wait one minute before trying again.  [Return to Day 9]
# ( You guessed 1797.)
