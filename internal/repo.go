package delta;

type Repository struct {
	Name		string
	Remote		string
	LocalPath	string
}

//TODO: impl
func (r Repository) GetDiff() (localAhead int, remoteAhead int) {
	return 0, 0
}
