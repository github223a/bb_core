package core

type Log struct {
	Type         string `json:"type"`
	FilePath     string `json:"filePath"`
	FileLevel    string `json:"fileLevel"`
	ConsoleLevel string `json:"consoleLevel"`
}
