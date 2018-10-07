package chessboard

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
	result string
	moves  []Move
}

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
	PiecePawn   Piece = "P"
	PieceRook   Piece = "R"
	PieceKnight Piece = "N"
	PieceBishop Piece = "B"
	PieceQueen  Piece = "Q"
	PieceKing   Piece = "K"
)

type MoveType string

const (
	TypeMove    MoveType = "-"
	TypeCapture MoveType = "x"
)

type MoveResult string

const (
	ResultCheck     MoveResult = "+"
	ResultCheckmate MoveResult = "#"
)

const (
	TagEvent  = "Event"
	TagSite   = "Site"
	TagDate   = "Date"
	TagRound  = "Round"
	TagWhite  = "White"
	TagBlack  = "Black"
	TagResult = "Result"
)
