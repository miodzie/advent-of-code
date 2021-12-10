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

sum = 0
incomplete = []
lines.each do |l|
  invalid = get_first_invalid(l.chars)
  if invalid.nil?
    incomplete.push(l)
  else
    sum += points[invalid]
  end
end

puts "Part 1: #{sum}"
