package unit

type Unit struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Number      int    `json:"number"`
	CourseID    int    `json:"courseID"`
	SequenceNum int    `json:"sequenceNumber"`
}
