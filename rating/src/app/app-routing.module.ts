import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AddEpisodeComponent } from './components/add-episode/add-episode.component';
import { AddProductComponent } from './components/add-product/add-product.component';
import { AnimeEpisodesComponent } from './components/anime-episodes/anime-episodes.component';
import { AnimesComponent } from './components/animes/animes.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { MovieEpisodesComponent } from './components/movie-episodes/movie-episodes.component';
import { MoviesComponent } from './components/movies/movies.component';
import { SportEpisodesComponent } from './components/sport-episodes/sport-episodes.component';
import { SportsComponent } from './components/sports/sports.component';

const routes: Routes = [
  {path:"",component:DashboardComponent},
  {path:"animes",component:AnimesComponent},
  {path:"movies",component:MoviesComponent},
  {path:"sports",component:SportsComponent},
  {path:"add-product",component:AddProductComponent},
  {path:"add-episode",component:AddEpisodeComponent},
  {path:"sport-episodes/:productId",component:SportEpisodesComponent},
  {path:"anime-episodes/:productId",component:AnimeEpisodesComponent},
  {path:"movie-episodes/:productId",component:MovieEpisodesComponent},
  {path:"delete-anime/:productId",component:DashboardComponent},
  {path:"delete-movie/:productId",component:DashboardComponent},
  {path:"delete-sport/:productId",component:DashboardComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
