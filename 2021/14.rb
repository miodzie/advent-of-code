input = File.readlines('14.in')

polymer = input.shift.chomp
rules = {}
input.slice(1, input.length).map do |l|
  l = l.chomp.split(' -> ')
  rules[l[0]] = l[1]
end

count = {}
polymer.chars.each do |c|
  count[c] = 0 unless count.has_key?(c)
  count[c] += 1
end

pairs = {}
polymer.chars.each_with_index do |c, i|
  next if polymer.chars[i + 1].nil?

  pair = c + polymer.chars[i + 1]
  pairs[pair] = 0 unless pairs.has_key?(pair)
  pairs[pair] += 1
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
    count[insert] = 0 unless count.has_key?(insert)
    count[insert] += removed_pairs

    prefix = rule[0] + insert
    # New pair is added
    tmp[prefix] = 0 unless tmp.has_key?(prefix)
    tmp[prefix] += removed_pairs

    suffix = insert + rule[1]
    # New pair is added
    tmp[suffix] = 0 unless tmp.has_key?(suffix)
    tmp[suffix] += removed_pairs
  end
  puts count.values.max - count.values.min if i == 9
  pairs = tmp.clone
end

puts count.values.max - count.values.min
