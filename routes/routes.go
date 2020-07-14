package routes

import (
	"net/http"

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

// Response  Represent response http
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
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
	u := new(repository.Student)
	t := new(repository.Teacher)
	c := new(repository.Course)

	return &API{
		studentRepo: u,
		teacherRepo: t,
		courseRepo:  c,
	}
}

// New  handdle http endpoinst
func New() Server {

	a := &API{}
	initServices()

	r := mux.NewRouter()

	// Routes Studens
	r.HandleFunc("/students", a.createStudent).Methods(http.MethodPost)
	r.HandleFunc("/students", a.getStudents).Methods(http.MethodGet)
	r.HandleFunc("/student/{id}", a.getStudentByID).Methods(http.MethodGet)

	// Routes Teacher
	r.HandleFunc("/teacher", a.createTeacher).Methods(http.MethodPost)
	r.HandleFunc("/teacher/{id}", a.getTeacherByID).Methods(http.MethodGet)
	r.HandleFunc("/teachers", a.getTeachers).Methods(http.MethodGet)

	//Course Student
	r.HandleFunc("/course", a.createCourse).Methods(http.MethodPost)
	r.HandleFunc("/course/{id}", a.getCourseByID).Methods(http.MethodGet)
	r.HandleFunc("/courses", a.getAllCourses).Methods(http.MethodGet)

	a.router = r
	return a
}
