
import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func FetchThenSaveFile(r *http.Request, sourceFileName string, destinationFileName string) (multipart.File, *multipart.FileHeader, error) {
	file, handler, err := r.FormFile(sourceFileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	f, err := os.OpenFile(destinationFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	io.Copy(f, file)

	return file, handler, nil
}

func (t *DataSourceController) SaveDataSource(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputJson
	r.Request.ParseMultipartForm(32 << 20)

	formData := "file"
	destinationPath := "/file/to/path"
	file, fileHeader, err := FetchThenSaveFile(r.Request, formData, destinationPath)
	if err != nil {
		// err
	}
}