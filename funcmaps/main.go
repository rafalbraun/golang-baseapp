package funcmaps

import (
	"baseapp/models"
	"baseapp/system"
	"time"
)

var FuncMap = map[string]interface{}{
	"MapPageData":                  MapPageData,
    "FormatDate":                   FormatDate,
    "MemberSinceDays":              MemberSinceDays,
    "mul":                          Mul,
    "min":                          Min,
    "max":                          Max,
    "abs":                          Abs,
    "intRange":                     IntRange,
}

func MapPageData(pd models.PageData, inner interface{}) map[string]interface{} {
	s := make(map[string]interface{}, 0)
	s["inner"] = inner
	s["pageData"] = pd
	s["loggedIn"] = pd.LoggedIn
	s["isAdmin"] = system.IsAdmin(pd.LoggedIn)
	return s
}

func Mul(fst, sec int) int {
	return fst * sec
}

func Min(fst, sec int) int {
	if fst < sec {
		return fst
	}
	return sec
}

func Max(fst, sec int) int {
	if fst < sec {
		return sec
	}
	return fst
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func removeDuplicates(slice []int) []int {
	var flag bool = false
	list := []int{}
	for _, entry := range slice {
		if (entry == 0 && flag == false) {
			list = append(list, entry)
			flag = true
		} else
		if (entry == 0 && flag == true) {
			// nothing
			flag = true
		} else
		if (entry != 0) {
			list = append(list, entry)
			flag = false
		}
	}
	return list
}

func IntRange(start, current, end int) []int {
	n := end - start + 1
	result := make([]int, n)
	for i := 0; i < n; i++ {
		if (Abs(i+1-start)<2 || Abs(i+1-current)<2 || Abs(end-i-1)<2) {
			result[i] = start + i
		} else {
			result[i] = 0
		}
	}
	return removeDuplicates(result)
}

func FormatDate(date *time.Time) string {
    if (date != nil) {
        return date.Format("Jan 02, 2006 15:04:05 UTC")
    }
    return ""
}

func MemberSinceDays(user models.User) int {
    now := time.Now()
    activatedAt := user.ActivatedAt
    if (activatedAt != nil) {
        return int(now.Sub(*activatedAt).Hours() / 24)
    }
    return 0
}