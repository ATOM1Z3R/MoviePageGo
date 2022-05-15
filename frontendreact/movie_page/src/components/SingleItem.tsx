import React, { FC } from 'react'
import { SimpleMovieModel } from '../model/movie'
interface Props {
    movie: SimpleMovieModel,
}

const SingleItem: FC<Props> = ({movie}) => {
  return (
    <div className="p-2.5 mt-2 flex items-center rounded-md px-4 duration-300 cursor-pointer  hover:bg-blue-600">
      <span className="text-[15px] ml-4 text-gray-200">{movie.title} {movie.premiereDate}</span>
    </div>
  )
}

export default SingleItem