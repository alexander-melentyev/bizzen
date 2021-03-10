package query

// Pagination -.
type Pagination struct {
	Limit  uint64 `form:"limit"`
	Offset uint64 `form:"offset"`
}
