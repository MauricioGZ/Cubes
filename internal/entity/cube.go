package entity

type Cube struct {
	ID      int64  `db:"id"`
	Name    string `db:"name"`
	Brand   string `db:"brand"`
	Shape   string `db:"shape"`
	OwnedBy int64  `db:"owned_by"`
}
