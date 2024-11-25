package ocr

import (
	"log"
	"openapi/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	ocr "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ocr/v1"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ocr/v1/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ocr/v1/region"
)

// func main() {
// 	fmt.Println("请输入识别图片：")
// 	var path string
// 	fmt.Scan(&path)
// 	rep := GetOcr(path)
// 	var result string
// 	for _, v := range rep.Result.WordsBlockList {
// 		result += v.Words
// 	}
// 	fmt.Println(result)
// }

func OCR(path string) string {
	rep := GetOcr(path)
	var result string
	for _, v := range rep.Result.WordsBlockList {
		result += v.Words
	}
	return result
}
func GetOcr(path string) *model.RecognizeGeneralTextResponse {
	// The AK and SK used for authentication are hard-coded or stored in plaintext, which has great security risks. It is recommended that the AK and SK be stored in ciphertext in configuration files or environment variables and decrypted during use to ensure security.
	// In this example, AK and SK are stored in environment variables for authentication. Before running this example, set environment variables CLOUD_SDK_AK and CLOUD_SDK_SK in the local environment
	// ak := os.Getenv("CLOUD_SDK_AK")
	// sk := os.Getenv("CLOUD_SDK_SK")
	ak := "W49XRXFNJKJ32BYSEEUL"
	sk := "MAGbgWpEHRjdKMsRGIa08SK4Ti0KzEDMpQ4iQXX1"

	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := ocr.NewOcrClient(
		ocr.OcrClientBuilder().
			WithRegion(region.ValueOf("cn-north-4")).
			WithCredential(auth).
			Build())

	image, err := utils.EncodeImageToBase64(path)
	if err != nil {
		log.Fatal(err)
	}

	request := &model.RecognizeGeneralTextRequest{}
	request.Body = &model.GeneralTextRequestBody{
		Image: &image,
	}
	response, err := client.RecognizeGeneralText(request)
	if err == nil {
		// fmt.Printf("[Response] %+v\n", response)
	} else {
		log.Fatal(err)
	}
	return response
}
