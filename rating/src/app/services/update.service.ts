import { Injectable } from '@angular/core';
import { Anime } from '../models/anime';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UpdateService {

  constructor(private httpClient:HttpClient) { }

  apiUrl = 'http://localhost:8000/'

  updateAnime(animeId:string,newAnime:Anime):Observable<any>{

    let newPath = this.apiUrl + "anime-update/" + animeId

    return this.httpClient.put(newPath,newAnime)

  }



}
