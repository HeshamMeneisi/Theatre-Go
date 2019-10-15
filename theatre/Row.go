package theatre

import (
	m "math"
	"strings"
	gn "theatre-mgr/generators"
)

type Row struct {
	ID                                                      string
	nextCenterRight, nextCenterLeft, nextXRight, rightStart int
	mid                                                     float64
	bookedCount                                             int
	Seats                                                   []*Seat
	AllBooked                                               bool
}

func NewRow(nSeats, rightStart int, id string, seatIDGen gn.IDGenerator) *Row {
	r := new(Row)
	r.ID = id
	r.mid = float64(nSeats-1) / 2
	r.rightStart = rightStart
	r.nextXRight = rightStart
	r.nextCenterRight = int(m.Ceil(r.mid))
	r.nextCenterLeft = r.nextCenterRight - 1
	r.Seats = make([]*Seat, nSeats)
	for i := range r.Seats {
		var ort Orientation
		if i < rightStart {
			ort = Left
		} else {
			ort = Right
		}
		r.Seats[i] = NewSeat(seatIDGen.Next(), r, ort)
	}
	return r
}

func (r *Row) getNextCenterSeat() *Seat {
	var seat *Seat
	// Choose closest to center, unless taken (right priority)
	if int(m.Abs(float64(r.nextCenterLeft)-r.mid)) >=
		int(m.Abs(float64(r.nextCenterRight)-r.mid)) &&
		r.nextCenterRight < len(r.Seats) &&
		!r.Seats[r.nextCenterRight].Booked {
		seat = r.Seats[r.nextCenterRight]
		r.nextCenterRight += 1
	} else if r.nextCenterLeft >= 0 {
		seat = r.Seats[r.nextCenterLeft]
		r.nextCenterLeft -= 1
	}

	return seat
}

func (r *Row) BookNext() *Seat {
	if r.AllBooked {
		return nil
	}

	var seat *Seat

	// If left section is bigger than right
	if r.nextCenterRight < r.rightStart &&
		r.nextXRight < len(r.Seats) &&
		!r.Seats[r.nextXRight].Booked {
		// Must fill the whole right section first
		seat = r.Seats[r.nextXRight]
		r.nextXRight += 1
	} else if r.nextCenterLeft >= r.rightStart || r.nextXRight >= len(r.Seats) {
		// If left section is empty/non-existent, or right side is full/non-existent
		// Get closest to center regardless of section
		seat = r.getNextCenterSeat()
	} else if r.nextCenterRight < len(r.Seats) {
		// Otherwise, fill right
		seat = r.Seats[r.nextCenterRight]
		r.nextCenterRight += 1
	} else if r.nextCenterLeft >= 0 {
		// Then left
		seat = r.Seats[r.nextCenterLeft]
		r.nextCenterLeft -= 1
	}

	seat.Book()
	r.bookedCount += 1
	if r.bookedCount == len(r.Seats) {
		r.AllBooked = true
	}

	return seat
}

func (r *Row) ToString() string {
	var sb strings.Builder
	sb.WriteString(r.ID + "|\t")
	for _, seat := range r.Seats {
		if seat.Orientation == Left {
			sb.WriteString("<")
		}
		if seat.Booked {
			sb.WriteString("X")
		} else {
			sb.WriteString("O")
		}
		if seat.Orientation == Right {
			sb.WriteString(">")
		}
	}
	return sb.String()
}
