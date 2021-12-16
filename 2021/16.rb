input = File.readlines('16.in')

parsed = input.pack('H*').unpack('B*').first
# pp parsed

# 10001010000000000100101010000000000110101000000000000010111101000111100010100000

# Every packet begins with a standard header: the first three bits encode the packet version,
# and the next three bits encode the packet type ID.

# Packets with type ID 4 represent a literal value. 

def dec(str)
  str.reverse.chars.map.with_index do |digit, i|
    digit.to_i * 2**i
  end.sum
end


def operate(type, packets)
  case type
  # sum
  when 0
    return packets.sum
  # product
  when 1
    return packets.inject(1) { |p, a| a * p }
  # minimum
  when 2
    return packets.min
  # maximum
  when 3
    return packets.max
  # greater-than
  when 5
    return packets.first > packets.last ? 1 : 0
  # less-than
  when 6
    return packets.last > packets.first ? 1 : 0 
  # equal to
  when 7
    return packets.first == packets.last ? 1 : 0
  end
end

$sum = 0
def parse_packet(str)
  return 0 if str.empty?

  version = dec(str.slice!(0, 3))
  $sum += version
  type = dec(str.slice!(0, 3))
  # puts "V: #{version} T: #{type}"

  # Literal value
  if type == 4
    slice = str.slice!(0, 5)
    groups = [slice]
    until slice[0] == '0'
      slice = str.slice!(0, 5)
      groups.push(slice)
    end
    bin =  groups.map { |x| 
      x.slice!(0)
      x
    }.join
    return dec(bin)
  end

  # Type must be an operator
  id = str.slice!(0).to_i

  packets = []
  if id.zero?
    # Length is a 15-bit number, representing the total bits of the subpackets
    length = dec(str.slice!(0, 15))
    sub_packet = str.slice!(0, length)
    until sub_packet.empty?
      # puts "sub_packet:"
      packets.push(parse_packet(sub_packet))
    end
  elsif id == 1
    # Length is a 11-bit number, representing the amount of subpackets
    n_sub_packets = dec(str.slice!(0, 11))
    until n_sub_packets.zero?
      packets.push(parse_packet(str))
      n_sub_packets -= 1
    end
  end

  return operate(type, packets)
end

p2 =  parse_packet(parsed)
puts $sum
puts p2
