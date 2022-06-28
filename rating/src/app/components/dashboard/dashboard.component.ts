import { Component, OnInit } from '@angular/core';
import { Movie } from 'src/app/models/movie';
import { MovieService } from 'src/app/services/movie.service';
import { Anime } from 'src/app/models/anime';
import { Sport } from 'src/app/models/sport';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {

  movies : Movie[] = [];
  animes : Anime[] = [];
  sports : Sport[] = [];

  constructor(private movieService:MovieService) { }

  ngOnInit(): void {
    this.getMovies();
    this.getAnimes();
    this.getSports();
  }



  getMovies(){
    this.movieService.getMovies().subscribe(response => {
      this.movies = response
    })
  }

  getAnimes(){
    this.movieService.getAnimes().subscribe(response => {
      this.animes = response
    })
  }

  getSports(){
    this.movieService.getSports().subscribe(response => {
      this.sports = response
    })
  }

}
