package main

// This file can be a nice home for your Battlesnake logic and related helper functions.
//
// We have started this for you, with a function to help remove the 'neck' direction
// from the list of possible moves!

import (
	"log"
	"math/rand"
)

// This function is called when you register your Battlesnake on play.battlesnake.com
// See https://docs.battlesnake.com/guides/getting-started#step-4-register-your-battlesnake
// It controls your Battlesnake appearance and author permissions.
// For customization options, see https://docs.battlesnake.com/references/personalization
// TIP: If you open your Battlesnake URL in browser you should see this data.
func info() BattlesnakeInfoResponse {
	log.Println("INFO")
	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "mpapale",
		Color:      "#888888", // TODO: Personalize
		Head:       "default", // TODO: Personalize
		Tail:       "default", // TODO: Personalize
	}
}

// This function is called everytime your Battlesnake is entered into a game.
// The provided GameState contains information about the game that's about to be played.
// It's purely for informational purposes, you don't have to make any decisions here.
func start(state GameState) {
	log.Printf("%s START\n", state.Game.ID)
}

// This function is called when a game your Battlesnake was in has ended.
// It's purely for informational purposes, you don't have to make any decisions here.
func end(state GameState) {
	log.Printf("%s END\n\n", state.Game.ID)
}

type direction string

const (
  up direction = "up"
  down direction = "down"
  left direction = "left"
  right direction = "right"
)

type directions map[direction]bool

// neighborDirection returns the Direction in which the
// target Coord is a neighbor of the source Coord.
// If the Coords are not neighbors, it returns nil.
func neighborDirection(source Coord, target Coord) *direction{
  dx := source.X - target.X
  dy := source.Y - target.Y

  var dir direction
  switch {
  case dx == -1 && dy == 0:
    dir = right
  case dx == 1 && dy == 0:
    dir = left
  case dx == 0 && dy == -1:
    dir = up
  case dx == 0 && dy == 1:
    dir = down
  }
  
  return &dir
}
// openDirections returns the directions that are unblocked
// for the given coordinate, given the set of blocked coordinates
func openDirections(pos Coord, blocked []Coord) directions{
  dirs := directions{
    up: true,
    down: true,
    left: true,
    right: true,
  }

  for _, block := range blocked {
    dir := neighborDirection(pos, block)
    if dir == nil {
      continue
    }
    dirs[*dir] = false
  }

  return dirs
}

// This function is called on every turn of a game. Use the provided GameState to decide
// where to move -- valid moves are "up", "down", "left", or "right".
// We've provided some code and comments to get you started.
func move(state GameState) BattlesnakeMoveResponse {
  myHead := state.You.Body[0]
  var blocked []Coord
  // My body is blocked, but we elide the Head
  blocked = append(blocked, state.You.Body[1:]...)
  // The walls are blocked:
  // Left wall:
  for i := 0; i < state.Board.Height; i++ {
    blocked = append(blocked, Coord{X: -1, Y: i})
  }
  // Right wall:
  for i := 0; i < state.Board.Height; i++ {
    blocked = append(blocked, Coord{X: state.Board.Width, Y: i})
  }
  // Top wall:
  for i := 0; i < state.Board.Width; i++ {
    blocked = append(blocked, Coord{X: i, Y: state.Board.Height})
  }
  // Bottom wall:
  for i := 0; i < state.Board.Width; i++ {
    blocked = append(blocked, Coord{X: i, Y: -1})
  }
  // Others
  for _, snake := range state.Board.Snakes {
    blocked = append(blocked, snake.Body...)
  }

  openDirs := openDirections(myHead, blocked)

  log.Println("myHead: ", myHead.X, myHead.Y)
  log.Println("board dims: ", state.Board.Width, state.Board.Height)
  log.Println("blocked:", blocked)

	// TODO: Step 4 - Find food.
	// Use information in GameState to seek out and find food.

	// Finally, choose a move from the available safe moves.
	// TODO: Step 5 - Select a move to make based on strategy, rather than random.
	var nextMove string

	safeMoves := []string{}
  for k, _ := range openDirs {
    if openDirs[k] {
      safeMoves = append(safeMoves, string(k))
    }
  }


  log.Println("safe moves:", safeMoves)
	if len(safeMoves) == 0 {
		nextMove = "down"
		log.Printf("%s MOVE %d: No safe moves detected! Moving %s\n", state.Game.ID, state.Turn, nextMove)
	} else {
		nextMove = safeMoves[rand.Intn(len(safeMoves))]
		log.Printf("%s MOVE %d: %s\n", state.Game.ID, state.Turn, nextMove)
	}
	return BattlesnakeMoveResponse{
		Move: nextMove,
	}
}
