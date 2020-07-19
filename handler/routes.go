package handler

import (
	"net/http"

	"github.com/AdrianMendez1199/simple-crud-go-sql/middleware"
	"github.com/AdrianMendez1199/simple-crud-go-sql/repository"
	"github.com/gorilla/mux"
)

// API Contains all repository reference and http.Handle
type API struct {
	router      http.Handler
	studentRepo *repository.Student
	teacherRepo *repository.Teacher
	courseRepo  *repository.Course
}

// Server interfce implements
type Server interface {
	Router() http.Handler
}

// Router return router from API struct
func (a *API) Router() http.Handler {
	return a.router
}

func initServices() *API {
	//User Repo
	u := &repository.Student{}
	t := &repository.Teacher{}
	c := &repository.Course{}

	return &API{
		studentRepo: u,
		teacherRepo: t,
		courseRepo:  c,
	}
}

// Setting default headers like a middleware
// func setDefaultHeaders(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT")
// 		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
// 		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 		next.ServeHTTP(w, r)
// 	})
// }

// New  handdle http endpoinst
func New() Server {

	a := &API{}
	initServices()

	r := mux.NewRouter()

	// handler Studens
	r.HandleFunc("/api/v1/students", a.createStudent).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/students", middleware.WriteLog(a.getStudents)).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/student/{id}", a.getStudentByID).Methods(http.MethodGet)

	// handler Teacher
	r.HandleFunc("/api/v1/teacher", a.createTeacher).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/teacher/{id}", a.getTeacherByID).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/teachers", a.getTeachers).Methods(http.MethodGet)

	//Course Student
	r.HandleFunc("/api/v1/course", a.createCourse).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/course/{id}", a.getCourseByID).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/courses", a.getAllCourses).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/course/", a.searchCourse).
		Queries("search", "{search}").
		Methods(http.MethodGet)

	// r.Use(setDefaultHeaders)
	a.router = r
	return a
}
