package models

type Share struct {
	Id     int
	Name   string
	Symbol string
}

type ShareRepo struct {
	Entities []Share
}

func (repo *ShareRepo) Initialize() {
	entities := make([]Share, 2)
	nokia := Share{Id: 1, Name: "Nokia", Symbol: "NOK"}
	ahtium := Share{Id: 2, Name: "Ahtium", Symbol: "AHT"}
	entities[0] = nokia
	entities[1] = ahtium
	repo.Entities = entities
}

func (repo *ShareRepo) GetOne(id int) interface{} {
	return repo.Entities[id-1]
}

func (repo *ShareRepo) GetAll() []interface{} {
	entities := make([]interface{}, len(repo.Entities))
	for i, _ := range repo.Entities {
		entities[i] = repo.Entities[i]
	}
	return entities
}
