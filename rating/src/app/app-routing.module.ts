import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AddProductComponent } from './components/add-product/add-product.component';
import { AnimesComponent } from './components/animes/animes.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { MoviesComponent } from './components/movies/movies.component';
import { SportsComponent } from './components/sports/sports.component';

const routes: Routes = [
  {path:"",component:DashboardComponent},
  {path:"animes",component:AnimesComponent},
  {path:"movies",component:MoviesComponent},
  {path:"sports",component:SportsComponent},
  {path:"add-anime",component:AddProductComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
