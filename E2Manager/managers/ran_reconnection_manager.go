package managers

import (
	"e2mgr/configuration"
	"e2mgr/logger"
	"e2mgr/rNibWriter"
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/common"
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/entities"
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/reader"
)

type RanReconnectionManager struct {
	logger             *logger.Logger
	config             *configuration.Configuration
	rnibReaderProvider func() reader.RNibReader
	rnibWriterProvider func() rNibWriter.RNibWriter
	ranSetupManager    *RanSetupManager
}

func NewRanReconnectionManager(logger *logger.Logger, config *configuration.Configuration, rnibReaderProvider func() reader.RNibReader, rnibWriterProvider func() rNibWriter.RNibWriter, ranSetupManager *RanSetupManager) *RanReconnectionManager {
	return &RanReconnectionManager{
		logger:             logger,
		config:             config,
		rnibReaderProvider: rnibReaderProvider,
		rnibWriterProvider: rnibWriterProvider,
		ranSetupManager:    ranSetupManager,
	}
}

func (m *RanReconnectionManager) ReconnectRan(inventoryName string) error {
	nodebInfo, rnibErr := m.rnibReaderProvider().GetNodeb(inventoryName)

	if rnibErr != nil {
		m.logger.Errorf("#ReconnectRan - RAN name: %s - Failed fetching RAN from rNib. Error: %v", inventoryName, rnibErr)
		return rnibErr
	}

	if !m.canReconnectRan(nodebInfo) {
		m.logger.Warnf("#ReconnectRan - RAN name: %s - Cannot reconnect RAN", inventoryName)
		return m.setConnectionStatusOfUnconnectableRan(nodebInfo)
	}

	err := m.ranSetupManager.ExecuteSetup(nodebInfo)

	if err != nil {
		m.logger.Errorf("#ReconnectRan - RAN name: %s - Failed executing setup. Error: %v", inventoryName, err)
		return err
	}

	m.logger.Infof("#ReconnectRan - RAN name: %s - Successfully done executing setup", inventoryName)
	return nil
}

func (m *RanReconnectionManager) canReconnectRan(nodebInfo *entities.NodebInfo) bool {
	connectionStatus := nodebInfo.GetConnectionStatus()
	return connectionStatus != entities.ConnectionStatus_SHUT_DOWN && connectionStatus != entities.ConnectionStatus_SHUTTING_DOWN &&
		int(nodebInfo.GetConnectionAttempts()) < m.config.MaxConnectionAttempts
}

func (m *RanReconnectionManager) updateNodebInfoStatus(nodebInfo *entities.NodebInfo, connectionStatus entities.ConnectionStatus) common.IRNibError {
	nodebInfo.ConnectionStatus = connectionStatus;
	err := m.rnibWriterProvider().UpdateNodebInfo(nodebInfo)

	if err != nil {
		m.logger.Errorf("#updateNodebInfoStatus - RAN name: %s - Failed updating RAN's connection status to %s in rNib. Error: %v", nodebInfo.RanName, connectionStatus, err)
		return err
	}

	m.logger.Infof("#updateNodebInfoStatus - RAN name: %s - Successfully updated RAN's connection status to %s in rNib", nodebInfo.RanName, connectionStatus)
	return nil
}

func (m *RanReconnectionManager) setConnectionStatusOfUnconnectableRan(nodebInfo *entities.NodebInfo) common.IRNibError {
	connectionStatus := nodebInfo.GetConnectionStatus()
	m.logger.Warnf("#setConnectionStatusOfUnconnectableRan - RAN name: %s, RAN's connection status: %s, RAN's connection attempts: %d", nodebInfo.RanName, nodebInfo.ConnectionStatus, nodebInfo.ConnectionAttempts)

	if connectionStatus == entities.ConnectionStatus_SHUTTING_DOWN {
		return m.updateNodebInfoStatus(nodebInfo, entities.ConnectionStatus_SHUT_DOWN)
	}

	if int(nodebInfo.GetConnectionAttempts()) >= m.config.MaxConnectionAttempts {
		m.logger.Warnf("#setConnectionStatusOfUnconnectableRan - RAN name: %s - RAN's connection attempts are greater than %d", nodebInfo.RanName, m.config.MaxConnectionAttempts)
		return m.updateNodebInfoStatus(nodebInfo, entities.ConnectionStatus_DISCONNECTED)
	}

	return nil
}