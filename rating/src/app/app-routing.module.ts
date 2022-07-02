import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AddEpisodeComponent } from './components/add-episode/add-episode.component';
import { AddProductComponent } from './components/add-product/add-product.component';
import { AnimesComponent } from './components/animes/animes.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { EpisodesComponent } from './components/episodes/episodes.component';
import { MoviesComponent } from './components/movies/movies.component';
import { SportsComponent } from './components/sports/sports.component';

const routes: Routes = [
  {path:"",component:DashboardComponent},
  {path:"animes",component:AnimesComponent},
  {path:"movies",component:MoviesComponent},
  {path:"sports",component:SportsComponent},
  {path:"add-product",component:AddProductComponent},
  {path:"add-episode",component:AddEpisodeComponent},
  {path:"sport-episodes/:productId",component:EpisodesComponent},
  {path:"anime-episodes/:productId",component:EpisodesComponent},
  {path:"movie-episodes/:productId",component:EpisodesComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
