import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Offers } from '../model/offers';
import { ActivatedRoute } from '@angular/router';
import { OffersService } from '../services/offers.service';
import { NgForm } from '@angular/forms';

@Component({
  selector: 'app-save-offer',
  templateUrl: './save-offer.component.html',
  styleUrls: ['./save-offer.component.css']
})
export class SaveOfferComponent implements OnInit {

  offer: Offers
  constructor(private route: ActivatedRoute, private router: Router, private offerService: OffersService) {
    this.offer = new Offers();
  }

  ngOnInit(): void {
  }

  onSubmit(){
    console.log(this.offer)
    this.offerService.save(this.offer).subscribe(result => this.goToIndex())
  }

  goToIndex(){
    this.router.navigate([''])
  }

}
