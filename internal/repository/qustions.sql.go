// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: qustions.sql

package repository

import (
	"context"
)

const getAllQuestions = `-- name: GetAllQuestions :many
SELECT id, question FROM quest_of_today.questions
`

func (q *Queries) GetAllQuestions(ctx context.Context) ([]QuestOfTodayQuestion, error) {
	rows, err := q.db.Query(ctx, getAllQuestions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []QuestOfTodayQuestion
	for rows.Next() {
		var i QuestOfTodayQuestion
		if err := rows.Scan(&i.ID, &i.Question); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}