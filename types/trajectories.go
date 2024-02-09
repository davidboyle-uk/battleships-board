package types

import "fmt"

type Trajectories map[Coord]Trajectory

type Trajectory map[Direction][]Coord

func (t Trajectory) Reverse() []Coord {
	fmt.Printf("%#v\n", t)
	return []Coord{}
}
