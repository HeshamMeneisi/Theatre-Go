package tests

import (
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"testing"
	gn "theatre-mgr/generators"
	th "theatre-mgr/theatre"
	"strconv"
)

func Test_SeatGrid_Book_One(t *testing.T) {
	// Arrange
	n := 3
	m := 5
	s := th.NewSeatGrid(n, m, m/2, gn.NewAlphabeticalGenerator(),
		gn.NewNumericalGenerator(0))

	// Act
	s.Book(1)
	c := 0
	for _, s := range s.Rows[0].Seats {
		if s.Booked {
			c += 1
		}
	}
	// Assert
	assert.Equal(t, 1, c)
}

func Test_SeatGrid_Book_All(t *testing.T) {
	// Arrange
	n := 3
	m := 5
	s := th.NewSeatGrid(n, m, m/2, gn.NewAlphabeticalGenerator(), gn.NewNumericalGenerator(0))

	// Act
	s.Book(n * m)
	c := 0
	for _, row := range s.Rows {
		for _, seat := range row.Seats {
			if seat.Booked {
				c += 1
			}
		}
	}
	// Assert
	assert.Equal(t, n*m, c)
}

func Test_SeatGrid_Overbook(t *testing.T) {
	// Arrange
	n := 31
	m := 5
	s := th.NewSeatGrid(n, m, m/2, gn.NewAlphabeticalGenerator(), gn.NewNumericalGenerator(0))

	// Act
	var seats = s.Book(n*m + 100)

	// Assert
	assert.Equal(t, n*m, len(seats))
}

func Test_SeatGrid_Minimal_Book(t *testing.T) {
	// Arrange
	s := th.NewSeatGrid(2, 2, 1, gn.NewAlphabeticalGenerator(), gn.NewNumericalGenerator(0))

	// Act
	var seats = s.Book(3)
	sids := make([]string, len(seats))
	for i, seat := range seats {
		sids[i] = seat.ID
	}

	// Assert
	assert.Assert(t, is.DeepEqual([]string{"1", "0", "1"}, sids))
}

func Test_Row_Book_No_Left(t *testing.T) {
	// Arrange
	r := th.NewRow(5, 0, "T", gn.NewNumericalGenerator(0))

	// Act
	var seats = []*th.Seat{r.BookNext(), r.BookNext(), r.BookNext()}
	sids := make([]string, len(seats))
	for i, seat := range seats {
		sids[i] = seat.ID
	}

	// Assert
	assert.Assert(t, is.DeepEqual([]string{"2", "3", "1"}, sids))
}

func Test_Row_Book_No_Right(t *testing.T) {
	// Arrange
	r := th.NewRow(5, 5, "T", gn.NewNumericalGenerator(0))

	// Act
	var seats = []*th.Seat{r.BookNext(), r.BookNext(), r.BookNext()}
	sids := make([]string, len(seats))
	for i, seat := range seats {
		sids[i] = seat.ID
	}

	// Assert
	assert.Assert(t, is.DeepEqual([]string{"2", "3", "1"}, sids))
}

func Test_Row_Book_Small_Right(t *testing.T) {
	// Arrange
	r := th.NewRow(9, 7, "T", gn.NewNumericalGenerator(0))

	// Act
	var seats = []*th.Seat{r.BookNext(), r.BookNext(), r.BookNext()}
	sids := make([]string, len(seats))
	for i, seat := range seats {
		sids[i] = seat.ID
	}

	// Assert
	assert.Assert(t, is.DeepEqual([]string{"7", "8", "4"}, sids))
}

func Test_Row_Book_General_Case(t *testing.T) {
	// Arrange
	n := 42
	r := th.NewRow(n, n/2, "T", gn.NewNumericalGenerator(0))

	var sIDs = make([]string, n)

	for i := n / 2; i < n; i++ {
		sIDs[i-n/2] = strconv.Itoa(i)
	}

	for i := 0; i < n/2; i++ {
		sIDs[i+n/2] = strconv.Itoa(n/2 - i - 1)
	}

	// Act
	sids := make([]string, n)
	for i := 0; i < n; i++ {
		sids[i] = r.BookNext().ID
	}

	// Assert
	assert.Assert(t, is.DeepEqual(sIDs, sids))
}
