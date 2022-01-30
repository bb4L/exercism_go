package bowling

import (
	"errors"
)

type frameType int

const (
	ongoing frameType = iota
	open
	spare
	strike
)

type frame struct {
	frameType   frameType
	PinsKnocked int
	rolls       int
	firstThrow  int
}

func (framePtr *frame) roll(pins int) error {
	framePtr.rolls++
	if framePtr.rolls == 1 {
		framePtr.firstThrow = pins
	}

	framePtr.PinsKnocked += pins

	if framePtr.rolls == 1 && framePtr.PinsKnocked == 10 {
		framePtr.frameType = strike
	}

	if framePtr.rolls == 2 {
		if framePtr.PinsKnocked == 10 {
			framePtr.frameType = spare
		} else {
			framePtr.frameType = open
		}
	}

	if framePtr.PinsKnocked > 10 {
		return errors.New("pin count exceeds pins on the lane")
	}

	return nil
}

// Game represents a bowling game
type Game struct {
	currentFrame frame
	frames       []frame
	rolls        int
}

// NewGame creates a new game
func NewGame() *Game {
	return &Game{}
}

// Roll rolls a given number of pins
func (game *Game) Roll(pins int) error {
	game.rolls++
	if pins < 0 {
		return errors.New("cannot roll negativ count of pins")
	}

	if game.currentFrame.frameType != ongoing {
		game.frames = append(game.frames, game.currentFrame)
		game.currentFrame = frame{}
	}

	err := game.currentFrame.roll(pins)
	if err != nil {
		return err
	}

	gameOver := errors.New("cannot roll after game is over")
	if len(game.frames) >= 10 {
		switch game.frames[9].frameType {
		case spare:
			if game.currentFrame.frameType == open || game.currentFrame.frameType == spare {
				return gameOver
			}
		case strike:
			if len(game.frames) > 10 && game.frames[10].frameType != strike && game.currentFrame.frameType != strike {
				return gameOver
			}
		default:
			return gameOver
		}
	}

	return nil
}

// Score returns the score of a game
func (game *Game) Score() (int, error) {
	frames := append(game.frames, game.currentFrame)
	if !game.gameFinished(frames) {
		return 0, errors.New("score cannot be taken until the end of the game")
	}
	score := 0
	for idx, currentFrame := range frames {
		score += currentFrame.PinsKnocked

		if idx < 9 {
			score += game.getNextThrows(idx, currentFrame.frameType)
		}
	}

	return score, nil
}

func (game *Game) getNextThrows(idx int, currentFrameType frameType) (result int) {
	if currentFrameType == spare && len(game.frames) > idx {
		result += game.frames[idx+1].firstThrow
	}
	if currentFrameType == strike && len(game.frames) > idx {
		nextFrame := game.frames[idx+1]
		result += nextFrame.PinsKnocked

		if nextFrame.frameType == strike && len(game.frames) > idx+1 {
			result += game.frames[idx+2].firstThrow
		}
	}

	return result
}

func (game *Game) gameFinished(frames []frame) bool {
	if len(frames) < 10 {
		return false
	}

	switch frames[9].frameType {
	case spare:
		return len(frames) == 11
	case strike:
		if len(frames) < 11 {
			return false
		}
		if frames[10].frameType == strike {
			return len(frames) == 12
		}

		return len(frames) == 11
	default:
		return len(frames) == 10
	}
}
