package gateway

type Gateway struct {
	Devices GMap[int64, Client]
}
