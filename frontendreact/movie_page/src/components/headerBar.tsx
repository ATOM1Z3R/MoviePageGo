import React from "react";
import { Link } from "react-router-dom";

const HeaderBar = () => {
  return (
        <nav className="flex bg-gray-800">
            <div className="px-8 py-4 text-white text-2xl">
                <h3>Movie Page</h3>
            </div>
            <div className="flex mr-0 ml-auto font-semibold">
                <Link to="/movies">
                    <h3 className="block px-10 py-5 no-underline text-gray-300 hover:bg-gray-200 hover:text-green-600">Movies</h3>
                </Link>
                <Link to="/directors">
                    <h3 className="block px-10 py-5 no-underline text-gray-300 hover:bg-gray-900 hover:text-green-600">Directors</h3>
                </Link>
                <Link to="/genres">
                    <h3 className="block px-10 py-5 no-underline text-gray-300 hover:bg-gray-900 hover:text-green-600">Genres</h3>
                </Link>
            </div>
            <div className="flex justify-between mr-0 ml-auto font-semibold">
                <Link to="/login">
                    <h3 className="block px-10 py-5 no-underline text-gray-300 hover:bg-gray-900 hover:text-green-600">Login</h3>
                </Link>
                <Link to="/signup">
                    <h3 className="block px-10 py-5 no-underline text-gray-300 hover:bg-gray-900 hover:text-green-600">Signup</h3>
                </Link>
            </div>
        </nav>
    );
}

export default HeaderBar