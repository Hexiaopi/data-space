package store

type Option interface {
	WithName(name string) Option
	WithId(id int64) Option
	WithState(state uint8) Option
	WithPage(pageNum, pageSize int) Option
	WithField(field string, value interface{}) Option
	WithLikeField(field string, value string) Option
	WithDepartmentId(departmentId int64) Option
}
