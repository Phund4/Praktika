package internal

import (
	"strings"
	"fmt"
)

type skill struct {
	Name string `json:"name"`
}

type metroStation struct {
	LineName    string  `json:"line_name"`
	StationName string  `json:"station_name"`
}

type vacancies struct {
	Vacancies []vacancy `json:"items"`
}

type vacancy struct {
	Address                 struct {
		Building      string  `json:"building"`
		City          string  `json:"city"`
		Description   string  `json:"description"`
		MetroStations []metroStation `json:"metro_stations"`
		Street string `json:"street"`
	} `json:"address"`
	Area              struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"area"`
	Contacts              struct {
		Email  string `json:"email"`
		Name   string `json:"name"`
		Phones []struct {
			City    string `json:"city"`
			Country string `json:"country"`
			Number  string `json:"number"`
		} `json:"phones"`
	} `json:"contacts"`
	CreatedAt  string `json:"created_at"`
	Description        string `json:"description"`
	Employment struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"employment"`
	Experience struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"experience"`
	ID               string `json:"id"`
	KeySkills []skill `json:"key_skills"`
	Languages []struct {
		ID    string `json:"id"`
		Name string `json:"name"`
	} `json:"languages"`
	Name              string `json:"name"`
	ProfessionalRoles []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"professional_roles"`
	Salary                 struct {
		Currency string `json:"currency"`
		From     int    `json:"from"`
		To       any    `json:"to"`
	} `json:"salary"`
	Schedule struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"schedule"`
}

func (v *vacancies) getInsertString() string {
	sb := strings.Builder{};
	sb.WriteString(`insert into "vacancies" values `);

	for _, vacancy := range v.Vacancies {
		metroLineNameStr := ``
		for _, metro := range vacancy.Address.MetroStations {
			metroLineNameStr += metro.LineName + " ";
		}
		metroLineNameStr, _ = strings.CutSuffix(metroLineNameStr, " ");

		metroStationNameStr := ``
		for _, metro := range vacancy.Address.MetroStations {
			metroStationNameStr += metro.StationName + " ";
		}
		metroStationNameStr, _ = strings.CutSuffix(metroStationNameStr, " ");

		phoneStr := ``;
		for _, phone := range vacancy.Contacts.Phones {
			phoneStr += phone.Country + phone.City + phone.Number + " ";
		}
		phoneStr, _ = strings.CutSuffix(phoneStr, " ");

		keySkillsStr := ``;
		for _, key := range vacancy.KeySkills {
			keySkillsStr += key.Name + " "
		}
		keySkillsStr, _ = strings.CutSuffix(keySkillsStr, " ");

		languagesStr := ``;
		for _, key := range vacancy.Languages {
			languagesStr += key.Name + " "
		}
		languagesStr, _ = strings.CutSuffix(languagesStr, " ");

		rolesStr := ``;
		for _, key := range vacancy.ProfessionalRoles {
			rolesStr += key.Name + " "
		}
		rolesStr, _ = strings.CutSuffix(rolesStr, " ");
		
		vacancyStr := fmt.Sprintf(`('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v'),`, 
		vacancy.Address.Building, vacancy.Address.City, vacancy.Address.Description, metroLineNameStr, metroStationNameStr, 
		vacancy.Address.Street, vacancy.Area.Name, vacancy.Contacts.Email, vacancy.Contacts.Name, phoneStr, vacancy.CreatedAt, 
		vacancy.Description, vacancy.Employment.Name, vacancy.Experience.Name, vacancy.ID, keySkillsStr, languagesStr, vacancy.Name, 
		rolesStr, vacancy.Salary.Currency, vacancy.Salary.From, vacancy.Salary.To, vacancy.Schedule.Name)

		sb.WriteString(vacancyStr);
	}

	insertStr, _ := strings.CutSuffix(sb.String(), ",");
	return insertStr
}