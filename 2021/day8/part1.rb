file = File.readlines('input')

outputs = file.map { |l| l.split('|').last.chomp.split(' ') }

digits = {
  1 => 'cf',
  4 => 'bcfd',
  7 => 'acf',
  8 => 'abcdefg'
}
sizes = digits.map { |_, x| x.size }

total = outputs.map do |segment|
  segment.select do |seg|
    sizes.include?(seg.size)
  end
end

pp total.flatten.size
