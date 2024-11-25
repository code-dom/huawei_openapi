package table

import (
	"fmt"
	"log"
	"openapi/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	ocr "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ocr/v1"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ocr/v1/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ocr/v1/region"
)

// func main() {
// 	fmt.Println("请输入识别表格：")
// 	var path string
// 	fmt.Scan(&path)
// 	Table(path)
// }

func Table(path string) (string, *model.RecognizeGeneralTableResponse) {
	rep := GetTable(path)
	var result string
	for _, v := range rep.Result.WordsRegionList {
		for _, vv := range v.WordsBlockList {
			result += *vv.Words
		}
		result += "\n"
	}
	return result, rep
}

func GetExcel(rep *model.RecognizeGeneralTableResponse, path string) {
	if rep.Result.Excel != nil {
		encoded := *rep.Result.Excel
		if err := utils.DecodeBase64ToFile(encoded, path); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Excel file has been saved as result.xlsx")
		}
	} else {
		fmt.Println("OCR response Excel is nil")
	}
}

func GetTable(path string) *model.RecognizeGeneralTableResponse {
	// The AK and SK used for authentication are hard-coded or stored in plaintext, which has great security risks. It is recommended that the AK and SK be stored in ciphertext in configuration files or environment variables and decrypted during use to ensure security.
	// In this example, AK and SK are stored in environment variables for authentication. Before running this example, set environment variables CLOUD_SDK_AK and CLOUD_SDK_SK in the local environment
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
	returnExcel := true
	request := &model.RecognizeGeneralTableRequest{}
	request.Body = &model.GeneralTableRequestBody{
		Image:       &image,
		ReturnExcel: &returnExcel,
	}
	response, err := client.RecognizeGeneralTable(request)
	if err == nil {
		// fmt.Printf("%+v\n", response)
		fmt.Println("response: ok")
	} else {
		fmt.Println(err)
	}
	return response
}
