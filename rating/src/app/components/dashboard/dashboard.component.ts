import { Component, OnInit } from '@angular/core';
import { Movie } from 'src/app/models/movie';
import { MovieService } from 'src/app/services/movie.service';
import { Anime } from 'src/app/models/anime';
import { Sport } from 'src/app/models/sport';
import { ActivatedRoute } from '@angular/router';
import { Router } from '@angular/router';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  
  deletedAnime : string;
  movies : Movie[] = [];
  animes : Anime[] = [];
  sports : Sport[] = [];

  constructor(private movieService:MovieService,private activatedRoot:ActivatedRoute,private router:Router) { }

  ngOnInit(): void {
    this.getMovies();
    this.getAnimes();
    this.getSports();

    this.activatedRoot.params.subscribe(params => {

      if(params["productId"]){
        this.deleteAnime(params["productId"]);
      }
      
    })
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

  deleteAnime(animeId:string){
    this.movieService.getAnimes().subscribe(response => {
      this.animes = response
      this.deletedAnime = this.animes.filter(item => item._id == animeId)[0].title;
      this.movieService.deleteAnime(animeId).subscribe(data =>{
        alert(this.deletedAnime + " successfully deleted")
        console.log(data)
      });
    })
    this.router.navigate(['']);
  }

}
