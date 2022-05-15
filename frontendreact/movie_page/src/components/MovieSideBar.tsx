import { Dispatch, FC, SetStateAction } from "react";
import { MovieApi } from "../api/movie-api";
import { MovieModel, SimpleMovieModel } from "../model/movie";
import SingleItem from "./SingleItem";

interface Props {
    movies: SimpleMovieModel[],
    setMovie: Dispatch<SetStateAction<MovieModel | undefined>>
}

const movieApi = new MovieApi();

const MovieSideBar:FC<Props> = ({movies, setMovie}) => {
    const selectMovie = (movieId: number) =>{
        movieApi.getMovieByIdApiCallAsync(movieId).then((movie) => {
            setMovie(movie);
          });
    };
    return (
        <div className="top-0 bottom-0 lg:left-0 left-[-300px] duration-1000
            p-2 w-[300px] overflow-y-auto text-center bg-gray-900 shadow h-screen">
            <div className="text-gray-100 text-xl">
                <div>
                    {movies.map((movie) => (
                        <span onClick={()=>selectMovie(movie.id ?? 0)}>
                            <SingleItem movie={movie} />
                        </span>
                    ))}
                </div>
            </div>
        </div>
    )
}

export default MovieSideBar;