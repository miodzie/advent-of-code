input = File.read('17.in').chomp
# https://github.com/mebeim/aoc/tree/master/2021#day-17---trick-shot

d = input.split(':').last.split(',')
xmin, xmax = d.first.slice(3, d.first.size).split('..').map(&:to_i)
ymin, ymax = d.last.slice(3, d.first.size).split('..').map(&:to_i)

puts ymin * (ymin + 1) / 2

total = 0

# for every reasonable (v0x, v0y)
for v0x in 1..xmax + 1
  for v0y in ymin..-ymin
    x = 0
    y = 0 
    vx = v0x
    vy = v0y

    # While we're not past the target (on either axis)
    while x <= xmax && y >= ymin
      # We're inside the target area
      if x >= xmin && y <= ymax
        total += 1
        break
      end

      x += vx
      y += vy
      vy -= 1

      vx -= 1 if vx > 0
    end

  end
end
puts total
