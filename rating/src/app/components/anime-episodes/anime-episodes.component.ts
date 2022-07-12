import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { animeEpisode } from 'src/app/models/anime_episode';
import { EpisodesService } from 'src/app/services/episodes.service';
@Component({
  selector: 'app-anime-episodes',
  templateUrl: './anime-episodes.component.html',
  styleUrls: ['./anime-episodes.component.css']
})
export class AnimeEpisodesComponent implements OnInit {
  anime_episodes : animeEpisode[] = [];

  constructor(private episodeService:EpisodesService,private activatedRoot:ActivatedRoute) { }

  ngOnInit(): void {

    
    this.activatedRoot.params.subscribe(params => {

      if(params["productId"]){
        this.getAnimeEpisodesById(params["productId"]);
      }
      
    })

  }

  getAnimeEpisodesById(productId:string){
    this.episodeService.getAnimesById(productId).subscribe(response => {
      this.anime_episodes = response
      
    })
  }

}
