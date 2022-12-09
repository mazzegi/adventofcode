package day_23

import (
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2019/intcode"
	"github.com/mazzegi/adventofcode/errutil"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := part1MainFunc(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

type NIC struct {
	id   int
	comp *intcode.Computer
	in   *intcode.NonBlockingIntChannelReader
	out  *intcode.NonBlockingIntChannelWriter
}

func NewNIC(id int, prg []int) *NIC {
	in := intcode.NewNonBlockingIntChannelReader(-1)
	out := intcode.NewNonBlockingIntChannelWriter()
	comp := intcode.NewComputer(prg, in, out)
	nic := &NIC{
		id:   id,
		comp: comp,
		in:   in,
		out:  out,
	}
	return nic
}

// func (nic *NIC) run(ctx context.Context, nics []*NIC) {
// 	go func() {
// 		nic.in.Provide(nic.id)
// 		err := nic.comp.Exec()
// 		log("nic[%02d]: exec ended with err=%v", nic.id, err)
// 	}()

// 	for {
// 		select {
// 		case addr := <-nic.out.C:
// 			x := <-nic.out.C
// 			y := <-nic.out.C
// 			if addr < 0 || addr >= len(nics) {
// 				panic(fmt.Errorf("invalid addr %d", addr))
// 			}
// 			log("nic[%02d]: out: (to %02d, x=%d, y=%d)", nic.id, addr, x, y)
// 			nics[addr].in.Provide(x, y)
// 			//nics[addr].in.Provide(y)
// 			log("nic[%02d]: provided inputs", nic.id)
// 		case <-ctx.Done():
// 			log("nic[%02d]: context is done", nic.id)
// 			return
// 		case <-done:
// 			log("nic[%02d]: exec is done", nic.id)
// 			return
// 		}
// 	}
// }

func (nic *NIC) io(nics []*NIC) {
	for {
		vals, ok := nic.out.Read(3)
		if !ok {
			return
		}
		addr, x, y := vals[0], vals[1], vals[2]
		if addr < 0 || addr >= len(nics) {
			panic(fmt.Errorf("invalid addr %d", addr))
		}
		log("nic[%02d]: out: (to %02d, x=%d, y=%d)", nic.id, addr, x, y)
		nics[addr].in.Provide(x, y)
		log("nic[%02d]: provided inputs", nic.id)
	}
}

func part1MainFunc(prg []int) (int, error) {
	nics := make([]*NIC, 50)
	for id := 0; id < 50; id++ {
		nics[id] = NewNIC(id, prg)
	}

	for _, nic := range nics {
		go func(nic *NIC) {
			nic.in.Provide(nic.id)
			err := nic.comp.Exec()
			log("nic[%02d] exec ended with err=%v", nic.id, err)
		}(nic)
	}
	// for _, nic := range nics {
	// 	log("nic[%02d] provide id ...", nic.id)
	// 	nic.in.Provide(nic.id)
	// 	log("nic[%02d] provided id", nic.id)
	// }
	for {
		for _, nic := range nics {
			nic.io(nics)
		}
	}

	return 0, nil
}

// func part1MainFunc(prg []int) (int, error) {
// 	nics := make([]*NIC, 50)
// 	for id := 0; id < 50; id++ {
// 		nics[id] = NewNIC(id, prg)
// 	}
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
// 	for _, nic := range nics {
// 		go func(nic *NIC) {
// 			nic.run(ctx, nics)
// 			cancel()
// 		}(nic)
// 	}
// 	<-ctx.Done()
// 	log("done")

// 	return 0, nil
// }

func part2MainFunc(in []int) (int, error) {
	return 0, nil
}
