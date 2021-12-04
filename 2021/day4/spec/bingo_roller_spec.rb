require 'bingo_roller'

RSpec.describe BingoRoller, '#next_num' do
  it 'it draws one by default' do
    # Arrange
    roller = BingoRoller.new Array(1..10)

    # Act
    roller.draw

    # Assert
    expect(roller.next_num).to eq 1
  end
end
