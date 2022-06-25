package interfaces

import "io"

type IDocument interface {
	AddPage(io.Reader)
	SaveDocumentToFile(filePath string) error
	GetFileContent() io.Reader
}
