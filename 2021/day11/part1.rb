raw = File.readlines('input')
require 'colorize'

dumbos = raw.map {|line| line.chomp.split('').map(&:to_i)}

def print_dumbos(dumbos, step = 0)
  puts "Step #{step}" unless step.zero?
  dumbos.each do |row|
    puts row.map { |c| c.zero? ? c.to_s.yellow : c }.join
  end
  puts
end

print_dumbos(dumbos)

def add_adjacent(dumbos, y, x, r)
  nearby = [
    [y - 1, x], # top
    [y, x - 1], # left
    [y, x + 1], # right
    [y + 1, x], # bottom
    [y - 1, x - 1], # top left
    [y - 1, x + 1], # top right
    [y + 1, x - 1], # bottom left
    [y + 1, x + 1] # bottom right
  ]
  dumbos[y][x] = 0
  flashed = 1
  nearby.each do |pair|
    y = pair[0]
    x = pair[1]
    next if y < 0 || y >= dumbos.length
    next if x < 0 || x >= r.length

    # print_dumbos(dumbos)
    # gets

    dumbos[y][x] += 1 unless dumbos[y][x].zero? # already flashed
    flashed += add_adjacent(dumbos, y, x, r) if dumbos[y][x] > 9
  end
  flashed
end

def grow(dumbos, steps = 1)
  flashed = 0
  steps.times do
    dumbos.each_with_index do |r, y|
      r.each_with_index do |_d, x|
        dumbos[y][x] += 1
      end
    end
    dumbos.each_with_index do |r, y|
      r.each_with_index do |_d, x|
        flashed += add_adjacent(dumbos, y, x, r) if dumbos[y][x] > 9
      end
    end
  end

  p flashed

  dumbos
end

print_dumbos(grow(dumbos, 100), 1)
