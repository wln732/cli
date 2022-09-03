package test

// Student
type Student struct {
	Name  string  // Name
	Age   int     // Age
	Score float32 /* Score */
}

/*
*
*获取Student.Name
 */
func (s *Student) GetName() string {
	// return s.Name
	/*
		1245
		//4541514
	*/
	return s.Name
}
