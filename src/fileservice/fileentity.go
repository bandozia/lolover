package fileservice

type FileEntity struct {
	Name    string `json:"name"`
	IsDir   bool   `json:"isDir"`
	IsError bool   `json:"isError"`
}
