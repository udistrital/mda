import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CommonModule } from '@angular/common';
import { [Entity]Component } from '../[entity]/view/[entity].component';
import { [Entity]NewComponent } from '../[entity]/new/[entity]-new.component';
import { [Entity]EditComponent } from '../[entity]/edit/[entity]-edit.component';


const routes: Routes = [
  { path: '[entity]', component: [Entity]Component },
  { path: '[entity]/new', component: [Entity]NewComponent },
  { path: '[entity]/edit/:id', component: [Entity]EditComponent },
];

@NgModule({
  imports: [
    CommonModule,
    RouterModule.forRoot(routes)
  ],
  exports: [RouterModule],
  declarations: []
})
export class RoutingModule { }
