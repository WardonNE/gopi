package event

type ProgressUpdated struct {
	Total     int64
	Pending   int64
	Executing int64
	Completed int64
}

func NewProgressUpdated(total, pending, executing, completed int64) *ProgressUpdated {
	return &ProgressUpdated{total, pending, executing, completed}
}

func (event *ProgressUpdated) Topic() string {
	return ProgressUpdatedTopic
}
