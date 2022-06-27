import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AddProductComponent } from './components/add-product/add-product.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { MoviesComponent } from './components/movies/movies.component';

const routes: Routes = [
  {path:"",component:DashboardComponent},
  {path:"animes",component:MoviesComponent},
  {path:"add-anime",component:AddProductComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
