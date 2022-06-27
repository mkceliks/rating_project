import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Movie } from '../models/movie';
@Injectable({
  providedIn: 'root'
})
export class ProductAddService {

  apiUrl = 'http://localhost:8000/add-anime'

  constructor(private httpClient: HttpClient) { }

  add(product:Movie){

    return this.httpClient.post(this.apiUrl,product)

  }

}
