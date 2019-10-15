package theatre

type Orientation int

const (
	Left Orientation = iota
	Right
)

type Seat struct {
	ID          string
	Booked      bool
	row         *Row
	Orientation Orientation
}

func NewSeat(id string, row *Row, ort Orientation) *Seat {
	s := new(Seat)
	s.ID = id
	s.row = row
	s.Orientation = ort
	return s
}

func (s *Seat) Book() {
	s.Booked = true
}

func (s *Seat) ToString() string {
	return s.row.ID + ":" + s.ID
}
