package db

import (
	"fmt"
	"launchdarklytest/dao"
	"testing"
)

func BuildTestData() {
	AddScore(`{"studentId":"studentA","exam":14102,"score":0.65}`)
	AddScore(`{"studentId":"studentB","exam":14102,"score":0.66}`)
	AddScore(`{"studentId":"studentC","exam":14103,"score":0.67}`)
	AddScore(`{"studentId":"studentC","exam":14104,"score":0.68}`)
	return
}

func TestAddScore(t *testing.T) {
	testCases := []struct {
		testResult    string
		expectedCount int
	}{
		{`{"studentId":"Hannah.Herman76","exam":14102,"score":0.6815787969197457}`, 1}, // happy path
		{`bad data`, 0}, // error path
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("AddScore(%s)", tc.testResult), func(t *testing.T) {
			dao.ResultMap = make(map[string]dao.Score) //reset map
			AddScore(tc.testResult)
			if len(dao.ResultMap) != tc.expectedCount {
				t.Errorf("AddScore(%s) returned %d, expected %d", tc.testResult, len(dao.ResultMap), tc.expectedCount)
			}
		})
	}
}

func TestGetStudentList(t *testing.T) {
	testCases := []struct {
		expectedResult []string
		expectedCount  int
		expectError    bool
	}{
		{[]string{"studentA", "studentB", "studentC"}, 3, false}, //happy path
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("GetStudentList(%s)", tc.expectedResult), func(t *testing.T) {
			dao.ResultMap = make(map[string]dao.Score) //reset map
			BuildTestData()
			studentList := GetStudentList()
			if CompareStringSlices(studentList, tc.expectedResult) != true {
				if tc.expectError == false {
					t.Errorf("GetStudentList() returned %d, expected %d", len(studentList), tc.expectedCount)
				}
			}
		})
	}
}

func TestGetStudentListResult(t *testing.T) {
	testCases := []struct {
		expectedResult []dao.Score
		testStudentID  string
		expectedCount  int
	}{
		{
			expectedResult: []dao.Score{
				{Exam: 14103, Score: 0.67, StudentID: "studentC"},
				{Exam: 14104, Score: 0.68, StudentID: "studentC"},
			},
			testStudentID: "studentC",
			expectedCount: 2,
		},
		{
			expectedResult: []dao.Score{
				{Exam: 14102, Score: 0.65, StudentID: "studentA"},
			},
			testStudentID: "studentA",
			expectedCount: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("GetStudentResultList: %s", tc.testStudentID), func(t *testing.T) {
			dao.ResultMap = make(map[string]dao.Score) //reset map
			BuildTestData()
			examList := GetStudentResultList(tc.testStudentID)
			if len(examList) != tc.expectedCount {
				t.Errorf("GetStudentList() returned %d, expected %d", len(examList), tc.expectedCount)
			}
			for i, v := range examList {
				if v.StudentID != tc.expectedResult[i].StudentID {
					t.Errorf("GetStudentList() returned %s, expected %s", v.StudentID, tc.expectedResult[i].StudentID)
				}
			}
		})
	}
}

func TestGetExamList(t *testing.T) {
	testCases := []struct {
		expectedResult []int
		expectedCount  int
		expectError    bool
	}{
		{[]int{14102, 14103, 14104}, 3, false}, //happy path
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("GetExamList(%d)", tc.expectedCount), func(t *testing.T) {
			dao.ResultMap = make(map[string]dao.Score) //reset map
			BuildTestData()
			examList := GetExamList()
			if CompareIntSlices(examList, tc.expectedResult) != true {
				if tc.expectError == false {
					t.Errorf("GetExamList() returned %d, expected %d", len(examList), tc.expectedCount)
				}
			}
		})
	}
}

func TestGetExamResultList(t *testing.T) {
	testCases := []struct {
		expectedResult []dao.Score
		testExam       string
		expectedCount  int
	}{
		{
			expectedResult: []dao.Score{
				{Exam: 14102, Score: 0.65, StudentID: "studentA"},
				{Exam: 14102, Score: 0.66, StudentID: "studentB"},
			},
			testExam:      "14102",
			expectedCount: 2,
		},
		{
			expectedResult: []dao.Score{
				{Exam: 14103, Score: 0.67, StudentID: "studentC"},
			},
			testExam:      "14103",
			expectedCount: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("GetExamResultList: %s", tc.testExam), func(t *testing.T) {
			dao.ResultMap = make(map[string]dao.Score) //reset map
			BuildTestData()
			examList := GetExamResultList(tc.testExam)
			if len(examList) != tc.expectedCount {
				t.Errorf("GetExamResultList() returned %d, expected %d", len(examList), tc.expectedCount)
			}
			for i, v := range examList {
				if v.Score != tc.expectedResult[i].Score {
					t.Errorf("GetExamResultList() returned %f, expected %f", v.Score, tc.expectedResult[i].Score)
				}
			}
		})
	}
}

// CompareStringSlices compares two string slices and returns true if they are equal, false otherwise.
func CompareStringSlices(slice1, slice2 []string) bool {
	// Check if the slices have the same length
	if len(slice1) != len(slice2) {
		return false
	}

	// Compare each element of the slices
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

func CompareIntSlices(slice1, slice2 []int) bool {
	// Check if the slices have the same length
	if len(slice1) != len(slice2) {
		return false
	}

	// Compare each element of the slices
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
