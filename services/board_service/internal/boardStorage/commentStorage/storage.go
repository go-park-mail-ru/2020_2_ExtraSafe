package commentStorage

import (
	"database/sql"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
)

type Storage interface {
	CreateComment(input models.CommentInput) (comment models.CommentOutside, err error)
	UpdateComment(input models.CommentInput) (comment models.CommentOutside, err error)
	DeleteComment(input models.CommentInput) (comment models.CommentOutside, err error)

	GetCommentsByTask(input models.TaskInput) (comments []models.CommentOutside, userIDS[] int64, err error)
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) CreateComment(input models.CommentInput) (comment models.CommentOutside, err error) {
	err = s.db.QueryRow("INSERT INTO comments (message, taskID, commentOrder, userID) VALUES ($1, $2, $3, $4) RETURNING commentID, taskID, cardID",
						input.Message, input.TaskID, input.Order, input.UserID).Scan(&comment.CommentID, &comment.TaskID, &comment.CardID)
	if err != nil {
		return comment, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "CreateComment"}
	}

	comment.Message = input.Message
	comment.Order = input.Order
	return
}

func (s *storage) UpdateComment(input models.CommentInput) (comment models.CommentOutside, err error) {
	err = s.db.QueryRow("UPDATE comments SET message = $1 WHERE commentID = $2 RETURNING taskID, cardID", input.Message, input.CommentID).
		Scan(&comment.TaskID, &comment.CardID)
	if err!= nil {
		return comment, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "UpdateComment"}
	}
	comment.Message = input.Message
	comment.CommentID = input.CommentID
	return
}

func (s *storage) DeleteComment(input models.CommentInput) (comment models.CommentOutside, err error) {
	comment = models.CommentOutside{
		CommentID: input.CommentID,
	}
	err = s.db.QueryRow("DELETE FROM comments WHERE commentID = $1 RETURNING taskID, cardID", input.CommentID).
		Scan(&comment.TaskID, &comment.CardID)
	if err!= nil {
		return models.CommentOutside{}, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "DeleteComment"}
	}
	return
}

func (s *storage) GetCommentsByTask(input models.TaskInput) (comments []models.CommentOutside, userIDS[] int64, err error) {
	rows, err := s.db.Query("SELECT commentID, message, commentOrder, userID FROM comments WHERE taskID = $1", input.TaskID)
	if err != nil {
		return comments, userIDS, models.ServeError{Codes: []string{"500"}, OriginalError: err,
			MethodName: "GetBoardTags"}
	}
	defer rows.Close()

	for rows.Next() {
		comment := models.CommentOutside{}
		var userID int64

		err = rows.Scan(&comment.CommentID, &comment.Message, &comment.Order, &userID)
		if err != nil {
			return comments, userIDS, models.ServeError{Codes: []string{"500"}, OriginalError: err,
				MethodName: "GetBoardTags"}
		}

		comments = append(comments, comment)
		userIDS = append(userIDS, userID)
	}

	return
}