package workerpool

// Progress is a struct contains progress info
type Progress struct {
	// Total is the count of jobs
	Total int64
	// Pending is the count of waiting jobs
	Pending int64
	// Executing is the count of executing jobs
	Executing int64
	// Completed is the count of completed jobs
	Completed int64
	// Success is the count of success jobs
	Success int64
	// Failed is the count of failed jobs
	Failed int64
}
