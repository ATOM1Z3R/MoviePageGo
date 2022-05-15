import { FormEvent } from "react";
import { AuthApi } from "../api/auth-api";
import { UserModel } from "../model/user";


const authApi = new AuthApi();

const LoginPage = () => {
    const handleLogin =(e: FormEvent) => {
        e.preventDefault();

        const data = new FormData(e.currentTarget);
        const values = Object.fromEntries(data.entries());

        authApi.loginApiCallAsync({
            email: values["email"].toString(),
            password: values["password"].toString(),
        }).then((res: UserModel) => {
            console.log(res.token);
        });

    };

    return (
        <div className="flex align-middle justify-center">
            <div className="px-14 py-4 w-[30rem] rounded-md bg-gray-800 my-20">
                <div className="justify-center">
                    <div>
                        <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-300">Log in</h2>
                    </div>
                    <form onSubmit={handleLogin} className="space-y-8 py-8 text-gray-300">
                        <div>
                            <label className="block pb-1 font-semibold text-1xl">
                                Email address
                                <input type="email" name="email" autoComplete="email"
                                    className="w-full rounded-sm px-3 block border text-gray-300 focus:outline-none 
                                    focus:ring-green-600 focus:border-green-600 focus:text-green-600 focus:z-10 bg-gray-800" />
                            </label>
                        </div>
                        <div>
                            <label className="block pb-1 font-semibold text-1xl">Password</label>
                            <input type="password" name="password" autoComplete="current-password"
                                className="w-full rounded-sm px-3 block border text-gray-300 focus:outline-none focus:ring-green-600 focus:border-green-600 
                                 focus:text-green-600 focus:z-10 bg-gray-800" />
                        </div>
                        <div>
                            <button type="submit" className="w-full flex justify-center py-3 border border-transparent text-sm font-medium rounded-sm
                                text-gray-300 bg-green-600 hover:bg-green-700 focus:ring-2 focus:ring-offset-2 focus:outline-none focus:ring-green-600">
                                Sign up
                            </button>
                        </div>
                        <div className="w-full justify-center">
                            <label className="block">Sign up</label>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    );
}

export default LoginPage;