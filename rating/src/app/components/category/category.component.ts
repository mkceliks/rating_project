import { Component, OnInit } from '@angular/core';
import { Anime } from 'src/app/models/anime';
import { Movie } from 'src/app/models/movie';
import { Sport } from 'src/app/models/sport';
import { MovieService } from 'src/app/services/movie.service';

@Component({
  selector: 'app-category',
  templateUrl: './category.component.html',
  styleUrls: ['./category.component.css']
})
export class CategoryComponent implements OnInit {

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
