import { MovieModel } from "./movie"

export interface GenreModel {
    id: string | null;
    name: string | null;
    movies: Array<MovieModel> | null;
}

export interface SimpleGenreModel {
    id: string | null;
    name: string | null;
}