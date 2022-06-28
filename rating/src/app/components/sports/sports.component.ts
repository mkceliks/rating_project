import { Component, OnInit } from '@angular/core';
import { Sport } from 'src/app/models/sport';
import { MovieService } from 'src/app/services/movie.service';

@Component({
  selector: 'app-sports',
  templateUrl: './sports.component.html',
  styleUrls: ['./sports.component.css']
})
export class SportsComponent implements OnInit {

  sports : Sport [] = [];

  constructor(private movieService:MovieService) { }

  ngOnInit(): void {
    this.getSports();
  }

  getSports(){
    this.movieService.getSports().subscribe(response => {
      this.sports = response
    })
  }

}
