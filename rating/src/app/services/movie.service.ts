import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Movie } from '../models/movie';

@Injectable({
  providedIn: 'root'
})
export class MovieService {

  apiUrl = 'http://localhost:8000/movies'

  constructor(private httpClient: HttpClient) { }

  getMovies():Observable<Movie[]> {

    return this.httpClient.get<Movie[]>(this.apiUrl);

  }

}
