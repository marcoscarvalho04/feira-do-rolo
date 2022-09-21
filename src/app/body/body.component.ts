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
      this.offers.forEach(m => {
        m.contato = m.contato.length < 20 ? m.contato :  m.contato.substring(0,17).concat('...')
        m.descricaoProduto = m.descricaoProduto.length< 20 ? m.descricaoProduto :  m.descricaoProduto.substring(0,17).concat('...')
        m.tituloProduto = m.tituloProduto.length < 20 ? m.tituloProduto : m.tituloProduto.substring(0,17).concat('...')
      })
    })
  }

}
