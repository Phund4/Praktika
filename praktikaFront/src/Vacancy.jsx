import React from "react";
import "./styles/vacancy.css";

export function Vacancy({name, description, employment, experience, schedule, createdAt, contacts, salary, area, address, skills, }) {
    return (
        <div className="job-posting">
            <h1 className="job-title">{name}</h1>
            <div className="job-details">
                <h2>Описание вакансии:</h2>
                <p>
                    Создано: {createdAt}
                    {description}
                </p>

                <h2>Обязанности:</h2>
                <ul>
                    {skills.length == 1 && skills[0].name == "" ? "Обязанности не указаны" : skills.map(el => <li key={createdAt + "-skill"}>{el.name}</li>)}
                </ul>

                <h2>Условия:</h2>
                <ul>
                    <li>{employment.name}</li>
                    <li>{experience.name}</li>
                    <li>{schedule.name == "" ? "График не указан" : schedule.name}</li>
                    <li>Зарплата: {salary.from != "<nil>" ? `${salary.from}${salary.currency}` : ""} -&nbsp;
                        {salary.to != "<nil>" ? `${salary.to}${salary.currency}` : ""}</li>
                </ul>

                <h2>Контакты:</h2>
                <p>Email: {contacts.email}</p>
                <p>Имя: {contacts.name}</p>
                <p>Адрес: {area.name}, {address.city}, {address.street}, {address.building}</p>
                <p>Метро: {address.metro_stations.map(el => `${el.station_name} - ${el.line_name}; `)}</p>
            </div>
        </div>
    );
}