raw = File.readlines('sample')

heightmap = Array.new([])
raw.each_with_index do |x, i|
  heightmap[i] = x.chomp.split('').map(&:to_i) if heightmap[i].nil?
end

Point = Struct.new :x, :y, :low

def find_low_points(heightmap, seen)
  # We've checked all the numbers, return the low points only.
  if seen.size == heightmap.flatten.size
    lows = seen.filter do |p| 
      if p.low
        puts "OK!"
      end
      p.low
    end
    return lows
  end

  # Get the point to check
  if seen.empty?
    point = Point.new 0, 0, false
  else
    point = seen.last
    # Find the point on the right first.
    if !heightmap[point.y][point.x + 1].nil?
      point.x += 1
    elsif !heightmap[point.y + 1][point.x].nil?
      # Next line
      point.y += 1
      point.x = 0
    else
      p 'I DONT KNOW WHAT TO DO'
      exit
    end
  end

  curP = heightmap[point.y][point.x]
  puts "Checking #{curP} nearby points"
  pp point

  nearbyLows = [
    [point.y - 1, point.x],
    [point.y + 1, point.x],
    [point.y, point.x - 1],
    [point.y, point.x + 1]
  ].filter do |p|
    next if heightmap[p[0]].nil? || p[0] < 0 || p[1] < 0

    near = heightmap[p[0]][p[1]]
    puts "Comparing #{near}"
    !near.nil? && near < curP
  end

  point.low = nearbyLows.size.zero?
  # if point.low
  #   puts 'WAS A LOW POINT!'
  # else
  #   puts "Not a low point."
  # end

  # gets
  seen.append(point)

  find_low_points(heightmap, seen)
end

p find_low_points(heightmap, [])
