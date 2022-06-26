import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Episode } from '../models/episode';

@Injectable({
  providedIn: 'root'
})
export class EpisodesService {

  apiUrl = 'http://localhost:8000/episodes'

  constructor(private httpClient: HttpClient) { }

  getEpisodes():Observable<Episode[]> {

    return this.httpClient.get<Episode[]>(this.apiUrl);

  }
}
