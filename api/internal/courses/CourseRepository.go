package courses

import (
    "context"
    "database/sql"
    "errors"
    "fmt"
    clogger "fullstackguru/pkg/logger"
    validator "fullstackguru/pkg/vaildator"
    "github.com/doug-martin/goqu/v9"
    _ "github.com/doug-martin/goqu/v9/dialect/postgres" // selecting goqu dialect
)

const (
    noPercentage = 100
    schema       = "fullstackguruapi"
)

var (
    ErrCreateCourse = errors.New("error while creating new course")
)

type Course struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Lessons     string `json:"lessons"`
    Duration    string `json:"duration"`
}

func ValidateCourse(v *validator.Validator, course *Course) {
    // v.Check(test.FirstName != "", "firstName", "must be provided")
    // v.Check(len(test.FirstName) <= 500, "firstName", "must not be more than 500 bytes long")
}

type CourseRepo struct {
    db               *sql.DB
    courseDetailRepo CourseDetailRepo
}

type CourseDetailRepo interface {
    InsertCourseDetailRecord(ctx context.Context) error
}

func goquBuilder() goqu.DialectWrapper {
    return goqu.Dialect("postgres")
}

func (r *CourseRepo) GetAll() ([]Course, error) {
    var courses []Course

    query, args, err := goquBuilder().From(
        goqu.S(schema).Table("courses")).Prepared(true).ToSQL()

    if err != nil {
        clogger.Error(err)
        return nil, err
    }

    rows, err := r.db.Query(query, args...)
    if err != nil {
        clogger.Error(err)

        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var c Course
        err := rows.Scan(&c.ID, &c.Title, &c.Description, &c.Lessons, &c.Duration)
        if err != nil {
            clogger.Error(err)

            return nil, err
        }
        courses = append(courses, c)
    }

    if err = rows.Err(); err != nil {
        clogger.Error(err)

        return nil, err
    }

    return courses, nil
}

func (r *CourseRepo) GetCourseById(id string) (Course, error) {
    var c Course

    query, args, err := goquBuilder().From(
        goqu.S(schema).Table("courses")).Where(goqu.Ex{"id": id}).Prepared(true).ToSQL()
    if err != nil {
        clogger.Error(err)

        return Course{}, err
    }

    if err = r.db.QueryRow(query, args...).Scan(
        &c.ID,
        &c.Title,
        &c.Description,
        &c.Lessons,
        &c.Duration); err != nil {
        if err == sql.ErrNoRows {
            clogger.Info("course id was not found")

            return c, err
        }
    }

    return c, nil
}

func (r *CourseRepo) CreateCourse(ctx context.Context, c *Course) error {

    var tx *sql.Tx

    tx, err := r.db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }

    err = r.saveCourseWithTransaction(ctx, c, tx)
    if err != nil {
        return err
    }

    return nil
}

func (r *CourseRepo) saveCourseWithTransaction(ctx context.Context, course *Course, tx *sql.Tx) error {

    err := r.createCourseRecord(ctx, course, tx)
    if err != nil {
        clogger.Error(ErrCreateCourse)
    }

    defer func() {
        err = r.commitOrRollback(err, tx)
    }()

    return err
}

func (r *CourseRepo) createCourseRecord(ctx context.Context, c *Course, tx *sql.Tx) error {

    query, params, err := goquBuilder().Insert(goqu.S(schema).Table("courses")).Rows(goqu.Record{
        "description": c.Description,
        "title":       c.Title,
        "lessons":     c.Lessons,
        "duration":    c.Duration,
    }).Prepared(true).Returning("id").ToSQL()
    if err != nil {
        return err
    }

    _, err = tx.ExecContext(ctx, query, params...)

    return err
}

func (r *CourseRepo) commitOrRollback(err error, tx *sql.Tx) error {
    if err == nil {
        if errT := tx.Commit(); errT != nil {
            return fmt.Errorf("error in tx Commit: %w", errT)
        }
    } else {
        if errT := tx.Rollback(); errT != nil {
            // choose losing rollback error type because we can use the type of the incoming error in the caller
            return fmt.Errorf("error in tx Rollback: %v : %w", errT, err)
        }
    }

    return err
}

func NewCourseRepository(db *sql.DB) *CourseRepo {
    return &CourseRepo{
        db: db,
    }
}
