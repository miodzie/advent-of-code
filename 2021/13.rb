input = File.read('13.in')

G = []
folds = []
input.split("\n").each do |line|
  x, y = line.split(',')
  if y.nil?
    next if x.nil?

    fold = x.split(' ').dig(2).split('=')
    fold[1] = fold[1].to_i
    folds.push(fold)
  else
    G.push([x, y].map(&:to_i))
  end
end

p1 = false
folds.each do |fold|
  axis, num = fold
  G.each_with_index do |p, i|
    if axis == 'y'
      G[i][1] -= (p[1] - num) * 2 if p[1] > num
    elsif p[0] > num
      G[i][0] -= (p[0] - num) * 2
    end
  end
  unless p1
    puts G.uniq.size
    p1 = true
  end
end

G.uniq!.sort!
graph = []
G.each do |p|
  graph[p[1]] = [] if graph[p[1]].nil?
  graph[p[1]].push(p[0])
end

# y
(0..graph.length).each do |y|
  (0..graph[y].sort.last).each do |x|
    if graph[y].include?(x)
      putc '*'
    else
      putc ' '
    end
  end
  putc "\n"
end

G.each_with_index do |row, _y|
  row.each_with_index do |p, x|
  end
  # puts
end
