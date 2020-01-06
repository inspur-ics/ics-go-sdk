package types

type SDKError struct {
    Code              string        `json:"code"`
    Message           string        `json:"message"`
    Params            string        `json:"params"`
}

func (e *SDKError) Error() string {
    return e.Message
}