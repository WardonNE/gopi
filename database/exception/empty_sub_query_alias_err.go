package exception

// EmptySubQueryAliasErr empty alias error
type EmptySubQueryAliasErr struct {
}

func (e *EmptySubQueryAliasErr) Error() string {
	return "An alias is required when table is an instance of *gorm.DB or *builder.Builder or builder.Callback"
}

// NewEmptySubQueryAliasErr creates an [EmptySubQueryAliasErr]
func NewEmptySubQueryAliasErr() *EmptySubQueryAliasErr {
	return new(EmptySubQueryAliasErr)
}

// ThrowEmptySubQueryAliasErr throw an [EmptySubQueryAliasErr]
func ThrowEmptySubQueryAliasErr() {
	panic(new(EmptySubQueryAliasErr))
}
