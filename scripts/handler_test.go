package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/AMao7/taskmanager/api/handlers"
	"github.com/AMao7/taskmanager/pkg/entity"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateTask(t *testing.T) {

	// Initialising a gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// We create a new SQL mock database. The sqlmock.New() function returns a *sql.DB which can be used just like a real SQL database, as well as a sqlmock.Sqlmock
	db, mock, err := sqlmock.New()

	// We check that no error occurred using require.NoError(t, err) which will stop the test immediately if an error occurred.
	require.NoError(t, err)

	// Next, we create a new GORM database using the SQL mock database as the underlying SQL database. Again, we check that no error occurred.

	gdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

	require.NoError(t, err)
	t.Log("Hello")

	// Set up the mock database behavior

	user := entity.User{
		ID:       1,
		Name:     "abdimao",
		Email:    "abdimao@example.com",
		Password: "password",
	}

	task := entity.Task{
		ID:        1,
		Title:     "Testing Task",
		Content:   "Sample Content",
		CreatedAt: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
		UserID:    user.ID,
		User:      user,
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users"`).WithArgs(user.Name, user.Email, user.Password, user.ID).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectQuery(`INSERT INTO "tasks"`).WithArgs(task.Title, task.Content, task.CreatedAt, task.UpdatedAt, task.UserID, task.ID).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	// We register the CreateTask handler on the router. We wrap the handler in another function that sets the db context value to our GORM database. This allows us to inject the mock database into the handler.

	router.Use(func(c *gin.Context) {
		c.Set("db", gdb)
	})

	router.POST("/tasks", handlers.CreateTask)

	// create a new HTTP request that we'll pass to the handler.

	// marshal the Task struct into JSON and put it in the body of the request. We also set the "Content-Type" header to "application/json" to indicate that we're sending JSON.

	body, err := json.Marshal(task)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	// create a ResponseRecorder which is an implementation of http.ResponseWriter that records its mutations for later inspection in tests.

	rr := httptest.NewRecorder()

	// We pass the request and the ResponseRecorder to the router, which will route the request to the appropriate handler.

	router.ServeHTTP(rr, req)

	// We check that the status code of the response is 200 OK.

	assert.Equal(t, http.StatusOK, rr.Code)

	var responseTask entity.Task

	err = json.Unmarshal(rr.Body.Bytes(), &responseTask)
	require.NoError(t, err)

	assert.Equal(t, task.ID, responseTask.ID)
	assert.Equal(t, task.Title, responseTask.Title)
}
