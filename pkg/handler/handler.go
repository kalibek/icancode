package handler

import (
	"icancode/pkg/model"
	"log"
)

const (
	Host = "127.0.0.1:8080"
	User = "ud88nh3l1ldj1x09bge2"
	Code = "8112415840513744246"
)

func Handle(board model.Board) (string, error) {
	me := board.GetMe()
	log.Printf("Robot at %v", me)
	if board.IsMeAlive() {

		return board.Go(model.Left), nil

	}
	return board.Reset(), nil
}
