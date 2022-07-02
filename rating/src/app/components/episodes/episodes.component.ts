import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
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

// sports : Sport[] = [];
sport_episodes : sportEpisode[] = [];

  constructor(private episodeService:EpisodesService,private activatedRoot:ActivatedRoute) { }

  ngOnInit(): void {


    this.activatedRoot.params.subscribe(params => {

      if(params["productId"]){
        this.getSportEpisodesById(params["productId"])
      }

    })
    // this.getSports();

    // this.episodesSport();
  }

  getSportEpisodesById(productId:string){
    this.episodeService.getSportsById(productId).subscribe(response => {
      this.sport_episodes = response
      
    })
  }

  // getSports(){
  //   this.movieService.getSports().subscribe(response => {
  //     this.sports = response
  //   })
  // }

  // episodesSport(){
  //   var j = 0;
  //   console.log(this.sport_episodes[0].sport_id)
  //   console.log(this.sports[0]._id)
  //   for(var i = 0; i< this.sports.length; i++ ){
  //     if (this.sport_episodes[j].sport_id == this.sports[i]._id){
  //       var game_title = this.sports[i].title
  //       console.log(game_title)
  //     }
  //   }
  //   return game_title
  // }

}
