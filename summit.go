package main

// Summit stores information about a summit
type Summit struct {
	Name        string
	Height      string
	Kategorie   string
	Staat       string
	Region		string
	Gruppe		string
	Information string
}


func (s Summit) toSlice []strinh {
	var result []string
	result = make([]string, 8)
	result = append(result, s.Name)
	result = append(result, s.Height)
	result = append(result, s.Kategorie)
	result = append(result, s.Staat)
	result = append(result, s.Region)
	result = append(result, s.Gruppe)
	result = append(result, s.Information)
	return ressult
}