package model

const(
	MainUrl = "http://localhost:4080/api/index"
	IndexUrl = "http://localhost:4080/api/index/enron_zinc_v03"
	SearchUrl = "http://localhost:4080/api/enron_zinc_v03"
)

type Email struct {
	File                    string `json:"File"`
	User                    string `json:"User"`
	MessageID               string `json:"Message-ID"`
	Date                    string `json:"Date"`
	From                    string `json:"From"`
	To                      string `json:"To"`
	Subject                 string `json:"Subject"`
	MimeVersion             string `json:"Mime-Version"`
	ContentType             string `json:"Content-Type"`
	ContentTransferEncoding string `json:"Content-Transfer-Encoding"`
	XFrom                   string `json:"X-From"`
	XTo                     string `json:"X-To"`
	Xcc                     string `json:"X-cc"`
	Xbcc                    string `json:"X-bcc"`
	XFolder                 string `json:"X-Folder"`
	XOrigin                 string `json:"X-Origin"`
	XFileName               string `json:"X-FileName"`
	Message                 string `json:"Message"`
}

type SearchInput struct {
	SearchType  string   `json:"search_type"`
	Query       Query    `json:"query"`
	SortFields  []string `json:"sort_fields"`
	From        int      `json:"from"`
	MaxResults  int      `json:"max_results"`
	Source      []string `json:"_source"`
}

type Query struct {
	Term  string `json:"term"`
	Field string `json:"field"`
}