package bowling

import (
	"errors"
	"fmt"
)

const (
	ONGOING = iota
	OPEN
	SPARE
	STRIKE
)

type Frame struct {
	frameType   int
	PinsKnocked int
	rolls       int
	firstThrow  int
}

func (frame *Frame) roll(pins int) (err error) {
	frame.rolls += 1
	if frame.rolls == 1 {
		frame.firstThrow = pins
	}

	frame.PinsKnocked += pins

	if frame.rolls == 1 && frame.PinsKnocked == 10 {
		frame.frameType = STRIKE
	}

	if frame.rolls == 2 {
		if frame.PinsKnocked == 10 {
			frame.frameType = SPARE
		} else {
			frame.frameType = OPEN
		}
	}

	if frame.PinsKnocked > 10 {
		return errors.New("Pin count exceeds pins on the lane")
	}

	return
}

type Game struct {
	currentFrame Frame
	frames       []Frame
	rolls        int
}

func NewGame() *Game {
	return &Game{}
}

func (game *Game) Roll(pins int) (err error) {
	game.rolls++
	if pins < 0 {
		return errors.New("Cannot roll negativ count of pins")
	}

	if game.currentFrame.frameType == ONGOING {
		err := game.currentFrame.roll(pins)
		if err != nil {
			return err
		}
	} else {
		game.frames = append(game.frames, game.currentFrame)
		game.currentFrame = Frame{}
		err := game.currentFrame.roll(pins)
		if err != nil {
			return err
		}
	}

	gameOver := errors.New("Cannot roll after game is over")
	if len(game.frames) >= 10 {
		switch game.frames[9].frameType {
		case SPARE:
			if game.currentFrame.frameType == OPEN || game.currentFrame.frameType == SPARE {
				return gameOver
			}
		case STRIKE:
			if len(game.frames) > 10 && game.frames[10].frameType != STRIKE && game.currentFrame.frameType != STRIKE {
				return gameOver
			}
		default:
			return gameOver
		}
	}

	return
}

func (game *Game) Score() (int, error) {
	frames := append(game.frames, game.currentFrame)
	if !game.gameFinished(frames) {
		return 0, errors.New("Score cannot be taken until the end of the game")
	}
	score := 0
	for idx, frame := range frames {
		score += frame.PinsKnocked

		if idx < 9 {
			score += game.getNextThrows(idx, frame.frameType)
		}
	}

	return score, nil
}

func (game *Game) getNextThrows(idx int, frameType int) (result int) {
	if frameType == SPARE && len(game.frames) > idx {
		result += game.frames[idx+1].firstThrow
	}
	if frameType == STRIKE && len(game.frames) > idx {
		nextFrame := game.frames[idx+1]
		result += nextFrame.PinsKnocked

		if nextFrame.frameType == STRIKE && len(game.frames) > idx+1 {
			result += game.frames[idx+2].firstThrow
		}
	}

	return
}

func (game *Game) gameFinished(frames []Frame) bool {
	fmt.Printf("frames len %d\n", len(game.frames))
	if len(frames) < 10 {
		return false
	}

	switch frames[9].frameType {
	case SPARE:
		return len(frames) == 11
	case STRIKE:
		if len(frames) < 11 {
			return false
		}
		if frames[10].frameType == STRIKE {
			return len(frames) == 12
		}

		return len(frames) == 11
	default:
		return len(frames) == 10
	}
}
