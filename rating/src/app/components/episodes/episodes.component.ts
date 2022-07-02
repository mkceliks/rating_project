import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { animeEpisode } from 'src/app/models/anime_episode';
import { movieEpisode } from 'src/app/models/movie_episode';
// import { Sport } from 'src/app/models/sport';
import { sportEpisode } from 'src/app/models/sport_episode';
import { EpisodesService } from 'src/app/services/episodes.service';
// import { MovieService } from 'src/app/services/movie.service';

@Component({
  selector: 'app-episodes',
  templateUrl: './episodes.component.html',
  styleUrls: ['./episodes.component.css']
})
export class EpisodesComponent implements OnInit {

sport_episodes : sportEpisode[] = [];
anime_episodes : animeEpisode[] = [];
movie_episodes : movieEpisode[] = [];

  constructor(private episodeService:EpisodesService,private activatedRoot:ActivatedRoute) { }

  ngOnInit(): void {


    this.activatedRoot.params.subscribe(params => {

      if(params["productId"]){
        this.getSportEpisodesById(params["productId"]);
        this.getAnimeEpisodesById(params["productId"]);
        this.getMovieEpisodesById(params["productId"]);
      }
      
    })
  }

  getSportEpisodesById(productId:string){
    this.episodeService.getSportsById(productId).subscribe(response => {
      this.sport_episodes = response
      
    })
  }

  getAnimeEpisodesById(productId:string){
    this.episodeService.getAnimesById(productId).subscribe(response => {
      this.anime_episodes = response
      
    })
  }

  getMovieEpisodesById(productId:string){
    this.episodeService.getMoviesById(productId).subscribe(response => {
      this.movie_episodes = response
      
    })
  }

}
