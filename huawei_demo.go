package main

import (
	"fmt"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	ocr "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ocr/v1"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ocr/v1/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ocr/v1/region"
)

func main() {
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

	request := &model.RecognizeGeneralTableRequest{}
	request.Body = &model.GeneralTableRequestBody{}
	response, err := client.RecognizeGeneralTable(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
