package boards

/*func TestService_CreateBoard(t *testing.T) {
	request := &protoBoard.BoardChangeInput{
		UserID:    1,
		BoardID:   1,
		BoardName: "name",
		Theme:     "dark",
		Star:      false,
	}

	input := models.BoardChangeInput{
		UserID:    request.UserID,
		BoardID:   request.BoardID,
		BoardName: request.BoardName,
		Theme:     request.Theme,
		Star:      request.Star,
	}

	internal := models.BoardInternal{
		BoardID: input.BoardID,
		Name:    input.BoardName,
		Theme:   input.Theme,
		Star:    input.Star,
	}

	expect := &protoProfile.BoardOutsideShort{
		BoardID: internal.BoardID,
		Name:    internal.Name,
		Theme:   internal.Theme,
		Star:    internal.Star,
	}

	ctrlBoard := gomock.NewController(t)
	defer ctrlBoard.Finish()
	mockBoardService := mock.NewMockBoardClient(ctrlBoard)

	validator := validation.NewService()

	service := NewService(mockBoardService, validator)

	mockBoardService.EXPECT().CreateBoard(context.Background(), request).Return(expect, nil)

	board, err := service.CreateBoard(input)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if !reflect.DeepEqual(board, internal) {
		t.Errorf("results not match, want \n%v, \nhave \n%v", internal, board)
		return
	}
}*/