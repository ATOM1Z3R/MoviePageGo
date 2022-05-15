import { GenreModel } from "../model/genre";

export class GenreApi {
    apiAddress = 'http://localhost:8080/genres/';

    async getAllGenresApiCallAsync(): Promise<GenreModel[]> {
        const request = {
            method: 'GET',
            headers: {'Content-Type': 'application/json'},
        };
        
        const jsonString = await fetch(`${this.apiAddress}login`, request).then(async (response) => {
            const body = (await response.body?.getReader().read())?.value
            return new TextDecoder().decode(body);
        });
        const jsonObj = JSON.parse(jsonString);
        return jsonObj.map((x: GenreModel)=><GenreModel>x);
    }
}