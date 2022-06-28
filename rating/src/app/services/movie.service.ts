import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Movie } from '../models/movie';
import { Anime } from '../models/anime';
import { Sport } from '../models/sport';

@Injectable({
  providedIn: 'root'
})
export class MovieService {

  apiUrl = 'http://localhost:8000/'

  constructor(private httpClient: HttpClient) { }

  getMovies():Observable<Movie[]> {

    let newPath = this.apiUrl + "movies"

    return this.httpClient.get<Movie[]>(newPath);

  }
  
  getAnimes():Observable<Anime[]> {

    let newPath = this.apiUrl + "animes"

    return this.httpClient.get<Anime[]>(newPath);

  }

  getSports():Observable<Sport[]> {

    let newPath = this.apiUrl + "sports"

    return this.httpClient.get<Sport[]>(newPath);

  }

}
