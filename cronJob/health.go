package cronJob

import (
	"network-chesswork/net"
)

func HealthCheck(host string) error {
	err := net.Ping(host)
	if err != nil {
		net.Reconnect("")
		return err
	}
	return nil
}
