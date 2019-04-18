package process

import (
	"encoding/json"
	"icancode/pkg/handler"
	"icancode/pkg/model"
)

const PREFIX = "board="

func Process(message []byte) ([]byte, error) {

	var board model.Board
	err := json.Unmarshal(message[len(PREFIX):], &board)
	if err != nil {
		return nil, err
	}
	s, err := handler.Handle(board)
	return []byte(s), err
}
