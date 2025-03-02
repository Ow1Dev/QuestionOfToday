-- name: GetTodaysQuestion :one
SELECT id, question FROM quest_of_today.questions WHERE dato = $1;

-- name: GetAnswerById :one
SELECT answer, source_url, source_text FROM quest_of_today.questions where id = $1;
