package models

type struct ClientPackage {
    Url string `json:"url"`
    Description bool `json:"description"`
    Image bool `json:"image"`
    Video bool `json:"image"`
    Prompt string `json:"prompt"`
}

func NewClientPackage() ClientPackage {
    return ClientPackage{}
}

