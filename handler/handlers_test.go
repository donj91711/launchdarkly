package handler

import (
	"encoding/json"
	"launchdarklytest/dao"
	"launchdarklytest/db"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
)

func BuildTestData() {
	db.AddScore(`{"studentId":"studentA","exam":14102,"score":0.65}`)
	db.AddScore(`{"studentId":"studentB","exam":14102,"score":0.66}`)
	db.AddScore(`{"studentId":"studentC","exam":14103,"score":0.67}`)
	db.AddScore(`{"studentId":"studentC","exam":14104,"score":0.68}`)
	return
}

func TestExamsHandler(t *testing.T) {
	BuildTestData()

	rr := httptest.NewRecorder()
	Exams(rr)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Parse the response JSON
	var response []int
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error parsing JSON response: %v", err)
	}

	// Sort the response slice for comparison
	sort.Ints(response)

	// Expected sorted exam list
	expected := []int{14102, 14103, 14104}

	// Compare the response with the expected result
	if !compareIntSlices(response, expected) {
		t.Errorf("Handler returned unexpected result: got %v, want %v", response, expected)
	}
}

// Helper function to compare two slices of integers
func compareIntSlices(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func TestExamHandler(t *testing.T) {
	BuildTestData()

	// Create a ResponseRecorder to capture the HTTP response
	rr := httptest.NewRecorder()

	// Call the Exam function with the mock request, mock DB, and the "example" exam
	Exam(rr, "14102")

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Parse the response JSON
	var response scoreList
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error parsing JSON response: %v", err)
	}

	// Verify the expected data based on the mock database
	expected := scoreList{
		Scores: []dao.Score{
			{StudentID: "studentA", Exam: 14102, Score: 0.65},
			{StudentID: "studentB", Exam: 14102, Score: 0.66},
		},
		Average: 0.655,
	}

	// Compare the response with the expected result
	if !compareScores(response, expected) {
		t.Errorf("Handler returned unexpected result: got %v, want %v", response, expected)
	}
}

// Helper function to compare two scoreList structs
func compareScores(list1, list2 scoreList) bool {
	if len(list1.Scores) != len(list2.Scores) {
		return false
	}
	for i := range list1.Scores {
		if list1.Scores[i] != list2.Scores[i] {
			return false
		}
	}
	if list1.Average != list2.Average {
		return false
	}
	return true
}

func TestStudentsHandler(t *testing.T) {
	BuildTestData()

	// Create a ResponseRecorder to capture the HTTP response
	rr := httptest.NewRecorder()

	// Call the Students function with the mock request and mock DB
	Students(rr)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Parse the response JSON
	var response []string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error parsing JSON response: %v", err)
	}

	// Verify the expected data based on the mock database
	expected := []string{"studentA", "studentB", "studentC"}

	// Compare the response with the expected result
	if !compareStringSlices(response, expected) {
		t.Errorf("Handler returned unexpected result: got %v, want %v", response, expected)
	}
}

// Helper function to compare two slices of strings
func compareStringSlices(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	sort.Strings(slice1)
	sort.Strings(slice2)
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func TestStudentHandler(t *testing.T) {
	BuildTestData()

	// Create a ResponseRecorder to capture the HTTP response
	rr := httptest.NewRecorder()

	// Call the Student function with the mock request, mock DB, and "student1" as studentID
	Student(rr, "studentC")

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Parse the response JSON
	var response scoreList
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error parsing JSON response: %v", err)
	}

	// Verify the expected data based on the mock database
	expected := scoreList{
		Scores: []dao.Score{
			{StudentID: "studentC", Exam: 14103, Score: 0.67},
			{StudentID: "studentC", Exam: 14104, Score: 0.68},
		},
		Average: 0.675, // Calculate the expected average based on the mock data
	}

	// Compare the response with the expected result
	if !compareScores(response, expected) {
		t.Errorf("Handler returned unexpected result: got %v, want %v", response, expected)
	}
}
