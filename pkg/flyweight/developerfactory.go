package flyweight

const (
	golangDev = "golangDeveloper"
	pascalDev = "pascalDeveloper"
)

type DeveloperFactory interface {
	GetDeveloper(devType string) Developer
}

type factory struct {
	developersMap map[string]Developer
}

func (f *factory) GetDeveloper(devType string) Developer {
	return f.developersMap[devType]
}

func NewDeveloperFactory() DeveloperFactory {
	return &factory{
		developersMap: map[string]Developer{
			golangDev: &GolangDeveloper{},
			pascalDev: &PascalDeveloper{},
		}}
}
