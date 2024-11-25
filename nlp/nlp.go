package nlp

import (
	"fmt"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	nlp "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/nlp/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/nlp/v2/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/nlp/v2/region"
)

// func main() {
// 	fmt.Println("请输入需要翻译的文字：")
// 	var text string
// 	fmt.Scan(&text)
// 	result := GetNLP(text)
// 	fmt.Println(result)
// }

func GetNLP(text string) string {
	rep := NLP(text)
	result := *rep.TranslatedText
	return result
}

func NLP(text string) *model.RunTextTranslationResponse {

	// The AK and SK used for authentication are hard-coded or stored in plaintext, which has great security risks. It is recommended that the AK and SK be stored in ciphertext in configuration files or environment variables and decrypted during use to ensure security.
	// In this example, AK and SK are stored in environment variables for authentication. Before running this example, set environment variables CLOUD_SDK_AK and CLOUD_SDK_SK in the local environment
	ak := "W49XRXFNJKJ32BYSEEUL"
	sk := "MAGbgWpEHRjdKMsRGIa08SK4Ti0KzEDMpQ4iQXX1"

	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := nlp.NewNlpClient(
		nlp.NlpClientBuilder().
			WithRegion(region.ValueOf("cn-north-4")).
			WithCredential(auth).
			Build())

	from := model.GetTextTranslationReqFromEnum()
	to := model.GetTextTranslationReqToEnum()
	request := &model.RunTextTranslationRequest{}
	request.Body = &model.TextTranslationReq{
		Text: text,
		From: from.AUTO,
		To:   to.ZH,
	}
	response, err := client.RunTextTranslation(request)
	if err == nil {
		// fmt.Printf("%+v\n", response)
		fmt.Println("respnose: ok")
	} else {
		fmt.Println(err)
	}
	return response
}
