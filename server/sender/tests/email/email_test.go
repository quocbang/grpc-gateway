package email

import (
	"context"

	"github.com/google/uuid"
	"github.com/quocbang/grpc-gateway/server/utils/html/activate"
)

func (s *Suite) TestSendVerifyEmail() {
	assertion := s.Assertions

	// good case
	{
		// Arrange
		to := "bangquoc9@gmail.com"
		subject := "WellCome to grpc gateway server designed by quocbang"
		id := "quocbang"
		secretCode := uuid.NewString()

		// content = html.GetActivateHTMLCode(id, secretCode)
		activateService := activate.NewHTMLActivateService(id, secretCode)
		content, err := activateService.GenerateHTML()
		assertion.NoError(err)
		assertion.NotNil(content)

		// Act
		err = s.Sender.Email().SendVerifyEmail(context.Background(), to, subject, content)

		// Assert
		assertion.NoError(err)
	}
}
