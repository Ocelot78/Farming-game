package state

type Shop struct {
	Tractors	[]Tractor
	Harvesters	[]Harvester
}
type Tractor struct {
	ID			int
	Name		string
	PowerIndex	int
	Price		int
}
type Harvester struct {
	ID				int
	Name			string
	PowerIndex		int
	WorkingWidth	int
	Price			int
}
