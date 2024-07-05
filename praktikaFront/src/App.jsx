import { useEffect, useState } from "react";
import { getVacancies } from "../api/vacancies";
import { Vacancy } from "./Vacancy";
import { SearchBar } from "./SearchBar";
import { SelectBox } from "./ SelectBox";

function App() {
    const [searchTerm, setSearchTerm] = useState("");
    const [filterEmployment, setFilterEmployment] = useState("");
    const [filterExperience, setFilterExperience] = useState("");
    const [filterSchedule, setFilterSchedule] = useState("");
    let [vacancies, setVacancies] = useState([]);

    useEffect(() => {
        getVacancies().then((resp) => {
            setVacancies(resp.items);
            console.log(resp.items);
        });
    }, []);

    const filteredVacancyPostings = vacancies.filter((vac) => {
        const matchesSearchTerm = vac.name
            .toLowerCase()
            .includes(searchTerm.toLowerCase());
        const matchesEmployment =
            filterEmployment === "" ||
            vac.employment.name.includes(filterEmployment);
        const matchesExperience = 
            filterExperience === "" ||
            vac.experience.name.includes(filterExperience)
        const matchesSchedule = 
            filterSchedule === "" ||
            vac.schedule.name.includes(filterSchedule)
        return matchesSearchTerm && matchesEmployment && matchesExperience && matchesSchedule;
    });

    return (
        <>
            <SearchBar searchTerm={searchTerm} setSearchTerm={setSearchTerm} />
            <SelectBox
                filter={filterEmployment}
                setFilter={setFilterEmployment}
                standartValue={"Выберите занятость"}
                valueArr={["Полная занятость", "Частичная занятость"]}
            />
            <SelectBox
                filter={filterExperience}
                setFilter={setFilterExperience}
                standartValue={"Выберите опыт работы"}
                valueArr={["Нет опыта", "От 1 года до 3 лет", "От 3 до 6 лет", "Более 6 лет"]}
            />
             <SelectBox
                filter={filterSchedule}
                setFilter={setFilterSchedule}
                standartValue={"Выберите график работы"}
                valueArr={["Удаленная работа", "Полный день", "Гибкий график", "Сменный график", "Стажировка"]}
            />
            <div className="vacancies">
                {filteredVacancyPostings.map((el, index) => (
                    <Vacancy
                        key={index}
                        name={el.name}
                        employment={el.employment}
                        experience={el.experience}
                        schedule={el.schedule}
                        contacts={el.contacts}
                        salary={el.salary}
                        address={el.address}
                        area={el.area}
                        createdAt={el.created_at}
                        skills={el.key_skills}
                    />
                ))}
            </div>
        </>
    );
}

export default App;
