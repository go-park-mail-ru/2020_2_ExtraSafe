package service

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/storage"
	protoBoard "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/board"
	protoProfile "github.com/go-park-mail-ru/2020_2_ExtraSafe/services/proto/profile"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	boardStorage boardStorage.Storage
}


func NewService(boardStorage boardStorage.Storage) *service {
	return &service{
		boardStorage: boardStorage,
	}
}

func (s *service) GetBoardsList(c context.Context, input *protoProfile.UserID) (output *protoProfile.BoardsOutsideShort, err error) {
	userInput := models.UserInput{ID: input.ID}

	boardsList, err := s.boardStorage.GetBoardsList(userInput)
	if err != nil {
		return output, nil
	}

	output = new(protoProfile.BoardsOutsideShort)

	for _, board := range boardsList {
		outputBoard := protoProfile.BoardOutsideShort{
			BoardID: board.BoardID,
			Name:    board.Name,
			Theme:   board.Theme,
			Star:    board.Star,
		}
		output.Boards = append(output.Boards, &outputBoard)
	}
	return output, nil
}

func (s *service) CreateBoard(c context.Context, input *protoBoard.BoardChangeInput) (output *protoProfile.BoardOutsideShort, err error) {
	userInput := models.BoardChangeInput{
		UserID:    input.UserID,
		BoardID:   input.BoardID,
		BoardName: input.BoardName,
		Theme:     input.Theme,
		Star:      input.Star,
	}

	boardInternal, err := s.boardStorage.CreateBoard(userInput)
	if err != nil {
		return output, err
	}

	output = &protoProfile.BoardOutsideShort{
		BoardID: boardInternal.BoardID,
		Name:    boardInternal.Name,
		Theme:   boardInternal.Theme,
		Star:    boardInternal.Star,
	}
	return output, nil
}

func (s *service) GetBoard(c context.Context, input *protoBoard.BoardInput) (output *protoBoard.BoardOutside, err error) {
	userInput := models.BoardInput{
		UserID:    input.UserID,
		BoardID:   input.BoardID,
	}

	boardInternal, err := s.boardStorage.GetBoard(userInput)
	if err != nil {
		return output, err
	}

	output = &protoBoard.BoardOutside{
		BoardID: boardInternal.BoardID,
		Admin:   ,
		Name:    boardInternal.Name,
		Theme:   boardInternal.Theme,
		Star:    boardInternal.Star,
		Users:   nil,
		Cards:   boardInternal.Cards,
	}

	return output, nil
}

func (s *service) ChangeBoard(c context.Context, input *protoBoard.BoardChangeInput) (output *protoBoard.BoardOutside, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeBoard not implemented")
}

func (s *service) DeleteBoard(c context.Context, input *protoBoard.BoardInput) (*protoBoard.Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBoard not implemented")
}

func (s *service) CreateCard(c context.Context, input *protoBoard.CardInput) (output *protoBoard.CardOutside, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCard not implemented")
}

func (s *service) GetCard(c context.Context, input *protoBoard.CardInput) (output *protoBoard.CardOutside, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCard not implemented")
}

func (s *service) ChangeCard(c context.Context, input *protoBoard.CardInput) (output *protoBoard.CardOutside, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeCard not implemented")
}

func (s *service) DeleteCard(c context.Context, input *protoBoard.CardInput) (*protoBoard.Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCard not implemented")
}

func (s *service) CardOrderChange(c context.Context, input *protoBoard.CardsOrderInput) (*protoBoard.Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CardOrderChange not implemented")
}

func (s *service) CreateTask(c context.Context, input *protoBoard.TaskInput) (output *protoBoard.TaskOutside, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (s *service) GetTask(c context.Context, input *protoBoard.TaskInput) (output *protoBoard.TaskOutside, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}

func (s *service) ChangeTask(c context.Context, input *protoBoard.TaskInput) (output *protoBoard.TaskOutside, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeTask not implemented")
}

func (s *service) DeleteTask(c context.Context, input *protoBoard.TaskInput) (*protoBoard.Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}

func (s *service) TasksOrderChange(c context.Context, input *protoBoard.TasksOrderInput) (*protoBoard.Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TasksOrderChange not implemented")
}