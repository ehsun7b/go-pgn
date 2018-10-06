package board

import (
	"time"
)

type State struct {
	event  string
	site   string
	date   time.Time
	round  string
	white  string
	black  string
	result Result
	moves  []Move
}

type Result string

const (
	RESULT_WHITE_WINS Result = "1-0"
	RESULT_BLACK_WINS Result = "0-1"
	RESULT_DRAW       Result = "1/2-1/2"
)

type Move struct {
	move        NormalMove
	kingCastle  KingCastle
	queenCastle QueenCastle
	MoveResult  MoveResult
}

type QueenCastle struct {
}
type KingCastle struct{}

type NormalMove struct {
	number   int
	piece    Piece
	cell1    Cell
	moveType MoveType
	cell2    Cell
}

type Cell struct {
	file string
	rank int
}

type Piece string

const (
	PAWN   Piece = "P"
	ROOK   Piece = "R"
	KNIGHT Piece = "N"
	BISHOP Piece = "B"
	QUEEN  Piece = "Q"
	KING   Piece = "K"
)

type MoveType string

const (
	MOVE    MoveType = "-"
	CAPTURE MoveType = "x"
)

type MoveResult string

const (
	CHECK     MoveResult = "+"
	CHECKMATE MoveResult = "#"
)
