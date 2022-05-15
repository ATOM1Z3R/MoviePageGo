import { ErrorModel } from "./error";

export interface UserModel extends ErrorModel {
    id: number;
    userName: string | null;
    firstName: string | null;
    lastName: string | null;
    email: string | null;
    password: string | null;
    token: string | null;
    refreshToken: string | null;
    createdDate: string | null;
}

export interface CreateUserModel {
    userName: string;
    firstName: string;
    lastName: string;
    email: string;
    password: string;
}

export interface AuthorModel extends ErrorModel {
    id: number;
    userName: string | null;
    firstName: string | null;
    lastName: string | null;
    email: string | null;
}

export interface LoginModel {
    email: string | null;
    password: string | null;
}