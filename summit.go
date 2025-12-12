package main

// Summit stores information about a summit
type Summit struct {
	Name        string
	Height      string
	Category    string
	Country     string
	Region		string
	Group		string
	Information string
}


func (s Summit) toSlice() []string {
	var result []string
	result = make([]string, 0)
	result = append(result, s.Name)
	result = append(result, s.Height)
	result = append(result, s.Category)
	result = append(result, s.Country)
	result = append(result, s.Region)
	result = append(result, s.Group)
	result = append(result, s.Information)
	return result
}