import { CreateMovieModel, MovieModel, SimpleMovieModel } from "../model/movie";

export class MovieApi {
    apiAddress = 'http://localhost:8080/';

    async getAllMoviesApiCallAsync(): Promise<SimpleMovieModel[]> {
        const request = {
            method: 'GET',
            headers: {'Content-Type': 'application/json'},
        };
        
        const jsonString = await fetch(`${this.apiAddress}movies/`, request).then(async (response) => {
            const body = (await response.body?.getReader().read())?.value
            return new TextDecoder().decode(body);
        });
        const jsonObj = JSON.parse(jsonString);
        return jsonObj.map((x: SimpleMovieModel)=><SimpleMovieModel>x);
    }

    async getMovieByIdApiCallAsync(id: number): Promise<MovieModel> {
        const request = {
            method: 'GET',
            headers: {'Content-Type': 'application/json'},
        };
        
        const jsonString = await fetch(`${this.apiAddress}movie/${id}`, request).then(async (response) => {
            const body = (await response.body?.getReader().read())?.value
            return new TextDecoder().decode(body);
        });

        return <MovieModel>JSON.parse(jsonString);
    }

    async getMoviesByGenreIdApiCallAsync(genreId: number): Promise<SimpleMovieModel[]> {
        const request = {
            method: 'GET',
            headers: {'Content-Type': 'application/json'},
        };
        
        const jsonString = await fetch(`${this.apiAddress}movie/genre/${genreId}`, request).then(async (response) => {
            const body = (await response.body?.getReader().read())?.value
            return new TextDecoder().decode(body);
        });
        const jsonObj = JSON.parse(jsonString);
        return jsonObj.map((x: SimpleMovieModel)=><SimpleMovieModel>x);
    }

    async createMovieApiCallAsync(movie: CreateMovieModel, token: string): Promise<MovieModel> {
        const request = {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'token': token,
            },
            body: JSON.stringify(movie),
        };
        
        const jsonString = await fetch(`${this.apiAddress}movie/add`, request).then(async (response) => {
            const body = (await response.body?.getReader().read())?.value
            return new TextDecoder().decode(body);
        });

        return <MovieModel>JSON.parse(jsonString);
    }
}