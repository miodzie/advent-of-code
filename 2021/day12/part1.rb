raw = File.readlines('input.in')

graph = {}

raw.each do |line|
  paths = line.chomp.split('-')
  graph[paths.first] = [] if graph[paths.first].nil?
  graph[paths.last] = [] if graph[paths.last].nil?
  graph[paths.first].push paths.last
  graph[paths.last].push  paths.first
end

start = ['start', ['start']]
ans = 0
queue = [start]
until queue.empty?
  node, seen = queue.pop
  if node == 'end'
    ans += 1
    next
  end
  neighbors = graph[node]
  neighbors.each do |n|
    next if seen.include?(n)

    new_seen = seen.clone
    # we can't go back to small caves
    new_seen.push(n) if n.downcase == n
    queue.push([n, new_seen])
  end
end

puts ans
