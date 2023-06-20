package courses

import (
    "context"
    "encoding/json"
    "errors"
    "fullstackguru/pkg"
    clogger "fullstackguru/pkg/logger"
    "github.com/gorilla/mux"
    "net/http"
)

var (
    ErrDecodingCourse = errors.New("error occurred while decoding the course object")
)

type Repository interface {
    GetAll() ([]Course, error)
    GetCourseById(id string) (Course, error)
    CreateCourse(ctx context.Context, c *Course) error
}

type Courses struct {
    helper      pkg.Helper
    repo        Repository
    detailsRepo CourseDetailRepo
}

type envelope map[string]any

func (app *Courses) CoursesAllHandler(w http.ResponseWriter, r *http.Request) {
    ctx := context.Background()

    clog := clogger.GetLoggerFromContext(ctx)

    clog.Info("Trying to make database call")

    res, err := app.repo.GetAll()
    if err != nil {
        clog.Error(err)

        return
    }

    err = app.helper.WriteJSON(w, http.StatusOK, envelope{"data": res, "metadata": "none"}, nil)
    if err != nil {
        clog.ErrorCtx(err, clogger.Ctx{
            "header":      w.Header(),
            "request_url": r.URL.String(),
        })
    }

    clog.Info("Complete")

}

func (app *Courses) CoursesIdHandler(w http.ResponseWriter, r *http.Request) {
    ctx := context.Background()

    clog := clogger.GetLoggerFromContext(ctx)

    vars := mux.Vars(r)
    id, ok := vars["id"]
    if !ok {
        clog.Warn("id is missing in parameters")
    }

    res, err := app.repo.GetCourseById(id)
    if err != nil {
        clog.Error(err)

        return
    }

    err = app.helper.WriteJSON(w, http.StatusOK, envelope{"data": res, "metadata": "none"}, nil)
    if err != nil {
        clogger.ErrorCtx(err, clogger.Ctx{
            "header":      w.Header(),
            "request_url": r.URL.String(),
        })
    }

    clog.Info("Complete")

}

func (app *Courses) CreateCourseHandler(w http.ResponseWriter, r *http.Request) {
    ctx := context.Background()

    clog := clogger.GetLoggerFromContext(ctx)

    var c Course

    err := json.NewDecoder(r.Body).Decode(&c)
    if err != nil {
        clog.ErrorCtx(err, clogger.Ctx{
            "msg": ErrDecodingCourse,
        })
    }

    clog.InfoCtx("course object", clogger.Ctx{
        "lessons": c.Lessons,
        "title":   c.Title,
    })

    //TODO: add a validator here for course

    course := &Course{
        Title:       c.Title,
        Duration:    c.Duration,
        Lessons:     c.Lessons,
        Description: c.Description,
    }

    err = app.repo.CreateCourse(ctx, course)
    if err != nil {
        clog.Error(err)
    }

}

func NewCoursesService(repo Repository) *Courses {
    return &Courses{
        repo: repo,
    }
}
