import { FormEvent } from "react";
import { AuthApi } from "../api/auth-api";
import { CreateUserModel } from "../model/user";

const authApi = new AuthApi();

const SignUpPage = () => {
    const handleSubmitForm = (e: FormEvent) => {
        e.preventDefault();

        const data = new FormData(e.currentTarget);
        const values = Object.fromEntries(data.entries());

        const user: CreateUserModel = {
            userName: values["userName"].toString(),
            firstName: values["firstName"].toString(),
            lastName: values["lastName"].toString(),
            email: values["email"].toString(),
            password: values["password"].toString()
        };

        authApi.signupApiCallAsync(user).then((res) =>{
            console.log(res);
        });
    }

    return (
        <div className="flex align-middle justify-center">
            <div className="px-14 py-4 w-[30rem] rounded-md bg-gray-800 my-20">
                <div className="justify-center">
                    <div>
                        <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-300">Sign up</h2>
                    </div>
                    <form onSubmit={handleSubmitForm} className="space-y-7 py-8 text-gray-300">
                        <div>
                            <label className="block pb-1 font-semibold text-1xl">
                                Username
                                <input type="text" name="userName" className="w-full rounded-sm px-3 block border text-gray-300 focus:outline-none 
                                    focus:ring-green-600 focus:border-green-600 focus:text-green-600 focus:z-10 bg-gray-800" />
                            </label>
                        </div>
                        <div>
                            <label className="block pb-1 font-semibold text-1xl">
                                Email address
                                <input type="email" name="email" className="w-full rounded-sm px-3 block border text-gray-300 focus:outline-none 
                                    focus:ring-green-600 focus:border-green-600 focus:text-green-600 focus:z-10 bg-gray-800" />
                            </label>
                        </div>
                        <div>
                            <label className="block pb-1 font-semibold text-1xl">
                                First name
                                <input type="text" name="firstName" className="w-full rounded-sm px-3 block border text-gray-300 focus:outline-none 
                                    focus:ring-green-600 focus:border-green-600 focus:text-green-600 focus:z-10 bg-gray-800" />
                            </label>
                        </div>
                        <div>
                            <label className="block pb-1 font-semibold text-1xl">
                                Last name
                                <input type="text" name="lastName" className="w-full rounded-sm px-3 block border text-gray-300 focus:outline-none 
                                    focus:ring-green-600 focus:border-green-600 focus:text-green-600 focus:z-10 bg-gray-800" />
                            </label>
                        </div>
                        <div>
                            <label className="block pb-1 font-semibold text-1xl">
                                Password
                                <input type="password" name="password" 
                                    className="w-full rounded-sm px-3 block border text-gray-300 focus:outline-none 
                                    focus:ring-green-600 focus:border-green-600 focus:text-green-600 focus:z-10 bg-gray-800" />
                            </label>
                        </div>
                        {/* <div>
                            <label className="block pb-1 font-semibold text-1xl">
                                Confirm password
                                <input type="password" name="password2" value={this.state.password2} onChange={this.handleChange}
                                    className="w-full rounded-sm px-3 block border text-gray-300 focus:outline-none 
                                    focus:ring-green-600 focus:border-green-600 focus:text-green-600 focus:z-10 bg-gray-800" />
                            </label>
                        </div> */}
                        <div>
                            <button type="submit" className="w-full flex justify-center py-3 border border-transparent text-sm font-medium rounded-sm
                            text-gray-300 bg-green-600 hover:bg-green-700 focus:ring-2 focus:ring-offset-2 focus:outline-none focus:ring-green-600">
                                Sign up
                            </button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    );
}

export default SignUpPage;
