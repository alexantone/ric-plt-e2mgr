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

package rmrmsghandlers

import (
	"e2mgr/logger"
	"e2mgr/managers"
	"e2mgr/models"
	"e2mgr/services"
	"encoding/json"
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/common"
)

type E2TermInitNotificationHandler struct {
	logger                 *logger.Logger
	rnibDataService        services.RNibDataService
	ranReconnectionManager *managers.RanReconnectionManager
	e2tInstancesManager    managers.IE2TInstancesManager
}

func NewE2TermInitNotificationHandler(logger *logger.Logger, ranReconnectionManager *managers.RanReconnectionManager, rnibDataService services.RNibDataService, e2tInstancesManager managers.IE2TInstancesManager) E2TermInitNotificationHandler {
	return E2TermInitNotificationHandler{
		logger:                 logger,
		rnibDataService:        rnibDataService,
		ranReconnectionManager: ranReconnectionManager,
		e2tInstancesManager:    e2tInstancesManager,
	}
}

func (h E2TermInitNotificationHandler) Handle(request *models.NotificationRequest) {
	unmarshalledPayload := models.E2TermInitPayload{}
	err :=  json.Unmarshal(request.Payload, &unmarshalledPayload)

	if err != nil {
		h.logger.Errorf("#E2TermInitNotificationHandler.Handle - Error unmarshaling E2 Term Init payload: %s", err)
		return
	}

	e2tAddress := unmarshalledPayload.Address

	if len(e2tAddress) == 0 {
		h.logger.Errorf("#E2TermInitNotificationHandler.Handle - Empty E2T address received")
		return
	}

	h.logger.Infof("#E2TermInitNotificationHandler.Handle - E2T address: %s - handling E2_TERM_INIT", e2tAddress)

	e2tInstance, err := h.e2tInstancesManager.GetE2TInstance(e2tAddress)

	if err != nil {
		_, ok := err.(*common.ResourceNotFoundError)

		if !ok {
			h.logger.Errorf("#E2TermInitNotificationHandler.Handle - Failed retrieving E2TInstance. error: %s", err)
			return
		}

		_ = h.e2tInstancesManager.AddE2TInstance(e2tAddress)
		return
	}

	if len(e2tInstance.AssociatedRanList) == 0 {
		h.logger.Infof("#E2TermInitNotificationHandler.Handle - E2T Address: %s - E2T instance has no associated RANs", e2tInstance.Address)
		return
	}

	for _, ranName := range e2tInstance.AssociatedRanList {

		if err := h.ranReconnectionManager.ReconnectRan(ranName); err != nil {
			h.logger.Errorf("#E2TermInitNotificationHandler.Handle - Ran name: %s - connection attempt failure, error: %s", ranName, err)
			_, ok := err.(*common.ResourceNotFoundError)
			if !ok {
				break
			}
		}
	}

	h.logger.Infof("#E2TermInitNotificationHandler.Handle - Completed handling of E2_TERM_INIT")
}
