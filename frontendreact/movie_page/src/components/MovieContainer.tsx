import React, { FC } from 'react'
import { MovieModel } from '../model/movie'

interface Props {
    movie: MovieModel | undefined
}

const MovieContainer: FC<Props> = ({movie}) => {
  return (
    <div className="top-0 bottom-0 p-4 mx-4 w-[1000px] overflow-y-auto text-center bg-gray-900 shadow h-screen">
        <div className="flex w-full h-[250px]">
          <div className="w-1/4 bg-red-500">
                sdfsdfsdf
          </div>
          <div className="px-8">
            <div className="text-4xl font-semibold">
              {movie?.title} ({movie?.premiereDate?.slice(-4)})
            </div>
            <div className="text-2xl">
              Gatunek: {movie?.genre?.name}
            </div>
            <div className="text-xl">
              Reżyseria:
              {movie?.directors?.map((director) => (
                  <div>
                    {director.firstName} {director.lastName}
                  </div>
              ))}
            </div>
            <div className="text-xl">
              Lokacja: {movie?.location}
            </div>
          </div>
        </div>
        <div className="descripton">
          {movie?.description}
        </div>
        <div className="teraser">

        </div>
        <div className="comments">
        
        </div>
      </div>
  )
}

export default MovieContainer