package controller

type struct NewsOutlet {
    Url string `json:"url"`
    Name string `json:"name"`
    Language string `json:"language"`
    credibility int `json:"credibility"`
}

func NewNewsOutlet() NewsOutlet {
    return NewsOutlet{}
}

