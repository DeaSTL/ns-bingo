package main

import (
	"log"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"jmhart.dev/ns-bingo/objects"
	"jmhart.dev/ns-bingo/views"
)


func stringIn(items []string,item string) bool{
  for _, cItem := range items {
    if cItem == item {
      return true
    }
  }
  return false
}



func getWins(board objects.Board) int {
  colCount := []int{0,0,0,0,0}
  rowCount := []int{0,0,0,0,0}
  diagonal := []int{0,0}
  wins := 0
  for i, row := range board.Tiles {
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


func main(){

  possibleTiles := []string{
    "Diesel deafens mid-discussion to take a customer call",
    "lvcky trolling germans",
    "israel/Gaza",
    "communism/ soviet union/ yugoslavia/ vietnam",
    "3m+ autstic monologue",
    "\"just use <technology>\"",
    "Nate shows gun collection on webcam",
    "lucas tells us how things actually literally work, actually",
    "Ukraine",
    "Roads party",
    "economics an/cap",
    "Rust",
    "Crypto is a scam",
    "shitting on focus bot",
    "editor war",
    "lvcky singing/ hot micing",
    "lvcky playing devils advocate",
    "Nate drops a hard R",
    "Geordi preaching ruby doctrine",
    "C++",
    "type theory",
    "ketsu tells us he's going to bed and then returns 20 minutes later",
    "Nate's discord disconnects",
    "Rex talks about manufacturing drugs",
    "Nate refutes being gay",
    "Geordi using light theme",
    "Ketsu complains about worthless javascript developers",
    "Ryan talks about how great helix is",
  }

  boards := map[string]objects.Board{}


  app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
    newBoard := objects.Board{}
    newBoard.New(possibleTiles)
    boards[newBoard.ID] = newBoard

		return Render(c, views.Index(newBoard.ID))
	})

  app.Get("/board/:id",func(c *fiber.Ctx)error{
    id := c.Params("id","")
    target := c.Get("HX-Target")

    if id == "" {
      return fiber.ErrNotFound
    }

    board,ok := boards[id]

    if !ok {
      return fiber.ErrNotFound
    }

    for _, row := range board.Tiles {
      for _, item := range row {
        if target == item.ID{
          item.Selected = true     
        }
      }
    }

    return Render(c,views.Board(board))
  })
  app.Get("/stats/:id", func(c *fiber.Ctx) error {
    id := c.Params("id","")

    if id == "" {
      return fiber.ErrNotFound
    }

    board,ok := boards[id]

    if !ok {
      return fiber.ErrNotFound
    }

    wins := getWins(board)

    if wins > 0 {
      log.Printf("Wins :%v",wins)
    }

    return Render(c,views.Stats(wins))
  })


	log.Fatal(app.Listen(":3000"))
}



func Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}
