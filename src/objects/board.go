package objects

import (
	"log"
	"math/rand"
	"time"

	"jmhart.dev/ns-bingo/utils"
)


type Board struct {
  Tiles [][]*Tile
  ID string
  TimeoutTimer *time.Timer
  Username string
}

func (b *Board) GetWins() int {
  colCount := []int{0,0,0,0,0}
  rowCount := []int{0,0,0,0,0}
  diagonal := []int{0,0}
  wins := 0
  for i, row := range b.Tiles {
    for j, item := range row {
      if item.Selected {
        colCount[j]++
        rowCount[i]++
        if j == i {
          diagonal[0]++
        }
        if j + i == 4 {
          diagonal[1]++
        }
      }
    }
  }

  if diagonal[0] == 5 {
    wins++
  }

  if diagonal[1] == 5 {
    wins++
  }

  for _, count := range rowCount  {
    if count == 5 {
      wins++
    } 
  }

  for _, count := range colCount {
    if count == 5 {
      wins++
    } 
  }


  return wins
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
  b.TimeoutTimer = time.NewTimer(time.Minute * 10)
}
