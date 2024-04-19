package objects

import (
	"math/rand"

	"jmhart.dev/ns-bingo/utils"
)


type Board struct {
  Tiles [][]*Tile
  ID string
}

func (b *Board) New(options []string){
  board := [][]*Tile{
    {},{},{},{},{},
  }


  for i := 0; i < 5; i++ {
    board[i] = []*Tile{{},{},{},{},{}}
    for j := 0; j < 5; j++ {
      if i == 2 && j == 2 {
        board[i][j].Value = "Actually"
        board[i][j].New(i,j)
        continue
      }
      randIndex := rand.Intn(len(options))
      selItem := options[randIndex]
      options = append(options[:randIndex],options[randIndex+1:]...)


      board[i][j].Value = selItem
      board[i][j].New(i,j)
    }
  }
  b.Tiles = board
  b.ID = utils.GenID(32)
}
