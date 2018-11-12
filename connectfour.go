package connectfour

import (
	"fmt"
)

/*These constants specify the concepts of "no player", the "red player"
  and the "yellow player".  Your code should use these constants in
  place of their hard-coded values.  Note that these values are not
  intended to be displayed to a user; they're simply intended as a
  universal way to distinguish which player we're talking about (e.g.,
  which player's turn it is, or which player (if any) has a piece in a
  particular cell on the board).*/

var NONE int = 0
var RED int = 1
var YELLOW int = 2

/*These constants specify the size of the game board (i.e., the number
  of rows and columns it has).  It should be possible to change these
  constants and re-run your program so that the board will have a
  different size; be sure that you're not hard-coding these values
  anywhere in your own code, because this is something we'll be
  attempting when we test your program.*/

var BoardColumns int = 7
var BoardRows int = 6

type GameState struct {
	turn  int
	board [][]int
}

type ColError struct {
	col     int
	message string
}

func (e *ColError) Error() string {
	return fmt.Sprintf("%d %s", e.col, e.message)
}

type GameOverError struct {
	message string
}

func (e *GameOverError) Error() string {
	return fmt.Sprintf("%s", e.message)
}

/*NewGame Returns a GameState representing a brand new game
  in which no moves have been made yet.*/
func (gs *GameState) NewGame() {
	gs.board = NewBoard()
	gs.turn = RED
}

/*NewBoard Creates a new game board.  Initially, a game board has the size
  BOARD_COLUMNS x BOARD_ROWS and is comprised only of integers with the
  value NONE'*/
func NewBoard() [][]int {
	var board [][]int

	for i := 0; i < BoardColumns; i++ {
		var row []int
		for j := 0; j < BoardRows; j++ {
			row = append(row, NONE)
		}

		board = append(board, row)
	}

	return board
}

func drop(gs *GameState, colNum int) error {
	errCol := RequiredValidColumnNumber(colNum)
	if errCol != nil {
		return errCol
	}

	errWin := RequiredGameNotOver(gs)
	if errWin != nil {
		return errWin
	}
}

func Winner(gs *GameState) int {
	winner := NONE

	for i := 0; i < BoardColumns; i++ {
		for j := 0; j < BoardRows; j++ {

		}
	}
}

func RequiredGameNotOver(gs *GameState) error {
	if Winner(gs) != NONE {
		return &GameOverError{"Game is over"}
	}
	return nil
}

func RequiredValidColumnNumber(col int) error {
	if !isValidColNumber(col) {
		return &ColError{col, "is invalid"}
	}
	return nil
}

func isValidColNumber(col int) bool {
	return 0 <= col && col < BoardColumns
}
