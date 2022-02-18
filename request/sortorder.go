package request

type SortOrder byte

const (
	SortOrderAscending SortOrder = iota
	SortOrderDescending
)

func (s SortOrder) String() string {
	switch s {
	case SortOrderAscending:
		return "ASC"
	case SortOrderDescending:
		return "DESC"
	default:
		return ""
	}
}
