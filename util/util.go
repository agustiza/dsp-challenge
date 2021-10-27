package util

import (
	"net"
	"net/http"
)

func ParseQueryImpId(r *http.Request) (impId string, bidId string) {
	impId = r.URL.Query().Get("id")
	bidId = r.URL.Query().Get("bidId")
	return impId, bidId
}

func IP(r *http.Request) (string, error) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	return ip, nil
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
