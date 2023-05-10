package dun

type Sort uint8

const (
	_        Sort = iota
	SortAsc       // Ascending order
	SortDesc      // Descending order
)

// Mysql sort keyword
// https://dev.mysql.com/doc/refman/8.0/en/select.html
func (s Sort) Mysql() (sort string) {

	if s == SortDesc {
		return "DESC"
	}

	return "ASC"
}

// Mongo sort keyword
// https://www.mongodb.com/docs/upcoming/reference/operator/aggregation/sort/#mongodb-pipeline-pipe.-sort
func (s Sort) Mongo() int {

	if s == SortDesc {
		return -1 // Descending order
	}

	return 1 // Ascending order
}
