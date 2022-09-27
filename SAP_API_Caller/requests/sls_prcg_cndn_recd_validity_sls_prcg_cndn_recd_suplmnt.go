package requests

type SlsPrcgCndnRecdValiditySlsPrcgCndnRecdSuplmnt struct {
	SlsPrcgCndnRecdValidity
	ToSlsPrcgCndnRecdSuplmnt `json:"to_SlsPrcgCndnRecdSuplmnt"`
}

type ToSlsPrcgCndnRecdSuplmnt struct {
	Results []SlsPrcgCndnRecdSuplmnt `json:"results"`
}
