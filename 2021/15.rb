input = File.readlines('15.in')
require 'algorithms'

G = []
input.each_with_index {|r,y|
  G[y] = r.chomp.split('').map(&:to_i)
}

INFINITY = 200000000000
vis = G.clone.map { |x| Array.new(G[0].size, false) }
minddist = vis.clone.map { |x| Array.new(x.size, INFINITY) }

q = Containers::MinHeap.new
q.push([0,0,0])
until q.empty?
  risk, y, x = q.pop

  next if vis[y][x]

  if x == G[0].size - 1 && G.size - 1 == y
    puts risk
    exit 0
  end

  vis[y][x] = true

  [[x + 1, y], [x - 1, y], [x, y - 1], [x, y + 1]].each do |nx, ny|
    next unless G.size > ny && ny >= 0 && 0 <= nx && nx < G[0].size
    next if vis[ny][nx]
    newdist = G[ny][nx] + risk

    if newdist < minddist[ny][nx]
      minddist[ny][nx] = newdist
      q.push([newdist, ny, nx])
    end
  end

end

def print_g(nodes, highlights = [])
  nodes.each_with_index do |r, y|
    r.each_with_index do |p, x|
      if highlights.include?([y,x])
        putc "*" 
      else
        putc p.to_s
      end
    end
    putc "\n"
  end
end
