package redis_socket

import (
	"github.com/giskook/mdps/base"
	"github.com/giskook/mdps/pb"
	"log"
)

func (socket *RedisSocket) FilterAlters(alters []*Report.DataCommand, alters_redis []*base.RouterAlterRedis) []*base.RouterAlter {
	if len(alters) != len(alters_redis) {
		log.Println("FilterAlters alters not equal")
		return nil
	}

	alter := make([]*base.RouterAlter, 0)
	for i, v := range alters {
		CompareAlter(v, alters_redis[i].Alters, alter, v.Tid, uint8(v.SerialPort))
	}

	return alter
}

func CompareAlter(data_command *Report.DataCommand, alter_redis []*base.Alter, result []*base.RouterAlter, router_id uint64, serial_port uint8) {
	alter_rep := data_command.Alters
	for _, a := range alter_rep {
		for _, b := range alter_redis {
			if a.ModusAddr == b.ModbusAddr &&
				a.DataType == uint32(b.DataType) &&
				a.DataLen == uint32(b.DataLen) &&
				a.Status != uint32(b.Status) {
				result = append(result, &base.RouterAlter{
					RouterID:   router_id,
					SerialPort: serial_port,
					ModbusAddr: b.ModbusAddr,
					DataType:   b.DataType,
					Data:       b.Data,
					Status:     b.Status,
				})
			}
		}
	}
}
