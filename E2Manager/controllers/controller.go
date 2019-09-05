//
// Copyright 2019 AT&T Intellectual Property
// Copyright 2019 Nokia
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package controllers

import (
	"e2mgr/configuration"
	"e2mgr/e2managererrors"
	"e2mgr/logger"
	"e2mgr/models"
	"e2mgr/providers/httpmsghandlerprovider"
	"e2mgr/rNibWriter"
	"e2mgr/services"
	"encoding/json"
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/reader"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	ParamRanName = "ranName"
)

type IController interface {
	ShutdownHandler(writer http.ResponseWriter, r *http.Request)
	X2ResetHandler(writer http.ResponseWriter, r *http.Request)
}

type Controller struct {
	logger          *logger.Logger
	handlerProvider *httpmsghandlerprovider.IncomingRequestHandlerProvider
}

func NewController(logger *logger.Logger, rmrService *services.RmrService, rNibReaderProvider func() reader.RNibReader, rNibWriterProvider func() rNibWriter.RNibWriter,
	config *configuration.Configuration) *Controller {

	provider := httpmsghandlerprovider.NewIncomingRequestHandlerProvider(logger, rmrService, config, rNibWriterProvider, rNibReaderProvider)
	return &Controller{
		logger:          logger,
		handlerProvider: provider,
	}
}

func (c *Controller) ShutdownHandler(writer http.ResponseWriter, r *http.Request) {
	c.logger.Infof("[Client -> E2 Manager] #controller.ShutdownHandler - request: %v", prettifyRequest(r))
	c.handleRequest(writer, &r.Header, httpmsghandlerprovider.ShutdownRequest, nil, false, http.StatusNoContent)
}

func (c *Controller) X2ResetHandler(writer http.ResponseWriter, r *http.Request) {
	c.logger.Infof("[Client -> E2 Manager] #controller.X2ResetHandler - request: %v", prettifyRequest(r))
	request := models.ResetRequest{}
	vars := mux.Vars(r)
	ranName := vars[ParamRanName]

	if !c.extractJsonBody(r, &request, writer) {
		return
	}
	request.RanName = ranName
	c.handleRequest(writer, &r.Header, httpmsghandlerprovider.ResetRequest, request, false, http.StatusNoContent)
}

func (c *Controller) extractJsonBody(r *http.Request, request models.Request, writer http.ResponseWriter) bool {
	if r.ContentLength <= 0 {
		return true
	}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(request)

	if err != nil {
		c.logger.Errorf("[Client -> E2 Manager] #controller.extractJsonBody - unable to extract json body - error: %s", err)
		c.handleErrorResponse(e2managererrors.NewInvalidJsonError(), writer)
		return false
	}

	return true
}

func (c *Controller) handleRequest(writer http.ResponseWriter, header *http.Header, requestName httpmsghandlerprovider.IncomingRequest,
	request models.Request, validateHeader bool, httpStatusResponse int) {

	if validateHeader {

		err := c.validateRequestHeader(header)
		if err != nil {
			c.handleErrorResponse(err, writer)
			return
		}
	}

	handler, err := c.handlerProvider.GetHandler(requestName)

	if err != nil {
		c.handleErrorResponse(err, writer)
		return
	}

	err = handler.Handle(c.logger, request)

	if err != nil {
		c.handleErrorResponse(err, writer)
		return
	}

	writer.WriteHeader(httpStatusResponse)
	c.logger.Infof("[E2 Manager -> Client] #controller.handleRequest - status response: %v", httpStatusResponse)
}

func (c *Controller) validateRequestHeader(header *http.Header) error {

	if header.Get("Content-Type") != "application/json" {
		c.logger.Errorf("#controller.validateRequestHeader - validation failure, incorrect content type")

		return e2managererrors.NewHeaderValidationError()
	}
	return nil
}

func (c *Controller) handleErrorResponse(err error, writer http.ResponseWriter) {

	var errorResponseDetails models.ErrorResponse
	var httpError int

	if err != nil {
		switch err.(type) {
		case *e2managererrors.RnibDbError:
			e2Error, _ := err.(*e2managererrors.RnibDbError)
			errorResponseDetails = models.ErrorResponse{Code: e2Error.Code, Message: e2Error.Message}
			httpError = http.StatusInternalServerError
		case *e2managererrors.CommandAlreadyInProgressError:
			e2Error, _ := err.(*e2managererrors.CommandAlreadyInProgressError)
			errorResponseDetails = models.ErrorResponse{Code: e2Error.Code, Message: e2Error.Message}
			httpError = http.StatusMethodNotAllowed
		case *e2managererrors.HeaderValidationError:
			e2Error, _ := err.(*e2managererrors.HeaderValidationError)
			errorResponseDetails = models.ErrorResponse{Code: e2Error.Code, Message: e2Error.Message}
			httpError = http.StatusUnsupportedMediaType
		case *e2managererrors.WrongStateError:
			e2Error, _ := err.(*e2managererrors.WrongStateError)
			errorResponseDetails = models.ErrorResponse{Code: e2Error.Code, Message: e2Error.Message}
			httpError = http.StatusBadRequest
		case *e2managererrors.RequestValidationError:
			e2Error, _ := err.(*e2managererrors.RequestValidationError)
			errorResponseDetails = models.ErrorResponse{Code: e2Error.Code, Message: e2Error.Message}
			httpError = http.StatusBadRequest
		case *e2managererrors.InvalidJsonError:
			e2Error, _ := err.(*e2managererrors.InvalidJsonError)
			errorResponseDetails = models.ErrorResponse{Code: e2Error.Code, Message: e2Error.Message}
			httpError = http.StatusBadRequest
		case *e2managererrors.RmrError:
			e2Error, _ := err.(*e2managererrors.RmrError)
			errorResponseDetails = models.ErrorResponse{Code: e2Error.Code, Message: e2Error.Message}
			httpError = http.StatusInternalServerError
		case *e2managererrors.ResourceNotFoundError:
			e2Error, _ := err.(*e2managererrors.ResourceNotFoundError)
			errorResponseDetails = models.ErrorResponse{Code: e2Error.Code, Message: e2Error.Message}
			httpError = http.StatusNotFound

		default:
			e2Error := e2managererrors.NewInternalError()
			errorResponseDetails = models.ErrorResponse{Code: e2Error.Code, Message: e2Error.Message}
			httpError = http.StatusInternalServerError
		}
	}
	errorResponse, _ := json.Marshal(errorResponseDetails)

	c.logger.Errorf("[E2 Manager -> Client] #controller.handleErrorResponse - http status: %d, error response: %+v", httpError, errorResponseDetails)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(httpError)
	_, err = writer.Write(errorResponse)

	if err != nil {
		c.logger.Errorf("#controller.handleErrorResponse - Cannot send response. writer:%v", writer)
	}
}
