import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppComponent } from './app.component';
import { BodyComponent } from './body/body.component';
import { HeaderComponent } from './header/header.component';
import { SaveOfferComponent } from './save-offer/save-offer.component';
import { AppRoutingModule } from './app.routing.module';

@NgModule({
  declarations: [
    AppComponent,
    BodyComponent,
    HeaderComponent,
    SaveOfferComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
