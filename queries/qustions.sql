-- name: GetTodaysQuestion :one
SELECT id, question FROM quest_of_today.questions WHERE dato = $1;
