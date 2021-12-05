def graph_point(point, graph)
  graph[point.y] = [] if graph[point.y].nil?
  graph[point.y][point.x] = 0 if graph[point.y][point.x].nil?
  graph[point.y][point.x] += 1
end

def graph_line(pair, graph)
  dx = pair.first.x - pair.second.x
  dy = pair.first.y - pair.second.y
  graph_point(pair.first, graph)
  point = pair.first.clone

  cx = dx.negative? ? 1 : -1
  cy = dy.negative? ? 1 : -1
  while dx != 0 || dy != 0
    unless dx.zero?
      point.x += cx
      dx += cx
    end
    unless dy.zero?
      point.y += cy
      dy += cy
    end
    graph_point(point, graph)
  end
end

def print_graph(graph)
  max = graph.map { |x| x.nil? ? 0 : x.size }.max
  graph.each do |row|
    if row.nil?
      max.times { putc '.' }
      putc "\n"
      next
    end
    row.each do |i|
      if i.nil?
        putc '.'
      else
        putc i.to_s
      end
    end
    putc "\n"
  end
end
