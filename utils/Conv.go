package utils

import "strconv"

func ParseFloat(valor string) (vRet float64) {

	if valor != "" {
		v, err := strconv.ParseFloat(valor, 64)
		if err != nil {
			panic(err.Error())
		}
		vRet = v
	} else {
		vRet = 0.0
	}

	return
}

func ParseInt(valor string) (vRet int64) {

	if valor != "" {
		v, err := strconv.ParseInt(valor, 10, 64)
		if err != nil {
			panic(err.Error())
		}
		vRet = v
	} else {
		vRet = 0
	}

	return
}
