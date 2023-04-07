package checkFinger

import (
	"Musscan/jsonRead"
	"Musscan/requests"
	"fmt"
)

func Tocheck(url string, fingerArr jsonRead.CmsFeature) { //思路是fofa里的rule逐个比对，抄的o2tukuxx的，k那里确实写的简单，避免了多个rule讨论
	data := requests.Requsets(url)
	body := data.Body
	//Rule是二位数组，如果一排存在多个条件都满足，或者满足其他行的条件
	for i := 0; i < len(fingerArr); i++ { //每个指纹的循环
		Product := fingerArr[i].Product
		Company := fingerArr[i].Company
		for j := 0; j < len(fingerArr[i].Rules); j++ { //[]大规则或关系

			flag := 0
			for k := 0; k < len(fingerArr[i].Rules[j]); k++ { //{}小规则且关系

				if requests.Check(fingerArr[i].Rules[j][k].Match, fingerArr[i].Rules[j][k].Content, body, data) {
					flag++

				}
			}
			if len(fingerArr[i].Rules[j]) == flag {
				fmt.Println("Poduct："+Product, "Company:"+Company)
			}
		}
	}
}
