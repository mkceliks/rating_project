import { Component, OnInit } from '@angular/core';
import { sportEpisode } from 'src/app/models/sport_episode';
import { EpisodesService } from 'src/app/services/episodes.service';
import { ActivatedRoute } from '@angular/router';
@Component({
  selector: 'app-sport-episodes',
  templateUrl: './sport-episodes.component.html',
  styleUrls: ['./sport-episodes.component.css']
})
export class SportEpisodesComponent implements OnInit {
  
  sport_episodes : sportEpisode[] = [];

  constructor(private episodeService:EpisodesService,private activatedRoot:ActivatedRoute) { }

  ngOnInit(): void {

    
    this.activatedRoot.params.subscribe(params => {

      if(params["productId"]){
        this.getSportEpisodesById(params["productId"]);
      }
      
    })

  }

  getSportEpisodesById(productId:string){
    this.episodeService.getSportsById(productId).subscribe(response => {
      this.sport_episodes = response
      
    })
  }

}
