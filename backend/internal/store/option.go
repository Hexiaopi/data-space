package store

type Option interface {
	WithUserName(username string) Option
	WithId(id int64) Option
	WithState(state uint8) Option
	WithPage(pageNum, pageSize int) Option
}
