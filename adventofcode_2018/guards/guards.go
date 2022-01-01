package guards

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type EventType string

const (
	BeginsShift EventType = "begins shift"
	FallsAsleep EventType = "falls asleep"
	WakesUp     EventType = "wakes up"
	EndsShift   EventType = "ends shift"
)

type Event struct {
	Timestamp time.Time
	Guard     int
	Type      EventType
}

func (e Event) String() string {
	return fmt.Sprintf("%q: %d: %q", e.Timestamp.Format("2006-01-02 15:04"), e.Guard, string(e.Type))
}

func ParseEvent(s string) (*Event, error) {
	if len(s) < 27 {
		return nil, errors.Errorf("invalid event %q", s)
	}
	ts := s[1:17]
	t, err := time.Parse("2006-01-02 15:04", ts)
	if err != nil {
		return nil, errors.Wrapf(err, "parse-time %q", ts)
	}
	e := &Event{
		Timestamp: t,
	}
	ty := strings.Trim(s[18:], " ")
	switch {
	case strings.HasSuffix(ty, string(BeginsShift)):
		gs := strings.Trim(strings.TrimPrefix(strings.TrimSuffix(ty, string(BeginsShift)), "Guard #"), " ")
		gid, err := strconv.ParseInt(gs, 10, 64)
		if err != nil {
			return nil, errors.Wrapf(err, "parse-guard-id %q", gs)
		}
		e.Guard = int(gid)
		e.Type = BeginsShift
	case strings.HasSuffix(ty, string(FallsAsleep)):
		e.Type = FallsAsleep
	case strings.HasSuffix(ty, string(WakesUp)):
		e.Type = WakesUp
	default:
		return nil, errors.Errorf("invalid event-type %q", ty)
	}
	return e, nil
}

func ParseEvents(r io.Reader) ([]*Event, error) {
	scanner := bufio.NewScanner(r)
	var es []*Event
	for scanner.Scan() {
		l := strings.Trim(scanner.Text(), " \r\n\t")
		if l == "" {
			continue
		}
		e, err := ParseEvent(l)
		if err != nil {
			return nil, err
		}
		es = append(es, e)
	}
	sort.SliceStable(es, func(i, j int) bool {
		return es[i].Timestamp.Before(es[j].Timestamp)
	})
	currGuard := -1
	for _, e := range es {
		if e.Type == BeginsShift {
			currGuard = e.Guard
		} else if currGuard < 0 {
			return nil, errors.Errorf("non shift event but current guard is zero")
		} else {
			e.Guard = currGuard
		}
		fmt.Printf("parsed: %s\n", e.String())
	}
	return es, nil
}
