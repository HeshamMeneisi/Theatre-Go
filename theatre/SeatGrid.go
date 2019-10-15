package theatre

import (
	"strings"
	gn "theatre-mgr/generators"
)

type SeatGrid struct {
	currentRow, perRow int
	Rows               []*Row
	AllBooked          bool
}

func NewSeatGrid(nRows, perRow, rightStart int, rowIDGen, seatIDGen gn.IDGenerator) *SeatGrid {
	sg := new(SeatGrid)
	sg.Rows = make([]*Row, nRows)
	sg.AllBooked = false
	sg.currentRow = 0
	sg.perRow = perRow

	for i := range sg.Rows {
		seatIDGen.Reset()
		sg.Rows[i] = NewRow(perRow, rightStart, rowIDGen.Next(), seatIDGen)
	}
	return sg
}

func (sg *SeatGrid) MoveToNextRow() {
	// Back to front (0 => len)
	sg.currentRow += 1
	if sg.currentRow >= len(sg.Rows) {
		sg.AllBooked = true
	}
}

func (sg *SeatGrid) Book(nSeats int) []*Seat {
	var booked_seats = make([]*Seat, 0, nSeats)

	if sg.AllBooked {
		return nil
	}

	for i := 0; i < nSeats && !sg.AllBooked; i++ {
		var next = sg.Rows[sg.currentRow].BookNext()
		booked_seats = append(booked_seats, next)
		if sg.Rows[sg.currentRow].AllBooked {
			// Can set AllBooked and exit loop on next iteration (if last row)
			sg.MoveToNextRow()
		}
	}

	return booked_seats
}

func (sg *SeatGrid) ToString() string {
	var sb strings.Builder
	for _, row := range sg.Rows {
		sb.WriteString(row.ToString() + "\n")
	}
	return sb.String()
}
