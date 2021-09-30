package config

import "time"

func parseTimeStr(s string) int64 {
	l := len(s)
	multiple1 := 0.0
	multiple2 := 0.0
	i := 0
	for ; i < l-1; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			multiple1 *= 10
			multiple1 += (float64)(s[i] - '0')
		}
	}
	if i == l-1 {
		if s[i] == 's' {
			return int64(multiple1 * (float64(time.Second)))
		}
		if s[i] == 'm' {
			return int64(multiple1 * (float64(time.Minute)))
		}
		if s[i] == 'h' {
			return int64(multiple1 * (float64(time.Hour)))
		}
		panic("时间格式非法")
	}
	if s[i] == '.' {
		i++
		count := 0
		for ; i < l; i++ {
			if s[i] > '0' && s[i] < '9' {
				count++
				temp := 0.0
				for j := 1; j <= count; j++ {
					temp = (float64)(s[i]-'0') * 0.1
				}
				multiple2 += temp
			}
		}
	}
	if i == l {
		panic("时间格式非法")
	}
	multiple := multiple1 + multiple2
	if s[i] == 's' {
		return int64(multiple * (float64(time.Second)))
	}
	if s[i] == 'm' {
		return int64(multiple * (float64(time.Minute)))
	}
	if s[i] == 'h' {
		return int64(multiple * (float64(time.Hour)))
	}
	panic("时间格式非法")
}
