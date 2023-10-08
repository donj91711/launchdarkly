package db

import (
	"encoding/json"
	"fmt"
	"launchdarklytest/dao"
	"strconv"
)

// AddScore: Add the exam results received from the listener into the database (ResultMap)
func AddScore(update string) {
	score := dao.Score{}

	err := json.Unmarshal([]byte(update), &score)
	if err != nil {
		fmt.Println("Error unmarshalling data:", err)
		return
	}
	fmt.Println(len(dao.ResultMap), score)
	key := score.StudentID + fmt.Sprintf("%d", score.Exam)
	dao.ResultMap[key] = score
}

// GetStudentList: Get a unique list of students from the master list of exam results
func GetStudentList() []string {
	students := make(map[string]string)
	var studentList []string

	//add each student to a map with student id as the key - this eliminates duplicates
	for _, v := range dao.ResultMap {
		if students[v.StudentID] != v.StudentID {
			studentList = append(studentList, v.StudentID)
			students[v.StudentID] = v.StudentID
		}
	}
	return studentList
}

// GetStudentResultList: Get a list of exam results for a given student id
func GetStudentResultList(studentID string) []dao.Score {
	resultList := []dao.Score{}

	for _, v := range dao.ResultMap {
		if v.StudentID == studentID {
			resultList = append(resultList, dao.Score{StudentID: v.StudentID, Exam: v.Exam, Score: v.Score})
		}
	}

	return resultList
}

// GetExamList: Get a unique list of exams from the master list of exam results
func GetExamList() []int {
	exams := make(map[int]int)
	var examList []int

	//add each exam to a map with exam number id as the key - this eliminates duplicates
	for _, v := range dao.ResultMap {
		if exams[v.Exam] != v.Exam {
			examList = append(examList, v.Exam)
			exams[v.Exam] = v.Exam
		}
	}
	return examList
}

// GetExamResultList: Get a list of exam results for a given exam number
func GetExamResultList(exam string) []dao.Score {
	resultList := []dao.Score{}

	for _, v := range dao.ResultMap {
		intExam, _ := strconv.Atoi(exam)
		if v.Exam == intExam {
			resultList = append(resultList, dao.Score{StudentID: v.StudentID, Exam: v.Exam, Score: v.Score})
		}
	}

	return resultList
}
