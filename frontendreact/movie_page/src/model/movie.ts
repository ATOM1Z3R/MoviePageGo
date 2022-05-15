import { ReadCommentModel } from "./comment";
import { DirectorModel } from "./director";
import { GenreModel } from "./genre";
import { UserModel } from "./user";

export interface MovieModel {
    id: number | null,
    title: string | null,
    premiereDate: string | null,
    description: string | null,
    location: string | null,
    user: UserModel | null,
    genre: GenreModel | null,
    comments: Array<ReadCommentModel> | null,
    directors: Array<DirectorModel> | null,
}

export interface SimpleMovieModel {
    id: number | null,
    title: string | null,
    premiereDate: string | null,
    description: string | null,
    location: string | null,
}

export interface CreateMovieModel {
    title: string | null,
    premiereDate: string | null,
    description: string | null,
    location: string | null,
    user: UserModel | null,
    genre: GenreModel | null,
    directors: Array<DirectorModel> | null,
}