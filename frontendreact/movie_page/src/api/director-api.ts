import { DirectorModel, SimpleDirectorModel } from "../model/director";

export class DirectorApi {
    apiAddress = 'http://localhost:8080/';

    async getAllDirectorsApiCallAsync(): Promise<SimpleDirectorModel[]> {
        const request = {
            method: 'GET',
            headers: {'Content-Type': 'application/json'},
        };
        
        const jsonString = await fetch(`${this.apiAddress}directors`, request).then(async (response) => {
            const body = (await response.body?.getReader().read())?.value
            return new TextDecoder().decode(body);
        });
        const jsonObj = JSON.parse(jsonString);
        return jsonObj.map((x: SimpleDirectorModel)=><SimpleDirectorModel>x);
    }

    async getDirectorByIdApiCallAsync(id: number): Promise<DirectorModel> {
        const request = {
            method: 'GET',
            headers: {'Content-Type': 'application/json'},
        };
        
        const jsonString = await fetch(`${this.apiAddress}director/${id}`, request).then(async (response) => {
            const body = (await response.body?.getReader().read())?.value
            return new TextDecoder().decode(body);
        });

        return <DirectorModel>JSON.parse(jsonString);
    }
}