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

package httpmsghandlerprovider

import (
	"e2mgr/configuration"
	"e2mgr/e2managererrors"
	"e2mgr/handlers"
	"e2mgr/logger"
	"e2mgr/mocks"
	"e2mgr/models"
	"e2mgr/rNibWriter"
	"e2mgr/rmrCgo"
	"e2mgr/services"
	"e2mgr/sessions"
	"e2mgr/tests"
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/reader"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func getRmrService(rmrMessengerMock *mocks.RmrMessengerMock, log *logger.Logger) *services.RmrService {
	rmrMessenger := rmrCgo.RmrMessenger(rmrMessengerMock)
	messageChannel := make(chan *models.NotificationResponse)
	rmrMessengerMock.On("Init", tests.GetPort(), tests.MaxMsgSize, tests.Flags, log).Return(&rmrMessenger)
	return services.NewRmrService(services.NewRmrConfig(tests.Port, tests.MaxMsgSize, tests.Flags, log), rmrMessenger, make(sessions.E2Sessions), messageChannel)
}

func TestNewIncomingRequestHandlerProvider(t *testing.T) {
	rmrMessengerMock := &mocks.RmrMessengerMock{}

	log := initLog(t)
	readerProvider := func() reader.RNibReader {
		return &mocks.RnibReaderMock{}
	}
	writerProvider := func() rNibWriter.RNibWriter {
		return &mocks.RnibWriterMock{}
	}

	provider := NewIncomingRequestHandlerProvider(log, getRmrService(rmrMessengerMock, log), configuration.ParseConfiguration(), writerProvider, readerProvider)
	/*if provider == nil {
		t.Errorf("want: provider, got: nil")
	}*/

	assert.NotNil(t, provider)
}

func TestShutdownRequestHandler(t *testing.T) {
	rmrMessengerMock := &mocks.RmrMessengerMock{}

	log := initLog(t)
	readerProvider := func() reader.RNibReader {
		return &mocks.RnibReaderMock{}
	}
	writerProvider := func() rNibWriter.RNibWriter {
		return &mocks.RnibWriterMock{}
	}

	provider := NewIncomingRequestHandlerProvider(log, getRmrService(rmrMessengerMock, log), configuration.ParseConfiguration(), writerProvider, readerProvider)

	handler, err := provider.GetHandler(ShutdownRequest)

	/*if handler == nil {
		t.Errorf("failed to get x2 setup handler")
	}*/
	assert.NotNil(t, provider)
	assert.Nil(t, err)

	_, ok := handler.(*handlers.DeleteAllRequestHandler)

	assert.True(t, ok)
	/*if !ok {
		t.Errorf("failed to delete all handler")
	}*/
}

func TestGetShutdownHandlerFailure(t *testing.T) {
	rmrMessengerMock := &mocks.RmrMessengerMock{}
	log := initLog(t)
	readerProvider := func() reader.RNibReader {
		return &mocks.RnibReaderMock{}
	}
	writerProvider := func() rNibWriter.RNibWriter {
		return &mocks.RnibWriterMock{}
	}

	provider := NewIncomingRequestHandlerProvider(log, getRmrService(rmrMessengerMock, log), configuration.ParseConfiguration(), writerProvider, readerProvider)

	_, actual := provider.GetHandler("test")
	expected := &e2managererrors.InternalError{}

	assert.NotNil(t, actual)
	if reflect.TypeOf(actual) != reflect.TypeOf(expected) {
		t.Errorf("Error actual = %v, and Expected = %v.", actual, expected)
	}
}

func initLog(t *testing.T) *logger.Logger {
	log, err := logger.InitLogger(logger.InfoLevel)
	if err != nil {
		t.Errorf("#delete_all_request_handler_test.TestHandleSuccessFlow - failed to initialize logger, error: %s", err)
	}
	return log
}