package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 55, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 59, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 70, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 70, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 60, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestAssignmentString(t *testing.T) {
	expected_value := "assignment"

	gradeType := Assignment

	actual_value := gradeType.String()

	if expected_value != actual_value {
		t.Errorf("Expected String to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestExamString(t *testing.T) {
	expected_value := "exam"

	gradeType := Exam

	actual_value := gradeType.String()

	if expected_value != actual_value {
		t.Errorf("Expected String to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestEssayString(t *testing.T) {
	expected_value := "essay"

	gradeType := Essay

	actual_value := gradeType.String()

	if expected_value != actual_value {
		t.Errorf("Expected String to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
