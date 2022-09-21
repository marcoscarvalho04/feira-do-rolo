import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Routes } from '@angular/router';

import {SaveOfferComponent} from './save-offer/save-offer.component'
import { BodyComponent } from './body/body.component';
import { HeaderComponent } from './header/header.component';

const appRoutes: Routes = [
  { path: 'save-offer', component: SaveOfferComponent},
  { path: '', component: BodyComponent}
]

@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    RouterModule.forRoot(appRoutes)
  ],
  exports: [RouterModule]
})
export class AppRoutingModule { }
