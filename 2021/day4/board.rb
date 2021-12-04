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

  def self.parse_boards(data)
    i = 0
    b = data.slice(2, data.size).each_with_object([]) do |b, boards|
      boards[i] = [] if boards[i].nil?

      if b == "\n"
        boards[i] = Board.new(boards[i])
        i += 1
        next
      end

      boards[i].push(b.chomp.split(' ').map(&:to_i))
    end
    # eh?
    b[-1] = Board.new b.last
    b
  end

  private

  def check_line(row)
    (row.select { |num| @seen_nums.include?(num) }).size == 5
  end
end
