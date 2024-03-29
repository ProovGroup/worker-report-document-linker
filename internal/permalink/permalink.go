package permalink

import (
	"fmt"
	"os"

	permalink "github.com/ProovGroup/lib-permalink"
)

var ENV = os.Getenv("ENV")

func GetPermalink(region string, bucket string, key string) string {
	if (key == "") {
		return ""
	}

	baseURL := "https://permalink2.weproov.com"
	link := permalink.NewPermalink(ENV, permalink.S3, 0)

	url, err := link.SetRegion(region).SetBucket(bucket).SetKey(key).AppendToURL(baseURL)
	if err != nil {
		fmt.Println("[ERROR] GetPermalink:", err)
		return ""
	}

	return url
}



