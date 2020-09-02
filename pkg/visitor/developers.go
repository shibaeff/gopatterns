package visitor

import "fmt"

const (
	juniorClusterMessage = "sent cluster to downtime"
	juniorDBMessage      = "dropping db"
	juniorTestsMessage   = "deleting senior's tests"

	seniorClusterMessage = "uptiming cluster"
	seniorDBMessage      = "rerolling db"
	seniorTestsMessage   = "writing more tests"
)

// Developer interface provides actions to be performed on project items
type Developer interface {
	SetupCluster()
	SetupDB()
	WriteTests()
}

type junior struct{}

func (j *junior) SetupCluster() {
	fmt.Println(juniorClusterMessage)
}

func (j *junior) SetupDB() {
	fmt.Println(juniorDBMessage)
}

func (j *junior) WriteTests() {
	fmt.Println(juniorTestsMessage)
}

type senior struct{}

func (s *senior) SetupCluster() {
	fmt.Println(seniorClusterMessage)
}

func (s *senior) SetupDB() {
	fmt.Println(seniorDBMessage)
}

func (s *senior) WriteTests() {
	fmt.Println(seniorTestsMessage)
}

func NewDeveloper(skilled bool) Developer {
	if skilled {
		return &senior{}
	}
	return &junior{}
}
