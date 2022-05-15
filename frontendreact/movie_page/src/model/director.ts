import { MovieModel } from "./movie";

export interface DirectorModel {
    id: string | null,
    firstName: string | null,
    lastName: string | null,
    birthDate: string | null,
    birthCountry: string | null,
    movies: Array<MovieModel> | null,
}

export interface SimpleDirectorModel {
    id: string | null,
    firstName: string | null,
    lastName: string | null,
    birthDate: string | null,
    birthCountry: string | null,
}