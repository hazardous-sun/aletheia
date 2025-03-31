package repositories_test

import (
	"database/sql"
	"errors"
	"fact-checker-server/src/models"
	"fact-checker-server/src/repositories"
	"fact-checker-server/src/errors"
	"github.com/lib/pq"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestLanguageRepository_AddLanguage(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := repositories.NewLanguageRepository(db)

	tests := []struct {
		name         string
		mockBehavior func()
		input        models.Language
		expectedId   int
		expectedErr  error
	}{
		{
			name: "Success",
			mockBehavior: func() {
				mock.ExpectPrepare("INSERT INTO languages (.+)").ExpectQuery().
					WithArgs("english").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			input:       models.Language{Name: "English"},
			expectedId:  1,
			expectedErr: nil,
		},
		{
			name: "Duplicate Language",
			mockBehavior: func() {
				mock.ExpectPrepare("INSERT INTO languages (.+)").ExpectQuery().
					WithArgs("english").
					WillReturnError(&pq.Error{Code: "23505", Constraint: "languages_name_key"})
			},
			input:       models.Language{Name: "English"},
			expectedId:  -1,
			expectedErr: errors.New(server_errors.LanguageAlreadyExists),
		},
		{
			name: "Table Missing",
			mockBehavior: func() {
				mock.ExpectPrepare("INSERT INTO languages (.+)").WillReturnError(sql.ErrNoRows)
			},
			input:       models.Language{Name: "English"},
			expectedId:  -1,
			expectedErr: errors.New(server_errors.LanguageTableMissing),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			id, err := repo.AddLanguage(tt.input)
			assert.Equal(t, tt.expectedId, id)
			assert.Equal(t, tt.expectedErr, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestLanguageRepository_GetLanguages(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := repositories.NewLanguageRepository(db)

	tests := []struct {
		name          string
		mockBehavior  func()
		expected      []models.Language
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(1, "english").
					AddRow(2, "spanish")
				mock.ExpectQuery("SELECT \\* FROM languages").WillReturnRows(rows)
			},
			expected: []models.Language{
				{Id: 1, Name: "english"},
				{Id: 2, Name: "spanish"},
			},
			expectedError: nil,
		},
		{
			name: "Table Missing",
			mockBehavior: func() {
				mock.ExpectQuery("SELECT \\* FROM languages").WillReturnError(sql.ErrNoRows)
			},
			expected:      []models.Language{},
			expectedError: errors.New(server_errors.LanguageTableMissing),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			result, err := repo.GetLanguages()
			assert.Equal(t, tt.expected, result)
			assert.Equal(t, tt.expectedError, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestLanguageRepository_GetLanguageById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := repositories.NewLanguageRepository(db)

	tests := []struct {
		name          string
		mockBehavior  func()
		input         int
		expected      *models.Language
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "english")
				mock.ExpectPrepare("SELECT \\* FROM languages WHERE id = \\$1").
					ExpectQuery().
					WithArgs(1).
					WillReturnRows(rows)
			},
			input:         1,
			expected:      &models.Language{Id: 1, Name: "english"},
			expectedError: nil,
		},
		{
			name: "Not Found",
			mockBehavior: func() {
				mock.ExpectPrepare("SELECT \\* FROM languages WHERE id = \\$1").
					ExpectQuery().
					WithArgs(1).
					WillReturnError(sql.ErrNoRows)
			},
			input:         1,
			expected:      nil,
			expectedError: errors.New(server_errors.LanguageNotFound),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			result, err := repo.GetLanguageById(tt.input)
			assert.Equal(t, tt.expected, result)
			assert.Equal(t, tt.expectedError, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestLanguageRepository_GetLanguageByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := repositories.NewLanguageRepository(db)

	tests := []struct {
		name          string
		mockBehavior  func()
		input         string
		expected      *models.Language
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "english")
				mock.ExpectPrepare("SELECT \\* FROM languages WHERE name = \\$1").
					ExpectQuery().
					WithArgs("english").
					WillReturnRows(rows)
			},
			input:         "english",
			expected:      &models.Language{Id: 1, Name: "english"},
			expectedError: nil,
		},
		{
			name: "Not Found",
			mockBehavior: func() {
				mock.ExpectPrepare("SELECT \\* FROM languages WHERE name = \\$1").
					ExpectQuery().
					WithArgs("english").
					WillReturnError(sql.ErrNoRows)
			},
			input:         "english",
			expected:      nil,
			expectedError: errors.New(server_errors.LanguageNotFound),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			result, err := repo.GetLanguageByName(tt.input)
			assert.Equal(t, tt.expected, result)
			assert.Equal(t, tt.expectedError, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
