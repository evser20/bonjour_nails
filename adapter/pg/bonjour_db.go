package pg

func NewBonjourDB(config DBConfig) (*DB, error) {
	return NewDB(config)
}
