import { useEffect, useState } from "react";
import { MovieApi } from "../api/movie-api";
import MovieContainer from "../components/MovieContainer";
import MovieSideBar from "../components/MovieSideBar";
import { MovieModel, SimpleMovieModel } from "../model/movie";

const movieApi = new MovieApi();

const MoviesPage = () => {
  const [movies, setMovies] = useState<SimpleMovieModel[]>([]);
  const [movie, setMovie] = useState<MovieModel | undefined>();
  
  useEffect(() => {
    movieApi.getAllMoviesApiCallAsync().then((movies) => {
      setMovies(movies);
    });
    // movieApi.getMovieByIdApiCallAsync(8).then((movie) => {
    //   setMovie(movie);
    // });
  });
  
  return (
    <div className="flex flex-row mx-auto w-max m-0 p-0 text-gray-200">
        <MovieSideBar movies={movies} setMovie={setMovie} />
        <MovieContainer movie={movie} />
    </div>
  )
}

export default MoviesPage;