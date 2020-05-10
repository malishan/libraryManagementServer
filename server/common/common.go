package common

type MetaStatus struct {
	Code int    `json:"statusCode"`
	Msg  string `json:"msg"`
}

// NewBook struct to contain book info when storing
// type newBook struct {
// 	SerialNos string `json:"serial_no"`
// 	BookName  string `json:"name"`
// 	Author    string `json:"author"`
// 	Category  string `json:"category"`
// }

const (
	Fail = iota
	Success
)

// var (
// 	categories = []string{"fiction", "thriller", "biography", "educational"} //to-do: implement using iota
// )
