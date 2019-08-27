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

package managers

import (
	"e2mgr/logger"
	"e2mgr/rNibWriter"
	"e2mgr/services"
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/entities"
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/reader"
)

type RanSetupManager struct {
	logger             *logger.Logger
	rnibReaderProvider func() reader.RNibReader
	rnibWriterProvider func() rNibWriter.RNibWriter
	rmrService         *services.RmrService
}

func NewRanSetupManager(logger *logger.Logger, rmrService *services.RmrService, rnibReaderProvider func() reader.RNibReader, rnibWriterProvider func() rNibWriter.RNibWriter) *RanSetupManager {
	return &RanSetupManager{
		logger:             logger,
		rnibReaderProvider: rnibReaderProvider,
		rnibWriterProvider: rnibWriterProvider,
	}
}

func (m *RanSetupManager) ExecuteSetup(nodebInfo *entities.NodebInfo) error {
	return nil
}
