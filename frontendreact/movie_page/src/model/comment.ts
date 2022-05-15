import { MovieModel } from "./movie";
import { UserModel } from "./user";

export interface CreateCommentModel {
    name: string | null,
    author: UserModel | null,
    movie: MovieModel | null,
    rating: number | null,
}

export interface ReadCommentModel {
    id: number | null,
    name: string | null,
    author: UserModel | null,
    createdDate: string | null,
    rating: number | null,
}