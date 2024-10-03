package therapist

import (
	"backend-mental-guardians/controllers/therapist/response"
	"backend-mental-guardians/entities/therapist"
	"backend-mental-guardians/utilities"
	"backend-mental-guardians/utilities/base"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TherapistController struct {
	therapistUseCase therapist.UseCaseInterface
}

func NewTherapistController (therapistUseCase therapist.UseCaseInterface) *TherapistController {
	return &TherapistController{
		therapistUseCase: therapistUseCase,
	}
}

func (tc *TherapistController) GetAll(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	metadata := utilities.GetMetadata(pageParam, limitParam)

	specialist := c.QueryParam("specialist")

	token := c.Request().Header.Get("Authorization")
	_, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	therapistData, err := tc.therapistUseCase.GetAll(*metadata, specialist)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	therapistResps := make([]response.TherapistResponse, len(therapistData))
	for i, therapist := range therapistData {
		therapistResps[i] = response.TherapistResponse{
			ID:  					therapist.ID,
			Name: 					therapist.Name,
			Age:                    therapist.Age,
			Specialist:             therapist.Specialist,
			PhotoURL:               therapist.PhotoURL,
			PhoneNumber:            therapist.PhoneNumber,
			Gender:                 therapist.Gender,
			Experience:             therapist.Experience,
			Fee:                    therapist.Fee,
			PracticeCity:           therapist.PracticeCity,
			PracticeLocation:       therapist.PracticeLocation,
			BachelorAlmamater:      therapist.BachelorAlmamater,
			BachelorGraduationYear: therapist.BachelorGraduationYear,
			ConsultationMode:       therapist.ConsultationMode,
		}
	}

	return c.JSON(http.StatusOK, base.NewMetadataSuccessResponse("Success Get All Therapist", metadata, therapistResps))
}