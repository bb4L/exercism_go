package bowling

import (
	"errors"
)

type FrameType int

const (
	ONGOING FrameType = iota
	OPEN
	SPARE
	STRIKE
)

type Frame struct {
	frameType   FrameType
	PinsKnocked int
	rolls       int
	firstThrow  int
}

func (frame *Frame) roll(pins int) error {
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
		return errors.New("pin count exceeds pins on the lane")
	}

	return nil
}

type Game struct {
	currentFrame Frame
	frames       []Frame
	rolls        int
}

func NewGame() *Game {
	return &Game{}
}

func (game *Game) Roll(pins int) error {
	game.rolls++
	if pins < 0 {
		return errors.New("cannot roll negativ count of pins")
	}

	if game.currentFrame.frameType != ONGOING {
		game.frames = append(game.frames, game.currentFrame)
		game.currentFrame = Frame{}
	}

	err := game.currentFrame.roll(pins)
	if err != nil {
		return err
	}

	gameOver := errors.New("cannot roll after game is over")
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

	return nil
}

func (game *Game) Score() (int, error) {
	frames := append(game.frames, game.currentFrame)
	if !game.gameFinished(frames) {
		return 0, errors.New("score cannot be taken until the end of the game")
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

func (game *Game) getNextThrows(idx int, frameType FrameType) (result int) {
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

	return result
}

func (game *Game) gameFinished(frames []Frame) bool {
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
