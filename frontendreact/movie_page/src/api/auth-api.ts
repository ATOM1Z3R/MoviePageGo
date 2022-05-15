import { CreateUserModel, UserModel, LoginModel } from "../model/user";

export class AuthApi {
    apiAddress = 'http://localhost:8080/users/';

    async loginApiCallAsync(login: LoginModel): Promise<UserModel>{
        const request = {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify(login),
        };
        
        const jsonString = await fetch(`${this.apiAddress}login`, request).then(async (response) => {
            const body = (await response.body?.getReader().read())?.value
            return new TextDecoder().decode(body);
        });
        const user: UserModel = JSON.parse(jsonString);
        return user;
    }

    async signupApiCallAsync(user: CreateUserModel): Promise<UserModel>{
        const request = {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify(user),
        };
        
        const jsonString = await fetch(`${this.apiAddress}signup`, request).then(async (response) => {
            const body = (await response.body?.getReader().read())?.value
            return new TextDecoder().decode(body);
        });
        return <UserModel>JSON.parse(jsonString);
    }
}