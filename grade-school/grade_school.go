package school

import "sort"

//Grade lists students in a class
type Grade struct {
	class    int
	students []string
}

//School is a grade school of students
type School map[int][]string

/*New constructs a new School.*/
func New() *School {
	return &School{}
}

/*Add puts a student in a grade.*/
func (s School) Add(student string, grade int) {
	g, ok := s[grade]
	if !ok {
		g = []string{}
	}
	g = append(g, student)
	s[grade] = g
}

/*Enrollment gets the class of students from all grades*/
func (s School) Enrollment() []Grade {
	var enrollment []Grade
	for grade, students := range s {
		sort.Strings(students)
		enrollment = append(enrollment, Grade{grade, students})
	}
	sort.Sort(byLevel(enrollment))
	return enrollment
}

/*Grade gets the students in class.*/
func (s School) Grade(grade int) []string {
	g, ok := s[grade]
	if !ok {
		return []string{}
	}
	return g
}

//byLevel sorts a List of grades by class
type byLevel []Grade

func (grades byLevel) Len() int           { return len(grades) }
func (grades byLevel) Swap(i, j int)      { grades[i], grades[j] = grades[j], grades[i] }
func (grades byLevel) Less(i, j int) bool { return grades[i].class < grades[j].class }
