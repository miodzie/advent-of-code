raw = File.readlines('input.in')

graph = {}

raw.each do |line|
  paths = line.chomp.split('-')
  graph[paths.first] = [] if graph[paths.first].nil?
  graph[paths.last] = [] if graph[paths.last].nil?
  graph[paths.first].push paths.last
  graph[paths.last].push  paths.first
end

start = ['start', ['start'], false]
ans = 0
queue = [start]
until queue.empty?
  node, small, used = queue.pop
  if node == 'end'
    ans += 1
    next
  end
  neighbors = graph[node]
  neighbors.each do |n|
    if !small.include?(n)
      # we can't go back to small caves
      new_small = small.clone
      new_small.push(n) if n.downcase == n
      queue.push([n, new_small, used])
    elsif !used && !%w[start end].include?(n) && small.include?(n)
      queue.push([n, small, true])
    end
  end
end

puts ans
