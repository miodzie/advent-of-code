input = File.readlines('14.in')

polymer = input.shift.chomp
rules = {}
input.slice(1, input.length).map do |l|
  l = l.chomp.split(' -> ')
  rules[l[0]] = l[1]
end

10.times do
  ignore = []
  new = polymer.clone
  chars = new.chars
  polymer.chars.each_with_index do |c, i|
    break if chars[i + 1].nil?

    pair = c + chars[i + 1]
    next unless rules.has_key?(pair) && !ignore.include?(i + 1)

    rule = rules[pair]
    new.insert(i + 1 + ignore.size, rule)
    ignore.push(i + 1)
  end

  polymer = new.clone
end
count = {}
polymer.chars.each do |c|
  count[c] = 0 unless count.has_key?(c)
  count[c] += 1
end

pp polymer.size
pp count.values.max - count.values.min
