package main

func DefaultForEmplyString(s, defaultString string) string{
	if len(s) == 0{
		return defaultString
	}
	return s
}

func YesNo(b bool) string{
	if b{
		return "Да"
	}
	return "Нет"
}