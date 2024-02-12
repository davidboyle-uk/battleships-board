package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/dbx123/battleships-board/pkg/ai"
	"github.com/dbx123/battleships-board/pkg/board"
	"github.com/dbx123/battleships-board/pkg/game"
	"github.com/dbx123/battleships-board/types"
)

const (
	ROW_NUMBER int = 1024
	COL_NUMBER int = 1024

	BLANK_SPACE string = " "
	PAUSE_MEX   string = ">>> press ENTER to go on..."
)

type shot struct {
	from  *types.Player
	to    *types.Player
	coord types.Coord
	rc    chan string
}

var (
	shots = make(chan shot)
)

func init() {
	go func() {
		for s := range shots {
			s.rc <- game.TakeShot(s.from, s.to, s.coord)
		}
	}()
}

func main() {
	newGame := game.Initialise(10)
	simulate(newGame)
}

func simulate(g types.Game) {
	for {
		winner, err := takeTurn(g)
		if err != nil {
			panic(err)
		}
		if winner != "" {
			fmt.Printf("winner %s\n", winner)
			break
		}
	}
}

func takeTurn(g types.Game) (string, error) {
	cleanScreen()
	shooter, opponent := determineWhosTurn(g)
	opponentBoardString := opponent.Board.ToString()
	theirBoardAsWeSeeIt, err := board.BoardFromString(opponentBoardString)
	if err != nil {
		panic(err)
	}
	move := ai.CalculateMove(theirBoardAsWeSeeIt)
	winner := shoot(shooter, opponent, move)
	if winner != "" {
		return winner, nil
	}
	fmt.Println(shooter.Name, " -> ", move)
	fmt.Println(opponent.Board.ToString())
	//consolePause(PAUSE_MEX)
	return "", nil
}

func shoot(from, to *types.Player, t types.Coord) string {
	res := make(chan string, 1)

	shots <- shot{
		from:  from,
		to:    to,
		coord: t,
		rc:    res,
	}
	r := <-res

	return r
}

func determineWhosTurn(g types.Game) (*types.Player, *types.Player) {
	p1 := g.Players[0]
	p2 := g.Players[1]

	fmt.Printf("p1 Moves %v Hits %v\n", len(p1.Moves), p1.Hits)
	fmt.Printf("p2 Moves %v Hits %v\n", len(p2.Moves), p2.Hits)

	if len(p1.Moves) == len(p2.Moves) {
		return p1, p2
	}

	return p2, p1
}

func cleanScreen() {
	r, c := 0, 0
	for r < ROW_NUMBER {
		for c < COL_NUMBER {
			fmt.Print(BLANK_SPACE)
			c++
		}
		fmt.Println(BLANK_SPACE)
		r++
	}
	fmt.Print("\033[0;0H")
}

func consolePause(m string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s\n", m)
	reader.ReadString('\n')
}

func prettyPrint(data interface{}) string {
	var p []byte
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(p)
}
