// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package repository

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type QuestOfTodayQuestion struct {
	ID         uuid.UUID
	Question   string
	Answer     string
	SourceText string
	SourceUrl  string
	Dato       pgtype.Date
}
