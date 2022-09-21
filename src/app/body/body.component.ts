import { Component, OnInit } from '@angular/core';
import { Offers } from '../model/offers';
import { OffersService } from '../services/offers.service';



@Component({
  selector: 'app-body',
  templateUrl: './body.component.html',
  styleUrls: ['./body.component.css']
})
export class BodyComponent implements OnInit {

  offers: Offers[] = [];

  constructor(private offersService:OffersService) { }

  ngOnInit(): void {
    this.offersService.findAll().subscribe(data => {
      this.offers = data
    })
  }

}
