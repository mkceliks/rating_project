import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { animeEpisode } from '../models/anime_episode';
import { sportEpisode } from '../models/sport_episode';
import { movieEpisode } from '../models/movie_episode';

@Injectable({
  providedIn: 'root'
})
export class EpisodesService {

  apiUrl = 'http://localhost:8000/'

  constructor(private httpClient: HttpClient) { }

  ////////////////////////////////////////////// Getters

  getAnimeEpisodes():Observable<animeEpisode[]> {

    let newPath = this.apiUrl + "anime-episodes"

    return this.httpClient.get<animeEpisode[]>(this.apiUrl);

  }

  getMovieEpisodes():Observable<movieEpisode[]> {

    let newPath = this.apiUrl + "movie-episodes"

    return this.httpClient.get<movieEpisode[]>(this.apiUrl);

  }

  getSportEpisodes():Observable<sportEpisode[]> {

    let newPath = this.apiUrl + "sport-episodes"

    return this.httpClient.get<sportEpisode[]>(this.apiUrl);

  }

  //////////////////////////////////////////////// Setters

  addAnimeEpisode(episode:animeEpisode){

    let newPath = this.apiUrl + "anime-episodes/add"

    return this.httpClient.post(newPath,episode)

  }

  addMovieEpisode(episode:movieEpisode){

    let newPath = this.apiUrl + "movie-episodes/add"

    return this.httpClient.post(newPath,episode)

  }

  addSportEpisode(episode:sportEpisode){

    let newPath = this.apiUrl + "sport-episodes/add"

    return this.httpClient.post(newPath,episode)

  }

  getSportsById(sportId:string):Observable<sportEpisode[]> {

    let newPath = this.apiUrl + "products/getepisodebyid?sportid="+ sportId

    return this.httpClient.get<sportEpisode[]>(newPath);

  }
}
