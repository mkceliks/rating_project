import { Component, OnInit } from '@angular/core';
import {
  FormGroup,
  FormBuilder,
  Validators,
  FormControl,
} from '@angular/forms';
import { Anime } from 'src/app/models/anime';
import { Movie } from 'src/app/models/movie';
import { Sport } from 'src/app/models/sport';
import { EpisodesService } from 'src/app/services/episodes.service';
import { MovieService } from 'src/app/services/movie.service';



@Component({
  selector: 'app-add-episode',
  templateUrl: './add-episode.component.html',
  styleUrls: ['./add-episode.component.css']
})
export class AddEpisodeComponent implements OnInit {

movies : Movie[] = [];
animes : Anime[] = [];
sports : Sport[] = [];


animeEpisodeAddForm : FormGroup;
movieEpisodeAddForm : FormGroup;
sportEpisodeAddForm : FormGroup;

  constructor(
    private formBuilder: FormBuilder,
    private episodesService:EpisodesService,
    private movieService:MovieService
   ) { }

  ngOnInit(): void {
    this.getAnimes();
    this.getMovies();
    this.getSports();
    this.createEpisodeAddForm();


  }

  createEpisodeAddForm() {
    this.animeEpisodeAddForm = this.formBuilder.group({
      anime_id: ['', Validators.required],
      title: ['', Validators.required],
      description: ['', Validators.required],
      duration: ['', Validators.required],
    });

    this.sportEpisodeAddForm = this.formBuilder.group({
      sport_id: ['', Validators.required],
      title: ['', Validators.required],
      description: ['', Validators.required],
      duration: ['', Validators.required],
    });

    this.movieEpisodeAddForm = this.formBuilder.group({
      movie_id: ['', Validators.required],
      title: ['', Validators.required],
      description: ['', Validators.required],
      duration: ['', Validators.required],
    });
  }

  addAnimeEpisode() {
    if (this.animeEpisodeAddForm.valid) {
      let episodeModel = Object.assign({}, this.animeEpisodeAddForm.value);
      this.episodesService.addAnimeEpisode(episodeModel).subscribe((response) => {
        alert('Addition is successful.');
      });
    } else {
      alert('Formunuz Geçersizdir.');
    }
  }
  addMovieEpisode() {
    if (this.movieEpisodeAddForm.valid) {
      let episodeModel = Object.assign({}, this.movieEpisodeAddForm.value);
      this.episodesService.addMovieEpisode(episodeModel).subscribe((response) => {
        alert('Addition is successful.');
      });
    } else {
      alert('Formunuz Geçersizdir.');
    }
  }
  addSportEpisode() {
    console.log(this.sportEpisodeAddForm)
    if (this.sportEpisodeAddForm.valid) {
      let episodeModel = Object.assign({}, this.sportEpisodeAddForm.value);
      this.episodesService.addSportEpisode(episodeModel).subscribe((response) => {
        alert('Addition is successful.');
      });
    } else {
      alert('Formunuz Geçersizdir.');
    }
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
}
