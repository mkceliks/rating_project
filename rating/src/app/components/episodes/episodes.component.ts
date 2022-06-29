import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { sportEpisode } from 'src/app/models/sport_episode';
import { EpisodesService } from 'src/app/services/episodes.service';

@Component({
  selector: 'app-episodes',
  templateUrl: './episodes.component.html',
  styleUrls: ['./episodes.component.css']
})
export class EpisodesComponent implements OnInit {

  
sport_episodes : sportEpisode[] = [];

  constructor(private episodeService:EpisodesService,private activatedRoot:ActivatedRoute) { }

  ngOnInit(): void {

    this.activatedRoot.params.subscribe(params => {

      if(params["productId"]){
        this.getSportEpisodesById(params["productId"])
      }

    })
  }

  getSportEpisodesById(productId:string){
    this.episodeService.getSportsById(productId).subscribe(response => {
      this.sport_episodes = response
      
    })
  }

}
