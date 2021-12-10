raw = File.readlines('input')

lines = raw.map(&:chomp)

pairs = [
  '()',
  '[]',
  '{}',
  '<>'
]
$closing_pairs = {
  '(' => ')',
  '[' => ']',
  '{' => '}',
  '<' => '>'
}

points = {
  ')' => 3,
  ']' => 57,
  '}' => 1197,
  '>' => 25_137
}

# start with the first found closing pair
# Expected ], but found } instead.
line = '{([(<{}[<>[]}>{[]{[(<()>'
chars = line.chars

# Keep removing pairs until there's none left?!?!

def get_first_invalid(chars)
  closing = []
  chars.each do |c|
    if $closing_pairs.keys.include?(c)
      closing.push($closing_pairs[c])
    else
      popped = closing.pop
      if popped != c
        # puts "Expected #{popped} but found #{c} instead."
        return c
      end
    end
  end
  nil
end

incomplete = lines.filter { |l| get_first_invalid(l.chars).nil? }

def get_missing(chars)
  closing = []
  chars.each do |c|
    if $closing_pairs.keys.include?(c)
      closing.push($closing_pairs[c])
    else
      closing.pop
    end
  end

  closing.reverse
end

closing_points = {
  ')' => 1,
  ']' => 2,
  '}' => 3,
  '>' => 4
}

scores = incomplete.map do |line|
  # puts "#{line} Missing: #{get_missing(line.chars).join}"
  missing = get_missing(line.chars)
  missing.inject(0) do |score, c|
    score *= 5
    score += closing_points[c]
  end
end

middle = (scores.length - 1) / 2
p middle
p scores.sort[middle]
