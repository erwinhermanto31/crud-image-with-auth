package mysql

const (
	// User
	QueryFindUser = `SELECT id, username, password FROM users`

	// Image
	QueryFindImage = `SELECT id, user_id, image_url, upload_time FROM images`

	QueryInsertImage = `INSERT INTO images (user_id, image_url, upload_time) VALUES (:user_id, :image_url, :upload_time)`

	QueryUpdateQuestion = `UPDATE images SET image_url=:image_url, upload_time=:upload_time WHERE id=:id`
)
