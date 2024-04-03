package permalink

import (
	"fmt"
	"os"

	permalink "github.com/ProovGroup/lib-permalink"
)

var (
	ENV = os.Getenv("ENV")
	BASE_URL = os.Getenv("BASE_URL")
)

func GetPermalink(region string, bucket string, key string) string {
	if (key == "") {
		return ""
	}

	link := permalink.NewPermalink(ENV, permalink.S3, 0)

	url, err := link.SetRegion(region).SetBucket(bucket).SetKey(key).AppendToURL(BASE_URL)
	if err != nil {
		fmt.Println("[ERROR] GetPermalink:", err)
		return ""
	}

	return url
}



