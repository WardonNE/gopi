package pipeline

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {
	str := "Hello, {name}"
	pipeline := NewPipeline[string, string]()
	finalStr := pipeline.Send(str).Through(
		AsPipe[string, string](func(passable string, next Callback[string, string]) string {
			passable = strings.ReplaceAll(passable, "{name}", "World")
			return next(passable)
		}),
		AsPipe[string, string](func(passable string, next Callback[string, string]) string {
			passable = passable + "\nLong time no see"
			return next(passable)
		}),
	).AppendThrough(AsPipe[string, string](func(passable string, next Callback[string, string]) string {
		passable = passable + "\nHow are you?"
		return next(passable)
	})).Then(func(passable string) string {
		return passable
	})
	assert.Equal(t, "Hello, World\nLong time no see\nHow are you?", finalStr)

	str = "Hello, {name}"
	pipeline = NewPipeline[string, string]()
	finalStr = pipeline.Send(str).ThroughCallbacks(
		func(passable string, next Callback[string, string]) string {
			passable = strings.ReplaceAll(passable, "{name}", "World")
			return next(passable)
		},
		func(passable string, next Callback[string, string]) string {
			passable = passable + "\nLong time no see"
			return next(passable)
		},
	).AppendThroughCallback(func(passable string, next Callback[string, string]) string {
		passable = passable + "\nHow are you?"
		return next(passable)
	}).Then(func(passable string) string {
		return passable
	})
	assert.Equal(t, "Hello, World\nLong time no see\nHow are you?", finalStr)
}
