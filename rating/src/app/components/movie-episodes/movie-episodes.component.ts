import { Component, OnInit } from '@angular/core';
import { movieEpisode } from 'src/app/models/movie_episode';
import { EpisodesService } from 'src/app/services/episodes.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-movie-episodes',
  templateUrl: './movie-episodes.component.html',
  styleUrls: ['./movie-episodes.component.css']
})
export class MovieEpisodesComponent implements OnInit {
  movie_episodes : movieEpisode[] = [];

  constructor(private episodeService:EpisodesService,private activatedRoot:ActivatedRoute) { }

  ngOnInit(): void {

    this.activatedRoot.params.subscribe(params => {

      if(params["productId"]){
        this.getMovieEpisodesById(params["productId"]); 
      }
      
    })
  }

  getMovieEpisodesById(productId:string){
    this.episodeService.getMoviesById(productId).subscribe(response => {
      this.movie_episodes = response
      
    })
  }

}
