package smodbus

import (
	"errors"
	"fmt"

	"github.com/goburrow/modbus"
	"github.com/sigurn/crc16"
)

type RtuOverTcpHandler struct {
	modbus.Packager
	modbus.Transporter
	SlaveId byte
}

func (r *RtuOverTcpHandler) Encode(pdu *modbus.ProtocolDataUnit) (adu []byte, err error) {
	adu = []byte{r.SlaveId, pdu.FunctionCode}
	adu = append(adu, pdu.Data...)
	table := crc16.MakeTable(crc16.CRC16_MODBUS)
	crc := crc16.Checksum(adu, table)
	// 3. 附加 CRC (低位在前)
	adu = append(adu, byte(crc), byte(crc>>8))
	return adu, nil
}

func (r *RtuOverTcpHandler) Verify(aduRequest []byte, aduResponse []byte) (err error) {

	// 1. 基本长度判断
	if len(aduResponse) < 5 {
		return errors.New("response too short")
	}

	// 2. 比较 SlaveId（第1字节）
	if aduResponse[0] != aduRequest[0] {
		return fmt.Errorf("slave ID mismatch: request 0x%X, response 0x%X", aduRequest[0], aduResponse[0])
	}

	// 3. 比较功能码（第2字节）
	if aduResponse[1] != aduRequest[1] {
		return fmt.Errorf("function code mismatch: request 0x%X, response 0x%X", aduRequest[1], aduResponse[1])
	}

	// 4. 验证 CRC16
	data := aduResponse[:len(aduResponse)-2]
	expectedCRC := crc16.Checksum(data, crc16.MakeTable(crc16.CRC16_MODBUS))
	responseCRC := uint16(aduResponse[len(aduResponse)-2]) | uint16(aduResponse[len(aduResponse)-1])<<8

	if expectedCRC != responseCRC {
		return fmt.Errorf("crc check failed: expected 0x%X, got 0x%X", expectedCRC, responseCRC)
	}
	return
}

func (mb *RtuOverTcpHandler) Decode(adu []byte) (pdu *modbus.ProtocolDataUnit, err error) {
	// RTU ADU 至少应为：1(SlaveId) + 1(FunctionCode) + 2(CRC)
	if len(adu) < 5 {
		err = errors.New("modbus: ADU too short for RTU")
		return
	}

	// 计算 PDU 长度：总长度 - 1(SlaveId) - 2(CRC)
	pduLength := len(adu) - 3
	if pduLength <= 0 {
		err = errors.New("modbus: invalid PDU length")
		return
	}

	pdu = &modbus.ProtocolDataUnit{
		FunctionCode: adu[1],              // 第2字节是功能码
		Data:         adu[2 : len(adu)-2], // 中间是数据（不含CRC）
	}

	return pdu, nil
}
