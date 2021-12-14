input = File.readlines('14.in')

polymer = input.shift.chomp
rules = {}
input.slice(1, input.length).map do |l|
  l = l.chomp.split(' -> ')
  rules[l[0]] = l[1]
end

count = Hash.new(0)
polymer.chars.each {|c| count[c] += 1 }

pairs = Hash.new(0)
polymer.chars.each_with_index do |c, i|
  next if polymer.chars[i + 1].nil?

  pairs[c + polymer.chars[i + 1]] += 1
end

40.times do |i|
  tmp = pairs.clone
  rules.each do |rule, insert|
    # pair doesn't exist
    next unless pairs.has_key?(rule)
    next if pairs[rule].zero?

    # remove the pairs
    removed_pairs = pairs[rule]
    tmp[rule] -= removed_pairs
    # for each removed_pair, add a count for the inserted char
    count[insert] += removed_pairs
    # prefix
    tmp[rule[0] + insert] += removed_pairs
    # suffix
    tmp[insert + rule[1]] += removed_pairs
  end
  puts count.values.max - count.values.min if i == 9
  pairs = tmp.clone
end

puts count.values.max - count.values.min
