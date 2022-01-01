package day_10

import (
	"adventofcode_2016/errutil"
	"adventofcode_2016/readutil"
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

func Part1() {
	fac, err := Parse(input)
	errutil.ExitOnErr(err)
	fac.Process()
}

func Part2() {
	fac, err := Parse(input)
	errutil.ExitOnErr(err)
	fac.Process()

	var cs []int
	cs = append(cs, fac.Output(0).chips...)
	cs = append(cs, fac.Output(1).chips...)
	cs = append(cs, fac.Output(2).chips...)
	p := 1
	for _, c := range cs {
		p *= c
	}
	fmt.Printf("part-2: chips %v: product: %d\n", cs, p)
}

func Parse(in string) (*Factory, error) {
	fac := NewFactory()
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "value "):
			var val int
			var botID int
			_, err := fmt.Sscanf(line, "value %d goes to bot %d", &val, &botID)
			if err != nil {
				return nil, errors.Wrapf(err, "scan %q", line)
			}
			bot := fac.Bot(botID)
			bot.chips = append(bot.chips, val)

		case strings.HasPrefix(line, "bot "):
			var botID int
			var lowRecType string
			var lowRecID int
			var highRecType string
			var highRecID int
			_, err := fmt.Sscanf(line, "bot %d gives low to %s %d and high to %s %d", &botID, &lowRecType, &lowRecID, &highRecType, &highRecID)
			if err != nil {
				return nil, errors.Wrapf(err, "scan %q", line)
			}
			bot := fac.Bot(botID)
			bot.giveLowType = lowRecType
			bot.giveLowID = lowRecID
			bot.giveHighType = highRecType
			bot.giveHighID = highRecID

			if bot.giveLowType == "output" {
				fac.Output(bot.giveLowID)
			} else if bot.giveLowType != "bot" {
				return nil, errors.Errorf("invalid give-type %q", bot.giveLowType)
			}

			if bot.giveHighType == "output" {
				fac.Output(bot.giveHighID)
			} else if bot.giveHighType != "bot" {
				return nil, errors.Errorf("invalid give-type %q", bot.giveHighType)
			}
		}
	}

	return fac, nil
}

type Bot struct {
	id           int
	chips        []int
	giveLowType  string
	giveLowID    int
	giveHighType string
	giveHighID   int
}

type Output struct {
	id    int
	chips []int
}

type Bots map[int]*Bot
type Outputs map[int]*Output

type Factory struct {
	bots    Bots
	outputs Outputs
}

func NewFactory() *Factory {
	return &Factory{
		bots:    Bots{},
		outputs: Outputs{},
	}
}

func (f *Factory) Bot(id int) *Bot {
	bot, ok := f.bots[id]
	if !ok {
		bot = &Bot{
			id: id,
		}
		f.bots[id] = bot
	}
	return bot
}

func (f *Factory) Output(id int) *Output {
	out, ok := f.outputs[id]
	if !ok {
		out = &Output{
			id: id,
		}
		f.outputs[id] = out
	}
	return out
}

type Event struct {
	botID         int
	giveLowType   string
	giveLowID     int
	giveLowValue  int
	giveHighType  string
	giveHighID    int
	giveHighValue int
}

func (e Event) String() string {
	return fmt.Sprintf("bot %d: %d to %s %d; %d to %s %d", e.botID,
		e.giveLowValue, e.giveLowType, e.giveLowID, e.giveHighValue, e.giveHighType, e.giveHighID)
}

func (f *Factory) Process() []Event {
	var es []Event
	for {
		processed := false
		for _, bot := range f.bots {
			//
			if len(bot.chips) == 2 {
				sort.Ints(bot.chips)

				e := Event{botID: bot.id}

				if bot.giveLowType == "output" {
					rec := f.Output(bot.giveLowID)
					rec.chips = append(rec.chips, bot.chips[0])
				} else {
					rec := f.Bot(bot.giveLowID)
					rec.chips = append(rec.chips, bot.chips[0])
				}
				e.giveLowValue = bot.chips[0]
				e.giveLowType = bot.giveLowType
				e.giveLowID = bot.giveLowID

				if bot.giveHighType == "output" {
					rec := f.Output(bot.giveHighID)
					rec.chips = append(rec.chips, bot.chips[1])
				} else {
					rec := f.Bot(bot.giveHighID)
					rec.chips = append(rec.chips, bot.chips[1])
				}
				e.giveHighValue = bot.chips[1]
				e.giveHighType = bot.giveHighType
				e.giveHighID = bot.giveHighID
				es = append(es, e)
				fmt.Printf("%s\n", e)

				bot.chips = []int{}
				processed = true
			}
		}

		if !processed {
			return es
		}
	}
}
