package data

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

const ContainerHistoricalName = "deviceanomalydata"

type HistorialData struct {
	DeviceID  string `json:"deviceid"`
	TimeStamp string `json:"timestamp"`
	FileName  string `json:"filename"`
	FilePath  string `json:"filepath"`
}

var DataList []HistorialData

func ShowHistoricalData() []HistorialData {

	accountName := os.Getenv("AZURE_STORAGE_NAME")
	accountKey := os.Getenv("AZURE_STORAGE_KEY")

	cred, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatalln("Not able to connect to storage account")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	p := azblob.NewPipeline(cred, azblob.PipelineOptions{})

	u, _ := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net", accountName))

	serviceURL := azblob.NewServiceURL(*u, p)

	cURL := serviceURL.NewContainerURL(ContainerHistoricalName)
	listBlob, err := cURL.ListBlobsFlatSegment(ctx, azblob.Marker{}, azblob.ListBlobsSegmentOptions{MaxResults: 50})
	
	if listBlob != nil {
		blobs := listBlob.Segment.BlobItems
		DataList = nil
		var h HistorialData
		for _, b := range blobs {
			s := strings.Split(b.Name, "/")
			h = HistorialData{
				DeviceID:  s[0],
				TimeStamp: s[1],
				FileName:  s[2],
				FilePath:  path.Join(cURL.String(), b.Name),
			}
			DataList = append(DataList, h)
		}

		sort.SliceStable(DataList, func(i, j int) bool {
			return DataList[i].TimeStamp > DataList[j].TimeStamp
		})
	}

	return DataList
}
