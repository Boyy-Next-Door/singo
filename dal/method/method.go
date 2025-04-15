package method

import "gorm.io/gen"

type TestTable interface {
	// select * from @@table
	// {{where}}
	//		status in @status
	// {{end}}
	// order by created_at desc
	// limit @limit offset @offset
	GetDataByStatus(status []int, limit, offset int32) ([]*gen.T, error)
}
