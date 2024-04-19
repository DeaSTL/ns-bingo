package objects

import (
	"log"
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
  rand.Shuffle(len(options),func(i,j int){options[i], options[j] = options[j], options[i] })
  log.Printf("Options list: %+v", options)
  log.Printf("options length :%+v", len(options))
  optionIndex := 0
  for i := 0; i < 5; i++ {
    board[i] = []*Tile{{},{},{},{},{}}
    for j := 0; j < 5; j++ {
      if i == 2 && j == 2 {
        board[i][j].Value = "Actually"
        board[i][j].New(i,j)
        continue
      }

      selItem := options[optionIndex]

      optionIndex++

      board[i][j].Value = selItem
      board[i][j].New(i,j)
    }
  }
  b.Tiles = board
  b.ID = utils.GenID(32)
}
