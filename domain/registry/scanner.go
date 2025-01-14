package registry

import (
	"fmt"

	"github.com/forta-protocol/forta-core-go/contracts/contract_scanner_registry"
	"github.com/forta-protocol/forta-core-go/utils"
	"github.com/goccy/go-json"
)

var SaveScanner = "SaveScanner"
var EnableScanner = "EnableScanner"
var DisableScanner = "DisableScanner"

type ScannerMessage struct {
	Message
	ScannerID string `json:"scannerId"`
}

type ScannerSaveMessage struct {
	ScannerMessage
	ChainID int64 `json:"chainId"`
}

func ParseScannerSave(msg string) (*ScannerSaveMessage, error) {
	var save ScannerSaveMessage
	err := json.Unmarshal([]byte(msg), &save)
	if err != nil {
		return nil, err
	}
	if save.Action != SaveScanner {
		return nil, fmt.Errorf("invalid action for ScannerSave: %s", save.Action)
	}
	return &save, nil
}

func ParseScannerMessage(msg string) (*ScannerMessage, error) {
	var m ScannerMessage
	err := json.Unmarshal([]byte(msg), &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func NewScannerMessage(evt *contract_scanner_registry.ScannerRegistryScannerEnabled) *ScannerMessage {
	scannerID := utils.HexAddr(evt.ScannerId)
	evtName := DisableScanner
	if evt.Enabled {
		evtName = EnableScanner
	}
	return &ScannerMessage{
		Message: Message{
			Action: evtName,
		},
		ScannerID: scannerID,
	}
}

func NewScannerSaveMessage(evt *contract_scanner_registry.ScannerRegistryScannerUpdated) *ScannerSaveMessage {
	scannerID := utils.HexAddr(evt.ScannerId)
	return &ScannerSaveMessage{
		ScannerMessage: ScannerMessage{
			ScannerID: scannerID,
			Message: Message{
				Action: SaveScanner,
			},
		},
		ChainID: evt.ChainId.Int64(),
	}
}
