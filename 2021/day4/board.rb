class Board
  def initialize(board)
    @board = board
    @seen_nums = []
  end

  def final_score
    # sum all unmarked numbers
    unmarked = @board.flatten.reject { |n| @seen_nums.include? n }
    unmarked.sum * @seen_nums.last
  end

  def winner?(nums = [])
    @seen_nums = (@seen_nums + nums).uniq
    # check rows
    @board.each do |row|
      return true if check_line(row)
    end

    # check columns
    (0..@board[0].size - 1).each do |c|
      col = @board.map { |x| x[c] }
      return true if check_line(col)
    end

    false
  end

  def print
    p @board
  end

  private

  def check_line(row)
    (row.select { |num| @seen_nums.include?(num) }).size == 5
  end
end
