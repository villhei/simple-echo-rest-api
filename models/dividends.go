package models

type Dividend struct {
	Id    int
	Value int
}

type DividendRepo struct {
	Entities []Dividend
}

func (repo *DividendRepo) Initialize() {
	entities := make([]Dividend, 2)
	nokia := Dividend{Id: 1, Value: 10}
	ahtium := Dividend{Id: 2, Value: 20}
	entities[0] = nokia
	entities[1] = ahtium
	repo.Entities = entities
}

func (repo *DividendRepo) GetOne(id int) interface{} {
	return repo.Entities[id-1]
}

func (repo *DividendRepo) GetAll() []interface{} {
	entities := make([]interface{}, len(repo.Entities))
	for i, _ := range repo.Entities {
		entities[i] = repo.Entities[i]
	}
	return entities
}
