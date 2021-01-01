// https://blog.learngoprogramming.com/golang-const-type-enums-iota-bc4befd096d3
package main

type Weekday int

const (
	Sunday Weekday = iota + 1
- // iota : 1 ~ iota increased
	// some comments  //iota: 1 ~skip: comment
	// iota: 1 ~skip: empty line
Monday // iota: 2 ~Monday 3
)

type Timezone int

const (
	//iota: 0, EST: -5
	EST Timezone = -(5 + iota)

	// _ is the blank identifier
	// iota: 1
	_

	//iota: 2, MST: -7
	MST
)

// Resetting iota
const (
	Zero = iota // Zero = 0
	One         // One = 1
)

// iota reset: will be 0 again
const (
	Two = iota // Two = 0
)

// iota: reset
const Three = iota // Three = 0

// Using iota in the middle
const (
	One = 1
	Two = 2

	// Three = 2 + 1 => 3
	// iota in the middle
	Three = iota + 1

	// Four = 3 + 1 => 4
	Four
)

// Multiple iotas in a single line
const (
	// Active = 0, Moving= 0, Running=0
	Active, Moving, Running = iota, iota, iota

	// Passive= 1, Stopped=1, Stale=1
	Passive, Stopped, Stale
)

const (
	// Active= 0, Running= 100
	Active, Running = iota, iota + 100

	// Passive= 1, Stopped= 101
	Passive, Stopped

	// You can't declare like this.
	// The last expression will be repeated
	//CantDeclare

	// But, you can reset the last expression
	Reset = iota

	// You can use any other expression even without iota
	AnyOther = 10
)

const (
	// iota: 0, One: 1 (type: int64)
	One int64 = iota + 1
	// iota: 1, Two: 2 (type: int64)
	// Two will be declared as if:
	// Two int64 = iota + 1
	Two
	// iota: 2, Four: 4 (type: int32)
	Four int32 = iota + 2
	// iota: 3, Five: 5 (type: int32)
	// Five will be declared as if:
	// Five int32 = iota + 2
	Five
	// (type: int)
	Six = 6
	// (type: int)
	// Seven will be declared as if:
	// Seven = 6
	Seven
)

// Evens and odds
type Even bool

const (
	// 0 % 2 == 0 ==> Even(true)
	a = Even(iota%2 == 0)
	// 1 % 2 == 0 ==> Even(false)
	b
	// 2 % 2 == 0 ==> Even(true)
	c
	// 3 % 2 == 0 ==> Even(false)
	d
)

// Counting backward
const (
	max = 10
)
const (
	a = (max - iota) // 10
	b                // 9
	c                // 8
)

// Producing Alphabets
const (
	// string will convert the
	// expression into string.
	//
	// or, it'll assign character
	// codes.
	a = string(iota + 'a') // a
	b                      // b
	c                      // c
	d                      // d
	e                      // e
)

// iota + 1 trick
type Activity int

const (
	Sleeping = iota + 1
	Walking
	Running
)

func main() {
	var activity Activity
	// activity will be zero,
	// so it's not initialized
	activity = Sleeping
	// now you know that it's been
	// initialized
}

// Unknown state pattern
const (
	Unknown = iota
	Sleeping
	Walking
	Running
)
