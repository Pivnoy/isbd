package service

import (
	"github.com/Pivnoy/isbd/models"
	"github.com/Pivnoy/isbd/server"
)

func GetCountryReptiloids(country_name string) (models.RepriloidResponseCountry, error) {
	var tvAll, tvReptiloid int
	rows, err := server.DbInstance.Query("select a.human_id from tv_work inner join anchorman a on tv_work.anchorman_id = a.id inner join human h on h.id = a.human_id where country_name = $1 group by a.human_id", country_name)
	if err != nil {
		panic("Error in going to tv work table 1")
	}
	tvHumAllList, err := models.CreatePeopleInSphereCollectionCollection(rows)
	if err != nil {
		return models.RepriloidResponseCountry{}, err
	}
	tvAll = len(tvHumAllList) // всего челвоек работающих на телевидиние

	rows, err = server.DbInstance.Query("select a.human_id from tv_work inner join anchorman a on tv_work.anchorman_id = a.id inner join reptiloid r on a.human_id = r.human_id inner join human h on h.id = a.human_id where country_name = $1 group by a.human_id", country_name)
	if err != nil {
		panic("Error in going to tv work table 2")
	}
	tvHumRepList, err := models.CreatePeopleInSphereCollectionCollection(rows)
	if err != nil {
		return models.RepriloidResponseCountry{}, err
	}
	tvReptiloid = len(tvHumRepList) // рептилоидов на телевединии

	var buisnessAll, buisnessRepriloid int

	rows, err = server.DbInstance.Query("select b.human_id from businessman_work inner join businessman b on b.id = businessman_work.businessman_id inner join human h on h.id = b.human_id where country_name = $1 group by human_id", country_name)
	if err != nil {
		panic("Error in going to tv work table 1")
	}
	buisnessAllList, err := models.CreatePeopleInSphereCollectionCollection(rows)
	if err != nil {
		return models.RepriloidResponseCountry{}, err
	}
	buisnessAll = len(buisnessAllList) // сумма людей бизнессменов в бизнессе

	rows, err = server.DbInstance.Query("select b.human_id from businessman_work inner join businessman b on b.id = businessman_work.businessman_id inner join reptiloid r on b.human_id = r.human_id inner join human h on h.id = b.human_id where country_name = $1 group by b.human_id", country_name)
	if err != nil {
		panic("Error in going to tv work table 2")
	}
	buisnessRepList, err := models.CreatePeopleInSphereCollectionCollection(rows)
	if err != nil {
		return models.RepriloidResponseCountry{}, err
	}
	buisnessRepriloid = len(buisnessRepList) // Сумма алигаторов в бизнесе

	var scienceAll, scienceReptil int
	rows, err = server.DbInstance.Query("select s.human_id from science_work inner join scientist s on s.id = science_work.scientist_id inner join human h on h.id = s.human_id where country_name = $1 group by human_id", country_name)
	if err != nil {
		panic("Error in going to tv work table 1")
	}
	scienceAllList, err := models.CreatePeopleInSphereCollectionCollection(rows)
	if err != nil {
		return models.RepriloidResponseCountry{}, err
	}
	scienceAll = len(scienceAllList) // сумма ученых людей в науке

	rows, err = server.DbInstance.Query("select s.human_id from science_work inner join scientist s on s.id = science_work.scientist_id inner join reptiloid r on s.human_id = r.human_id inner join human h on h.id = s.human_id where country_name = $1 group by s.human_id", country_name)
	if err != nil {
		panic("Error in going to tv work table 2")
	}
	scienceReptilList, err := models.CreatePeopleInSphereCollectionCollection(rows)
	if err != nil {
		return models.RepriloidResponseCountry{}, err
	}
	scienceReptil = len(scienceReptilList) // сумма алигаторв в науке

	return models.RepriloidResponseCountry{Tv: float32(tvReptiloid) / float32(tvAll), Business: float32(buisnessRepriloid) / float32(buisnessAll), Science: float32(scienceReptil) / float32(scienceAll)}, nil
}


