package utils

import (
	"go-core/config"
	"regexp"
	"strconv"
	"strings"
)

func ToSlug(str string) string {
	str = strings.ToLower(str)

	r := regexp.MustCompile("á|à|ả|ạ|ã|ă|ắ|ằ|ẳ|ẵ|ặ|â|ấ|ầ|ẩ|ẫ|ậ")
	str = r.ReplaceAllString(str, "a")

	r = regexp.MustCompile("é|è|ẻ|ẽ|ẹ|ê|ế|ề|ể|ễ|ệ")
	str = r.ReplaceAllString(str, "e")

	r = regexp.MustCompile("i|í|ì|ỉ|ĩ|ị")
	str = r.ReplaceAllString(str, "i")

	r = regexp.MustCompile("ó|ò|ỏ|õ|ọ|ô|ố|ồ|ổ|ỗ|ộ|ơ|ớ|ờ|ở|ỡ|ợ")
	str = r.ReplaceAllString(str, "o")

	r = regexp.MustCompile("ú|ù|ủ|ũ|ụ|ư|ứ|ừ|ử|ữ|ự")
	str = r.ReplaceAllString(str, "u")

	r = regexp.MustCompile("ý|ỳ|ỷ|ỹ|ỵ")
	str = r.ReplaceAllString(str, "y")

	r = regexp.MustCompile("đ")
	str = r.ReplaceAllString(str, "d")

	r = regexp.MustCompile(`[^\w]`)
	str = r.ReplaceAllString(str, " ")

	r = regexp.MustCompile(`[\s]+`)
	str = r.ReplaceAllString(str, "-")

	str = strings.Trim(str, "-")

	return str
}

func ParseUint(strNum string) uint64 {
	num, err := strconv.ParseUint(strNum, 10, 64)
	if err != nil {
		return 0
	}

	return num
}

func ParseUInt(value string) uint {
	num := ParseInt(value)

	return uint(num)
}

func ParseInt(strNum string) int {
	num, err := strconv.Atoi(strNum)
	if err != nil {
		return 0
	}

	return num
}

func DeleteRedisKey(str string) error {
	redisClient := config.RedisClient
	pipeLine := redisClient.Pipeline()
	pipeLine.Del(str)
	_, err := pipeLine.Exec()
	LogError(str, err)
	return err
}

func ParsePageSize(page int, size int) (int, int) {
	if page <= 0 {
		page = 1
	}

	switch {
	case size > 100:
		size = 100
	case size <= 0:
		size = 20
	}

	return page, size
}
