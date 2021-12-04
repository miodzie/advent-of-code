class BingoRoller
  attr_reader :numbers, :next_num

  def initialize(nums)
    @numbers = nums
    @next_num = 0
  end

  def draw(amount = 1)
    raise 'I UH UM' if @next_num + amount > @numbers.size

    nums = []
    amount.times do |_|
      nums.push(@numbers[@next_num])
      @next_num += 1
    end

    nums
  end

  def finished?
    @next_num == @numbers.size
  end
end
