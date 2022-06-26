import { Component, OnInit } from '@angular/core';
import { Episode } from 'src/app/models/episode';
import { Movie } from 'src/app/models/movie';
import { EpisodesService } from 'src/app/services/episodes.service';
import { MovieService } from 'src/app/services/movie.service';

@Component({
  selector: 'app-movies',
  templateUrl: './movies.component.html',
  styleUrls: ['./movies.component.css']
})
export class MoviesComponent implements OnInit {

  movies : Movie[] = [];
  episodes : Episode[] = [];

  constructor(private movieService:MovieService,private episodeService:EpisodesService) { }

  ngOnInit(): void {
    this.getMovies();
    this.getEpisodes();
  }

  getMovies(){
    this.movieService.getMovies().subscribe(response => {
      this.movies = response
    })
  }

  getEpisodes(){
    this.episodeService.getEpisodes().subscribe(response => {
      this.episodes = response
    })
  }

}
