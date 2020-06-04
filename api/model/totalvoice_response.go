package model

// TotalVoiceResponse - Response padrão de retorno da TotalVoice
type TotalVoiceResponse struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		IDS []int `json:"id"`
	} `json:"dados"`
}
