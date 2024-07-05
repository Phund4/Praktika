import axios from "axios"

export async function getVacancies() {
    const url = "http://localhost:8080/vacancies"
    return axios.get(url).then(resp => {
        return resp.data;
    })
}