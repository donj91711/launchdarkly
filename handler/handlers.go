package handler

import (
	"encoding/json"
	"fmt"
	"launchdarklytest/dao"
	"launchdarklytest/db"
	"net/http"
	"sort"
)

type scoreList struct {
	Scores  []dao.Score
	Average float64
}

func ResultCount(w http.ResponseWriter) {
	fmt.Fprintf(w, "There are %d test results", len(dao.ResultMap))
}

func Students(w http.ResponseWriter) {
	studentList := db.GetStudentList()
	sort.Strings(studentList)

	jsonData, err := json.Marshal(studentList)
	if err != nil {
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(jsonData))
}

func Student(w http.ResponseWriter, studentID string) {
	resultList := db.GetStudentResultList(studentID)
	tot := 0.00
	cnt := 0

	scoreList := scoreList{}
	for _, v := range resultList {
		scoreList.Scores = append(scoreList.Scores, v)
		tot += v.Score
		cnt += 1
	}
	if tot > 0 && cnt > 0 {
		scoreList.Average = tot / float64(cnt)
	} else {
		scoreList.Average = 0.0
	}

	jsonData, err := json.Marshal(scoreList)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonData))
}

func Exams(w http.ResponseWriter) {
	examList := db.GetExamList()
	sort.Ints(examList)

	jsonData, err := json.Marshal(examList)
	if err != nil {
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonData))
}

func Exam(w http.ResponseWriter, exam string) {
	examResultList := db.GetExamResultList(exam)
	tot := 0.00
	cnt := 0

	examList := scoreList{}
	for _, v := range examResultList {
		examList.Scores = append(examList.Scores, v)
		tot += v.Score
		cnt += 1
	}
	if tot > 0 && cnt > 0 {
		examList.Average = tot / float64(cnt)
	} else {
		examList.Average = 0.0
	}

	jsonData, err := json.Marshal(examList)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonData))
}
