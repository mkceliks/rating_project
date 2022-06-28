import { Component, OnInit } from '@angular/core';
import { Anime } from 'src/app/models/anime';
import { MovieService } from 'src/app/services/movie.service';

@Component({
  selector: 'app-animes',
  templateUrl: './animes.component.html',
  styleUrls: ['./animes.component.css']
})
export class AnimesComponent implements OnInit {

  animes : Anime [] = []

  constructor(private movieService:MovieService) { }

  ngOnInit(): void {
    this.getAnimes();
  }

  getAnimes(){
    this.movieService.getAnimes().subscribe(response => {
      this.animes = response
    })
  }

}
