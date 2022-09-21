import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Offers } from '../model/offers';
import { ActivatedRoute } from '@angular/router';
import { OffersService } from '../services/offers.service';
import { FormGroup, FormControl, Validators, FormGroupDirective } from '@angular/forms';



@Component({
  selector: 'app-save-offer',
  templateUrl: './save-offer.component.html',
  styleUrls: ['./save-offer.component.css']

})
export class SaveOfferComponent implements OnInit {

  offer: Offers
  offerForm: FormGroup

  constructor(private route: ActivatedRoute, private router: Router, private offerService: OffersService) {
    this.offer = new Offers();
    this.offerForm = new FormGroup({
      tituloProduto: new FormControl(null, [Validators.required, Validators.maxLength(200)]),
      descricaoProduto: new FormControl(null, [Validators.required, Validators.maxLength(200)]),
      contato: new FormControl(null, [Validators.required, Validators.maxLength(200)])
    })
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
