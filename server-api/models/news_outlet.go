package models

type NewsOutlet struct {
    Url string `json:"url"`
    Name string `json:"name"`
    Language string `json:"language"`
    Credibility int `json:"credibility"`
}
