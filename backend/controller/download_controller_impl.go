package controller

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	myErrors "github.com/proyek2-git/HIMA-TI-e-Election/backend/errors"
	"github.com/proyek2-git/HIMA-TI-e-Election/backend/helper"
	"github.com/proyek2-git/HIMA-TI-e-Election/backend/model/web"
	"github.com/proyek2-git/HIMA-TI-e-Election/backend/service"
)

func NewDownloadController(downloadService service.DownloadService) DownloadController {
	return &DownloadControllerImpl{
		DownloadService: downloadService,
	}
}

type DownloadControllerImpl struct {
	DownloadService service.DownloadService
}

func (controller *DownloadControllerImpl) GetPresignedUrl(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Get query paramaters
	filename := r.URL.Query().Get("filename")

	if filename != "" {
		data, err := controller.DownloadService.CreatePresignedURL(r.Context(), filename)
		if err != nil {
			var customError *myErrors.AppError

			if errors.As(err, &customError) {
				myErrors.LogError(err, "failed to get download presigned URL")

				w.WriteHeader(customError.StatusCode)
				helper.WriteToResponseBody(w, map[string]web.WebFailedResponse{
					"error": {
						Message: customError.Message,
						Details: customError.Details,
					},
				})
				return
			} else {
				myErrors.LogError(err, "unexpected error")

				w.WriteHeader(http.StatusInternalServerError)
				helper.WriteToResponseBody(w, map[string]web.WebFailedResponse{
					"error": {
						Message: "Internal Server Error",
						Details: "Internal Server Error. Please try again later.",
					},
				})
				return
			}
		}

		// Write and send the response
		w.WriteHeader(http.StatusOK)
		helper.WriteToResponseBody(w, web.WebSuccessResponse{
			Message: "Success download log file",
			Data:    data,
		})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		helper.WriteToResponseBody(w, map[string]web.WebFailedResponse{
			"error": {
				Message: "Invalid request query parameters",
				Details: "Query parameter 'filename' is required",
			},
		})
		return
	}
}
