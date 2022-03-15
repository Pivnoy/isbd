package service

import (
	"github.com/Pivnoy/isbd/models"
	"github.com/Pivnoy/isbd/server"
)

func GetAllReptiloids() (models.ReptiloidsResponseAll, error) {
	var tvAll, tvReptiloid int
	rows, err := server.DbInstance.Query("select a.human_id from tv_work inner join anchorman a on tv_work.anchorman_id = a.id group by human_id")
	if err != nil {
		panic("Error in going to tv work table 1")
	}
	tvHumAllList, err := models.CreatePeopleInSphereCollectionCollection(rows)
	if err != nil {
		return models.ReptiloidsResponseAll{}, err
	}
	tvAll = len(tvHumAllList) // всего челвоек работающих на телевидиние

	rows, err = server.DbInstance.Query("select a.human_id from tv_work inner join anchorman a on tv_work.anchorman_id = a.id inner join reptiloid r on a.human_id = r.human_id group by a.human_id")
	if err != nil {
		panic("Error in going to tv work table 2")
	}
	tvHumRepList, err := models.CreatePeopleInSphereCollectionCollection(rows)
	if err != nil {
		return models.ReptiloidsResponseAll{}, err
	}
	tvReptiloid = len(tvHumRepList) // рептилоидов на телевединии

	var buisnessAll, buisnessRepriloid int

	rows, err = server.DbInstance.Query("select b.human_id from businessman_work inner join businessman b on b.id = businessman_work.businessman_id group by human_id")
	if err != nil {
		panic("Error in going to tv work table 1")
	}
	buisnessAllList, err := models.CreatePeopleInSphereCollectionCollection(rows)
	if err != nil {
		return models.ReptiloidsResponseAll{}, err
	}
	buisnessAll = len(buisnessAllList) // сумма людей бизнессменов в бизнессе

	rows, err = server.DbInstance.Query("select b.human_id from businessman_work inner join businessman b on b.id = businessman_work.businessman_id inner join reptiloid r on b.human_id = r.human_id group by b.human_id")
	if err != nil {
		panic("Error in going to tv work table 2")
	}
	buisnessRepList, err := models.CreatePeopleInSphereCollectionCollection(rows)
	if err != nil {
		return models.ReptiloidsResponseAll{}, err
	}
	buisnessRepriloid = len(buisnessRepList) // Сумма алигаторов в бизнесе

	var scienceAll, scienceReptil int
	rows, err = server.DbInstance.Query("select s.human_id from science_work inner join scientist s on s.id = science_work.scientist_id group by human_id")
	if err != nil {
		panic("Error in going to tv work table 1")
	}
	scienceAllList, err := models.CreatePeopleInSphereCollectionCollection(rows)
	if err != nil {
		return models.ReptiloidsResponseAll{}, err
	}
	scienceAll = len(scienceAllList) // сумма ученых людей в науке

	rows, err = server.DbInstance.Query("select s.human_id from science_work inner join scientist s on s.id = science_work.scientist_id inner join reptiloid r on s.human_id = r.human_id group by s.human_id")
	if err != nil {
		panic("Error in going to tv work table 2")
	}
	scienceReptilList, err := models.CreatePeopleInSphereCollectionCollection(rows)
	if err != nil {
		return models.ReptiloidsResponseAll{}, err
	}
	scienceReptil = len(scienceReptilList) // сумма алигаторв в науке

	// список названи стран где президент рептизоид
	rows, err = server.DbInstance.Query("select p.country_name from country inner join president p on country.name = p.country_name inner join reptiloid r on p.human_id = r.human_id group by p.country_name")
	if err != nil {
		panic("Error in going to tv work table 2")
	}
	countryNameList, err := models.CreateCountryReptiloidNameCollection(rows)
	if err != nil {
		return models.ReptiloidsResponseAll{}, err
	}

	// процент стран где рептилоид президент
	rows, err = server.DbInstance.Query("select country.id from country inner join president p on country.name = p.country_name")
	if err != nil {
		panic("Error in going to tv work table 2")
	}
	countryIdList, err := models.CreateCountryReptiloidIdCollection(rows)
	if err != nil {
		return models.ReptiloidsResponseAll{}, err
	}

	return models.ReptiloidsResponseAll{Tv: float32(tvReptiloid) / float32(tvAll), Business: float32(buisnessRepriloid) / float32(buisnessAll), Science: float32(scienceReptil) / float32(scienceAll), Countries: countryNameList, CountriesP: float32(len(countryNameList)) / float32(len(countryIdList))}, nil
}
