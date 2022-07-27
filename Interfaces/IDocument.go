package interfaces

import "io"

type IDocument interface {
	AddPage(io.Reader) error
	SaveDocumentToFile(filePath string) error
	GetFileContent() (io.Reader, error)
	Copy() IDocument
}
