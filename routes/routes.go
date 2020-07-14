package routes

import (
	"net/http"

	"github.com/AdrianMendez1199/simple-crud-go-sql/repository"
	"github.com/gorilla/mux"
)

type API struct {
	router      http.Handler
	studentRepo *repository.Student
	teacherRepo *repository.Teacher
	courseRepo  *repository.Course
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Server interface {
	Router() http.Handler
}

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

func New() Server {

	a := &API{}
	initServices()

	r := mux.NewRouter()

	// Routes Studens
	r.HandleFunc("/students", a.CreateStudent).Methods(http.MethodPost)
	r.HandleFunc("/students", a.GetStudents).Methods(http.MethodGet)
	r.HandleFunc("/student/{id}", a.GetStudentById).Methods(http.MethodGet)

	// Routes Teacher
	r.HandleFunc("/teacher", a.createTeacher).Methods(http.MethodPost)
	r.HandleFunc("/teacher/{id}", a.getTeacherById).Methods(http.MethodGet)
	r.HandleFunc("/teachers", a.getTeachers).Methods(http.MethodGet)

	//Course Student
	r.HandleFunc("/course", a.createCourse).Methods(http.MethodPost)

	a.router = r
	return a
}
