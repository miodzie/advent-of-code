class Bingo
  attr_reader :boards, :winning_nums

  def initialize(boards, winning_nums)
    @boards = boards
    @winning_nums = winning_nums
  end

  def find_winner
    winner = nil
    while winner.nil?
      nums = winning_nums.next_num.zero? ? winning_nums.draw(5) : winning_nums.draw
      @boards.each do |b|
        winner = b if b.winner? nums
      end
    end

    winner
  end

  # The Easy Winners
  def all_winners
    winners = []
    boards = @boards.map(&:clone)
    until winning_nums.finished?
      nums = winning_nums.next_num.zero? ? winning_nums.draw(5) : winning_nums.draw
      boards.each_with_index do |b, i|
        next if b.winner?

        if b.winner?(nums)
          winners.push(b)
        end
      end
    end
    winners
  end
end
