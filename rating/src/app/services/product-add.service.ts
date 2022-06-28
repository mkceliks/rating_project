import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Movie } from '../models/movie';
import { Anime } from '../models/anime';
import { Sport } from '../models/sport';
@Injectable({
  providedIn: 'root'
})
export class ProductAddService {

  apiUrl = 'http://localhost:8000/'

  constructor(private httpClient: HttpClient) { }

  addMovie(product:Movie){

    let newPath = this.apiUrl + "add-movie"

    return this.httpClient.post(newPath,product)

  }

  addAnime(product:Anime){

    let newPath = this.apiUrl + "add-anime"

    return this.httpClient.post(newPath,product)

  }

  addSport(product:Sport){

    let newPath = this.apiUrl + "add-sport"

    return this.httpClient.post(newPath,product)

  }

}
