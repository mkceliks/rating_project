import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { MoviesComponent } from './components/movies/movies.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { NaviComponent } from './components/navi/navi.component';
import { AddProductComponent } from './components/add-product/add-product.component';
import { FormsModule , ReactiveFormsModule } from '@angular/forms';
import { CategoryComponent } from './components/category/category.component';
import { AnimesComponent } from './components/animes/animes.component';
import { SportsComponent } from './components/sports/sports.component';
import { AddEpisodeComponent } from './components/add-episode/add-episode.component';
import { AnimeEpisodesComponent } from './components/anime-episodes/anime-episodes.component';
import { SportEpisodesComponent } from './components/sport-episodes/sport-episodes.component';
import { MovieEpisodesComponent } from './components/movie-episodes/movie-episodes.component';

@NgModule({
  declarations: [
    AppComponent,
    MoviesComponent,
    DashboardComponent,
    NaviComponent,
    AddProductComponent,
    CategoryComponent,
    AnimesComponent,
    SportsComponent,
    AddEpisodeComponent,
    AnimeEpisodesComponent,
    SportEpisodesComponent,
    MovieEpisodesComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule,
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule { }
