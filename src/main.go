package main

import (
	"log"
	"os"
	"slices"
	"sort"
	"time"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/recover"
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

type GreeterForm struct{
  Username string `json:"username"`
}


func main(){

  var possibleTiles = []string{}

  tileFile,err := os.ReadFile("./tiles.txt")

  if err != nil {
    panic("Could not read ./tiles.txt")
  }
  tileBuffer := ""
  for _, char := range tileFile {
    tileBuffer += string(char)
    if char == byte('\n'){
      possibleTiles = append(possibleTiles,tileBuffer)
      tileBuffer = "" 
    }
  }


  boards := map[string]*objects.Board{}

  go func(boards *map[string]*objects.Board) {
    for{
      log.Printf("Attempting to prune unused boards")
      for index, board := range *boards {
        select{
        case <-board.TimeoutTimer.C:
          log.Printf("Deleting board %v", index)
          delete(*boards,index)
        default:
          continue
        }
      }

      time.Sleep(time.Minute * 5)
    }
    
  }(&boards)

  app := fiber.New()

  app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
    
    newBoard := objects.Board{}
    newBoard.New(possibleTiles)
    boards[newBoard.ID] = &newBoard

    log.Printf("Creating new board with id: %v", newBoard.ID)

		return Render(c, views.Greeter(newBoard,boards))
	})

  app.Post("/pulse/:id", func(c *fiber.Ctx) error {
    id := c.Params("id","")

    if id == "" {
      return fiber.NewError(404,"Game not found")
    }

    board,ok := boards[id]

    if !ok {
      return fiber.NewError(404,"Game not found")
    }
    
    board.TimeoutTimer.Reset(time.Minute * 10)

    return c.SendString("")
  })

  app.Get("/game/:id",func(c *fiber.Ctx)error{

    username := c.Query("username","")


    if len(username) < 3 {
       
      return fiber.NewError(400,"Username must be at least 4 characters, get fucked")
    }
    
    id := c.Params("id","")

    if id == "" {
      return fiber.NewError(404,"Game not found")
    }

    board,ok := boards[id]

    if !ok {
      return fiber.NewError(404,"Game not found")
    }

    board.Username = username
    return Render(c, views.Index(board.ID))
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

    if target == "board" {
      return Render(c,views.Board(board))
    }

    for _, row := range board.Tiles {
      for _, item := range row {
        if target == item.ID{
          item.Selected = true     
        }
      }
    }

    log.Printf("Selecting tile: %v on board: %v", target,id)

    return Render(c,views.Board(board))
  })

  app.Delete("/exit/:id", func(c *fiber.Ctx) error {
    
    id := c.Params("id","")

    if id == "" {
      return fiber.ErrNotFound
    }

    _,ok := boards[id]

    if !ok {
      return fiber.ErrNotFound
    }

    delete(boards,id)

    log.Printf("Deleting board: %+v", id)

    return c.SendString("")
  })
  app.Get("/players", func(c *fiber.Ctx) error {
    players := []*objects.Board{} 

    for _, player := range boards{
      players = append(players,player)
    }

    sort.Slice(players[:],func(i,j int) bool {
      return players[i].GetWins() < players[j].GetWins()
    })

    slices.Reverse(players)



    return Render(c,views.Players(players))
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

    wins := board.GetWins()

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
