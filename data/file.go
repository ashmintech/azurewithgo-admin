package data

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

const (
	ContainerName      = "devicefiles"
	DeviceFileCollName = "devicefiles"
)

type File struct {
	FileName   string `json:"filename"`
	FileDesc   string `json:"filedesc"`
	FileType   string `json:"filetype"`
	FileDate   string `json:"filedate"`
	FilePath   string `json:"filepath"`
	FileHeader string `json:"fileHeader"`
}

type Files []File

var FileList = []File{}

func GetFiles() Files {

	mcoll := GetCollection(DeviceFileCollName)

	err = mcoll.Find(nil).Iter().All(&FileList)
	if err != nil {
		log.Println("Error while querying the Collection:\n", err)
		return nil
	}
	return FileList
}

func PutFiles(r *http.Request) (string, bool) {

	//	log.Println("Into Put files")

	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("uploadfile")

	if err != nil {
		log.Println("Error Retrieving the File")
		return "", false
	}

	defer file.Close()

	accountName := os.Getenv("AZURE_STORAGE_NAME")
	accountKey := os.Getenv("AZURE_STORAGE_KEY")

	cred, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatalln("Not able to connect to storage account")
	}

	ctx := context.Background()

	p := azblob.NewPipeline(cred, azblob.PipelineOptions{})

	u, _ := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net", accountName))

	serviceURL := azblob.NewServiceURL(*u, p)

	cURL := serviceURL.NewContainerURL(ContainerName)
	bURL := cURL.NewBlockBlobURL(handler.Filename)

	resp, err := azblob.UploadStreamToBlockBlob(ctx, file, bURL, azblob.UploadStreamToBlockBlobOptions{})
	if err != nil {
		log.Println("Cannot upload the file:", err)
	}
	
	f := File{
		FileName:   r.FormValue("filename"),
		FileDesc:   r.FormValue("filedesc"),
		FileType:   r.FormValue("filetype"),
		FileDate:   time.Now().Format("January 2 2006"),
		FilePath:   bURL.String(),
		FileHeader: fmt.Sprintf("%v", handler.Header),
	}

	mcoll := GetCollection(DeviceFileCollName)

	err = mcoll.Insert(f)
	if err != nil {
		log.Println("Error while Inserting the File record: ", err)
		return "", false
	}

	return resp.Response().Status, true

}
