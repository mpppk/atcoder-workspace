package lib

func ToYesNo(yes bool) string {
	return TernaryOPString(yes, "Yes", "No")
}
