package main

import (
	"fmt"
)

const WHITE string = "w"
const BLACK string = "b"

type piece struct {
	color    string
	position string
	name     string
}

type position struct {
	piece *piece
	name  string
}

type board struct {
	spaces [][]*position
}

func newBoard() *board {
	s := [][]*position{
		{
			newPosition(newRook(BLACK)),
			newPosition(newKnight(BLACK)),
			newPosition(newBishop(BLACK)),
			newPosition(newQueen(BLACK)),
			newPosition(newKing(BLACK)),
			newPosition(newBishop(BLACK)),
			newPosition(newKnight(BLACK)),
			newPosition(newRook(BLACK)),
		},
		{
			newPosition(newPawn(BLACK)),
			newPosition(newPawn(BLACK)),
			newPosition(newPawn(BLACK)),
			newPosition(newPawn(BLACK)),
			newPosition(newPawn(BLACK)),
			newPosition(newPawn(BLACK)),
			newPosition(newPawn(BLACK)),
			newPosition(newPawn(BLACK)),
		},
		{
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
		},
		{
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
		},
		{
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
		},
		{
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
			newEmptyPosition(),
		},
		{
			newPosition(newPawn(WHITE)),
			newPosition(newPawn(WHITE)),
			newPosition(newPawn(WHITE)),
			newPosition(newPawn(WHITE)),
			newPosition(newPawn(WHITE)),
			newPosition(newPawn(WHITE)),
			newPosition(newPawn(WHITE)),
			newPosition(newPawn(WHITE)),
		},
		{
			newPosition(newRook(WHITE)),
			newPosition(newKnight(WHITE)),
			newPosition(newBishop(WHITE)),
			newPosition(newQueen(WHITE)),
			newPosition(newKing(WHITE)),
			newPosition(newBishop(WHITE)),
			newPosition(newKnight(WHITE)),
			newPosition(newRook(WHITE)),
		},
	}
	board := board{spaces: s}
	return &board
}

func newPosition(piece *piece) *position {
	return &position{piece: piece}
}

func newEmptyPosition() *position {
	return newPosition(newNoPiece())
}

func newPiece(name string, color string) *piece {
	return &piece{name: name, color: color}
}

func newNoPiece() *piece {
	return &piece{name: "", color: ""}
}

func newPawn(color string) *piece {
	return newPiece("pawn", color)
}

func newRook(color string) *piece {
	return newPiece("rook", color)
}

func newBishop(color string) *piece {
	return newPiece("bishop", color)
}

func newKing(color string) *piece {
	return newPiece("king", color)
}

func newQueen(color string) *piece {
	return newPiece("queen", color)
}

func newKnight(color string) *piece {
	return newPiece("knight", color)
}

func renderBoard(board *board) {
	var p *piece
	fmt.Print("     a    b    c    d    e    f    g    h  \n")
	fmt.Print("   ----------------------------------------\n")
	row := 8
	for i := range board.spaces {
		fmt.Printf("%d |", row)
		for j := range board.spaces[i] {
			p = board.spaces[i][j].piece
			if p.name != "" {
				fmt.Printf(" %c", p.name[0])
				fmt.Print(p.color)
				fmt.Print(" ")
			} else {
				fmt.Print("  * ")
			}
			fmt.Print(" ")
		}
		fmt.Printf("| %d \n", row)
		row--
	}
	fmt.Print("   ----------------------------------------\n")
	fmt.Print("     a    b    c    d    e    f    g    h  \n")
}

func main() {
	myBoard := newBoard()
	renderBoard(myBoard)
}
