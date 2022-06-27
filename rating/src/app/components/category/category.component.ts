import { Component, OnInit } from '@angular/core';
import { Movie } from 'src/app/models/movie';
import { MovieService } from 'src/app/services/movie.service';

@Component({
  selector: 'app-category',
  templateUrl: './category.component.html',
  styleUrls: ['./category.component.css']
})
export class CategoryComponent implements OnInit {

  movies : Movie[] = [];

  constructor(private movieService:MovieService) { }

  ngOnInit(): void {
    this.getMovies();
  }


  getMovies(){
    this.movieService.getMovies().subscribe(response => {
      this.movies = response
    })
  }

}
