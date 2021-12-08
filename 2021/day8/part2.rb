file = File.readlines('input')

lines = file.map do |l|
  h = l.chomp.split('|').map { |x| x.split(' ') }
  { patterns: h[0], output: h[1] }
end

Rating = Struct.new :number, :rule

# two: [0, 0, 7, 4, [1, 4, 7]]
# [0, 7, 4, 12, 0]
$two_rating = [0, 0, 7, 4, 12].sort.freeze
# [4, [1, 4, 7], 0, [1, 4, 7], 7]
$three_rating = [4, 12, 0, 12, 7].sort.freeze
# three: [0, 7, 4, [1, 4, 7], 0]
$five_rating = [0, 7, 4, 12, 0].sort.freeze
# six: [7, 0, 4, 0, [1, 4, 7], 0]
$six_rating = [7, 0, 4, 0, 12, 0].sort.freeze

$two_rating = [0, 0, 7, 4, 12].sort.freeze
# [4, [1, 4, 7], 0, [1, 4, 7], 7]
$three_rating = [4, 12, 0, 12, 7].sort.freeze
# three: [0, 7, 4, [1, 4, 7], 0]
$five_rating = [0, 7, 4, 12, 4].sort.freeze
# six: [7, 4, 4, 0, [1, 4, 7], 0]
$six_rating = [7, 4, 4, 0, 12, 0].sort.freeze
# nine: [0,4,4,[1,4,7], [1,4,7], 7]
$nine_rating = [0, 4, 4, 12, 12, 7].sort.freeze
# zero: [0, [1,4,7], 0, 4, 7, [1,4,7]]
$zero_rating = [0, 12, 0, 4, 7, 12].sort.freeze

# 2 and 5 are duplicate ratings

$ratings = [
  Rating.new(2, $two_rating),
  Rating.new(3, $three_rating),
  Rating.new(5, $five_rating),
  Rating.new(6, $six_rating),
  Rating.new(9, $nine_rating),
  Rating.new(0, $zero_rating),
]

def has_rule?(pattern, key, rule)
  rating = []
  pattern.chars do |c|
    cur = []
    key.each do |k, v|
      cur.append(k) if v.include?(c)
    end
    rating.append(cur.sum)
    cur = []
  end
  rating.sort == rule
end

def get_output(line)
  # First we need to find the uniq digits to start
  key = {
    1 => 'cf',
    4 => 'bcfd',
    7 => 'acf'
    # 8 => 'abcdefg'
  }
  finished_key = key.clone
  # Update the key
  line[:patterns].map do |pattern|
    if pattern.size == 7
      finished_key[8] = pattern
      next
    end
    match = key.select { |k, v| v.size == pattern.size ? k : false }.first
    key[match[0]] = pattern unless match.nil?
    finished_key[match[0]] = pattern unless match.nil?
  end

  # Find the remaining keys
  line[:patterns].each do |pattern|
    $ratings.each do |rating|
      finished_key[rating.number] = pattern if rating.rule.size == pattern.size && has_rule?(pattern, key, rating.rule)
    end
  end

  # Find the output number
  final_key = {}
  finished_key.each { |k, v| final_key[v.chars.sort.join] = k}
  line[:output].inject('') do |output, line|
    output += final_key[line.chars.sort.join].to_s
  end.to_i
end

# a = lines[4]
# p a
# p get_output a

a = lines.sum do |line|
  get_output(line)
end
p a
