package responses

type HistoryResponse struct {
	History string `json:"History" example:"1c8d54df80c03a56b5470d164c49f823108f96a67d020e4c677810c9a10b1ca7"`
}
type TxResponse struct {
	Tx string `json:"Tx" example:1"`
}
type TxsResponse struct {
	Txs []string `json:"Txs" example:"1,2,3"`
}
type SuccessResponse struct {
	Status string `json:"Status" example:"Ok"`
}
