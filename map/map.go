package _map

var actions = []struct {
	name     string
	category string
	aType    string
}{
	{name: "create user", category: "zero", aType: "read"},
	{name: "edit user", category: "zero", aType: "update"},
	{name: "get account", category: "account", aType: "read"},
	{name: "create account", category: "account", aType: "create"},
	{name: "delete disk", category: "neonsan", aType: "delete"},
	{name: "read disk", category: "neonsan", aType: "read"},
}

var ActionMap = make(map[string]map[string]string)

func InitActions() error {
	for _, action := range actions {
		if _, ok := ActionMap[action.category]; !ok {
			ActionMap[action.category] = make(map[string]string)
		}
		ActionMap[action.category][action.name] = action.aType
	}
	return nil
}
